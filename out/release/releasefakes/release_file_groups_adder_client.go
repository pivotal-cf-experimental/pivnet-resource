// This file was generated by counterfeiter
package releasefakes

import (
	"sync"

	"github.com/pivotal-cf/go-pivnet"
)

type ReleaseFileGroupsAdderClient struct {
	AddFileGroupStub        func(productSlug string, releaseID int, fileGroupID int) error
	addFileGroupMutex       sync.RWMutex
	addFileGroupArgsForCall []struct {
		productSlug string
		releaseID   int
		fileGroupID int
	}
	addFileGroupReturns struct {
		result1 error
	}
	addFileGroupReturnsOnCall map[int]struct {
		result1 error
	}
	CreateFileGroupStub        func(config pivnet.CreateFileGroupConfig) (pivnet.FileGroup, error)
	createFileGroupMutex       sync.RWMutex
	createFileGroupArgsForCall []struct {
		config pivnet.CreateFileGroupConfig
	}
	createFileGroupReturns struct {
		result1 pivnet.FileGroup
		result2 error
	}
	createFileGroupReturnsOnCall map[int]struct {
		result1 pivnet.FileGroup
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ReleaseFileGroupsAdderClient) AddFileGroup(productSlug string, releaseID int, fileGroupID int) error {
	fake.addFileGroupMutex.Lock()
	ret, specificReturn := fake.addFileGroupReturnsOnCall[len(fake.addFileGroupArgsForCall)]
	fake.addFileGroupArgsForCall = append(fake.addFileGroupArgsForCall, struct {
		productSlug string
		releaseID   int
		fileGroupID int
	}{productSlug, releaseID, fileGroupID})
	fake.recordInvocation("AddFileGroup", []interface{}{productSlug, releaseID, fileGroupID})
	fake.addFileGroupMutex.Unlock()
	if fake.AddFileGroupStub != nil {
		return fake.AddFileGroupStub(productSlug, releaseID, fileGroupID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addFileGroupReturns.result1
}

func (fake *ReleaseFileGroupsAdderClient) AddFileGroupCallCount() int {
	fake.addFileGroupMutex.RLock()
	defer fake.addFileGroupMutex.RUnlock()
	return len(fake.addFileGroupArgsForCall)
}

func (fake *ReleaseFileGroupsAdderClient) AddFileGroupArgsForCall(i int) (string, int, int) {
	fake.addFileGroupMutex.RLock()
	defer fake.addFileGroupMutex.RUnlock()
	return fake.addFileGroupArgsForCall[i].productSlug, fake.addFileGroupArgsForCall[i].releaseID, fake.addFileGroupArgsForCall[i].fileGroupID
}

func (fake *ReleaseFileGroupsAdderClient) AddFileGroupReturns(result1 error) {
	fake.AddFileGroupStub = nil
	fake.addFileGroupReturns = struct {
		result1 error
	}{result1}
}

func (fake *ReleaseFileGroupsAdderClient) AddFileGroupReturnsOnCall(i int, result1 error) {
	fake.AddFileGroupStub = nil
	if fake.addFileGroupReturnsOnCall == nil {
		fake.addFileGroupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addFileGroupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ReleaseFileGroupsAdderClient) CreateFileGroup(config pivnet.CreateFileGroupConfig) (pivnet.FileGroup, error) {
	fake.createFileGroupMutex.Lock()
	ret, specificReturn := fake.createFileGroupReturnsOnCall[len(fake.createFileGroupArgsForCall)]
	fake.createFileGroupArgsForCall = append(fake.createFileGroupArgsForCall, struct {
		config pivnet.CreateFileGroupConfig
	}{config})
	fake.recordInvocation("CreateFileGroup", []interface{}{config})
	fake.createFileGroupMutex.Unlock()
	if fake.CreateFileGroupStub != nil {
		return fake.CreateFileGroupStub(config)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createFileGroupReturns.result1, fake.createFileGroupReturns.result2
}

func (fake *ReleaseFileGroupsAdderClient) CreateFileGroupCallCount() int {
	fake.createFileGroupMutex.RLock()
	defer fake.createFileGroupMutex.RUnlock()
	return len(fake.createFileGroupArgsForCall)
}

func (fake *ReleaseFileGroupsAdderClient) CreateFileGroupArgsForCall(i int) pivnet.CreateFileGroupConfig {
	fake.createFileGroupMutex.RLock()
	defer fake.createFileGroupMutex.RUnlock()
	return fake.createFileGroupArgsForCall[i].config
}

func (fake *ReleaseFileGroupsAdderClient) CreateFileGroupReturns(result1 pivnet.FileGroup, result2 error) {
	fake.CreateFileGroupStub = nil
	fake.createFileGroupReturns = struct {
		result1 pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *ReleaseFileGroupsAdderClient) CreateFileGroupReturnsOnCall(i int, result1 pivnet.FileGroup, result2 error) {
	fake.CreateFileGroupStub = nil
	if fake.createFileGroupReturnsOnCall == nil {
		fake.createFileGroupReturnsOnCall = make(map[int]struct {
			result1 pivnet.FileGroup
			result2 error
		})
	}
	fake.createFileGroupReturnsOnCall[i] = struct {
		result1 pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *ReleaseFileGroupsAdderClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addFileGroupMutex.RLock()
	defer fake.addFileGroupMutex.RUnlock()
	fake.createFileGroupMutex.RLock()
	defer fake.createFileGroupMutex.RUnlock()
	return fake.invocations
}

func (fake *ReleaseFileGroupsAdderClient) recordInvocation(key string, args []interface{}) {
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
