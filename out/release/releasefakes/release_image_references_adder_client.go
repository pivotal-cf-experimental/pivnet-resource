// Code generated by counterfeiter. DO NOT EDIT.
package releasefakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v2"
)

type ReleaseImageReferencesAdderClient struct {
	AddImageReferenceStub        func(productSlug string, releaseID int, imageReferenceID int) error
	addImageReferenceMutex       sync.RWMutex
	addImageReferenceArgsForCall []struct {
		productSlug      string
		releaseID        int
		imageReferenceID int
	}
	addImageReferenceReturns struct {
		result1 error
	}
	addImageReferenceReturnsOnCall map[int]struct {
		result1 error
	}
	CreateImageReferenceStub        func(config pivnet.CreateImageReferenceConfig) (pivnet.ImageReference, error)
	createImageReferenceMutex       sync.RWMutex
	createImageReferenceArgsForCall []struct {
		config pivnet.CreateImageReferenceConfig
	}
	createImageReferenceReturns struct {
		result1 pivnet.ImageReference
		result2 error
	}
	createImageReferenceReturnsOnCall map[int]struct {
		result1 pivnet.ImageReference
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ReleaseImageReferencesAdderClient) AddImageReference(productSlug string, releaseID int, imageReferenceID int) error {
	fake.addImageReferenceMutex.Lock()
	ret, specificReturn := fake.addImageReferenceReturnsOnCall[len(fake.addImageReferenceArgsForCall)]
	fake.addImageReferenceArgsForCall = append(fake.addImageReferenceArgsForCall, struct {
		productSlug      string
		releaseID        int
		imageReferenceID int
	}{productSlug, releaseID, imageReferenceID})
	fake.recordInvocation("AddImageReference", []interface{}{productSlug, releaseID, imageReferenceID})
	fake.addImageReferenceMutex.Unlock()
	if fake.AddImageReferenceStub != nil {
		return fake.AddImageReferenceStub(productSlug, releaseID, imageReferenceID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addImageReferenceReturns.result1
}

func (fake *ReleaseImageReferencesAdderClient) AddImageReferenceCallCount() int {
	fake.addImageReferenceMutex.RLock()
	defer fake.addImageReferenceMutex.RUnlock()
	return len(fake.addImageReferenceArgsForCall)
}

func (fake *ReleaseImageReferencesAdderClient) AddImageReferenceArgsForCall(i int) (string, int, int) {
	fake.addImageReferenceMutex.RLock()
	defer fake.addImageReferenceMutex.RUnlock()
	return fake.addImageReferenceArgsForCall[i].productSlug, fake.addImageReferenceArgsForCall[i].releaseID, fake.addImageReferenceArgsForCall[i].imageReferenceID
}

func (fake *ReleaseImageReferencesAdderClient) AddImageReferenceReturns(result1 error) {
	fake.AddImageReferenceStub = nil
	fake.addImageReferenceReturns = struct {
		result1 error
	}{result1}
}

func (fake *ReleaseImageReferencesAdderClient) AddImageReferenceReturnsOnCall(i int, result1 error) {
	fake.AddImageReferenceStub = nil
	if fake.addImageReferenceReturnsOnCall == nil {
		fake.addImageReferenceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addImageReferenceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ReleaseImageReferencesAdderClient) CreateImageReference(config pivnet.CreateImageReferenceConfig) (pivnet.ImageReference, error) {
	fake.createImageReferenceMutex.Lock()
	ret, specificReturn := fake.createImageReferenceReturnsOnCall[len(fake.createImageReferenceArgsForCall)]
	fake.createImageReferenceArgsForCall = append(fake.createImageReferenceArgsForCall, struct {
		config pivnet.CreateImageReferenceConfig
	}{config})
	fake.recordInvocation("CreateImageReference", []interface{}{config})
	fake.createImageReferenceMutex.Unlock()
	if fake.CreateImageReferenceStub != nil {
		return fake.CreateImageReferenceStub(config)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createImageReferenceReturns.result1, fake.createImageReferenceReturns.result2
}

func (fake *ReleaseImageReferencesAdderClient) CreateImageReferenceCallCount() int {
	fake.createImageReferenceMutex.RLock()
	defer fake.createImageReferenceMutex.RUnlock()
	return len(fake.createImageReferenceArgsForCall)
}

func (fake *ReleaseImageReferencesAdderClient) CreateImageReferenceArgsForCall(i int) pivnet.CreateImageReferenceConfig {
	fake.createImageReferenceMutex.RLock()
	defer fake.createImageReferenceMutex.RUnlock()
	return fake.createImageReferenceArgsForCall[i].config
}

func (fake *ReleaseImageReferencesAdderClient) CreateImageReferenceReturns(result1 pivnet.ImageReference, result2 error) {
	fake.CreateImageReferenceStub = nil
	fake.createImageReferenceReturns = struct {
		result1 pivnet.ImageReference
		result2 error
	}{result1, result2}
}

func (fake *ReleaseImageReferencesAdderClient) CreateImageReferenceReturnsOnCall(i int, result1 pivnet.ImageReference, result2 error) {
	fake.CreateImageReferenceStub = nil
	if fake.createImageReferenceReturnsOnCall == nil {
		fake.createImageReferenceReturnsOnCall = make(map[int]struct {
			result1 pivnet.ImageReference
			result2 error
		})
	}
	fake.createImageReferenceReturnsOnCall[i] = struct {
		result1 pivnet.ImageReference
		result2 error
	}{result1, result2}
}

func (fake *ReleaseImageReferencesAdderClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addImageReferenceMutex.RLock()
	defer fake.addImageReferenceMutex.RUnlock()
	fake.createImageReferenceMutex.RLock()
	defer fake.createImageReferenceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ReleaseImageReferencesAdderClient) recordInvocation(key string, args []interface{}) {
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