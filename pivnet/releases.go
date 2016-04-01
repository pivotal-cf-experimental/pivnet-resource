package pivnet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type createReleaseBody struct {
	Release Release `json:"release"`
}

type CreateReleaseConfig struct {
	ProductSlug     string
	ProductVersion  string
	ReleaseType     string
	ReleaseDate     string
	EulaSlug        string
	Description     string
	ReleaseNotesURL string
	ECCN            string
}

func (c client) ReleasesForProductSlug(productSlug string) ([]Release, error) {
	url := fmt.Sprintf(
		"%s/products/%s/releases",
		c.url,
		productSlug,
	)

	var response ReleasesResponse
	err := c.makeRequest("GET", url, http.StatusOK, nil, &response)
	if err != nil {
		return nil, err
	}
	return response.Releases, nil
}

func (c client) GetRelease(productSlug, version string) (Release, error) {
	url := fmt.Sprintf("%s/products/%s/releases", c.url, productSlug)

	var response ReleasesResponse
	err := c.makeRequest("GET", url, http.StatusOK, nil, &response)
	if err != nil {
		return Release{}, err
	}

	var matchingRelease Release
	for i, r := range response.Releases {
		if r.Version == version {
			matchingRelease = r
			break
		}

		if i == len(response.Releases)-1 {
			return Release{}, fmt.Errorf(
				"The requested version: %s - could not be found", version)
		}
	}

	return matchingRelease, nil
}

func (c client) CreateRelease(config CreateReleaseConfig) (Release, error) {
	url := fmt.Sprintf("%s/products/%s/releases", c.url, config.ProductSlug)

	body := createReleaseBody{
		Release: Release{
			Availability: "Admins Only",
			Eula: &Eula{
				Slug: config.EulaSlug,
			},
			OSSCompliant:    "confirm",
			ReleaseDate:     config.ReleaseDate,
			ReleaseType:     config.ReleaseType,
			Version:         config.ProductVersion,
			Description:     config.Description,
			ReleaseNotesURL: config.ReleaseNotesURL,
			ECCN:            config.ECCN,
		},
	}

	if config.ReleaseDate == "" {
		body.Release.ReleaseDate = time.Now().Format("2006-01-02")
		c.logger.Debugf(
			"No release date found - defaulting to %s\n",
			body.Release.ReleaseDate,
		)
	}

	b, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	var response CreateReleaseResponse
	err = c.makeRequest("POST", url, http.StatusCreated, bytes.NewReader(b), &response)
	if err != nil {
		return Release{}, err
	}

	return response.Release, nil
}

func (c client) UpdateRelease(productSlug string, release Release) (Release, error) {
	url := fmt.Sprintf("%s/products/%s/releases/%d", c.url, productSlug, release.ID)

	release.OSSCompliant = "confirm"

	var updatedRelease = createReleaseBody{
		Release: release,
	}

	body, err := json.Marshal(updatedRelease)
	if err != nil {
		panic(err)
	}

	var response CreateReleaseResponse
	err = c.makeRequest("PATCH", url, http.StatusOK, bytes.NewReader(body), &response)
	if err != nil {
		return Release{}, err
	}

	return response.Release, nil
}

func (c client) ReleaseETag(productSlug string, release Release) (string, error) {
	url := fmt.Sprintf("%s/products/%s/releases/%d", c.url, productSlug, release.ID)

	var response Release
	resp, err := c.makeRequestWithHTTPResponse("GET", url, http.StatusOK, nil, &response)
	if err != nil {
		return "", err
	}

	rawEtag := resp.Header.Get("ETag")

	if rawEtag == "" {
		c.logger.Debugf("Missing ETag")
		return "", fmt.Errorf("ETag header not present")
	}

	c.logger.Debugf("Received ETag: %v\n", rawEtag)

	// Weak ETag looks like: W/"my-etag"; strong ETag looks like: "my-etag"
	splitRawEtag := strings.SplitN(rawEtag, `"`, -1)

	if len(splitRawEtag) < 2 {
		c.logger.Debugf("Malformed ETag: %s\n", rawEtag)
		return "", fmt.Errorf("ETag header malformed: %s", rawEtag)
	}

	etag := splitRawEtag[1]
	return etag, nil
}

func (c client) DeleteRelease(release Release, productSlug string) error {
	url := fmt.Sprintf(
		"%s/products/%s/releases/%d",
		c.url,
		productSlug,
		release.ID,
	)

	err := c.makeRequest(
		"DELETE",
		url,
		http.StatusNoContent,
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}
