// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeAllowSpaceSSHActor struct {
	AllowSpaceSSHStub        func(string, string) (v7action.Warnings, error)
	allowSpaceSSHMutex       sync.RWMutex
	allowSpaceSSHArgsForCall []struct {
		arg1 string
		arg2 string
	}
	allowSpaceSSHReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	allowSpaceSSHReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	GetSpaceFeatureStub        func(string, string, string) (bool, v7action.Warnings, error)
	getSpaceFeatureMutex       sync.RWMutex
	getSpaceFeatureArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	getSpaceFeatureReturns struct {
		result1 bool
		result2 v7action.Warnings
		result3 error
	}
	getSpaceFeatureReturnsOnCall map[int]struct {
		result1 bool
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAllowSpaceSSHActor) AllowSpaceSSH(arg1 string, arg2 string) (v7action.Warnings, error) {
	fake.allowSpaceSSHMutex.Lock()
	ret, specificReturn := fake.allowSpaceSSHReturnsOnCall[len(fake.allowSpaceSSHArgsForCall)]
	fake.allowSpaceSSHArgsForCall = append(fake.allowSpaceSSHArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("AllowSpaceSSH", []interface{}{arg1, arg2})
	fake.allowSpaceSSHMutex.Unlock()
	if fake.AllowSpaceSSHStub != nil {
		return fake.AllowSpaceSSHStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.allowSpaceSSHReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAllowSpaceSSHActor) AllowSpaceSSHCallCount() int {
	fake.allowSpaceSSHMutex.RLock()
	defer fake.allowSpaceSSHMutex.RUnlock()
	return len(fake.allowSpaceSSHArgsForCall)
}

func (fake *FakeAllowSpaceSSHActor) AllowSpaceSSHCalls(stub func(string, string) (v7action.Warnings, error)) {
	fake.allowSpaceSSHMutex.Lock()
	defer fake.allowSpaceSSHMutex.Unlock()
	fake.AllowSpaceSSHStub = stub
}

func (fake *FakeAllowSpaceSSHActor) AllowSpaceSSHArgsForCall(i int) (string, string) {
	fake.allowSpaceSSHMutex.RLock()
	defer fake.allowSpaceSSHMutex.RUnlock()
	argsForCall := fake.allowSpaceSSHArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAllowSpaceSSHActor) AllowSpaceSSHReturns(result1 v7action.Warnings, result2 error) {
	fake.allowSpaceSSHMutex.Lock()
	defer fake.allowSpaceSSHMutex.Unlock()
	fake.AllowSpaceSSHStub = nil
	fake.allowSpaceSSHReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeAllowSpaceSSHActor) AllowSpaceSSHReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.allowSpaceSSHMutex.Lock()
	defer fake.allowSpaceSSHMutex.Unlock()
	fake.AllowSpaceSSHStub = nil
	if fake.allowSpaceSSHReturnsOnCall == nil {
		fake.allowSpaceSSHReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.allowSpaceSSHReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeAllowSpaceSSHActor) GetSpaceFeature(arg1 string, arg2 string, arg3 string) (bool, v7action.Warnings, error) {
	fake.getSpaceFeatureMutex.Lock()
	ret, specificReturn := fake.getSpaceFeatureReturnsOnCall[len(fake.getSpaceFeatureArgsForCall)]
	fake.getSpaceFeatureArgsForCall = append(fake.getSpaceFeatureArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetSpaceFeature", []interface{}{arg1, arg2, arg3})
	fake.getSpaceFeatureMutex.Unlock()
	if fake.GetSpaceFeatureStub != nil {
		return fake.GetSpaceFeatureStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getSpaceFeatureReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeAllowSpaceSSHActor) GetSpaceFeatureCallCount() int {
	fake.getSpaceFeatureMutex.RLock()
	defer fake.getSpaceFeatureMutex.RUnlock()
	return len(fake.getSpaceFeatureArgsForCall)
}

func (fake *FakeAllowSpaceSSHActor) GetSpaceFeatureCalls(stub func(string, string, string) (bool, v7action.Warnings, error)) {
	fake.getSpaceFeatureMutex.Lock()
	defer fake.getSpaceFeatureMutex.Unlock()
	fake.GetSpaceFeatureStub = stub
}

func (fake *FakeAllowSpaceSSHActor) GetSpaceFeatureArgsForCall(i int) (string, string, string) {
	fake.getSpaceFeatureMutex.RLock()
	defer fake.getSpaceFeatureMutex.RUnlock()
	argsForCall := fake.getSpaceFeatureArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAllowSpaceSSHActor) GetSpaceFeatureReturns(result1 bool, result2 v7action.Warnings, result3 error) {
	fake.getSpaceFeatureMutex.Lock()
	defer fake.getSpaceFeatureMutex.Unlock()
	fake.GetSpaceFeatureStub = nil
	fake.getSpaceFeatureReturns = struct {
		result1 bool
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAllowSpaceSSHActor) GetSpaceFeatureReturnsOnCall(i int, result1 bool, result2 v7action.Warnings, result3 error) {
	fake.getSpaceFeatureMutex.Lock()
	defer fake.getSpaceFeatureMutex.Unlock()
	fake.GetSpaceFeatureStub = nil
	if fake.getSpaceFeatureReturnsOnCall == nil {
		fake.getSpaceFeatureReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getSpaceFeatureReturnsOnCall[i] = struct {
		result1 bool
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAllowSpaceSSHActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allowSpaceSSHMutex.RLock()
	defer fake.allowSpaceSSHMutex.RUnlock()
	fake.getSpaceFeatureMutex.RLock()
	defer fake.getSpaceFeatureMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAllowSpaceSSHActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.AllowSpaceSSHActor = new(FakeAllowSpaceSSHActor)