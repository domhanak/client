// Code generated by counterfeiter. DO NOT EDIT.
package typesfakes

import (
	"sync"

	"github.com/spf13/cobra"
	"k8s.io/client-go/rest"
	"knative.dev/client-pkg/pkg/kn-source-pkg/pkg/types"
)

type FakeRunEFactory struct {
	CreateRunEStub        func() func(cmd *cobra.Command, args []string) error
	createRunEMutex       sync.RWMutex
	createRunEArgsForCall []struct {
	}
	createRunEReturns struct {
		result1 func(cmd *cobra.Command, args []string) error
	}
	createRunEReturnsOnCall map[int]struct {
		result1 func(cmd *cobra.Command, args []string) error
	}
	DeleteRunEStub        func() func(cmd *cobra.Command, args []string) error
	deleteRunEMutex       sync.RWMutex
	deleteRunEArgsForCall []struct {
	}
	deleteRunEReturns struct {
		result1 func(cmd *cobra.Command, args []string) error
	}
	deleteRunEReturnsOnCall map[int]struct {
		result1 func(cmd *cobra.Command, args []string) error
	}
	DescribeRunEStub        func() func(cmd *cobra.Command, args []string) error
	describeRunEMutex       sync.RWMutex
	describeRunEArgsForCall []struct {
	}
	describeRunEReturns struct {
		result1 func(cmd *cobra.Command, args []string) error
	}
	describeRunEReturnsOnCall map[int]struct {
		result1 func(cmd *cobra.Command, args []string) error
	}
	KnSourceClientStub        func(*rest.Config, string) types.KnSourceClient
	knSourceClientMutex       sync.RWMutex
	knSourceClientArgsForCall []struct {
		arg1 *rest.Config
		arg2 string
	}
	knSourceClientReturns struct {
		result1 types.KnSourceClient
	}
	knSourceClientReturnsOnCall map[int]struct {
		result1 types.KnSourceClient
	}
	KnSourceFactoryStub        func() types.KnSourceFactory
	knSourceFactoryMutex       sync.RWMutex
	knSourceFactoryArgsForCall []struct {
	}
	knSourceFactoryReturns struct {
		result1 types.KnSourceFactory
	}
	knSourceFactoryReturnsOnCall map[int]struct {
		result1 types.KnSourceFactory
	}
	ListRunEStub        func() func(cmd *cobra.Command, args []string) error
	listRunEMutex       sync.RWMutex
	listRunEArgsForCall []struct {
	}
	listRunEReturns struct {
		result1 func(cmd *cobra.Command, args []string) error
	}
	listRunEReturnsOnCall map[int]struct {
		result1 func(cmd *cobra.Command, args []string) error
	}
	UpdateRunEStub        func() func(cmd *cobra.Command, args []string) error
	updateRunEMutex       sync.RWMutex
	updateRunEArgsForCall []struct {
	}
	updateRunEReturns struct {
		result1 func(cmd *cobra.Command, args []string) error
	}
	updateRunEReturnsOnCall map[int]struct {
		result1 func(cmd *cobra.Command, args []string) error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRunEFactory) CreateRunE() func(cmd *cobra.Command, args []string) error {
	fake.createRunEMutex.Lock()
	ret, specificReturn := fake.createRunEReturnsOnCall[len(fake.createRunEArgsForCall)]
	fake.createRunEArgsForCall = append(fake.createRunEArgsForCall, struct {
	}{})
	stub := fake.CreateRunEStub
	fakeReturns := fake.createRunEReturns
	fake.recordInvocation("CreateRunE", []interface{}{})
	fake.createRunEMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRunEFactory) CreateRunECallCount() int {
	fake.createRunEMutex.RLock()
	defer fake.createRunEMutex.RUnlock()
	return len(fake.createRunEArgsForCall)
}

func (fake *FakeRunEFactory) CreateRunECalls(stub func() func(cmd *cobra.Command, args []string) error) {
	fake.createRunEMutex.Lock()
	defer fake.createRunEMutex.Unlock()
	fake.CreateRunEStub = stub
}

func (fake *FakeRunEFactory) CreateRunEReturns(result1 func(cmd *cobra.Command, args []string) error) {
	fake.createRunEMutex.Lock()
	defer fake.createRunEMutex.Unlock()
	fake.CreateRunEStub = nil
	fake.createRunEReturns = struct {
		result1 func(cmd *cobra.Command, args []string) error
	}{result1}
}

func (fake *FakeRunEFactory) CreateRunEReturnsOnCall(i int, result1 func(cmd *cobra.Command, args []string) error) {
	fake.createRunEMutex.Lock()
	defer fake.createRunEMutex.Unlock()
	fake.CreateRunEStub = nil
	if fake.createRunEReturnsOnCall == nil {
		fake.createRunEReturnsOnCall = make(map[int]struct {
			result1 func(cmd *cobra.Command, args []string) error
		})
	}
	fake.createRunEReturnsOnCall[i] = struct {
		result1 func(cmd *cobra.Command, args []string) error
	}{result1}
}

func (fake *FakeRunEFactory) DeleteRunE() func(cmd *cobra.Command, args []string) error {
	fake.deleteRunEMutex.Lock()
	ret, specificReturn := fake.deleteRunEReturnsOnCall[len(fake.deleteRunEArgsForCall)]
	fake.deleteRunEArgsForCall = append(fake.deleteRunEArgsForCall, struct {
	}{})
	stub := fake.DeleteRunEStub
	fakeReturns := fake.deleteRunEReturns
	fake.recordInvocation("DeleteRunE", []interface{}{})
	fake.deleteRunEMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRunEFactory) DeleteRunECallCount() int {
	fake.deleteRunEMutex.RLock()
	defer fake.deleteRunEMutex.RUnlock()
	return len(fake.deleteRunEArgsForCall)
}

func (fake *FakeRunEFactory) DeleteRunECalls(stub func() func(cmd *cobra.Command, args []string) error) {
	fake.deleteRunEMutex.Lock()
	defer fake.deleteRunEMutex.Unlock()
	fake.DeleteRunEStub = stub
}

func (fake *FakeRunEFactory) DeleteRunEReturns(result1 func(cmd *cobra.Command, args []string) error) {
	fake.deleteRunEMutex.Lock()
	defer fake.deleteRunEMutex.Unlock()
	fake.DeleteRunEStub = nil
	fake.deleteRunEReturns = struct {
		result1 func(cmd *cobra.Command, args []string) error
	}{result1}
}

func (fake *FakeRunEFactory) DeleteRunEReturnsOnCall(i int, result1 func(cmd *cobra.Command, args []string) error) {
	fake.deleteRunEMutex.Lock()
	defer fake.deleteRunEMutex.Unlock()
	fake.DeleteRunEStub = nil
	if fake.deleteRunEReturnsOnCall == nil {
		fake.deleteRunEReturnsOnCall = make(map[int]struct {
			result1 func(cmd *cobra.Command, args []string) error
		})
	}
	fake.deleteRunEReturnsOnCall[i] = struct {
		result1 func(cmd *cobra.Command, args []string) error
	}{result1}
}

func (fake *FakeRunEFactory) DescribeRunE() func(cmd *cobra.Command, args []string) error {
	fake.describeRunEMutex.Lock()
	ret, specificReturn := fake.describeRunEReturnsOnCall[len(fake.describeRunEArgsForCall)]
	fake.describeRunEArgsForCall = append(fake.describeRunEArgsForCall, struct {
	}{})
	stub := fake.DescribeRunEStub
	fakeReturns := fake.describeRunEReturns
	fake.recordInvocation("DescribeRunE", []interface{}{})
	fake.describeRunEMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRunEFactory) DescribeRunECallCount() int {
	fake.describeRunEMutex.RLock()
	defer fake.describeRunEMutex.RUnlock()
	return len(fake.describeRunEArgsForCall)
}

func (fake *FakeRunEFactory) DescribeRunECalls(stub func() func(cmd *cobra.Command, args []string) error) {
	fake.describeRunEMutex.Lock()
	defer fake.describeRunEMutex.Unlock()
	fake.DescribeRunEStub = stub
}

func (fake *FakeRunEFactory) DescribeRunEReturns(result1 func(cmd *cobra.Command, args []string) error) {
	fake.describeRunEMutex.Lock()
	defer fake.describeRunEMutex.Unlock()
	fake.DescribeRunEStub = nil
	fake.describeRunEReturns = struct {
		result1 func(cmd *cobra.Command, args []string) error
	}{result1}
}

func (fake *FakeRunEFactory) DescribeRunEReturnsOnCall(i int, result1 func(cmd *cobra.Command, args []string) error) {
	fake.describeRunEMutex.Lock()
	defer fake.describeRunEMutex.Unlock()
	fake.DescribeRunEStub = nil
	if fake.describeRunEReturnsOnCall == nil {
		fake.describeRunEReturnsOnCall = make(map[int]struct {
			result1 func(cmd *cobra.Command, args []string) error
		})
	}
	fake.describeRunEReturnsOnCall[i] = struct {
		result1 func(cmd *cobra.Command, args []string) error
	}{result1}
}

func (fake *FakeRunEFactory) KnSourceClient(arg1 *rest.Config, arg2 string) types.KnSourceClient {
	fake.knSourceClientMutex.Lock()
	ret, specificReturn := fake.knSourceClientReturnsOnCall[len(fake.knSourceClientArgsForCall)]
	fake.knSourceClientArgsForCall = append(fake.knSourceClientArgsForCall, struct {
		arg1 *rest.Config
		arg2 string
	}{arg1, arg2})
	stub := fake.KnSourceClientStub
	fakeReturns := fake.knSourceClientReturns
	fake.recordInvocation("KnSourceClient", []interface{}{arg1, arg2})
	fake.knSourceClientMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRunEFactory) KnSourceClientCallCount() int {
	fake.knSourceClientMutex.RLock()
	defer fake.knSourceClientMutex.RUnlock()
	return len(fake.knSourceClientArgsForCall)
}

func (fake *FakeRunEFactory) KnSourceClientCalls(stub func(*rest.Config, string) types.KnSourceClient) {
	fake.knSourceClientMutex.Lock()
	defer fake.knSourceClientMutex.Unlock()
	fake.KnSourceClientStub = stub
}

func (fake *FakeRunEFactory) KnSourceClientArgsForCall(i int) (*rest.Config, string) {
	fake.knSourceClientMutex.RLock()
	defer fake.knSourceClientMutex.RUnlock()
	argsForCall := fake.knSourceClientArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRunEFactory) KnSourceClientReturns(result1 types.KnSourceClient) {
	fake.knSourceClientMutex.Lock()
	defer fake.knSourceClientMutex.Unlock()
	fake.KnSourceClientStub = nil
	fake.knSourceClientReturns = struct {
		result1 types.KnSourceClient
	}{result1}
}

func (fake *FakeRunEFactory) KnSourceClientReturnsOnCall(i int, result1 types.KnSourceClient) {
	fake.knSourceClientMutex.Lock()
	defer fake.knSourceClientMutex.Unlock()
	fake.KnSourceClientStub = nil
	if fake.knSourceClientReturnsOnCall == nil {
		fake.knSourceClientReturnsOnCall = make(map[int]struct {
			result1 types.KnSourceClient
		})
	}
	fake.knSourceClientReturnsOnCall[i] = struct {
		result1 types.KnSourceClient
	}{result1}
}

func (fake *FakeRunEFactory) KnSourceFactory() types.KnSourceFactory {
	fake.knSourceFactoryMutex.Lock()
	ret, specificReturn := fake.knSourceFactoryReturnsOnCall[len(fake.knSourceFactoryArgsForCall)]
	fake.knSourceFactoryArgsForCall = append(fake.knSourceFactoryArgsForCall, struct {
	}{})
	stub := fake.KnSourceFactoryStub
	fakeReturns := fake.knSourceFactoryReturns
	fake.recordInvocation("KnSourceFactory", []interface{}{})
	fake.knSourceFactoryMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRunEFactory) KnSourceFactoryCallCount() int {
	fake.knSourceFactoryMutex.RLock()
	defer fake.knSourceFactoryMutex.RUnlock()
	return len(fake.knSourceFactoryArgsForCall)
}

func (fake *FakeRunEFactory) KnSourceFactoryCalls(stub func() types.KnSourceFactory) {
	fake.knSourceFactoryMutex.Lock()
	defer fake.knSourceFactoryMutex.Unlock()
	fake.KnSourceFactoryStub = stub
}

func (fake *FakeRunEFactory) KnSourceFactoryReturns(result1 types.KnSourceFactory) {
	fake.knSourceFactoryMutex.Lock()
	defer fake.knSourceFactoryMutex.Unlock()
	fake.KnSourceFactoryStub = nil
	fake.knSourceFactoryReturns = struct {
		result1 types.KnSourceFactory
	}{result1}
}

func (fake *FakeRunEFactory) KnSourceFactoryReturnsOnCall(i int, result1 types.KnSourceFactory) {
	fake.knSourceFactoryMutex.Lock()
	defer fake.knSourceFactoryMutex.Unlock()
	fake.KnSourceFactoryStub = nil
	if fake.knSourceFactoryReturnsOnCall == nil {
		fake.knSourceFactoryReturnsOnCall = make(map[int]struct {
			result1 types.KnSourceFactory
		})
	}
	fake.knSourceFactoryReturnsOnCall[i] = struct {
		result1 types.KnSourceFactory
	}{result1}
}

func (fake *FakeRunEFactory) ListRunE() func(cmd *cobra.Command, args []string) error {
	fake.listRunEMutex.Lock()
	ret, specificReturn := fake.listRunEReturnsOnCall[len(fake.listRunEArgsForCall)]
	fake.listRunEArgsForCall = append(fake.listRunEArgsForCall, struct {
	}{})
	stub := fake.ListRunEStub
	fakeReturns := fake.listRunEReturns
	fake.recordInvocation("ListRunE", []interface{}{})
	fake.listRunEMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRunEFactory) ListRunECallCount() int {
	fake.listRunEMutex.RLock()
	defer fake.listRunEMutex.RUnlock()
	return len(fake.listRunEArgsForCall)
}

func (fake *FakeRunEFactory) ListRunECalls(stub func() func(cmd *cobra.Command, args []string) error) {
	fake.listRunEMutex.Lock()
	defer fake.listRunEMutex.Unlock()
	fake.ListRunEStub = stub
}

func (fake *FakeRunEFactory) ListRunEReturns(result1 func(cmd *cobra.Command, args []string) error) {
	fake.listRunEMutex.Lock()
	defer fake.listRunEMutex.Unlock()
	fake.ListRunEStub = nil
	fake.listRunEReturns = struct {
		result1 func(cmd *cobra.Command, args []string) error
	}{result1}
}

func (fake *FakeRunEFactory) ListRunEReturnsOnCall(i int, result1 func(cmd *cobra.Command, args []string) error) {
	fake.listRunEMutex.Lock()
	defer fake.listRunEMutex.Unlock()
	fake.ListRunEStub = nil
	if fake.listRunEReturnsOnCall == nil {
		fake.listRunEReturnsOnCall = make(map[int]struct {
			result1 func(cmd *cobra.Command, args []string) error
		})
	}
	fake.listRunEReturnsOnCall[i] = struct {
		result1 func(cmd *cobra.Command, args []string) error
	}{result1}
}

func (fake *FakeRunEFactory) UpdateRunE() func(cmd *cobra.Command, args []string) error {
	fake.updateRunEMutex.Lock()
	ret, specificReturn := fake.updateRunEReturnsOnCall[len(fake.updateRunEArgsForCall)]
	fake.updateRunEArgsForCall = append(fake.updateRunEArgsForCall, struct {
	}{})
	stub := fake.UpdateRunEStub
	fakeReturns := fake.updateRunEReturns
	fake.recordInvocation("UpdateRunE", []interface{}{})
	fake.updateRunEMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRunEFactory) UpdateRunECallCount() int {
	fake.updateRunEMutex.RLock()
	defer fake.updateRunEMutex.RUnlock()
	return len(fake.updateRunEArgsForCall)
}

func (fake *FakeRunEFactory) UpdateRunECalls(stub func() func(cmd *cobra.Command, args []string) error) {
	fake.updateRunEMutex.Lock()
	defer fake.updateRunEMutex.Unlock()
	fake.UpdateRunEStub = stub
}

func (fake *FakeRunEFactory) UpdateRunEReturns(result1 func(cmd *cobra.Command, args []string) error) {
	fake.updateRunEMutex.Lock()
	defer fake.updateRunEMutex.Unlock()
	fake.UpdateRunEStub = nil
	fake.updateRunEReturns = struct {
		result1 func(cmd *cobra.Command, args []string) error
	}{result1}
}

func (fake *FakeRunEFactory) UpdateRunEReturnsOnCall(i int, result1 func(cmd *cobra.Command, args []string) error) {
	fake.updateRunEMutex.Lock()
	defer fake.updateRunEMutex.Unlock()
	fake.UpdateRunEStub = nil
	if fake.updateRunEReturnsOnCall == nil {
		fake.updateRunEReturnsOnCall = make(map[int]struct {
			result1 func(cmd *cobra.Command, args []string) error
		})
	}
	fake.updateRunEReturnsOnCall[i] = struct {
		result1 func(cmd *cobra.Command, args []string) error
	}{result1}
}

func (fake *FakeRunEFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createRunEMutex.RLock()
	defer fake.createRunEMutex.RUnlock()
	fake.deleteRunEMutex.RLock()
	defer fake.deleteRunEMutex.RUnlock()
	fake.describeRunEMutex.RLock()
	defer fake.describeRunEMutex.RUnlock()
	fake.knSourceClientMutex.RLock()
	defer fake.knSourceClientMutex.RUnlock()
	fake.knSourceFactoryMutex.RLock()
	defer fake.knSourceFactoryMutex.RUnlock()
	fake.listRunEMutex.RLock()
	defer fake.listRunEMutex.RUnlock()
	fake.updateRunEMutex.RLock()
	defer fake.updateRunEMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRunEFactory) recordInvocation(key string, args []interface{}) {
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

var _ types.RunEFactory = new(FakeRunEFactory)
