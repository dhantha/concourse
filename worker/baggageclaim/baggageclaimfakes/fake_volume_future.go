// Code generated by counterfeiter. DO NOT EDIT.
package baggageclaimfakes

import (
	"sync"

	"github.com/concourse/concourse/worker/baggageclaim"
)

type FakeVolumeFuture struct {
	DestroyStub        func() error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
	}
	destroyReturns struct {
		result1 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 error
	}
	WaitStub        func() (baggageclaim.Volume, error)
	waitMutex       sync.RWMutex
	waitArgsForCall []struct {
	}
	waitReturns struct {
		result1 baggageclaim.Volume
		result2 error
	}
	waitReturnsOnCall map[int]struct {
		result1 baggageclaim.Volume
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVolumeFuture) Destroy() error {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
	}{})
	fake.recordInvocation("Destroy", []interface{}{})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.destroyReturns
	return fakeReturns.result1
}

func (fake *FakeVolumeFuture) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeVolumeFuture) DestroyCalls(stub func() error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = stub
}

func (fake *FakeVolumeFuture) DestroyReturns(result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolumeFuture) DestroyReturnsOnCall(i int, result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	if fake.destroyReturnsOnCall == nil {
		fake.destroyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolumeFuture) Wait() (baggageclaim.Volume, error) {
	fake.waitMutex.Lock()
	ret, specificReturn := fake.waitReturnsOnCall[len(fake.waitArgsForCall)]
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct {
	}{})
	fake.recordInvocation("Wait", []interface{}{})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.waitReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolumeFuture) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeVolumeFuture) WaitCalls(stub func() (baggageclaim.Volume, error)) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = stub
}

func (fake *FakeVolumeFuture) WaitReturns(result1 baggageclaim.Volume, result2 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 baggageclaim.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFuture) WaitReturnsOnCall(i int, result1 baggageclaim.Volume, result2 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	if fake.waitReturnsOnCall == nil {
		fake.waitReturnsOnCall = make(map[int]struct {
			result1 baggageclaim.Volume
			result2 error
		})
	}
	fake.waitReturnsOnCall[i] = struct {
		result1 baggageclaim.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFuture) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVolumeFuture) recordInvocation(key string, args []interface{}) {
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

var _ baggageclaim.VolumeFuture = new(FakeVolumeFuture)
