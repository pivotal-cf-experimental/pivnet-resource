// Code generated by counterfeiter. DO NOT EDIT.
package downloaderfakes

import (
	"io"
	"sync"

	"github.com/pivotal-cf/go-pivnet/v6/download"
)

type FakeClient struct {
	DownloadProductFileStub        func(*download.FileInfo, string, int, int, io.Writer) error
	downloadProductFileMutex       sync.RWMutex
	downloadProductFileArgsForCall []struct {
		arg1 *download.FileInfo
		arg2 string
		arg3 int
		arg4 int
		arg5 io.Writer
	}
	downloadProductFileReturns struct {
		result1 error
	}
	downloadProductFileReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) DownloadProductFile(arg1 *download.FileInfo, arg2 string, arg3 int, arg4 int, arg5 io.Writer) error {
	fake.downloadProductFileMutex.Lock()
	ret, specificReturn := fake.downloadProductFileReturnsOnCall[len(fake.downloadProductFileArgsForCall)]
	fake.downloadProductFileArgsForCall = append(fake.downloadProductFileArgsForCall, struct {
		arg1 *download.FileInfo
		arg2 string
		arg3 int
		arg4 int
		arg5 io.Writer
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("DownloadProductFile", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.downloadProductFileMutex.Unlock()
	if fake.DownloadProductFileStub != nil {
		return fake.DownloadProductFileStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.downloadProductFileReturns
	return fakeReturns.result1
}

func (fake *FakeClient) DownloadProductFileCallCount() int {
	fake.downloadProductFileMutex.RLock()
	defer fake.downloadProductFileMutex.RUnlock()
	return len(fake.downloadProductFileArgsForCall)
}

func (fake *FakeClient) DownloadProductFileCalls(stub func(*download.FileInfo, string, int, int, io.Writer) error) {
	fake.downloadProductFileMutex.Lock()
	defer fake.downloadProductFileMutex.Unlock()
	fake.DownloadProductFileStub = stub
}

func (fake *FakeClient) DownloadProductFileArgsForCall(i int) (*download.FileInfo, string, int, int, io.Writer) {
	fake.downloadProductFileMutex.RLock()
	defer fake.downloadProductFileMutex.RUnlock()
	argsForCall := fake.downloadProductFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeClient) DownloadProductFileReturns(result1 error) {
	fake.downloadProductFileMutex.Lock()
	defer fake.downloadProductFileMutex.Unlock()
	fake.DownloadProductFileStub = nil
	fake.downloadProductFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DownloadProductFileReturnsOnCall(i int, result1 error) {
	fake.downloadProductFileMutex.Lock()
	defer fake.downloadProductFileMutex.Unlock()
	fake.DownloadProductFileStub = nil
	if fake.downloadProductFileReturnsOnCall == nil {
		fake.downloadProductFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadProductFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downloadProductFileMutex.RLock()
	defer fake.downloadProductFileMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
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
