// This file was generated by counterfeiter
package gpfakes

import (
	"io"
	"net/http"
	"sync"

	"github.com/pivotal-cf/go-pivnet"
	"github.com/pivotal-cf/pivnet-resource/gp"
)

type FakeClient struct {
	ReleaseTypesStub        func() ([]string, error)
	releaseTypesMutex       sync.RWMutex
	releaseTypesArgsForCall []struct{}
	releaseTypesReturns     struct {
		result1 []string
		result2 error
	}
	GetReleaseStub        func(productSlug string, productVersion string) (pivnet.Release, error)
	getReleaseMutex       sync.RWMutex
	getReleaseArgsForCall []struct {
		productSlug    string
		productVersion string
	}
	getReleaseReturns struct {
		result1 pivnet.Release
		result2 error
	}
	ReleasesForProductSlugStub        func(string) ([]pivnet.Release, error)
	releasesForProductSlugMutex       sync.RWMutex
	releasesForProductSlugArgsForCall []struct {
		arg1 string
	}
	releasesForProductSlugReturns struct {
		result1 []pivnet.Release
		result2 error
	}
	UpdateReleaseStub        func(productslug string, release pivnet.Release) (pivnet.Release, error)
	updateReleaseMutex       sync.RWMutex
	updateReleaseArgsForCall []struct {
		productslug string
		release     pivnet.Release
	}
	updateReleaseReturns struct {
		result1 pivnet.Release
		result2 error
	}
	CreateReleaseStub        func(pivnet.CreateReleaseConfig) (pivnet.Release, error)
	createReleaseMutex       sync.RWMutex
	createReleaseArgsForCall []struct {
		arg1 pivnet.CreateReleaseConfig
	}
	createReleaseReturns struct {
		result1 pivnet.Release
		result2 error
	}
	AcceptEULAStub        func(productSlug string, releaseID int) error
	acceptEULAMutex       sync.RWMutex
	acceptEULAArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	acceptEULAReturns struct {
		result1 error
	}
	EULAsStub        func() ([]pivnet.EULA, error)
	eULAsMutex       sync.RWMutex
	eULAsArgsForCall []struct{}
	eULAsReturns     struct {
		result1 []pivnet.EULA
		result2 error
	}
	FindProductForSlugStub        func(slug string) (pivnet.Product, error)
	findProductForSlugMutex       sync.RWMutex
	findProductForSlugArgsForCall []struct {
		slug string
	}
	findProductForSlugReturns struct {
		result1 pivnet.Product
		result2 error
	}
	CreateProductFileStub        func(pivnet.CreateProductFileConfig) (pivnet.ProductFile, error)
	createProductFileMutex       sync.RWMutex
	createProductFileArgsForCall []struct {
		arg1 pivnet.CreateProductFileConfig
	}
	createProductFileReturns struct {
		result1 pivnet.ProductFile
		result2 error
	}
	AddProductFileStub        func(productSlug string, releaseID int, productFileID int) error
	addProductFileMutex       sync.RWMutex
	addProductFileArgsForCall []struct {
		productSlug   string
		releaseID     int
		productFileID int
	}
	addProductFileReturns struct {
		result1 error
	}
	GetProductFilesStub        func(productSlug string, releaseID int) ([]pivnet.ProductFile, error)
	getProductFilesMutex       sync.RWMutex
	getProductFilesArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	getProductFilesReturns struct {
		result1 []pivnet.ProductFile
		result2 error
	}
	GetProductFileStub        func(productSlug string, releaseID int, productFileID int) (pivnet.ProductFile, error)
	getProductFileMutex       sync.RWMutex
	getProductFileArgsForCall []struct {
		productSlug   string
		releaseID     int
		productFileID int
	}
	getProductFileReturns struct {
		result1 pivnet.ProductFile
		result2 error
	}
	AddUserGroupStub        func(productSlug string, releaseID int, userGroupID int) error
	addUserGroupMutex       sync.RWMutex
	addUserGroupArgsForCall []struct {
		productSlug string
		releaseID   int
		userGroupID int
	}
	addUserGroupReturns struct {
		result1 error
	}
	ReleaseDependenciesStub        func(productSlug string, releaseID int) ([]pivnet.ReleaseDependency, error)
	releaseDependenciesMutex       sync.RWMutex
	releaseDependenciesArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	releaseDependenciesReturns struct {
		result1 []pivnet.ReleaseDependency
		result2 error
	}
	MakeRequestStub        func(method string, url string, expectedResponseCode int, body io.Reader, data interface{}) (*http.Response, error)
	makeRequestMutex       sync.RWMutex
	makeRequestArgsForCall []struct {
		method               string
		url                  string
		expectedResponseCode int
		body                 io.Reader
		data                 interface{}
	}
	makeRequestReturns struct {
		result1 *http.Response
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) ReleaseTypes() ([]string, error) {
	fake.releaseTypesMutex.Lock()
	fake.releaseTypesArgsForCall = append(fake.releaseTypesArgsForCall, struct{}{})
	fake.recordInvocation("ReleaseTypes", []interface{}{})
	fake.releaseTypesMutex.Unlock()
	if fake.ReleaseTypesStub != nil {
		return fake.ReleaseTypesStub()
	} else {
		return fake.releaseTypesReturns.result1, fake.releaseTypesReturns.result2
	}
}

func (fake *FakeClient) ReleaseTypesCallCount() int {
	fake.releaseTypesMutex.RLock()
	defer fake.releaseTypesMutex.RUnlock()
	return len(fake.releaseTypesArgsForCall)
}

func (fake *FakeClient) ReleaseTypesReturns(result1 []string, result2 error) {
	fake.ReleaseTypesStub = nil
	fake.releaseTypesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetRelease(productSlug string, productVersion string) (pivnet.Release, error) {
	fake.getReleaseMutex.Lock()
	fake.getReleaseArgsForCall = append(fake.getReleaseArgsForCall, struct {
		productSlug    string
		productVersion string
	}{productSlug, productVersion})
	fake.recordInvocation("GetRelease", []interface{}{productSlug, productVersion})
	fake.getReleaseMutex.Unlock()
	if fake.GetReleaseStub != nil {
		return fake.GetReleaseStub(productSlug, productVersion)
	} else {
		return fake.getReleaseReturns.result1, fake.getReleaseReturns.result2
	}
}

func (fake *FakeClient) GetReleaseCallCount() int {
	fake.getReleaseMutex.RLock()
	defer fake.getReleaseMutex.RUnlock()
	return len(fake.getReleaseArgsForCall)
}

func (fake *FakeClient) GetReleaseArgsForCall(i int) (string, string) {
	fake.getReleaseMutex.RLock()
	defer fake.getReleaseMutex.RUnlock()
	return fake.getReleaseArgsForCall[i].productSlug, fake.getReleaseArgsForCall[i].productVersion
}

func (fake *FakeClient) GetReleaseReturns(result1 pivnet.Release, result2 error) {
	fake.GetReleaseStub = nil
	fake.getReleaseReturns = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ReleasesForProductSlug(arg1 string) ([]pivnet.Release, error) {
	fake.releasesForProductSlugMutex.Lock()
	fake.releasesForProductSlugArgsForCall = append(fake.releasesForProductSlugArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ReleasesForProductSlug", []interface{}{arg1})
	fake.releasesForProductSlugMutex.Unlock()
	if fake.ReleasesForProductSlugStub != nil {
		return fake.ReleasesForProductSlugStub(arg1)
	} else {
		return fake.releasesForProductSlugReturns.result1, fake.releasesForProductSlugReturns.result2
	}
}

func (fake *FakeClient) ReleasesForProductSlugCallCount() int {
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	return len(fake.releasesForProductSlugArgsForCall)
}

func (fake *FakeClient) ReleasesForProductSlugArgsForCall(i int) string {
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	return fake.releasesForProductSlugArgsForCall[i].arg1
}

func (fake *FakeClient) ReleasesForProductSlugReturns(result1 []pivnet.Release, result2 error) {
	fake.ReleasesForProductSlugStub = nil
	fake.releasesForProductSlugReturns = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) UpdateRelease(productslug string, release pivnet.Release) (pivnet.Release, error) {
	fake.updateReleaseMutex.Lock()
	fake.updateReleaseArgsForCall = append(fake.updateReleaseArgsForCall, struct {
		productslug string
		release     pivnet.Release
	}{productslug, release})
	fake.recordInvocation("UpdateRelease", []interface{}{productslug, release})
	fake.updateReleaseMutex.Unlock()
	if fake.UpdateReleaseStub != nil {
		return fake.UpdateReleaseStub(productslug, release)
	} else {
		return fake.updateReleaseReturns.result1, fake.updateReleaseReturns.result2
	}
}

func (fake *FakeClient) UpdateReleaseCallCount() int {
	fake.updateReleaseMutex.RLock()
	defer fake.updateReleaseMutex.RUnlock()
	return len(fake.updateReleaseArgsForCall)
}

func (fake *FakeClient) UpdateReleaseArgsForCall(i int) (string, pivnet.Release) {
	fake.updateReleaseMutex.RLock()
	defer fake.updateReleaseMutex.RUnlock()
	return fake.updateReleaseArgsForCall[i].productslug, fake.updateReleaseArgsForCall[i].release
}

func (fake *FakeClient) UpdateReleaseReturns(result1 pivnet.Release, result2 error) {
	fake.UpdateReleaseStub = nil
	fake.updateReleaseReturns = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CreateRelease(arg1 pivnet.CreateReleaseConfig) (pivnet.Release, error) {
	fake.createReleaseMutex.Lock()
	fake.createReleaseArgsForCall = append(fake.createReleaseArgsForCall, struct {
		arg1 pivnet.CreateReleaseConfig
	}{arg1})
	fake.recordInvocation("CreateRelease", []interface{}{arg1})
	fake.createReleaseMutex.Unlock()
	if fake.CreateReleaseStub != nil {
		return fake.CreateReleaseStub(arg1)
	} else {
		return fake.createReleaseReturns.result1, fake.createReleaseReturns.result2
	}
}

func (fake *FakeClient) CreateReleaseCallCount() int {
	fake.createReleaseMutex.RLock()
	defer fake.createReleaseMutex.RUnlock()
	return len(fake.createReleaseArgsForCall)
}

func (fake *FakeClient) CreateReleaseArgsForCall(i int) pivnet.CreateReleaseConfig {
	fake.createReleaseMutex.RLock()
	defer fake.createReleaseMutex.RUnlock()
	return fake.createReleaseArgsForCall[i].arg1
}

func (fake *FakeClient) CreateReleaseReturns(result1 pivnet.Release, result2 error) {
	fake.CreateReleaseStub = nil
	fake.createReleaseReturns = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) AcceptEULA(productSlug string, releaseID int) error {
	fake.acceptEULAMutex.Lock()
	fake.acceptEULAArgsForCall = append(fake.acceptEULAArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("AcceptEULA", []interface{}{productSlug, releaseID})
	fake.acceptEULAMutex.Unlock()
	if fake.AcceptEULAStub != nil {
		return fake.AcceptEULAStub(productSlug, releaseID)
	} else {
		return fake.acceptEULAReturns.result1
	}
}

func (fake *FakeClient) AcceptEULACallCount() int {
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	return len(fake.acceptEULAArgsForCall)
}

func (fake *FakeClient) AcceptEULAArgsForCall(i int) (string, int) {
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	return fake.acceptEULAArgsForCall[i].productSlug, fake.acceptEULAArgsForCall[i].releaseID
}

func (fake *FakeClient) AcceptEULAReturns(result1 error) {
	fake.AcceptEULAStub = nil
	fake.acceptEULAReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) EULAs() ([]pivnet.EULA, error) {
	fake.eULAsMutex.Lock()
	fake.eULAsArgsForCall = append(fake.eULAsArgsForCall, struct{}{})
	fake.recordInvocation("EULAs", []interface{}{})
	fake.eULAsMutex.Unlock()
	if fake.EULAsStub != nil {
		return fake.EULAsStub()
	} else {
		return fake.eULAsReturns.result1, fake.eULAsReturns.result2
	}
}

func (fake *FakeClient) EULAsCallCount() int {
	fake.eULAsMutex.RLock()
	defer fake.eULAsMutex.RUnlock()
	return len(fake.eULAsArgsForCall)
}

func (fake *FakeClient) EULAsReturns(result1 []pivnet.EULA, result2 error) {
	fake.EULAsStub = nil
	fake.eULAsReturns = struct {
		result1 []pivnet.EULA
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FindProductForSlug(slug string) (pivnet.Product, error) {
	fake.findProductForSlugMutex.Lock()
	fake.findProductForSlugArgsForCall = append(fake.findProductForSlugArgsForCall, struct {
		slug string
	}{slug})
	fake.recordInvocation("FindProductForSlug", []interface{}{slug})
	fake.findProductForSlugMutex.Unlock()
	if fake.FindProductForSlugStub != nil {
		return fake.FindProductForSlugStub(slug)
	} else {
		return fake.findProductForSlugReturns.result1, fake.findProductForSlugReturns.result2
	}
}

func (fake *FakeClient) FindProductForSlugCallCount() int {
	fake.findProductForSlugMutex.RLock()
	defer fake.findProductForSlugMutex.RUnlock()
	return len(fake.findProductForSlugArgsForCall)
}

func (fake *FakeClient) FindProductForSlugArgsForCall(i int) string {
	fake.findProductForSlugMutex.RLock()
	defer fake.findProductForSlugMutex.RUnlock()
	return fake.findProductForSlugArgsForCall[i].slug
}

func (fake *FakeClient) FindProductForSlugReturns(result1 pivnet.Product, result2 error) {
	fake.FindProductForSlugStub = nil
	fake.findProductForSlugReturns = struct {
		result1 pivnet.Product
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CreateProductFile(arg1 pivnet.CreateProductFileConfig) (pivnet.ProductFile, error) {
	fake.createProductFileMutex.Lock()
	fake.createProductFileArgsForCall = append(fake.createProductFileArgsForCall, struct {
		arg1 pivnet.CreateProductFileConfig
	}{arg1})
	fake.recordInvocation("CreateProductFile", []interface{}{arg1})
	fake.createProductFileMutex.Unlock()
	if fake.CreateProductFileStub != nil {
		return fake.CreateProductFileStub(arg1)
	} else {
		return fake.createProductFileReturns.result1, fake.createProductFileReturns.result2
	}
}

func (fake *FakeClient) CreateProductFileCallCount() int {
	fake.createProductFileMutex.RLock()
	defer fake.createProductFileMutex.RUnlock()
	return len(fake.createProductFileArgsForCall)
}

func (fake *FakeClient) CreateProductFileArgsForCall(i int) pivnet.CreateProductFileConfig {
	fake.createProductFileMutex.RLock()
	defer fake.createProductFileMutex.RUnlock()
	return fake.createProductFileArgsForCall[i].arg1
}

func (fake *FakeClient) CreateProductFileReturns(result1 pivnet.ProductFile, result2 error) {
	fake.CreateProductFileStub = nil
	fake.createProductFileReturns = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) AddProductFile(productSlug string, releaseID int, productFileID int) error {
	fake.addProductFileMutex.Lock()
	fake.addProductFileArgsForCall = append(fake.addProductFileArgsForCall, struct {
		productSlug   string
		releaseID     int
		productFileID int
	}{productSlug, releaseID, productFileID})
	fake.recordInvocation("AddProductFile", []interface{}{productSlug, releaseID, productFileID})
	fake.addProductFileMutex.Unlock()
	if fake.AddProductFileStub != nil {
		return fake.AddProductFileStub(productSlug, releaseID, productFileID)
	} else {
		return fake.addProductFileReturns.result1
	}
}

func (fake *FakeClient) AddProductFileCallCount() int {
	fake.addProductFileMutex.RLock()
	defer fake.addProductFileMutex.RUnlock()
	return len(fake.addProductFileArgsForCall)
}

func (fake *FakeClient) AddProductFileArgsForCall(i int) (string, int, int) {
	fake.addProductFileMutex.RLock()
	defer fake.addProductFileMutex.RUnlock()
	return fake.addProductFileArgsForCall[i].productSlug, fake.addProductFileArgsForCall[i].releaseID, fake.addProductFileArgsForCall[i].productFileID
}

func (fake *FakeClient) AddProductFileReturns(result1 error) {
	fake.AddProductFileStub = nil
	fake.addProductFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) GetProductFiles(productSlug string, releaseID int) ([]pivnet.ProductFile, error) {
	fake.getProductFilesMutex.Lock()
	fake.getProductFilesArgsForCall = append(fake.getProductFilesArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("GetProductFiles", []interface{}{productSlug, releaseID})
	fake.getProductFilesMutex.Unlock()
	if fake.GetProductFilesStub != nil {
		return fake.GetProductFilesStub(productSlug, releaseID)
	} else {
		return fake.getProductFilesReturns.result1, fake.getProductFilesReturns.result2
	}
}

func (fake *FakeClient) GetProductFilesCallCount() int {
	fake.getProductFilesMutex.RLock()
	defer fake.getProductFilesMutex.RUnlock()
	return len(fake.getProductFilesArgsForCall)
}

func (fake *FakeClient) GetProductFilesArgsForCall(i int) (string, int) {
	fake.getProductFilesMutex.RLock()
	defer fake.getProductFilesMutex.RUnlock()
	return fake.getProductFilesArgsForCall[i].productSlug, fake.getProductFilesArgsForCall[i].releaseID
}

func (fake *FakeClient) GetProductFilesReturns(result1 []pivnet.ProductFile, result2 error) {
	fake.GetProductFilesStub = nil
	fake.getProductFilesReturns = struct {
		result1 []pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetProductFile(productSlug string, releaseID int, productFileID int) (pivnet.ProductFile, error) {
	fake.getProductFileMutex.Lock()
	fake.getProductFileArgsForCall = append(fake.getProductFileArgsForCall, struct {
		productSlug   string
		releaseID     int
		productFileID int
	}{productSlug, releaseID, productFileID})
	fake.recordInvocation("GetProductFile", []interface{}{productSlug, releaseID, productFileID})
	fake.getProductFileMutex.Unlock()
	if fake.GetProductFileStub != nil {
		return fake.GetProductFileStub(productSlug, releaseID, productFileID)
	} else {
		return fake.getProductFileReturns.result1, fake.getProductFileReturns.result2
	}
}

func (fake *FakeClient) GetProductFileCallCount() int {
	fake.getProductFileMutex.RLock()
	defer fake.getProductFileMutex.RUnlock()
	return len(fake.getProductFileArgsForCall)
}

func (fake *FakeClient) GetProductFileArgsForCall(i int) (string, int, int) {
	fake.getProductFileMutex.RLock()
	defer fake.getProductFileMutex.RUnlock()
	return fake.getProductFileArgsForCall[i].productSlug, fake.getProductFileArgsForCall[i].releaseID, fake.getProductFileArgsForCall[i].productFileID
}

func (fake *FakeClient) GetProductFileReturns(result1 pivnet.ProductFile, result2 error) {
	fake.GetProductFileStub = nil
	fake.getProductFileReturns = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) AddUserGroup(productSlug string, releaseID int, userGroupID int) error {
	fake.addUserGroupMutex.Lock()
	fake.addUserGroupArgsForCall = append(fake.addUserGroupArgsForCall, struct {
		productSlug string
		releaseID   int
		userGroupID int
	}{productSlug, releaseID, userGroupID})
	fake.recordInvocation("AddUserGroup", []interface{}{productSlug, releaseID, userGroupID})
	fake.addUserGroupMutex.Unlock()
	if fake.AddUserGroupStub != nil {
		return fake.AddUserGroupStub(productSlug, releaseID, userGroupID)
	} else {
		return fake.addUserGroupReturns.result1
	}
}

func (fake *FakeClient) AddUserGroupCallCount() int {
	fake.addUserGroupMutex.RLock()
	defer fake.addUserGroupMutex.RUnlock()
	return len(fake.addUserGroupArgsForCall)
}

func (fake *FakeClient) AddUserGroupArgsForCall(i int) (string, int, int) {
	fake.addUserGroupMutex.RLock()
	defer fake.addUserGroupMutex.RUnlock()
	return fake.addUserGroupArgsForCall[i].productSlug, fake.addUserGroupArgsForCall[i].releaseID, fake.addUserGroupArgsForCall[i].userGroupID
}

func (fake *FakeClient) AddUserGroupReturns(result1 error) {
	fake.AddUserGroupStub = nil
	fake.addUserGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) ReleaseDependencies(productSlug string, releaseID int) ([]pivnet.ReleaseDependency, error) {
	fake.releaseDependenciesMutex.Lock()
	fake.releaseDependenciesArgsForCall = append(fake.releaseDependenciesArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("ReleaseDependencies", []interface{}{productSlug, releaseID})
	fake.releaseDependenciesMutex.Unlock()
	if fake.ReleaseDependenciesStub != nil {
		return fake.ReleaseDependenciesStub(productSlug, releaseID)
	} else {
		return fake.releaseDependenciesReturns.result1, fake.releaseDependenciesReturns.result2
	}
}

func (fake *FakeClient) ReleaseDependenciesCallCount() int {
	fake.releaseDependenciesMutex.RLock()
	defer fake.releaseDependenciesMutex.RUnlock()
	return len(fake.releaseDependenciesArgsForCall)
}

func (fake *FakeClient) ReleaseDependenciesArgsForCall(i int) (string, int) {
	fake.releaseDependenciesMutex.RLock()
	defer fake.releaseDependenciesMutex.RUnlock()
	return fake.releaseDependenciesArgsForCall[i].productSlug, fake.releaseDependenciesArgsForCall[i].releaseID
}

func (fake *FakeClient) ReleaseDependenciesReturns(result1 []pivnet.ReleaseDependency, result2 error) {
	fake.ReleaseDependenciesStub = nil
	fake.releaseDependenciesReturns = struct {
		result1 []pivnet.ReleaseDependency
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) MakeRequest(method string, url string, expectedResponseCode int, body io.Reader, data interface{}) (*http.Response, error) {
	fake.makeRequestMutex.Lock()
	fake.makeRequestArgsForCall = append(fake.makeRequestArgsForCall, struct {
		method               string
		url                  string
		expectedResponseCode int
		body                 io.Reader
		data                 interface{}
	}{method, url, expectedResponseCode, body, data})
	fake.recordInvocation("MakeRequest", []interface{}{method, url, expectedResponseCode, body, data})
	fake.makeRequestMutex.Unlock()
	if fake.MakeRequestStub != nil {
		return fake.MakeRequestStub(method, url, expectedResponseCode, body, data)
	} else {
		return fake.makeRequestReturns.result1, fake.makeRequestReturns.result2
	}
}

func (fake *FakeClient) MakeRequestCallCount() int {
	fake.makeRequestMutex.RLock()
	defer fake.makeRequestMutex.RUnlock()
	return len(fake.makeRequestArgsForCall)
}

func (fake *FakeClient) MakeRequestArgsForCall(i int) (string, string, int, io.Reader, interface{}) {
	fake.makeRequestMutex.RLock()
	defer fake.makeRequestMutex.RUnlock()
	return fake.makeRequestArgsForCall[i].method, fake.makeRequestArgsForCall[i].url, fake.makeRequestArgsForCall[i].expectedResponseCode, fake.makeRequestArgsForCall[i].body, fake.makeRequestArgsForCall[i].data
}

func (fake *FakeClient) MakeRequestReturns(result1 *http.Response, result2 error) {
	fake.MakeRequestStub = nil
	fake.makeRequestReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.releaseTypesMutex.RLock()
	defer fake.releaseTypesMutex.RUnlock()
	fake.getReleaseMutex.RLock()
	defer fake.getReleaseMutex.RUnlock()
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	fake.updateReleaseMutex.RLock()
	defer fake.updateReleaseMutex.RUnlock()
	fake.createReleaseMutex.RLock()
	defer fake.createReleaseMutex.RUnlock()
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	fake.eULAsMutex.RLock()
	defer fake.eULAsMutex.RUnlock()
	fake.findProductForSlugMutex.RLock()
	defer fake.findProductForSlugMutex.RUnlock()
	fake.createProductFileMutex.RLock()
	defer fake.createProductFileMutex.RUnlock()
	fake.addProductFileMutex.RLock()
	defer fake.addProductFileMutex.RUnlock()
	fake.getProductFilesMutex.RLock()
	defer fake.getProductFilesMutex.RUnlock()
	fake.getProductFileMutex.RLock()
	defer fake.getProductFileMutex.RUnlock()
	fake.addUserGroupMutex.RLock()
	defer fake.addUserGroupMutex.RUnlock()
	fake.releaseDependenciesMutex.RLock()
	defer fake.releaseDependenciesMutex.RUnlock()
	fake.makeRequestMutex.RLock()
	defer fake.makeRequestMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ gp.Client = new(FakeClient)
