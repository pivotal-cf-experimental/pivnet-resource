package release_test

import (
	"errors"
	"log"
	"time"

	"github.com/pivotal-cf/go-pivnet"
	"github.com/pivotal-cf/go-pivnet/logger"
	"github.com/pivotal-cf/go-pivnet/logshim"
	"github.com/pivotal-cf/pivnet-resource/metadata"
	"github.com/pivotal-cf/pivnet-resource/out/release"
	"github.com/pivotal-cf/pivnet-resource/out/release/releasefakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReleaseUploader", func() {
	var (
		fakeLogger logger.Logger

		s3Client      *releasefakes.S3Client
		uploadClient  *releasefakes.UploadClient
		sha256Summer  *releasefakes.Sha256Summer
		md5Summer     *releasefakes.Md5Summer
		pivnetRelease pivnet.Release
		uploader      release.ReleaseUploader
		asyncTimeout  time.Duration
		pollFrequency time.Duration

		productSlug string

		mdata metadata.Metadata

		existingProductFiles []pivnet.ProductFile
		actualSHA256Sum      string
		actualMD5Sum         string
		newAWSObjectKey      string

		existingProductFilesErr error
		createProductFileErr    error
		createFileGroupErr    	error
		uploadFileErr           error
		sha256SumFileErr        error
		md5SumFileErr           error
		productFileErr          error
	)

	BeforeEach(func() {
		logger := log.New(GinkgoWriter, "", log.LstdFlags)
		fakeLogger = logshim.NewLogShim(logger, logger, true)

		s3Client = &releasefakes.S3Client{}
		uploadClient = &releasefakes.UploadClient{}
		sha256Summer = &releasefakes.Sha256Summer{}
		md5Summer = &releasefakes.Md5Summer{}

		productSlug = "some-product-slug"

		asyncTimeout = 450 * time.Millisecond
		pollFrequency = 15 * time.Millisecond

		pivnetRelease = pivnet.Release{
			ID:      1111,
			Version: "some-release-version",
		}

		mdata = metadata.Metadata{
			ProductFiles: []metadata.ProductFile{
				{
					File:               "some/file",
					Description:        "a description",
					UploadAs:           "a file",
					FileType:           "something",
					DocsURL:            "some-docs-url",
					SystemRequirements: []string{"req1", "req2"},
					Platforms:          []string{"Linux"},
					IncludedFiles:      []string{"include1", "include2"},
				},
			},

			FileGroups: []metadata.FileGroup{
				{
					Name: "file-group-1",
				},
			},
		}

		existingProductFiles = []pivnet.ProductFile{
			{
				ID:           1234,
				AWSObjectKey: "some-existing-aws-object-key",
			},
		}

		actualSHA256Sum = "madeupsha256"
		actualMD5Sum = "madeupmd5"
		newAWSObjectKey = "s3-remote-path"

		existingProductFilesErr = nil
		createProductFileErr = nil
		uploadFileErr = nil
		sha256SumFileErr = nil
		md5SumFileErr = nil
		productFileErr = nil
	})

	JustBeforeEach(func() {
		uploader = release.NewReleaseUploader(
			s3Client,
			uploadClient,
			fakeLogger,
			sha256Summer,
			md5Summer,
			mdata,
			"/some/sources/dir",
			productSlug,
			asyncTimeout,
			pollFrequency,
		)

		sha256Summer.SumFileReturns(actualSHA256Sum, sha256SumFileErr)
		md5Summer.SumFileReturns(actualMD5Sum, md5SumFileErr)
		s3Client.UploadFileReturns(newAWSObjectKey, uploadFileErr)
		uploadClient.CreateProductFileReturns(pivnet.ProductFile{ID: 13367}, createProductFileErr)
		uploadClient.CreateFileGroupReturns(pivnet.FileGroup{ID: 8675309}, createFileGroupErr)

		uploadClient.ProductFilesReturns(existingProductFiles, existingProductFilesErr)

		invokeCount := 0
		uploadClient.ProductFileStub = func(string, int) (pivnet.ProductFile, error) {
			if productFileErr != nil {
				return pivnet.ProductFile{}, productFileErr
			}

			productFile := existingProductFiles[0]

			invokeCount += 1

			if invokeCount == 1 {
				return productFile, nil
			}

			productFile.FileTransferStatus = "complete"
			return productFile, nil
		}
	})

	Describe("Upload", func() {
		It("uploads a release to s3 and adds metadata to pivnet", func() {
			err := uploader.Upload(pivnetRelease, []string{"some/file"})
			Expect(err).NotTo(HaveOccurred())

			Expect(sha256Summer.SumFileArgsForCall(0)).To(Equal("/some/sources/dir/some/file"))
			Expect(md5Summer.SumFileArgsForCall(0)).To(Equal("/some/sources/dir/some/file"))
			Expect(s3Client.UploadFileArgsForCall(0)).To(Equal("some/file"))

			Expect(uploadClient.CreateProductFileArgsForCall(0)).To(Equal(pivnet.CreateProductFileConfig{
				ProductSlug:        productSlug,
				AWSObjectKey:       newAWSObjectKey,
				SHA256:             actualSHA256Sum,
				MD5:                actualMD5Sum,
				FileVersion:        pivnetRelease.Version,
				Name:               mdata.ProductFiles[0].UploadAs,
				Description:        mdata.ProductFiles[0].Description,
				FileType:           mdata.ProductFiles[0].FileType,
				DocsURL:            mdata.ProductFiles[0].DocsURL,
				SystemRequirements: mdata.ProductFiles[0].SystemRequirements,
				Platforms:          mdata.ProductFiles[0].Platforms,
				IncludedFiles:      mdata.ProductFiles[0].IncludedFiles,
			}))

			Expect(uploadClient.AddProductFileCallCount()).To(Equal(1))
			invokedProductSlug, releaseID, productFileID := uploadClient.AddProductFileArgsForCall(0)
			Expect(invokedProductSlug).To(Equal(productSlug))
			Expect(releaseID).To(Equal(1111))
			Expect(productFileID).To(Equal(13367))

			Expect(uploadClient.CreateFileGroupArgsForCall(0)).To(Equal(pivnet.CreateFileGroupConfig{
				ProductSlug:        productSlug,
				Name:               mdata.FileGroups[0].Name,
			}))

			Expect(uploadClient.AddFileGroupCallCount()).To(Equal(1))
			invokedProductSlug, releaseID, fileGroupID := uploadClient.AddFileGroupArgsForCall(0)
			Expect(invokedProductSlug).To(Equal(productSlug))
			Expect(releaseID).To(Equal(1111))
			Expect(fileGroupID).To(Equal(8675309))
		})

		Context("when a product file already exists with AWSObjectKey", func() {
			BeforeEach(func() {
				newAWSObjectKey = existingProductFiles[0].AWSObjectKey
			})

			It("Deletes the product file before recreating", func() {
				err := uploader.Upload(pivnetRelease, []string{""})
				Expect(err).NotTo(HaveOccurred())

				Expect(uploadClient.DeleteProductFileCallCount()).To(Equal(1))

				invokedProductSlug, invokedProductFileID := uploadClient.DeleteProductFileArgsForCall(0)
				Expect(invokedProductSlug).To(Equal(productSlug))
				Expect(invokedProductFileID).To(Equal(existingProductFiles[0].ID))

				Expect(uploadClient.CreateProductFileCallCount()).To(Equal(1))
				Expect(uploadClient.AddProductFileCallCount()).To(Equal(1))
			})

			Context("when there is an error deleting the product file", func() {
				var (
					deleteProductFileErr error
				)

				BeforeEach(func() {
					deleteProductFileErr = errors.New("delete product file error")
					uploadClient.DeleteProductFileReturns(pivnet.ProductFile{}, deleteProductFileErr)
				})

				It("returns the error", func() {
					err := uploader.Upload(pivnetRelease, []string{""})
					Expect(err).To(Equal(deleteProductFileErr))
				})
			})
		})

		Context("when the file sha256 cannot be computed", func() {
			BeforeEach(func() {
				sha256SumFileErr = errors.New("sha256 error")
			})

			It("returns an error", func() {
				err := uploader.Upload(pivnetRelease, []string{""})
				Expect(err).To(MatchError(errors.New("sha256 error")))
			})
		})

		Context("when the file md5 cannot be computed", func() {
			BeforeEach(func() {
				md5SumFileErr = errors.New("md5 error")
			})

			It("returns an error", func() {
				err := uploader.Upload(pivnetRelease, []string{""})
				Expect(err).To(MatchError(errors.New("md5 error")))
			})
		})

		Context("when the s3 upload fails", func() {
			BeforeEach(func() {
				uploadFileErr = errors.New("s3 failed")
			})

			It("returns an error", func() {
				err := uploader.Upload(pivnetRelease, []string{""})
				Expect(err).To(Equal(uploadFileErr))
			})
		})

		Context("when pivnet fails to find a product", func() {
			BeforeEach(func() {
				createProductFileErr = errors.New("some product files error")
			})

			It("returns an error", func() {
				err := uploader.Upload(pivnetRelease, []string{""})
				Expect(err).To(Equal(createProductFileErr))
			})
		})

		Context("when pivnet fails to get existing product files", func() {
			BeforeEach(func() {
				existingProductFilesErr = errors.New("some product files error")
			})

			It("returns an error", func() {
				err := uploader.Upload(pivnetRelease, []string{""})
				Expect(err).To(Equal(existingProductFilesErr))
			})
		})

		Context("when pivnet cannot add a product file", func() {
			BeforeEach(func() {
				uploadClient.AddProductFileReturns(errors.New("error adding product"))
			})

			It("returns an error", func() {
				err := uploader.Upload(pivnetRelease, []string{""})
				Expect(err).To(MatchError(errors.New("error adding product")))
			})
		})

		Context("when polling for the product file returns an error", func() {
			BeforeEach(func() {
				productFileErr = errors.New("product file error")
			})

			It("returns an error", func() {
				err := uploader.Upload(pivnetRelease, []string{""})
				Expect(err).To(Equal(productFileErr))
			})
		})

		Context("when polling for the product file times out", func() {
			BeforeEach(func() {
				asyncTimeout = pollFrequency / 2
			})

			It("returns an error", func() {
				err := uploader.Upload(pivnetRelease, []string{""})
				Expect(err).To(HaveOccurred())

				Expect(err.Error()).To(ContainSubstring("timed out"))
			})
		})
	})
})
