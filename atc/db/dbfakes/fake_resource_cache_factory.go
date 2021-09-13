// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
)

type FakeResourceCacheFactory struct {
	FindOrCreateResourceCacheStub        func(db.ResourceCacheUser, string, atc.Version, atc.Source, atc.Params, db.ResourceCache) (db.ResourceCache, error)
	findOrCreateResourceCacheMutex       sync.RWMutex
	findOrCreateResourceCacheArgsForCall []struct {
		arg1 db.ResourceCacheUser
		arg2 string
		arg3 atc.Version
		arg4 atc.Source
		arg5 atc.Params
		arg6 db.ResourceCache
	}
	findOrCreateResourceCacheReturns struct {
		result1 db.ResourceCache
		result2 error
	}
	findOrCreateResourceCacheReturnsOnCall map[int]struct {
		result1 db.ResourceCache
		result2 error
	}
	FindResourceCacheByIDStub        func(int) (db.ResourceCache, bool, error)
	findResourceCacheByIDMutex       sync.RWMutex
	findResourceCacheByIDArgsForCall []struct {
		arg1 int
	}
	findResourceCacheByIDReturns struct {
		result1 db.ResourceCache
		result2 bool
		result3 error
	}
	findResourceCacheByIDReturnsOnCall map[int]struct {
		result1 db.ResourceCache
		result2 bool
		result3 error
	}
	ResourceCacheMetadataStub        func(db.ResourceCache) (db.ResourceConfigMetadataFields, error)
	resourceCacheMetadataMutex       sync.RWMutex
	resourceCacheMetadataArgsForCall []struct {
		arg1 db.ResourceCache
	}
	resourceCacheMetadataReturns struct {
		result1 db.ResourceConfigMetadataFields
		result2 error
	}
	resourceCacheMetadataReturnsOnCall map[int]struct {
		result1 db.ResourceConfigMetadataFields
		result2 error
	}
	UpdateResourceCacheMetadataStub        func(db.ResourceCache, []atc.MetadataField) error
	updateResourceCacheMetadataMutex       sync.RWMutex
	updateResourceCacheMetadataArgsForCall []struct {
		arg1 db.ResourceCache
		arg2 []atc.MetadataField
	}
	updateResourceCacheMetadataReturns struct {
		result1 error
	}
	updateResourceCacheMetadataReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceCacheFactory) FindOrCreateResourceCache(arg1 db.ResourceCacheUser, arg2 string, arg3 atc.Version, arg4 atc.Source, arg5 atc.Params, arg6 db.ResourceCache) (db.ResourceCache, error) {
	fake.findOrCreateResourceCacheMutex.Lock()
	ret, specificReturn := fake.findOrCreateResourceCacheReturnsOnCall[len(fake.findOrCreateResourceCacheArgsForCall)]
	fake.findOrCreateResourceCacheArgsForCall = append(fake.findOrCreateResourceCacheArgsForCall, struct {
		arg1 db.ResourceCacheUser
		arg2 string
		arg3 atc.Version
		arg4 atc.Source
		arg5 atc.Params
		arg6 db.ResourceCache
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	stub := fake.FindOrCreateResourceCacheStub
	fakeReturns := fake.findOrCreateResourceCacheReturns
	fake.recordInvocation("FindOrCreateResourceCache", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.findOrCreateResourceCacheMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResourceCacheFactory) FindOrCreateResourceCacheCallCount() int {
	fake.findOrCreateResourceCacheMutex.RLock()
	defer fake.findOrCreateResourceCacheMutex.RUnlock()
	return len(fake.findOrCreateResourceCacheArgsForCall)
}

func (fake *FakeResourceCacheFactory) FindOrCreateResourceCacheCalls(stub func(db.ResourceCacheUser, string, atc.Version, atc.Source, atc.Params, db.ResourceCache) (db.ResourceCache, error)) {
	fake.findOrCreateResourceCacheMutex.Lock()
	defer fake.findOrCreateResourceCacheMutex.Unlock()
	fake.FindOrCreateResourceCacheStub = stub
}

func (fake *FakeResourceCacheFactory) FindOrCreateResourceCacheArgsForCall(i int) (db.ResourceCacheUser, string, atc.Version, atc.Source, atc.Params, db.ResourceCache) {
	fake.findOrCreateResourceCacheMutex.RLock()
	defer fake.findOrCreateResourceCacheMutex.RUnlock()
	argsForCall := fake.findOrCreateResourceCacheArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeResourceCacheFactory) FindOrCreateResourceCacheReturns(result1 db.ResourceCache, result2 error) {
	fake.findOrCreateResourceCacheMutex.Lock()
	defer fake.findOrCreateResourceCacheMutex.Unlock()
	fake.FindOrCreateResourceCacheStub = nil
	fake.findOrCreateResourceCacheReturns = struct {
		result1 db.ResourceCache
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceCacheFactory) FindOrCreateResourceCacheReturnsOnCall(i int, result1 db.ResourceCache, result2 error) {
	fake.findOrCreateResourceCacheMutex.Lock()
	defer fake.findOrCreateResourceCacheMutex.Unlock()
	fake.FindOrCreateResourceCacheStub = nil
	if fake.findOrCreateResourceCacheReturnsOnCall == nil {
		fake.findOrCreateResourceCacheReturnsOnCall = make(map[int]struct {
			result1 db.ResourceCache
			result2 error
		})
	}
	fake.findOrCreateResourceCacheReturnsOnCall[i] = struct {
		result1 db.ResourceCache
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceCacheFactory) FindResourceCacheByID(arg1 int) (db.ResourceCache, bool, error) {
	fake.findResourceCacheByIDMutex.Lock()
	ret, specificReturn := fake.findResourceCacheByIDReturnsOnCall[len(fake.findResourceCacheByIDArgsForCall)]
	fake.findResourceCacheByIDArgsForCall = append(fake.findResourceCacheByIDArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.FindResourceCacheByIDStub
	fakeReturns := fake.findResourceCacheByIDReturns
	fake.recordInvocation("FindResourceCacheByID", []interface{}{arg1})
	fake.findResourceCacheByIDMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeResourceCacheFactory) FindResourceCacheByIDCallCount() int {
	fake.findResourceCacheByIDMutex.RLock()
	defer fake.findResourceCacheByIDMutex.RUnlock()
	return len(fake.findResourceCacheByIDArgsForCall)
}

func (fake *FakeResourceCacheFactory) FindResourceCacheByIDCalls(stub func(int) (db.ResourceCache, bool, error)) {
	fake.findResourceCacheByIDMutex.Lock()
	defer fake.findResourceCacheByIDMutex.Unlock()
	fake.FindResourceCacheByIDStub = stub
}

func (fake *FakeResourceCacheFactory) FindResourceCacheByIDArgsForCall(i int) int {
	fake.findResourceCacheByIDMutex.RLock()
	defer fake.findResourceCacheByIDMutex.RUnlock()
	argsForCall := fake.findResourceCacheByIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeResourceCacheFactory) FindResourceCacheByIDReturns(result1 db.ResourceCache, result2 bool, result3 error) {
	fake.findResourceCacheByIDMutex.Lock()
	defer fake.findResourceCacheByIDMutex.Unlock()
	fake.FindResourceCacheByIDStub = nil
	fake.findResourceCacheByIDReturns = struct {
		result1 db.ResourceCache
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceCacheFactory) FindResourceCacheByIDReturnsOnCall(i int, result1 db.ResourceCache, result2 bool, result3 error) {
	fake.findResourceCacheByIDMutex.Lock()
	defer fake.findResourceCacheByIDMutex.Unlock()
	fake.FindResourceCacheByIDStub = nil
	if fake.findResourceCacheByIDReturnsOnCall == nil {
		fake.findResourceCacheByIDReturnsOnCall = make(map[int]struct {
			result1 db.ResourceCache
			result2 bool
			result3 error
		})
	}
	fake.findResourceCacheByIDReturnsOnCall[i] = struct {
		result1 db.ResourceCache
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceCacheFactory) ResourceCacheMetadata(arg1 db.ResourceCache) (db.ResourceConfigMetadataFields, error) {
	fake.resourceCacheMetadataMutex.Lock()
	ret, specificReturn := fake.resourceCacheMetadataReturnsOnCall[len(fake.resourceCacheMetadataArgsForCall)]
	fake.resourceCacheMetadataArgsForCall = append(fake.resourceCacheMetadataArgsForCall, struct {
		arg1 db.ResourceCache
	}{arg1})
	stub := fake.ResourceCacheMetadataStub
	fakeReturns := fake.resourceCacheMetadataReturns
	fake.recordInvocation("ResourceCacheMetadata", []interface{}{arg1})
	fake.resourceCacheMetadataMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResourceCacheFactory) ResourceCacheMetadataCallCount() int {
	fake.resourceCacheMetadataMutex.RLock()
	defer fake.resourceCacheMetadataMutex.RUnlock()
	return len(fake.resourceCacheMetadataArgsForCall)
}

func (fake *FakeResourceCacheFactory) ResourceCacheMetadataCalls(stub func(db.ResourceCache) (db.ResourceConfigMetadataFields, error)) {
	fake.resourceCacheMetadataMutex.Lock()
	defer fake.resourceCacheMetadataMutex.Unlock()
	fake.ResourceCacheMetadataStub = stub
}

func (fake *FakeResourceCacheFactory) ResourceCacheMetadataArgsForCall(i int) db.ResourceCache {
	fake.resourceCacheMetadataMutex.RLock()
	defer fake.resourceCacheMetadataMutex.RUnlock()
	argsForCall := fake.resourceCacheMetadataArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeResourceCacheFactory) ResourceCacheMetadataReturns(result1 db.ResourceConfigMetadataFields, result2 error) {
	fake.resourceCacheMetadataMutex.Lock()
	defer fake.resourceCacheMetadataMutex.Unlock()
	fake.ResourceCacheMetadataStub = nil
	fake.resourceCacheMetadataReturns = struct {
		result1 db.ResourceConfigMetadataFields
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceCacheFactory) ResourceCacheMetadataReturnsOnCall(i int, result1 db.ResourceConfigMetadataFields, result2 error) {
	fake.resourceCacheMetadataMutex.Lock()
	defer fake.resourceCacheMetadataMutex.Unlock()
	fake.ResourceCacheMetadataStub = nil
	if fake.resourceCacheMetadataReturnsOnCall == nil {
		fake.resourceCacheMetadataReturnsOnCall = make(map[int]struct {
			result1 db.ResourceConfigMetadataFields
			result2 error
		})
	}
	fake.resourceCacheMetadataReturnsOnCall[i] = struct {
		result1 db.ResourceConfigMetadataFields
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceCacheFactory) UpdateResourceCacheMetadata(arg1 db.ResourceCache, arg2 []atc.MetadataField) error {
	var arg2Copy []atc.MetadataField
	if arg2 != nil {
		arg2Copy = make([]atc.MetadataField, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.updateResourceCacheMetadataMutex.Lock()
	ret, specificReturn := fake.updateResourceCacheMetadataReturnsOnCall[len(fake.updateResourceCacheMetadataArgsForCall)]
	fake.updateResourceCacheMetadataArgsForCall = append(fake.updateResourceCacheMetadataArgsForCall, struct {
		arg1 db.ResourceCache
		arg2 []atc.MetadataField
	}{arg1, arg2Copy})
	stub := fake.UpdateResourceCacheMetadataStub
	fakeReturns := fake.updateResourceCacheMetadataReturns
	fake.recordInvocation("UpdateResourceCacheMetadata", []interface{}{arg1, arg2Copy})
	fake.updateResourceCacheMetadataMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeResourceCacheFactory) UpdateResourceCacheMetadataCallCount() int {
	fake.updateResourceCacheMetadataMutex.RLock()
	defer fake.updateResourceCacheMetadataMutex.RUnlock()
	return len(fake.updateResourceCacheMetadataArgsForCall)
}

func (fake *FakeResourceCacheFactory) UpdateResourceCacheMetadataCalls(stub func(db.ResourceCache, []atc.MetadataField) error) {
	fake.updateResourceCacheMetadataMutex.Lock()
	defer fake.updateResourceCacheMetadataMutex.Unlock()
	fake.UpdateResourceCacheMetadataStub = stub
}

func (fake *FakeResourceCacheFactory) UpdateResourceCacheMetadataArgsForCall(i int) (db.ResourceCache, []atc.MetadataField) {
	fake.updateResourceCacheMetadataMutex.RLock()
	defer fake.updateResourceCacheMetadataMutex.RUnlock()
	argsForCall := fake.updateResourceCacheMetadataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeResourceCacheFactory) UpdateResourceCacheMetadataReturns(result1 error) {
	fake.updateResourceCacheMetadataMutex.Lock()
	defer fake.updateResourceCacheMetadataMutex.Unlock()
	fake.UpdateResourceCacheMetadataStub = nil
	fake.updateResourceCacheMetadataReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceCacheFactory) UpdateResourceCacheMetadataReturnsOnCall(i int, result1 error) {
	fake.updateResourceCacheMetadataMutex.Lock()
	defer fake.updateResourceCacheMetadataMutex.Unlock()
	fake.UpdateResourceCacheMetadataStub = nil
	if fake.updateResourceCacheMetadataReturnsOnCall == nil {
		fake.updateResourceCacheMetadataReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateResourceCacheMetadataReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceCacheFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findOrCreateResourceCacheMutex.RLock()
	defer fake.findOrCreateResourceCacheMutex.RUnlock()
	fake.findResourceCacheByIDMutex.RLock()
	defer fake.findResourceCacheByIDMutex.RUnlock()
	fake.resourceCacheMetadataMutex.RLock()
	defer fake.resourceCacheMetadataMutex.RUnlock()
	fake.updateResourceCacheMetadataMutex.RLock()
	defer fake.updateResourceCacheMetadataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResourceCacheFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.ResourceCacheFactory = new(FakeResourceCacheFactory)
