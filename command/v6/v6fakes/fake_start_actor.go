// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v2action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeStartActor struct {
	GetApplicationByNameAndSpaceStub        func(string, string) (v2action.Application, v2action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}
	GetApplicationSummaryByNameAndSpaceStub        func(string, string) (v2action.ApplicationSummary, v2action.Warnings, error)
	getApplicationSummaryByNameAndSpaceMutex       sync.RWMutex
	getApplicationSummaryByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getApplicationSummaryByNameAndSpaceReturns struct {
		result1 v2action.ApplicationSummary
		result2 v2action.Warnings
		result3 error
	}
	getApplicationSummaryByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v2action.ApplicationSummary
		result2 v2action.Warnings
		result3 error
	}
	GetStreamingLogsStub        func(string, sharedaction.LogCacheClient) (<-chan sharedaction.LogMessage, <-chan error, context.CancelFunc)
	getStreamingLogsMutex       sync.RWMutex
	getStreamingLogsArgsForCall []struct {
		arg1 string
		arg2 sharedaction.LogCacheClient
	}
	getStreamingLogsReturns struct {
		result1 <-chan sharedaction.LogMessage
		result2 <-chan error
		result3 context.CancelFunc
	}
	getStreamingLogsReturnsOnCall map[int]struct {
		result1 <-chan sharedaction.LogMessage
		result2 <-chan error
		result3 context.CancelFunc
	}
	StartApplicationStub        func(v2action.Application) (<-chan v2action.ApplicationStateChange, <-chan string, <-chan error)
	startApplicationMutex       sync.RWMutex
	startApplicationArgsForCall []struct {
		arg1 v2action.Application
	}
	startApplicationReturns struct {
		result1 <-chan v2action.ApplicationStateChange
		result2 <-chan string
		result3 <-chan error
	}
	startApplicationReturnsOnCall map[int]struct {
		result1 <-chan v2action.ApplicationStateChange
		result2 <-chan string
		result3 <-chan error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStartActor) GetApplicationByNameAndSpace(arg1 string, arg2 string) (v2action.Application, v2action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{arg1, arg2})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeStartActor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeStartActor) GetApplicationByNameAndSpaceCalls(stub func(string, string) (v2action.Application, v2action.Warnings, error)) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = stub
}

func (fake *FakeStartActor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getApplicationByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStartActor) GetApplicationByNameAndSpaceReturns(result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStartActor) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v2action.Application, result2 v2action.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v2action.Application
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v2action.Application
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStartActor) GetApplicationSummaryByNameAndSpace(arg1 string, arg2 string) (v2action.ApplicationSummary, v2action.Warnings, error) {
	fake.getApplicationSummaryByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationSummaryByNameAndSpaceReturnsOnCall[len(fake.getApplicationSummaryByNameAndSpaceArgsForCall)]
	fake.getApplicationSummaryByNameAndSpaceArgsForCall = append(fake.getApplicationSummaryByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetApplicationSummaryByNameAndSpace", []interface{}{arg1, arg2})
	fake.getApplicationSummaryByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationSummaryByNameAndSpaceStub != nil {
		return fake.GetApplicationSummaryByNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationSummaryByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeStartActor) GetApplicationSummaryByNameAndSpaceCallCount() int {
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationSummaryByNameAndSpaceArgsForCall)
}

func (fake *FakeStartActor) GetApplicationSummaryByNameAndSpaceCalls(stub func(string, string) (v2action.ApplicationSummary, v2action.Warnings, error)) {
	fake.getApplicationSummaryByNameAndSpaceMutex.Lock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.Unlock()
	fake.GetApplicationSummaryByNameAndSpaceStub = stub
}

func (fake *FakeStartActor) GetApplicationSummaryByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getApplicationSummaryByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStartActor) GetApplicationSummaryByNameAndSpaceReturns(result1 v2action.ApplicationSummary, result2 v2action.Warnings, result3 error) {
	fake.getApplicationSummaryByNameAndSpaceMutex.Lock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.Unlock()
	fake.GetApplicationSummaryByNameAndSpaceStub = nil
	fake.getApplicationSummaryByNameAndSpaceReturns = struct {
		result1 v2action.ApplicationSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStartActor) GetApplicationSummaryByNameAndSpaceReturnsOnCall(i int, result1 v2action.ApplicationSummary, result2 v2action.Warnings, result3 error) {
	fake.getApplicationSummaryByNameAndSpaceMutex.Lock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.Unlock()
	fake.GetApplicationSummaryByNameAndSpaceStub = nil
	if fake.getApplicationSummaryByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationSummaryByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v2action.ApplicationSummary
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getApplicationSummaryByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v2action.ApplicationSummary
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStartActor) GetStreamingLogs(arg1 string, arg2 sharedaction.LogCacheClient) (<-chan sharedaction.LogMessage, <-chan error, context.CancelFunc) {
	fake.getStreamingLogsMutex.Lock()
	ret, specificReturn := fake.getStreamingLogsReturnsOnCall[len(fake.getStreamingLogsArgsForCall)]
	fake.getStreamingLogsArgsForCall = append(fake.getStreamingLogsArgsForCall, struct {
		arg1 string
		arg2 sharedaction.LogCacheClient
	}{arg1, arg2})
	fake.recordInvocation("GetStreamingLogs", []interface{}{arg1, arg2})
	fake.getStreamingLogsMutex.Unlock()
	if fake.GetStreamingLogsStub != nil {
		return fake.GetStreamingLogsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getStreamingLogsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeStartActor) GetStreamingLogsCallCount() int {
	fake.getStreamingLogsMutex.RLock()
	defer fake.getStreamingLogsMutex.RUnlock()
	return len(fake.getStreamingLogsArgsForCall)
}

func (fake *FakeStartActor) GetStreamingLogsCalls(stub func(string, sharedaction.LogCacheClient) (<-chan sharedaction.LogMessage, <-chan error, context.CancelFunc)) {
	fake.getStreamingLogsMutex.Lock()
	defer fake.getStreamingLogsMutex.Unlock()
	fake.GetStreamingLogsStub = stub
}

func (fake *FakeStartActor) GetStreamingLogsArgsForCall(i int) (string, sharedaction.LogCacheClient) {
	fake.getStreamingLogsMutex.RLock()
	defer fake.getStreamingLogsMutex.RUnlock()
	argsForCall := fake.getStreamingLogsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStartActor) GetStreamingLogsReturns(result1 <-chan sharedaction.LogMessage, result2 <-chan error, result3 context.CancelFunc) {
	fake.getStreamingLogsMutex.Lock()
	defer fake.getStreamingLogsMutex.Unlock()
	fake.GetStreamingLogsStub = nil
	fake.getStreamingLogsReturns = struct {
		result1 <-chan sharedaction.LogMessage
		result2 <-chan error
		result3 context.CancelFunc
	}{result1, result2, result3}
}

func (fake *FakeStartActor) GetStreamingLogsReturnsOnCall(i int, result1 <-chan sharedaction.LogMessage, result2 <-chan error, result3 context.CancelFunc) {
	fake.getStreamingLogsMutex.Lock()
	defer fake.getStreamingLogsMutex.Unlock()
	fake.GetStreamingLogsStub = nil
	if fake.getStreamingLogsReturnsOnCall == nil {
		fake.getStreamingLogsReturnsOnCall = make(map[int]struct {
			result1 <-chan sharedaction.LogMessage
			result2 <-chan error
			result3 context.CancelFunc
		})
	}
	fake.getStreamingLogsReturnsOnCall[i] = struct {
		result1 <-chan sharedaction.LogMessage
		result2 <-chan error
		result3 context.CancelFunc
	}{result1, result2, result3}
}

func (fake *FakeStartActor) StartApplication(arg1 v2action.Application) (<-chan v2action.ApplicationStateChange, <-chan string, <-chan error) {
	fake.startApplicationMutex.Lock()
	ret, specificReturn := fake.startApplicationReturnsOnCall[len(fake.startApplicationArgsForCall)]
	fake.startApplicationArgsForCall = append(fake.startApplicationArgsForCall, struct {
		arg1 v2action.Application
	}{arg1})
	fake.recordInvocation("StartApplication", []interface{}{arg1})
	fake.startApplicationMutex.Unlock()
	if fake.StartApplicationStub != nil {
		return fake.StartApplicationStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.startApplicationReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeStartActor) StartApplicationCallCount() int {
	fake.startApplicationMutex.RLock()
	defer fake.startApplicationMutex.RUnlock()
	return len(fake.startApplicationArgsForCall)
}

func (fake *FakeStartActor) StartApplicationCalls(stub func(v2action.Application) (<-chan v2action.ApplicationStateChange, <-chan string, <-chan error)) {
	fake.startApplicationMutex.Lock()
	defer fake.startApplicationMutex.Unlock()
	fake.StartApplicationStub = stub
}

func (fake *FakeStartActor) StartApplicationArgsForCall(i int) v2action.Application {
	fake.startApplicationMutex.RLock()
	defer fake.startApplicationMutex.RUnlock()
	argsForCall := fake.startApplicationArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStartActor) StartApplicationReturns(result1 <-chan v2action.ApplicationStateChange, result2 <-chan string, result3 <-chan error) {
	fake.startApplicationMutex.Lock()
	defer fake.startApplicationMutex.Unlock()
	fake.StartApplicationStub = nil
	fake.startApplicationReturns = struct {
		result1 <-chan v2action.ApplicationStateChange
		result2 <-chan string
		result3 <-chan error
	}{result1, result2, result3}
}

func (fake *FakeStartActor) StartApplicationReturnsOnCall(i int, result1 <-chan v2action.ApplicationStateChange, result2 <-chan string, result3 <-chan error) {
	fake.startApplicationMutex.Lock()
	defer fake.startApplicationMutex.Unlock()
	fake.StartApplicationStub = nil
	if fake.startApplicationReturnsOnCall == nil {
		fake.startApplicationReturnsOnCall = make(map[int]struct {
			result1 <-chan v2action.ApplicationStateChange
			result2 <-chan string
			result3 <-chan error
		})
	}
	fake.startApplicationReturnsOnCall[i] = struct {
		result1 <-chan v2action.ApplicationStateChange
		result2 <-chan string
		result3 <-chan error
	}{result1, result2, result3}
}

func (fake *FakeStartActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.getApplicationSummaryByNameAndSpaceMutex.RLock()
	defer fake.getApplicationSummaryByNameAndSpaceMutex.RUnlock()
	fake.getStreamingLogsMutex.RLock()
	defer fake.getStreamingLogsMutex.RUnlock()
	fake.startApplicationMutex.RLock()
	defer fake.startApplicationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStartActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.StartActor = new(FakeStartActor)
