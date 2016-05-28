package out

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/pivotal-cf-experimental/pivnet-resource/concourse"
	"github.com/pivotal-cf-experimental/pivnet-resource/globs"
	"github.com/pivotal-cf-experimental/pivnet-resource/logger"
	"github.com/pivotal-cf-experimental/pivnet-resource/md5sum"
	"github.com/pivotal-cf-experimental/pivnet-resource/metadata"
	"github.com/pivotal-cf-experimental/pivnet-resource/pivnet"
	"github.com/pivotal-cf-experimental/pivnet-resource/s3"
	"github.com/pivotal-cf-experimental/pivnet-resource/uploader"
	"github.com/pivotal-cf-experimental/pivnet-resource/useragent"
	"github.com/pivotal-cf-experimental/pivnet-resource/validator"
	"github.com/pivotal-cf-experimental/pivnet-resource/versions"
)

const (
	defaultBucket = "pivotalnetwork"
	defaultRegion = "eu-west-1"
)

type OutCommand struct {
	binaryVersion   string
	logger          logger.Logger
	outDir          string
	sourcesDir      string
	logFilePath     string
	s3OutBinaryName string
	screenWriter    *log.Logger
}

type OutCommandConfig struct {
	BinaryVersion   string
	Logger          logger.Logger
	OutDir          string
	SourcesDir      string
	LogFilePath     string
	S3OutBinaryName string
	ScreenWriter    *log.Logger
}

func NewOutCommand(config OutCommandConfig) *OutCommand {
	return &OutCommand{
		binaryVersion:   config.BinaryVersion,
		logger:          config.Logger,
		outDir:          config.OutDir,
		sourcesDir:      config.SourcesDir,
		logFilePath:     config.LogFilePath,
		s3OutBinaryName: config.S3OutBinaryName,
		screenWriter:    config.ScreenWriter,
	}
}

func (c *OutCommand) Run(input concourse.OutRequest) (concourse.OutResponse, error) {
	if c.outDir == "" {
		return concourse.OutResponse{}, fmt.Errorf("%s must be provided", "out dir")
	}

	var m metadata.Metadata
	var skipFileCheck bool
	if input.Params.MetadataFile != "" {
		metadataFilepath := filepath.Join(c.sourcesDir, input.Params.MetadataFile)
		metadataBytes, err := ioutil.ReadFile(metadataFilepath)
		if err != nil {
			return concourse.OutResponse{}, fmt.Errorf("metadata_file could not be read: %s", err.Error())
		}

		err = yaml.Unmarshal(metadataBytes, &m)
		if err != nil {
			return concourse.OutResponse{}, fmt.Errorf("metadata_file could not be parsed: %s", err.Error())
		}

		err = m.Validate()
		if err != nil {
			return concourse.OutResponse{}, fmt.Errorf("metadata_file is invalid: %s", err.Error())
		}

		skipFileCheck = true
	}

	c.logger.Debugf("metadata product_files parsed; contents: %+v\n", m.ProductFiles)

	if m.Release != nil {
		c.logger.Debugf("metadata release parsed; contents: %+v\n", *m.Release)
	}

	warnIfDeprecatedFilesFound(input.Params, c.logger, c.screenWriter)

	err := validator.NewOutValidator(input, skipFileCheck).Validate()
	if err != nil {
		return concourse.OutResponse{}, err
	}

	c.logger.Debugf("Received input: %+v\n", input)

	globber := globs.NewGlobber(globs.GlobberConfig{
		FileGlob:   input.Params.FileGlob,
		SourcesDir: c.sourcesDir,
		Logger:     c.logger,
	})

	exactGlobs, err := globber.ExactGlobs()
	if err != nil {
		return concourse.OutResponse{}, err
	}

	var missingFiles []string
	for _, f := range m.ProductFiles {
		if !contains(exactGlobs, f.File) {
			missingFiles = append(missingFiles, f.File)
		}
	}

	if len(missingFiles) > 0 {
		return concourse.OutResponse{},
			fmt.Errorf("product_files were provided in metadata that match no globs: %v", missingFiles)
	}

	var endpoint string
	if input.Source.Endpoint != "" {
		endpoint = input.Source.Endpoint
	} else {
		endpoint = pivnet.Endpoint
	}

	productSlug := input.Source.ProductSlug

	clientConfig := pivnet.NewClientConfig{
		Endpoint:  endpoint,
		Token:     input.Source.APIToken,
		UserAgent: useragent.UserAgent(c.binaryVersion, "put", productSlug),
	}
	pivnetClient := pivnet.NewClient(
		clientConfig,
		c.logger,
	)

	c.logger.Debugf("Getting all valid eulas\n")

	eulas, err := pivnetClient.EULAs()
	if err != nil {
		return concourse.OutResponse{}, err
	}

	eulaSlugs := make([]string, len(eulas))
	for i, e := range eulas {
		eulaSlugs[i] = e.Slug
	}

	eulaSlugsPrintable := fmt.Sprintf(
		"['%s']",
		strings.Join(eulaSlugs, "', '"),
	)

	c.logger.Debugf("All valid eula slugs: %s\n", eulaSlugsPrintable)

	eulaSlug := fetchFromMetadataOrFile("EULASlug", m, skipFileCheck, c.sourcesDir, input.Params.EULASlugFile)
	if !containsString(eulaSlugs, eulaSlug) {
		return concourse.OutResponse{}, fmt.Errorf(
			"provided eula_slug: '%s' must be one of: %s",
			eulaSlug,
			eulaSlugsPrintable,
		)
	}

	c.logger.Debugf("Getting all valid release types\n")

	releaseTypes, err := pivnetClient.ReleaseTypes()
	if err != nil {
		return concourse.OutResponse{}, err
	}

	releaseTypesPrintable := fmt.Sprintf(
		"['%s']",
		strings.Join(releaseTypes, "', '"),
	)

	c.logger.Debugf("All release types: %s\n", releaseTypesPrintable)

	releaseType := fetchFromMetadataOrFile("ReleaseType", m, skipFileCheck, c.sourcesDir, input.Params.ReleaseTypeFile)
	if !containsString(releaseTypes, releaseType) {
		return concourse.OutResponse{}, fmt.Errorf(
			"provided release_type: '%s' must be one of: %s",
			releaseType,
			releaseTypesPrintable,
		)
	}

	productVersion := fetchFromMetadataOrFile("Version", m, skipFileCheck, c.sourcesDir, input.Params.VersionFile)

	releases, err := pivnetClient.ReleasesForProductSlug(productSlug)
	if err != nil {
		return concourse.OutResponse{}, err
	}

	existingVersions, err := pivnetClient.ProductVersions(productSlug, releases)
	if err != nil {
		return concourse.OutResponse{}, err
	}

	for _, v := range existingVersions {
		if v == productVersion {
			return concourse.OutResponse{}, fmt.Errorf("release already exists with version: %s", productVersion)
		}
	}

	config := pivnet.CreateReleaseConfig{
		ProductSlug:     productSlug,
		ReleaseType:     releaseType,
		EULASlug:        fetchFromMetadataOrFile("EULASlug", m, skipFileCheck, c.sourcesDir, input.Params.EULASlugFile),
		ProductVersion:  productVersion,
		Description:     fetchFromMetadataOrFile("Description", m, skipFileCheck, c.sourcesDir, input.Params.DescriptionFile),
		ReleaseNotesURL: fetchFromMetadataOrFile("ReleaseNotesURL", m, skipFileCheck, c.sourcesDir, input.Params.ReleaseNotesURLFile),
		ReleaseDate:     fetchFromMetadataOrFile("ReleaseDate", m, skipFileCheck, c.sourcesDir, input.Params.ReleaseDateFile),
	}
	if m.Release != nil {
		config.Controlled = m.Release.Controlled
		config.ECCN = m.Release.ECCN
		config.LicenseException = m.Release.LicenseException
		config.EndOfSupportDate = m.Release.EndOfSupportDate
		config.EndOfGuidanceDate = m.Release.EndOfGuidanceDate
		config.EndOfAvailabilityDate = m.Release.EndOfAvailabilityDate
	}

	c.logger.Debugf("config used to create pivnet release: %+v\n", config)

	release, err := pivnetClient.CreateRelease(config)
	if err != nil {
		log.Fatalln(err)
	}

	skipUpload := input.Params.FileGlob == "" && input.Params.FilepathPrefix == ""
	if skipUpload {
		c.logger.Debugf("File glob and s3_filepath_prefix not provided - skipping upload to s3")
	} else {
		bucket := input.Source.Bucket
		if bucket == "" {
			bucket = defaultBucket
		}

		region := input.Source.Region
		if region == "" {
			region = defaultRegion
		}

		logFile, err := os.OpenFile(c.logFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			panic(err)
		}

		s3Client := s3.NewClient(s3.NewClientConfig{
			AccessKeyID:     input.Source.AccessKeyID,
			SecretAccessKey: input.Source.SecretAccessKey,
			RegionName:      region,
			Bucket:          bucket,

			Logger: c.logger,

			Stdout: os.Stdout,
			Stderr: logFile,

			OutBinaryPath: filepath.Join(c.outDir, c.s3OutBinaryName),
		})

		uploaderClient := uploader.NewClient(uploader.Config{
			FilepathPrefix: input.Params.FilepathPrefix,
			SourcesDir:     c.sourcesDir,

			Logger: c.logger,

			Transport: s3Client,
		})

		for _, exactGlob := range exactGlobs {
			fullFilepath := filepath.Join(c.sourcesDir, exactGlob)
			fileContentsMD5, err := md5sum.NewFileSummer().SumFile(fullFilepath)
			if err != nil {
				log.Fatalln(err)
			}

			remotePath, err := uploaderClient.UploadFile(exactGlob)
			if err != nil {
				return concourse.OutResponse{}, err
			}

			product, err := pivnetClient.FindProductForSlug(productSlug)
			if err != nil {
				log.Fatalln(err)
			}

			filename := filepath.Base(exactGlob)

			var description string
			uploadAs := filename
			fileType := "Software"
			for _, f := range m.ProductFiles {
				if f.File == exactGlob {
					c.logger.Debugf("exact glob '%s' matches metadata file: '%s'\n", exactGlob, f.File)
					description = f.Description
					if f.FileType != "" {
						fileType = f.FileType
					}
					if f.UploadAs != "" {
						c.logger.Debugf("upload_as provided for exact glob: '%s' - uploading to remote filename: '%s' instead\n", exactGlob, f.UploadAs)
						uploadAs = f.UploadAs
					}
				} else {
					c.logger.Debugf("exact glob %s does not match metadata file: %s\n", exactGlob, f.File)
				}
			}

			c.logger.Debugf(
				"Creating product file: {product_slug: %s, filename: %s, file_type: %s, aws_object_key: %s, file_version: %s, description: %s}\n",
				productSlug,
				uploadAs,
				fileType,
				remotePath,
				release.Version,
				description,
			)

			productFile, err := pivnetClient.CreateProductFile(pivnet.CreateProductFileConfig{
				ProductSlug:  productSlug,
				Name:         uploadAs,
				FileType:     fileType,
				AWSObjectKey: remotePath,
				FileVersion:  release.Version,
				MD5:          fileContentsMD5,
				Description:  description,
			})
			if err != nil {
				return concourse.OutResponse{}, err
			}

			c.logger.Debugf(
				"Adding product file: {product_slug: %s, product_id: %d, filename: %s, product_file_id: %d, release_id: %d}\n",
				productSlug,
				product.ID,
				filename,
				productFile.ID,
				release.ID,
			)

			err = pivnetClient.AddProductFile(product.ID, release.ID, productFile.ID)
			if err != nil {
				log.Fatalln(err)
			}
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	availability := fetchFromMetadataOrFile("Availability", m, skipFileCheck, c.sourcesDir, input.Params.AvailabilityFile)
	if availability != "Admins Only" {
		releaseUpdate := pivnet.Release{
			ID:           release.ID,
			Availability: availability,
		}
		release, err = pivnetClient.UpdateRelease(productSlug, releaseUpdate)
		if err != nil {
			log.Fatalln(err)
		}

		if availability == "Selected User Groups Only" {
			userGroupIDs := strings.Split(
				fetchFromMetadataOrFile("UserGroupIDs", m, skipFileCheck, c.sourcesDir, input.Params.UserGroupIDsFile),
				",",
			)

			for _, userGroupIDString := range userGroupIDs {
				userGroupID, err := strconv.Atoi(userGroupIDString)
				if err != nil {
					log.Fatalln(err)
				}

				pivnetClient.AddUserGroup(productSlug, release.ID, userGroupID)
			}
		}
	}

	releaseETag, err := pivnetClient.ReleaseETag(productSlug, release)
	//TODO this should not panic
	if err != nil {
		panic(err)
	}

	outputVersion, err := versions.CombineVersionAndETag(release.Version, releaseETag)
	//TODO this should not panic
	if err != nil {
		panic(err)
	}

	metadata := []concourse.Metadata{
		{Name: "version", Value: release.Version},
		{Name: "release_type", Value: release.ReleaseType},
		{Name: "release_date", Value: release.ReleaseDate},
		{Name: "description", Value: release.Description},
		{Name: "release_notes_url", Value: release.ReleaseNotesURL},
		{Name: "eula_slug", Value: release.EULA.Slug},
		{Name: "availability", Value: release.Availability},
		{Name: "controlled", Value: fmt.Sprintf("%t", release.Controlled)},
		{Name: "eccn", Value: release.ECCN},
		{Name: "license_exception", Value: release.LicenseException},
		{Name: "end_of_support_date", Value: release.EndOfSupportDate},
		{Name: "end_of_guidance_date", Value: release.EndOfGuidanceDate},
		{Name: "end_of_availability_date", Value: release.EndOfAvailabilityDate},
	}
	if release.EULA != nil {
		metadata = append(metadata, concourse.Metadata{Name: "eula_slug", Value: release.EULA.Slug})
	}

	out := concourse.OutResponse{
		Version: concourse.Version{
			ProductVersion: outputVersion,
		},
		Metadata: metadata,
	}

	return out, nil
}

func fetchFromMetadataOrFile(yamlKey string, m metadata.Metadata, skipFileCheck bool, dir, file string) string {
	if skipFileCheck && m.Release != nil {
		metadataValue := reflect.ValueOf(m.Release).Elem()
		fieldValue := metadataValue.FieldByName(yamlKey)

		if yamlKey == "UserGroupIDs" {
			var ids []string
			for i := 0; i < fieldValue.Len(); i++ {
				ids = append(ids, fieldValue.Index(i).String())
			}

			return strings.Join(ids, ",")
		}

		return fieldValue.String()
	}

	return readStringContents(dir, file)
}

func readStringContents(sourcesDir, file string) string {
	if file == "" {
		return ""
	}
	fullPath := filepath.Join(sourcesDir, file)
	contents, err := ioutil.ReadFile(fullPath)
	if err != nil {
		log.Fatal(err)
	}
	return string(contents)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func warnIfDeprecatedFilesFound(
	params concourse.OutParams,
	logger logger.Logger,
	screenWriter *log.Logger,
) {
	files := map[string]string{
		"version_file":        params.VersionFile,
		"eula_slug_file":      params.EULASlugFile,
		"release_date_file":   params.ReleaseDateFile,
		"description_file":    params.DescriptionFile,
		"release_type_file":   params.ReleaseTypeFile,
		"user_group_ids_file": params.UserGroupIDsFile,
		"availability_file":   params.AvailabilityFile,
		"release_notes_file":  params.ReleaseNotesURLFile,
	}
	for key, value := range files {
		if value == "" {
			continue
		}

		logger.Debugf("\x1b[31mDEPRECATION WARNING: %q is deprecated and will be removed in a future release\x1b[0m\n", key)

		if screenWriter != nil {
			screenWriter.Printf("\x1b[31mDEPRECATION WARNING: %q is deprecated and will be removed in a future release\x1b[0m", key)
		}
	}
}

func containsString(strings []string, str string) bool {
	for _, s := range strings {
		if str == s {
			return true
		}
	}
	return false
}
