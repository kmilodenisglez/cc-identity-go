// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
)

type ContractInterface struct {
	GetAfterTransactionStub        func() interface{}
	getAfterTransactionMutex       sync.RWMutex
	getAfterTransactionArgsForCall []struct {
	}
	getAfterTransactionReturns struct {
		result1 interface{}
	}
	getAfterTransactionReturnsOnCall map[int]struct {
		result1 interface{}
	}
	GetBeforeTransactionStub        func() interface{}
	getBeforeTransactionMutex       sync.RWMutex
	getBeforeTransactionArgsForCall []struct {
	}
	getBeforeTransactionReturns struct {
		result1 interface{}
	}
	getBeforeTransactionReturnsOnCall map[int]struct {
		result1 interface{}
	}
	GetInfoStub        func() metadata.InfoMetadata
	getInfoMutex       sync.RWMutex
	getInfoArgsForCall []struct {
	}
	getInfoReturns struct {
		result1 metadata.InfoMetadata
	}
	getInfoReturnsOnCall map[int]struct {
		result1 metadata.InfoMetadata
	}
	GetNameStub        func() string
	getNameMutex       sync.RWMutex
	getNameArgsForCall []struct {
	}
	getNameReturns struct {
		result1 string
	}
	getNameReturnsOnCall map[int]struct {
		result1 string
	}
	GetTransactionContextHandlerStub        func() contractapi.SettableTransactionContextInterface
	getTransactionContextHandlerMutex       sync.RWMutex
	getTransactionContextHandlerArgsForCall []struct {
	}
	getTransactionContextHandlerReturns struct {
		result1 contractapi.SettableTransactionContextInterface
	}
	getTransactionContextHandlerReturnsOnCall map[int]struct {
		result1 contractapi.SettableTransactionContextInterface
	}
	GetUnknownTransactionStub        func() interface{}
	getUnknownTransactionMutex       sync.RWMutex
	getUnknownTransactionArgsForCall []struct {
	}
	getUnknownTransactionReturns struct {
		result1 interface{}
	}
	getUnknownTransactionReturnsOnCall map[int]struct {
		result1 interface{}
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ContractInterface) GetAfterTransaction() interface{} {
	fake.getAfterTransactionMutex.Lock()
	ret, specificReturn := fake.getAfterTransactionReturnsOnCall[len(fake.getAfterTransactionArgsForCall)]
	fake.getAfterTransactionArgsForCall = append(fake.getAfterTransactionArgsForCall, struct {
	}{})
	stub := fake.GetAfterTransactionStub
	fakeReturns := fake.getAfterTransactionReturns
	fake.recordInvocation("GetAfterTransaction", []interface{}{})
	fake.getAfterTransactionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ContractInterface) GetAfterTransactionCallCount() int {
	fake.getAfterTransactionMutex.RLock()
	defer fake.getAfterTransactionMutex.RUnlock()
	return len(fake.getAfterTransactionArgsForCall)
}

func (fake *ContractInterface) GetAfterTransactionCalls(stub func() interface{}) {
	fake.getAfterTransactionMutex.Lock()
	defer fake.getAfterTransactionMutex.Unlock()
	fake.GetAfterTransactionStub = stub
}

func (fake *ContractInterface) GetAfterTransactionReturns(result1 interface{}) {
	fake.getAfterTransactionMutex.Lock()
	defer fake.getAfterTransactionMutex.Unlock()
	fake.GetAfterTransactionStub = nil
	fake.getAfterTransactionReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *ContractInterface) GetAfterTransactionReturnsOnCall(i int, result1 interface{}) {
	fake.getAfterTransactionMutex.Lock()
	defer fake.getAfterTransactionMutex.Unlock()
	fake.GetAfterTransactionStub = nil
	if fake.getAfterTransactionReturnsOnCall == nil {
		fake.getAfterTransactionReturnsOnCall = make(map[int]struct {
			result1 interface{}
		})
	}
	fake.getAfterTransactionReturnsOnCall[i] = struct {
		result1 interface{}
	}{result1}
}

func (fake *ContractInterface) GetBeforeTransaction() interface{} {
	fake.getBeforeTransactionMutex.Lock()
	ret, specificReturn := fake.getBeforeTransactionReturnsOnCall[len(fake.getBeforeTransactionArgsForCall)]
	fake.getBeforeTransactionArgsForCall = append(fake.getBeforeTransactionArgsForCall, struct {
	}{})
	stub := fake.GetBeforeTransactionStub
	fakeReturns := fake.getBeforeTransactionReturns
	fake.recordInvocation("GetBeforeTransaction", []interface{}{})
	fake.getBeforeTransactionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ContractInterface) GetBeforeTransactionCallCount() int {
	fake.getBeforeTransactionMutex.RLock()
	defer fake.getBeforeTransactionMutex.RUnlock()
	return len(fake.getBeforeTransactionArgsForCall)
}

func (fake *ContractInterface) GetBeforeTransactionCalls(stub func() interface{}) {
	fake.getBeforeTransactionMutex.Lock()
	defer fake.getBeforeTransactionMutex.Unlock()
	fake.GetBeforeTransactionStub = stub
}

func (fake *ContractInterface) GetBeforeTransactionReturns(result1 interface{}) {
	fake.getBeforeTransactionMutex.Lock()
	defer fake.getBeforeTransactionMutex.Unlock()
	fake.GetBeforeTransactionStub = nil
	fake.getBeforeTransactionReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *ContractInterface) GetBeforeTransactionReturnsOnCall(i int, result1 interface{}) {
	fake.getBeforeTransactionMutex.Lock()
	defer fake.getBeforeTransactionMutex.Unlock()
	fake.GetBeforeTransactionStub = nil
	if fake.getBeforeTransactionReturnsOnCall == nil {
		fake.getBeforeTransactionReturnsOnCall = make(map[int]struct {
			result1 interface{}
		})
	}
	fake.getBeforeTransactionReturnsOnCall[i] = struct {
		result1 interface{}
	}{result1}
}

func (fake *ContractInterface) GetInfo() metadata.InfoMetadata {
	fake.getInfoMutex.Lock()
	ret, specificReturn := fake.getInfoReturnsOnCall[len(fake.getInfoArgsForCall)]
	fake.getInfoArgsForCall = append(fake.getInfoArgsForCall, struct {
	}{})
	stub := fake.GetInfoStub
	fakeReturns := fake.getInfoReturns
	fake.recordInvocation("GetInfo", []interface{}{})
	fake.getInfoMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ContractInterface) GetInfoCallCount() int {
	fake.getInfoMutex.RLock()
	defer fake.getInfoMutex.RUnlock()
	return len(fake.getInfoArgsForCall)
}

func (fake *ContractInterface) GetInfoCalls(stub func() metadata.InfoMetadata) {
	fake.getInfoMutex.Lock()
	defer fake.getInfoMutex.Unlock()
	fake.GetInfoStub = stub
}

func (fake *ContractInterface) GetInfoReturns(result1 metadata.InfoMetadata) {
	fake.getInfoMutex.Lock()
	defer fake.getInfoMutex.Unlock()
	fake.GetInfoStub = nil
	fake.getInfoReturns = struct {
		result1 metadata.InfoMetadata
	}{result1}
}

func (fake *ContractInterface) GetInfoReturnsOnCall(i int, result1 metadata.InfoMetadata) {
	fake.getInfoMutex.Lock()
	defer fake.getInfoMutex.Unlock()
	fake.GetInfoStub = nil
	if fake.getInfoReturnsOnCall == nil {
		fake.getInfoReturnsOnCall = make(map[int]struct {
			result1 metadata.InfoMetadata
		})
	}
	fake.getInfoReturnsOnCall[i] = struct {
		result1 metadata.InfoMetadata
	}{result1}
}

func (fake *ContractInterface) GetName() string {
	fake.getNameMutex.Lock()
	ret, specificReturn := fake.getNameReturnsOnCall[len(fake.getNameArgsForCall)]
	fake.getNameArgsForCall = append(fake.getNameArgsForCall, struct {
	}{})
	stub := fake.GetNameStub
	fakeReturns := fake.getNameReturns
	fake.recordInvocation("GetName", []interface{}{})
	fake.getNameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ContractInterface) GetNameCallCount() int {
	fake.getNameMutex.RLock()
	defer fake.getNameMutex.RUnlock()
	return len(fake.getNameArgsForCall)
}

func (fake *ContractInterface) GetNameCalls(stub func() string) {
	fake.getNameMutex.Lock()
	defer fake.getNameMutex.Unlock()
	fake.GetNameStub = stub
}

func (fake *ContractInterface) GetNameReturns(result1 string) {
	fake.getNameMutex.Lock()
	defer fake.getNameMutex.Unlock()
	fake.GetNameStub = nil
	fake.getNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *ContractInterface) GetNameReturnsOnCall(i int, result1 string) {
	fake.getNameMutex.Lock()
	defer fake.getNameMutex.Unlock()
	fake.GetNameStub = nil
	if fake.getNameReturnsOnCall == nil {
		fake.getNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *ContractInterface) GetTransactionContextHandler() contractapi.SettableTransactionContextInterface {
	fake.getTransactionContextHandlerMutex.Lock()
	ret, specificReturn := fake.getTransactionContextHandlerReturnsOnCall[len(fake.getTransactionContextHandlerArgsForCall)]
	fake.getTransactionContextHandlerArgsForCall = append(fake.getTransactionContextHandlerArgsForCall, struct {
	}{})
	stub := fake.GetTransactionContextHandlerStub
	fakeReturns := fake.getTransactionContextHandlerReturns
	fake.recordInvocation("GetTransactionContextHandler", []interface{}{})
	fake.getTransactionContextHandlerMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ContractInterface) GetTransactionContextHandlerCallCount() int {
	fake.getTransactionContextHandlerMutex.RLock()
	defer fake.getTransactionContextHandlerMutex.RUnlock()
	return len(fake.getTransactionContextHandlerArgsForCall)
}

func (fake *ContractInterface) GetTransactionContextHandlerCalls(stub func() contractapi.SettableTransactionContextInterface) {
	fake.getTransactionContextHandlerMutex.Lock()
	defer fake.getTransactionContextHandlerMutex.Unlock()
	fake.GetTransactionContextHandlerStub = stub
}

func (fake *ContractInterface) GetTransactionContextHandlerReturns(result1 contractapi.SettableTransactionContextInterface) {
	fake.getTransactionContextHandlerMutex.Lock()
	defer fake.getTransactionContextHandlerMutex.Unlock()
	fake.GetTransactionContextHandlerStub = nil
	fake.getTransactionContextHandlerReturns = struct {
		result1 contractapi.SettableTransactionContextInterface
	}{result1}
}

func (fake *ContractInterface) GetTransactionContextHandlerReturnsOnCall(i int, result1 contractapi.SettableTransactionContextInterface) {
	fake.getTransactionContextHandlerMutex.Lock()
	defer fake.getTransactionContextHandlerMutex.Unlock()
	fake.GetTransactionContextHandlerStub = nil
	if fake.getTransactionContextHandlerReturnsOnCall == nil {
		fake.getTransactionContextHandlerReturnsOnCall = make(map[int]struct {
			result1 contractapi.SettableTransactionContextInterface
		})
	}
	fake.getTransactionContextHandlerReturnsOnCall[i] = struct {
		result1 contractapi.SettableTransactionContextInterface
	}{result1}
}

func (fake *ContractInterface) GetUnknownTransaction() interface{} {
	fake.getUnknownTransactionMutex.Lock()
	ret, specificReturn := fake.getUnknownTransactionReturnsOnCall[len(fake.getUnknownTransactionArgsForCall)]
	fake.getUnknownTransactionArgsForCall = append(fake.getUnknownTransactionArgsForCall, struct {
	}{})
	stub := fake.GetUnknownTransactionStub
	fakeReturns := fake.getUnknownTransactionReturns
	fake.recordInvocation("GetUnknownTransaction", []interface{}{})
	fake.getUnknownTransactionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ContractInterface) GetUnknownTransactionCallCount() int {
	fake.getUnknownTransactionMutex.RLock()
	defer fake.getUnknownTransactionMutex.RUnlock()
	return len(fake.getUnknownTransactionArgsForCall)
}

func (fake *ContractInterface) GetUnknownTransactionCalls(stub func() interface{}) {
	fake.getUnknownTransactionMutex.Lock()
	defer fake.getUnknownTransactionMutex.Unlock()
	fake.GetUnknownTransactionStub = stub
}

func (fake *ContractInterface) GetUnknownTransactionReturns(result1 interface{}) {
	fake.getUnknownTransactionMutex.Lock()
	defer fake.getUnknownTransactionMutex.Unlock()
	fake.GetUnknownTransactionStub = nil
	fake.getUnknownTransactionReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *ContractInterface) GetUnknownTransactionReturnsOnCall(i int, result1 interface{}) {
	fake.getUnknownTransactionMutex.Lock()
	defer fake.getUnknownTransactionMutex.Unlock()
	fake.GetUnknownTransactionStub = nil
	if fake.getUnknownTransactionReturnsOnCall == nil {
		fake.getUnknownTransactionReturnsOnCall = make(map[int]struct {
			result1 interface{}
		})
	}
	fake.getUnknownTransactionReturnsOnCall[i] = struct {
		result1 interface{}
	}{result1}
}

func (fake *ContractInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAfterTransactionMutex.RLock()
	defer fake.getAfterTransactionMutex.RUnlock()
	fake.getBeforeTransactionMutex.RLock()
	defer fake.getBeforeTransactionMutex.RUnlock()
	fake.getInfoMutex.RLock()
	defer fake.getInfoMutex.RUnlock()
	fake.getNameMutex.RLock()
	defer fake.getNameMutex.RUnlock()
	fake.getTransactionContextHandlerMutex.RLock()
	defer fake.getTransactionContextHandlerMutex.RUnlock()
	fake.getUnknownTransactionMutex.RLock()
	defer fake.getUnknownTransactionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ContractInterface) recordInvocation(key string, args []interface{}) {
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