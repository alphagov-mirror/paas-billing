// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/alphagov/paas-billing/cfstore"
	"github.com/cloudfoundry-community/go-cfclient"
)

type FakeCFDataClient struct {
	ListServicePlansStub        func() ([]cfstore.ServicePlan, error)
	listServicePlansMutex       sync.RWMutex
	listServicePlansArgsForCall []struct{}
	listServicePlansReturns     struct {
		result1 []cfstore.ServicePlan
		result2 error
	}
	listServicePlansReturnsOnCall map[int]struct {
		result1 []cfstore.ServicePlan
		result2 error
	}
	ListServicesStub        func() ([]cfstore.Service, error)
	listServicesMutex       sync.RWMutex
	listServicesArgsForCall []struct{}
	listServicesReturns     struct {
		result1 []cfstore.Service
		result2 error
	}
	listServicesReturnsOnCall map[int]struct {
		result1 []cfstore.Service
		result2 error
	}
	ListOrgsStub        func() ([]cfclient.Org, error)
	listOrgsMutex       sync.RWMutex
	listOrgsArgsForCall []struct{}
	listOrgsReturns     struct {
		result1 []cfclient.Org
		result2 error
	}
	listOrgsReturnsOnCall map[int]struct {
		result1 []cfclient.Org
		result2 error
	}
	ListSpacesStub        func() ([]cfclient.Space, error)
	listSpacesMutex       sync.RWMutex
	listSpacesArgsForCall []struct{}
	listSpacesReturns     struct {
		result1 []cfclient.Space
		result2 error
	}
	listSpacesReturnsOnCall map[int]struct {
		result1 []cfclient.Space
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCFDataClient) ListServicePlans() ([]cfstore.ServicePlan, error) {
	fake.listServicePlansMutex.Lock()
	ret, specificReturn := fake.listServicePlansReturnsOnCall[len(fake.listServicePlansArgsForCall)]
	fake.listServicePlansArgsForCall = append(fake.listServicePlansArgsForCall, struct{}{})
	fake.recordInvocation("ListServicePlans", []interface{}{})
	fake.listServicePlansMutex.Unlock()
	if fake.ListServicePlansStub != nil {
		return fake.ListServicePlansStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listServicePlansReturns.result1, fake.listServicePlansReturns.result2
}

func (fake *FakeCFDataClient) ListServicePlansCallCount() int {
	fake.listServicePlansMutex.RLock()
	defer fake.listServicePlansMutex.RUnlock()
	return len(fake.listServicePlansArgsForCall)
}

func (fake *FakeCFDataClient) ListServicePlansReturns(result1 []cfstore.ServicePlan, result2 error) {
	fake.ListServicePlansStub = nil
	fake.listServicePlansReturns = struct {
		result1 []cfstore.ServicePlan
		result2 error
	}{result1, result2}
}

func (fake *FakeCFDataClient) ListServicePlansReturnsOnCall(i int, result1 []cfstore.ServicePlan, result2 error) {
	fake.ListServicePlansStub = nil
	if fake.listServicePlansReturnsOnCall == nil {
		fake.listServicePlansReturnsOnCall = make(map[int]struct {
			result1 []cfstore.ServicePlan
			result2 error
		})
	}
	fake.listServicePlansReturnsOnCall[i] = struct {
		result1 []cfstore.ServicePlan
		result2 error
	}{result1, result2}
}

func (fake *FakeCFDataClient) ListServices() ([]cfstore.Service, error) {
	fake.listServicesMutex.Lock()
	ret, specificReturn := fake.listServicesReturnsOnCall[len(fake.listServicesArgsForCall)]
	fake.listServicesArgsForCall = append(fake.listServicesArgsForCall, struct{}{})
	fake.recordInvocation("ListServices", []interface{}{})
	fake.listServicesMutex.Unlock()
	if fake.ListServicesStub != nil {
		return fake.ListServicesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listServicesReturns.result1, fake.listServicesReturns.result2
}

func (fake *FakeCFDataClient) ListServicesCallCount() int {
	fake.listServicesMutex.RLock()
	defer fake.listServicesMutex.RUnlock()
	return len(fake.listServicesArgsForCall)
}

func (fake *FakeCFDataClient) ListServicesReturns(result1 []cfstore.Service, result2 error) {
	fake.ListServicesStub = nil
	fake.listServicesReturns = struct {
		result1 []cfstore.Service
		result2 error
	}{result1, result2}
}

func (fake *FakeCFDataClient) ListServicesReturnsOnCall(i int, result1 []cfstore.Service, result2 error) {
	fake.ListServicesStub = nil
	if fake.listServicesReturnsOnCall == nil {
		fake.listServicesReturnsOnCall = make(map[int]struct {
			result1 []cfstore.Service
			result2 error
		})
	}
	fake.listServicesReturnsOnCall[i] = struct {
		result1 []cfstore.Service
		result2 error
	}{result1, result2}
}

func (fake *FakeCFDataClient) ListOrgs() ([]cfclient.Org, error) {
	fake.listOrgsMutex.Lock()
	ret, specificReturn := fake.listOrgsReturnsOnCall[len(fake.listOrgsArgsForCall)]
	fake.listOrgsArgsForCall = append(fake.listOrgsArgsForCall, struct{}{})
	fake.recordInvocation("ListOrgs", []interface{}{})
	fake.listOrgsMutex.Unlock()
	if fake.ListOrgsStub != nil {
		return fake.ListOrgsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listOrgsReturns.result1, fake.listOrgsReturns.result2
}

func (fake *FakeCFDataClient) ListOrgsCallCount() int {
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	return len(fake.listOrgsArgsForCall)
}

func (fake *FakeCFDataClient) ListOrgsReturns(result1 []cfclient.Org, result2 error) {
	fake.ListOrgsStub = nil
	fake.listOrgsReturns = struct {
		result1 []cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeCFDataClient) ListOrgsReturnsOnCall(i int, result1 []cfclient.Org, result2 error) {
	fake.ListOrgsStub = nil
	if fake.listOrgsReturnsOnCall == nil {
		fake.listOrgsReturnsOnCall = make(map[int]struct {
			result1 []cfclient.Org
			result2 error
		})
	}
	fake.listOrgsReturnsOnCall[i] = struct {
		result1 []cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeCFDataClient) ListSpaces() ([]cfclient.Space, error) {
	fake.listSpacesMutex.Lock()
	ret, specificReturn := fake.listSpacesReturnsOnCall[len(fake.listSpacesArgsForCall)]
	fake.listSpacesArgsForCall = append(fake.listSpacesArgsForCall, struct{}{})
	fake.recordInvocation("ListSpaces", []interface{}{})
	fake.listSpacesMutex.Unlock()
	if fake.ListSpacesStub != nil {
		return fake.ListSpacesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listSpacesReturns.result1, fake.listSpacesReturns.result2
}

func (fake *FakeCFDataClient) ListSpacesCallCount() int {
	fake.listSpacesMutex.RLock()
	defer fake.listSpacesMutex.RUnlock()
	return len(fake.listSpacesArgsForCall)
}

func (fake *FakeCFDataClient) ListSpacesReturns(result1 []cfclient.Space, result2 error) {
	fake.ListSpacesStub = nil
	fake.listSpacesReturns = struct {
		result1 []cfclient.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeCFDataClient) ListSpacesReturnsOnCall(i int, result1 []cfclient.Space, result2 error) {
	fake.ListSpacesStub = nil
	if fake.listSpacesReturnsOnCall == nil {
		fake.listSpacesReturnsOnCall = make(map[int]struct {
			result1 []cfclient.Space
			result2 error
		})
	}
	fake.listSpacesReturnsOnCall[i] = struct {
		result1 []cfclient.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeCFDataClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listServicePlansMutex.RLock()
	defer fake.listServicePlansMutex.RUnlock()
	fake.listServicesMutex.RLock()
	defer fake.listServicesMutex.RUnlock()
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	fake.listSpacesMutex.RLock()
	defer fake.listSpacesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCFDataClient) recordInvocation(key string, args []interface{}) {
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

var _ cfstore.CFDataClient = new(FakeCFDataClient)
