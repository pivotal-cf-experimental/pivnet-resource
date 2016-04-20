package pivnet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateProductFileConfig struct {
	ProductSlug  string
	FileVersion  string
	FileType     string
	AWSObjectKey string
	Name         string
	MD5          string
	Description  string
}

func (c client) GetProductFiles(release Release) (ProductFiles, error) {
	links := release.Links
	if links == nil {
		return ProductFiles{}, fmt.Errorf("No links found")
	}

	productFiles := ProductFiles{}

	link := links.ProductFiles["href"]
	c.logger.Debugf("link: %s\n", link)

	err := c.makeRequest(
		"GET",
		link,
		http.StatusOK,
		nil,
		&productFiles,
	)
	if err != nil {
		return ProductFiles{}, err
	}

	return productFiles, nil
}

func (c client) GetProductFile(productSlug string, releaseID int, productID int) (ProductFile, error) {
	url := fmt.Sprintf("%s/products/%s/releases/%d/product_files/%d",
		c.url,
		productSlug,
		releaseID,
		productID,
	)
	response := ProductFileResponse{}

	err := c.makeRequest(
		"GET",
		url,
		http.StatusOK,
		nil,
		&response,
	)
	if err != nil {
		return ProductFile{}, err
	}

	return response.ProductFile, nil
}

func (c client) CreateProductFile(config CreateProductFileConfig) (ProductFile, error) {
	if config.AWSObjectKey == "" {
		return ProductFile{}, fmt.Errorf("AWS object key must not be empty")
	}

	url := c.url + "/products/" + config.ProductSlug + "/product_files"

	body := createProductFileBody{
		ProductFile: ProductFile{
			MD5:          config.MD5,
			FileType:     config.FileType,
			FileVersion:  config.FileVersion,
			AWSObjectKey: config.AWSObjectKey,
			Name:         config.Name,
			Description:  config.Description,
		},
	}

	b, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	var response ProductFileResponse
	err = c.makeRequest(
		"POST",
		url,
		http.StatusCreated,
		bytes.NewReader(b),
		&response,
	)
	if err != nil {
		return ProductFile{}, err
	}

	return response.ProductFile, nil
}

type createProductFileBody struct {
	ProductFile ProductFile `json:"product_file"`
}

func (c client) DeleteProductFile(productSlug string, id int) (ProductFile, error) {
	url := fmt.Sprintf(
		"%s/products/%s/product_files/%d",
		c.url,
		productSlug,
		id,
	)

	var response ProductFileResponse
	err := c.makeRequest(
		"DELETE",
		url,
		http.StatusOK,
		nil,
		&response,
	)
	if err != nil {
		return ProductFile{}, err
	}

	return response.ProductFile, nil
}

func (c client) AddProductFile(
	productID int,
	releaseID int,
	productFileID int,
) error {
	url := fmt.Sprintf(
		"%s/products/%d/releases/%d/add_product_file",
		c.url,
		productID,
		releaseID,
	)

	body := createProductFileBody{
		ProductFile: ProductFile{
			ID: productFileID,
		},
	}

	b, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	err = c.makeRequest(
		"PATCH",
		url,
		http.StatusNoContent,
		bytes.NewReader(b),
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}
