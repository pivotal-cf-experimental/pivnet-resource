// Code generated by counterfeiter. DO NOT EDIT.
package releasefakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v6"
)

type UploadClient struct {
	AddProductFileStub        func(string, int, int) error
	addProductFileMutex       sync.RWMutex
	addProductFileArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 int
	}
	addProductFileReturns struct {
		result1 error
	}
	addProductFileReturnsOnCall map[int]struct {
		result1 error
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
	createProductFileReturnsOnCall map[int]struct {
		result1 pivnet.ProductFile
		result2 error
	}
	DeleteProductFileStub        func(string, int) (pivnet.ProductFile, error)
	deleteProductFileMutex       sync.RWMutex
	deleteProductFileArgsForCall []struct {
		arg1 string
		arg2 int
	}
	deleteProductFileReturns struct {
		result1 pivnet.ProductFile
		result2 error
	}
	deleteProductFileReturnsOnCall map[int]struct {
		result1 pivnet.ProductFile
		result2 error
	}
	FindProductForSlugStub        func(string) (pivnet.Product, error)
	findProductForSlugMutex       sync.RWMutex
	findProductForSlugArgsForCall []struct {
		arg1 string
	}
	findProductForSlugReturns struct {
		result1 pivnet.Product
		result2 error
	}
	findProductForSlugReturnsOnCall map[int]struct {
		result1 pivnet.Product
		result2 error
	}
	ProductFileStub        func(string, int) (pivnet.ProductFile, error)
	productFileMutex       sync.RWMutex
	productFileArgsForCall []struct {
		arg1 string
		arg2 int
	}
	productFileReturns struct {
		result1 pivnet.ProductFile
		result2 error
	}
	productFileReturnsOnCall map[int]struct {
		result1 pivnet.ProductFile
		result2 error
	}
	ProductFilesStub        func(string) ([]pivnet.ProductFile, error)
	productFilesMutex       sync.RWMutex
	productFilesArgsForCall []struct {
		arg1 string
	}
	productFilesReturns struct {
		result1 []pivnet.ProductFile
		result2 error
	}
	productFilesReturnsOnCall map[int]struct {
		result1 []pivnet.ProductFile
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *UploadClient) AddProductFile(arg1 string, arg2 int, arg3 int) error {
	fake.addProductFileMutex.Lock()
	ret, specificReturn := fake.addProductFileReturnsOnCall[len(fake.addProductFileArgsForCall)]
	fake.addProductFileArgsForCall = append(fake.addProductFileArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("AddProductFile", []interface{}{arg1, arg2, arg3})
	fake.addProductFileMutex.Unlock()
	if fake.AddProductFileStub != nil {
		return fake.AddProductFileStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addProductFileReturns
	return fakeReturns.result1
}

func (fake *UploadClient) AddProductFileCallCount() int {
	fake.addProductFileMutex.RLock()
	defer fake.addProductFileMutex.RUnlock()
	return len(fake.addProductFileArgsForCall)
}

func (fake *UploadClient) AddProductFileCalls(stub func(string, int, int) error) {
	fake.addProductFileMutex.Lock()
	defer fake.addProductFileMutex.Unlock()
	fake.AddProductFileStub = stub
}

func (fake *UploadClient) AddProductFileArgsForCall(i int) (string, int, int) {
	fake.addProductFileMutex.RLock()
	defer fake.addProductFileMutex.RUnlock()
	argsForCall := fake.addProductFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *UploadClient) AddProductFileReturns(result1 error) {
	fake.addProductFileMutex.Lock()
	defer fake.addProductFileMutex.Unlock()
	fake.AddProductFileStub = nil
	fake.addProductFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *UploadClient) AddProductFileReturnsOnCall(i int, result1 error) {
	fake.addProductFileMutex.Lock()
	defer fake.addProductFileMutex.Unlock()
	fake.AddProductFileStub = nil
	if fake.addProductFileReturnsOnCall == nil {
		fake.addProductFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addProductFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *UploadClient) CreateProductFile(arg1 pivnet.CreateProductFileConfig) (pivnet.ProductFile, error) {
	fake.createProductFileMutex.Lock()
	ret, specificReturn := fake.createProductFileReturnsOnCall[len(fake.createProductFileArgsForCall)]
	fake.createProductFileArgsForCall = append(fake.createProductFileArgsForCall, struct {
		arg1 pivnet.CreateProductFileConfig
	}{arg1})
	fake.recordInvocation("CreateProductFile", []interface{}{arg1})
	fake.createProductFileMutex.Unlock()
	if fake.CreateProductFileStub != nil {
		return fake.CreateProductFileStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createProductFileReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *UploadClient) CreateProductFileCallCount() int {
	fake.createProductFileMutex.RLock()
	defer fake.createProductFileMutex.RUnlock()
	return len(fake.createProductFileArgsForCall)
}

func (fake *UploadClient) CreateProductFileCalls(stub func(pivnet.CreateProductFileConfig) (pivnet.ProductFile, error)) {
	fake.createProductFileMutex.Lock()
	defer fake.createProductFileMutex.Unlock()
	fake.CreateProductFileStub = stub
}

func (fake *UploadClient) CreateProductFileArgsForCall(i int) pivnet.CreateProductFileConfig {
	fake.createProductFileMutex.RLock()
	defer fake.createProductFileMutex.RUnlock()
	argsForCall := fake.createProductFileArgsForCall[i]
	return argsForCall.arg1
}

func (fake *UploadClient) CreateProductFileReturns(result1 pivnet.ProductFile, result2 error) {
	fake.createProductFileMutex.Lock()
	defer fake.createProductFileMutex.Unlock()
	fake.CreateProductFileStub = nil
	fake.createProductFileReturns = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *UploadClient) CreateProductFileReturnsOnCall(i int, result1 pivnet.ProductFile, result2 error) {
	fake.createProductFileMutex.Lock()
	defer fake.createProductFileMutex.Unlock()
	fake.CreateProductFileStub = nil
	if fake.createProductFileReturnsOnCall == nil {
		fake.createProductFileReturnsOnCall = make(map[int]struct {
			result1 pivnet.ProductFile
			result2 error
		})
	}
	fake.createProductFileReturnsOnCall[i] = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *UploadClient) DeleteProductFile(arg1 string, arg2 int) (pivnet.ProductFile, error) {
	fake.deleteProductFileMutex.Lock()
	ret, specificReturn := fake.deleteProductFileReturnsOnCall[len(fake.deleteProductFileArgsForCall)]
	fake.deleteProductFileArgsForCall = append(fake.deleteProductFileArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("DeleteProductFile", []interface{}{arg1, arg2})
	fake.deleteProductFileMutex.Unlock()
	if fake.DeleteProductFileStub != nil {
		return fake.DeleteProductFileStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteProductFileReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *UploadClient) DeleteProductFileCallCount() int {
	fake.deleteProductFileMutex.RLock()
	defer fake.deleteProductFileMutex.RUnlock()
	return len(fake.deleteProductFileArgsForCall)
}

func (fake *UploadClient) DeleteProductFileCalls(stub func(string, int) (pivnet.ProductFile, error)) {
	fake.deleteProductFileMutex.Lock()
	defer fake.deleteProductFileMutex.Unlock()
	fake.DeleteProductFileStub = stub
}

func (fake *UploadClient) DeleteProductFileArgsForCall(i int) (string, int) {
	fake.deleteProductFileMutex.RLock()
	defer fake.deleteProductFileMutex.RUnlock()
	argsForCall := fake.deleteProductFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *UploadClient) DeleteProductFileReturns(result1 pivnet.ProductFile, result2 error) {
	fake.deleteProductFileMutex.Lock()
	defer fake.deleteProductFileMutex.Unlock()
	fake.DeleteProductFileStub = nil
	fake.deleteProductFileReturns = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *UploadClient) DeleteProductFileReturnsOnCall(i int, result1 pivnet.ProductFile, result2 error) {
	fake.deleteProductFileMutex.Lock()
	defer fake.deleteProductFileMutex.Unlock()
	fake.DeleteProductFileStub = nil
	if fake.deleteProductFileReturnsOnCall == nil {
		fake.deleteProductFileReturnsOnCall = make(map[int]struct {
			result1 pivnet.ProductFile
			result2 error
		})
	}
	fake.deleteProductFileReturnsOnCall[i] = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *UploadClient) FindProductForSlug(arg1 string) (pivnet.Product, error) {
	fake.findProductForSlugMutex.Lock()
	ret, specificReturn := fake.findProductForSlugReturnsOnCall[len(fake.findProductForSlugArgsForCall)]
	fake.findProductForSlugArgsForCall = append(fake.findProductForSlugArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FindProductForSlug", []interface{}{arg1})
	fake.findProductForSlugMutex.Unlock()
	if fake.FindProductForSlugStub != nil {
		return fake.FindProductForSlugStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findProductForSlugReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *UploadClient) FindProductForSlugCallCount() int {
	fake.findProductForSlugMutex.RLock()
	defer fake.findProductForSlugMutex.RUnlock()
	return len(fake.findProductForSlugArgsForCall)
}

func (fake *UploadClient) FindProductForSlugCalls(stub func(string) (pivnet.Product, error)) {
	fake.findProductForSlugMutex.Lock()
	defer fake.findProductForSlugMutex.Unlock()
	fake.FindProductForSlugStub = stub
}

func (fake *UploadClient) FindProductForSlugArgsForCall(i int) string {
	fake.findProductForSlugMutex.RLock()
	defer fake.findProductForSlugMutex.RUnlock()
	argsForCall := fake.findProductForSlugArgsForCall[i]
	return argsForCall.arg1
}

func (fake *UploadClient) FindProductForSlugReturns(result1 pivnet.Product, result2 error) {
	fake.findProductForSlugMutex.Lock()
	defer fake.findProductForSlugMutex.Unlock()
	fake.FindProductForSlugStub = nil
	fake.findProductForSlugReturns = struct {
		result1 pivnet.Product
		result2 error
	}{result1, result2}
}

func (fake *UploadClient) FindProductForSlugReturnsOnCall(i int, result1 pivnet.Product, result2 error) {
	fake.findProductForSlugMutex.Lock()
	defer fake.findProductForSlugMutex.Unlock()
	fake.FindProductForSlugStub = nil
	if fake.findProductForSlugReturnsOnCall == nil {
		fake.findProductForSlugReturnsOnCall = make(map[int]struct {
			result1 pivnet.Product
			result2 error
		})
	}
	fake.findProductForSlugReturnsOnCall[i] = struct {
		result1 pivnet.Product
		result2 error
	}{result1, result2}
}

func (fake *UploadClient) ProductFile(arg1 string, arg2 int) (pivnet.ProductFile, error) {
	fake.productFileMutex.Lock()
	ret, specificReturn := fake.productFileReturnsOnCall[len(fake.productFileArgsForCall)]
	fake.productFileArgsForCall = append(fake.productFileArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("ProductFile", []interface{}{arg1, arg2})
	fake.productFileMutex.Unlock()
	if fake.ProductFileStub != nil {
		return fake.ProductFileStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.productFileReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *UploadClient) ProductFileCallCount() int {
	fake.productFileMutex.RLock()
	defer fake.productFileMutex.RUnlock()
	return len(fake.productFileArgsForCall)
}

func (fake *UploadClient) ProductFileCalls(stub func(string, int) (pivnet.ProductFile, error)) {
	fake.productFileMutex.Lock()
	defer fake.productFileMutex.Unlock()
	fake.ProductFileStub = stub
}

func (fake *UploadClient) ProductFileArgsForCall(i int) (string, int) {
	fake.productFileMutex.RLock()
	defer fake.productFileMutex.RUnlock()
	argsForCall := fake.productFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *UploadClient) ProductFileReturns(result1 pivnet.ProductFile, result2 error) {
	fake.productFileMutex.Lock()
	defer fake.productFileMutex.Unlock()
	fake.ProductFileStub = nil
	fake.productFileReturns = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *UploadClient) ProductFileReturnsOnCall(i int, result1 pivnet.ProductFile, result2 error) {
	fake.productFileMutex.Lock()
	defer fake.productFileMutex.Unlock()
	fake.ProductFileStub = nil
	if fake.productFileReturnsOnCall == nil {
		fake.productFileReturnsOnCall = make(map[int]struct {
			result1 pivnet.ProductFile
			result2 error
		})
	}
	fake.productFileReturnsOnCall[i] = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *UploadClient) ProductFiles(arg1 string) ([]pivnet.ProductFile, error) {
	fake.productFilesMutex.Lock()
	ret, specificReturn := fake.productFilesReturnsOnCall[len(fake.productFilesArgsForCall)]
	fake.productFilesArgsForCall = append(fake.productFilesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ProductFiles", []interface{}{arg1})
	fake.productFilesMutex.Unlock()
	if fake.ProductFilesStub != nil {
		return fake.ProductFilesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.productFilesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *UploadClient) ProductFilesCallCount() int {
	fake.productFilesMutex.RLock()
	defer fake.productFilesMutex.RUnlock()
	return len(fake.productFilesArgsForCall)
}

func (fake *UploadClient) ProductFilesCalls(stub func(string) ([]pivnet.ProductFile, error)) {
	fake.productFilesMutex.Lock()
	defer fake.productFilesMutex.Unlock()
	fake.ProductFilesStub = stub
}

func (fake *UploadClient) ProductFilesArgsForCall(i int) string {
	fake.productFilesMutex.RLock()
	defer fake.productFilesMutex.RUnlock()
	argsForCall := fake.productFilesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *UploadClient) ProductFilesReturns(result1 []pivnet.ProductFile, result2 error) {
	fake.productFilesMutex.Lock()
	defer fake.productFilesMutex.Unlock()
	fake.ProductFilesStub = nil
	fake.productFilesReturns = struct {
		result1 []pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *UploadClient) ProductFilesReturnsOnCall(i int, result1 []pivnet.ProductFile, result2 error) {
	fake.productFilesMutex.Lock()
	defer fake.productFilesMutex.Unlock()
	fake.ProductFilesStub = nil
	if fake.productFilesReturnsOnCall == nil {
		fake.productFilesReturnsOnCall = make(map[int]struct {
			result1 []pivnet.ProductFile
			result2 error
		})
	}
	fake.productFilesReturnsOnCall[i] = struct {
		result1 []pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *UploadClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addProductFileMutex.RLock()
	defer fake.addProductFileMutex.RUnlock()
	fake.createProductFileMutex.RLock()
	defer fake.createProductFileMutex.RUnlock()
	fake.deleteProductFileMutex.RLock()
	defer fake.deleteProductFileMutex.RUnlock()
	fake.findProductForSlugMutex.RLock()
	defer fake.findProductForSlugMutex.RUnlock()
	fake.productFileMutex.RLock()
	defer fake.productFileMutex.RUnlock()
	fake.productFilesMutex.RLock()
	defer fake.productFilesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *UploadClient) recordInvocation(key string, args []interface{}) {
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
