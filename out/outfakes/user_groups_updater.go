// Code generated by counterfeiter. DO NOT EDIT.
package outfakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v6"
)

type UserGroupsUpdater struct {
	UpdateUserGroupsStub        func(pivnet.Release) (pivnet.Release, error)
	updateUserGroupsMutex       sync.RWMutex
	updateUserGroupsArgsForCall []struct {
		arg1 pivnet.Release
	}
	updateUserGroupsReturns struct {
		result1 pivnet.Release
		result2 error
	}
	updateUserGroupsReturnsOnCall map[int]struct {
		result1 pivnet.Release
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *UserGroupsUpdater) UpdateUserGroups(arg1 pivnet.Release) (pivnet.Release, error) {
	fake.updateUserGroupsMutex.Lock()
	ret, specificReturn := fake.updateUserGroupsReturnsOnCall[len(fake.updateUserGroupsArgsForCall)]
	fake.updateUserGroupsArgsForCall = append(fake.updateUserGroupsArgsForCall, struct {
		arg1 pivnet.Release
	}{arg1})
	fake.recordInvocation("UpdateUserGroups", []interface{}{arg1})
	fake.updateUserGroupsMutex.Unlock()
	if fake.UpdateUserGroupsStub != nil {
		return fake.UpdateUserGroupsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateUserGroupsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *UserGroupsUpdater) UpdateUserGroupsCallCount() int {
	fake.updateUserGroupsMutex.RLock()
	defer fake.updateUserGroupsMutex.RUnlock()
	return len(fake.updateUserGroupsArgsForCall)
}

func (fake *UserGroupsUpdater) UpdateUserGroupsCalls(stub func(pivnet.Release) (pivnet.Release, error)) {
	fake.updateUserGroupsMutex.Lock()
	defer fake.updateUserGroupsMutex.Unlock()
	fake.UpdateUserGroupsStub = stub
}

func (fake *UserGroupsUpdater) UpdateUserGroupsArgsForCall(i int) pivnet.Release {
	fake.updateUserGroupsMutex.RLock()
	defer fake.updateUserGroupsMutex.RUnlock()
	argsForCall := fake.updateUserGroupsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *UserGroupsUpdater) UpdateUserGroupsReturns(result1 pivnet.Release, result2 error) {
	fake.updateUserGroupsMutex.Lock()
	defer fake.updateUserGroupsMutex.Unlock()
	fake.UpdateUserGroupsStub = nil
	fake.updateUserGroupsReturns = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *UserGroupsUpdater) UpdateUserGroupsReturnsOnCall(i int, result1 pivnet.Release, result2 error) {
	fake.updateUserGroupsMutex.Lock()
	defer fake.updateUserGroupsMutex.Unlock()
	fake.UpdateUserGroupsStub = nil
	if fake.updateUserGroupsReturnsOnCall == nil {
		fake.updateUserGroupsReturnsOnCall = make(map[int]struct {
			result1 pivnet.Release
			result2 error
		})
	}
	fake.updateUserGroupsReturnsOnCall[i] = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *UserGroupsUpdater) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.updateUserGroupsMutex.RLock()
	defer fake.updateUserGroupsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *UserGroupsUpdater) recordInvocation(key string, args []interface{}) {
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
