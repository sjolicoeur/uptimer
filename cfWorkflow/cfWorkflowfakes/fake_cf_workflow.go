// Code generated by counterfeiter. DO NOT EDIT.
package cfWorkflowfakes

import (
	"context"
	"sync"

	"github.com/cloudfoundry/uptimer/cfCmdGenerator"
	"github.com/cloudfoundry/uptimer/cfWorkflow"
	"github.com/cloudfoundry/uptimer/cmdStartWaiter"
)

type FakeCfWorkflow struct {
	AppUrlStub        func() string
	appUrlMutex       sync.RWMutex
	appUrlArgsForCall []struct{}
	appUrlReturns     struct {
		result1 string
	}
	appUrlReturnsOnCall map[int]struct {
		result1 string
	}
	SetupStub        func(cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter
	setupMutex       sync.RWMutex
	setupArgsForCall []struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}
	setupReturns struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	setupReturnsOnCall map[int]struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	PushStub        func(cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter
	pushMutex       sync.RWMutex
	pushArgsForCall []struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}
	pushReturns struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	pushReturnsOnCall map[int]struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	DeleteStub        func(cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}
	deleteReturns struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	deleteReturnsOnCall map[int]struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	TearDownStub        func(cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter
	tearDownMutex       sync.RWMutex
	tearDownArgsForCall []struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}
	tearDownReturns struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	tearDownReturnsOnCall map[int]struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	RecentLogsStub        func(cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter
	recentLogsMutex       sync.RWMutex
	recentLogsArgsForCall []struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}
	recentLogsReturns struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	recentLogsReturnsOnCall map[int]struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	StreamLogsStub        func(context.Context, cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter
	streamLogsMutex       sync.RWMutex
	streamLogsArgsForCall []struct {
		arg1 context.Context
		arg2 cfCmdGenerator.CfCmdGenerator
	}
	streamLogsReturns struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	streamLogsReturnsOnCall map[int]struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	MapRouteStub        func(cfCmdGenerator cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter
	mapRouteMutex       sync.RWMutex
	mapRouteArgsForCall []struct {
		cfCmdGenerator cfCmdGenerator.CfCmdGenerator
	}
	mapRouteReturns struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	mapRouteReturnsOnCall map[int]struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	CreateAndBindSyslogDrainServiceStub        func(cfCmdGenerator cfCmdGenerator.CfCmdGenerator, serviceName string) []cmdStartWaiter.CmdStartWaiter
	createAndBindSyslogDrainServiceMutex       sync.RWMutex
	createAndBindSyslogDrainServiceArgsForCall []struct {
		cfCmdGenerator cfCmdGenerator.CfCmdGenerator
		serviceName    string
	}
	createAndBindSyslogDrainServiceReturns struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	createAndBindSyslogDrainServiceReturnsOnCall map[int]struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCfWorkflow) AppUrl() string {
	fake.appUrlMutex.Lock()
	ret, specificReturn := fake.appUrlReturnsOnCall[len(fake.appUrlArgsForCall)]
	fake.appUrlArgsForCall = append(fake.appUrlArgsForCall, struct{}{})
	fake.recordInvocation("AppUrl", []interface{}{})
	fake.appUrlMutex.Unlock()
	if fake.AppUrlStub != nil {
		return fake.AppUrlStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.appUrlReturns.result1
}

func (fake *FakeCfWorkflow) AppUrlCallCount() int {
	fake.appUrlMutex.RLock()
	defer fake.appUrlMutex.RUnlock()
	return len(fake.appUrlArgsForCall)
}

func (fake *FakeCfWorkflow) AppUrlReturns(result1 string) {
	fake.AppUrlStub = nil
	fake.appUrlReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCfWorkflow) AppUrlReturnsOnCall(i int, result1 string) {
	fake.AppUrlStub = nil
	if fake.appUrlReturnsOnCall == nil {
		fake.appUrlReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.appUrlReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCfWorkflow) Setup(arg1 cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter {
	fake.setupMutex.Lock()
	ret, specificReturn := fake.setupReturnsOnCall[len(fake.setupArgsForCall)]
	fake.setupArgsForCall = append(fake.setupArgsForCall, struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}{arg1})
	fake.recordInvocation("Setup", []interface{}{arg1})
	fake.setupMutex.Unlock()
	if fake.SetupStub != nil {
		return fake.SetupStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setupReturns.result1
}

func (fake *FakeCfWorkflow) SetupCallCount() int {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return len(fake.setupArgsForCall)
}

func (fake *FakeCfWorkflow) SetupArgsForCall(i int) cfCmdGenerator.CfCmdGenerator {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return fake.setupArgsForCall[i].arg1
}

func (fake *FakeCfWorkflow) SetupReturns(result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.SetupStub = nil
	fake.setupReturns = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) SetupReturnsOnCall(i int, result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.SetupStub = nil
	if fake.setupReturnsOnCall == nil {
		fake.setupReturnsOnCall = make(map[int]struct {
			result1 []cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.setupReturnsOnCall[i] = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) Push(arg1 cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter {
	fake.pushMutex.Lock()
	ret, specificReturn := fake.pushReturnsOnCall[len(fake.pushArgsForCall)]
	fake.pushArgsForCall = append(fake.pushArgsForCall, struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}{arg1})
	fake.recordInvocation("Push", []interface{}{arg1})
	fake.pushMutex.Unlock()
	if fake.PushStub != nil {
		return fake.PushStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pushReturns.result1
}

func (fake *FakeCfWorkflow) PushCallCount() int {
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	return len(fake.pushArgsForCall)
}

func (fake *FakeCfWorkflow) PushArgsForCall(i int) cfCmdGenerator.CfCmdGenerator {
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	return fake.pushArgsForCall[i].arg1
}

func (fake *FakeCfWorkflow) PushReturns(result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.PushStub = nil
	fake.pushReturns = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) PushReturnsOnCall(i int, result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.PushStub = nil
	if fake.pushReturnsOnCall == nil {
		fake.pushReturnsOnCall = make(map[int]struct {
			result1 []cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.pushReturnsOnCall[i] = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) Delete(arg1 cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}{arg1})
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteReturns.result1
}

func (fake *FakeCfWorkflow) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeCfWorkflow) DeleteArgsForCall(i int) cfCmdGenerator.CfCmdGenerator {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].arg1
}

func (fake *FakeCfWorkflow) DeleteReturns(result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) DeleteReturnsOnCall(i int, result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 []cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) TearDown(arg1 cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter {
	fake.tearDownMutex.Lock()
	ret, specificReturn := fake.tearDownReturnsOnCall[len(fake.tearDownArgsForCall)]
	fake.tearDownArgsForCall = append(fake.tearDownArgsForCall, struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}{arg1})
	fake.recordInvocation("TearDown", []interface{}{arg1})
	fake.tearDownMutex.Unlock()
	if fake.TearDownStub != nil {
		return fake.TearDownStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.tearDownReturns.result1
}

func (fake *FakeCfWorkflow) TearDownCallCount() int {
	fake.tearDownMutex.RLock()
	defer fake.tearDownMutex.RUnlock()
	return len(fake.tearDownArgsForCall)
}

func (fake *FakeCfWorkflow) TearDownArgsForCall(i int) cfCmdGenerator.CfCmdGenerator {
	fake.tearDownMutex.RLock()
	defer fake.tearDownMutex.RUnlock()
	return fake.tearDownArgsForCall[i].arg1
}

func (fake *FakeCfWorkflow) TearDownReturns(result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.TearDownStub = nil
	fake.tearDownReturns = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) TearDownReturnsOnCall(i int, result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.TearDownStub = nil
	if fake.tearDownReturnsOnCall == nil {
		fake.tearDownReturnsOnCall = make(map[int]struct {
			result1 []cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.tearDownReturnsOnCall[i] = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) RecentLogs(arg1 cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter {
	fake.recentLogsMutex.Lock()
	ret, specificReturn := fake.recentLogsReturnsOnCall[len(fake.recentLogsArgsForCall)]
	fake.recentLogsArgsForCall = append(fake.recentLogsArgsForCall, struct {
		arg1 cfCmdGenerator.CfCmdGenerator
	}{arg1})
	fake.recordInvocation("RecentLogs", []interface{}{arg1})
	fake.recentLogsMutex.Unlock()
	if fake.RecentLogsStub != nil {
		return fake.RecentLogsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.recentLogsReturns.result1
}

func (fake *FakeCfWorkflow) RecentLogsCallCount() int {
	fake.recentLogsMutex.RLock()
	defer fake.recentLogsMutex.RUnlock()
	return len(fake.recentLogsArgsForCall)
}

func (fake *FakeCfWorkflow) RecentLogsArgsForCall(i int) cfCmdGenerator.CfCmdGenerator {
	fake.recentLogsMutex.RLock()
	defer fake.recentLogsMutex.RUnlock()
	return fake.recentLogsArgsForCall[i].arg1
}

func (fake *FakeCfWorkflow) RecentLogsReturns(result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.RecentLogsStub = nil
	fake.recentLogsReturns = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) RecentLogsReturnsOnCall(i int, result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.RecentLogsStub = nil
	if fake.recentLogsReturnsOnCall == nil {
		fake.recentLogsReturnsOnCall = make(map[int]struct {
			result1 []cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.recentLogsReturnsOnCall[i] = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) StreamLogs(arg1 context.Context, arg2 cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter {
	fake.streamLogsMutex.Lock()
	ret, specificReturn := fake.streamLogsReturnsOnCall[len(fake.streamLogsArgsForCall)]
	fake.streamLogsArgsForCall = append(fake.streamLogsArgsForCall, struct {
		arg1 context.Context
		arg2 cfCmdGenerator.CfCmdGenerator
	}{arg1, arg2})
	fake.recordInvocation("StreamLogs", []interface{}{arg1, arg2})
	fake.streamLogsMutex.Unlock()
	if fake.StreamLogsStub != nil {
		return fake.StreamLogsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.streamLogsReturns.result1
}

func (fake *FakeCfWorkflow) StreamLogsCallCount() int {
	fake.streamLogsMutex.RLock()
	defer fake.streamLogsMutex.RUnlock()
	return len(fake.streamLogsArgsForCall)
}

func (fake *FakeCfWorkflow) StreamLogsArgsForCall(i int) (context.Context, cfCmdGenerator.CfCmdGenerator) {
	fake.streamLogsMutex.RLock()
	defer fake.streamLogsMutex.RUnlock()
	return fake.streamLogsArgsForCall[i].arg1, fake.streamLogsArgsForCall[i].arg2
}

func (fake *FakeCfWorkflow) StreamLogsReturns(result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.StreamLogsStub = nil
	fake.streamLogsReturns = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) StreamLogsReturnsOnCall(i int, result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.StreamLogsStub = nil
	if fake.streamLogsReturnsOnCall == nil {
		fake.streamLogsReturnsOnCall = make(map[int]struct {
			result1 []cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.streamLogsReturnsOnCall[i] = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) MapRoute(cfCmdGenerator cfCmdGenerator.CfCmdGenerator) []cmdStartWaiter.CmdStartWaiter {
	fake.mapRouteMutex.Lock()
	ret, specificReturn := fake.mapRouteReturnsOnCall[len(fake.mapRouteArgsForCall)]
	fake.mapRouteArgsForCall = append(fake.mapRouteArgsForCall, struct {
		cfCmdGenerator cfCmdGenerator.CfCmdGenerator
	}{cfCmdGenerator})
	fake.recordInvocation("MapRoute", []interface{}{cfCmdGenerator})
	fake.mapRouteMutex.Unlock()
	if fake.MapRouteStub != nil {
		return fake.MapRouteStub(cfCmdGenerator)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.mapRouteReturns.result1
}

func (fake *FakeCfWorkflow) MapRouteCallCount() int {
	fake.mapRouteMutex.RLock()
	defer fake.mapRouteMutex.RUnlock()
	return len(fake.mapRouteArgsForCall)
}

func (fake *FakeCfWorkflow) MapRouteArgsForCall(i int) cfCmdGenerator.CfCmdGenerator {
	fake.mapRouteMutex.RLock()
	defer fake.mapRouteMutex.RUnlock()
	return fake.mapRouteArgsForCall[i].cfCmdGenerator
}

func (fake *FakeCfWorkflow) MapRouteReturns(result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.MapRouteStub = nil
	fake.mapRouteReturns = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) MapRouteReturnsOnCall(i int, result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.MapRouteStub = nil
	if fake.mapRouteReturnsOnCall == nil {
		fake.mapRouteReturnsOnCall = make(map[int]struct {
			result1 []cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.mapRouteReturnsOnCall[i] = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) CreateAndBindSyslogDrainService(cfCmdGenerator cfCmdGenerator.CfCmdGenerator, serviceName string) []cmdStartWaiter.CmdStartWaiter {
	fake.createAndBindSyslogDrainServiceMutex.Lock()
	ret, specificReturn := fake.createAndBindSyslogDrainServiceReturnsOnCall[len(fake.createAndBindSyslogDrainServiceArgsForCall)]
	fake.createAndBindSyslogDrainServiceArgsForCall = append(fake.createAndBindSyslogDrainServiceArgsForCall, struct {
		cfCmdGenerator cfCmdGenerator.CfCmdGenerator
		serviceName    string
	}{cfCmdGenerator, serviceName})
	fake.recordInvocation("CreateAndBindSyslogDrainService", []interface{}{cfCmdGenerator, serviceName})
	fake.createAndBindSyslogDrainServiceMutex.Unlock()
	if fake.CreateAndBindSyslogDrainServiceStub != nil {
		return fake.CreateAndBindSyslogDrainServiceStub(cfCmdGenerator, serviceName)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createAndBindSyslogDrainServiceReturns.result1
}

func (fake *FakeCfWorkflow) CreateAndBindSyslogDrainServiceCallCount() int {
	fake.createAndBindSyslogDrainServiceMutex.RLock()
	defer fake.createAndBindSyslogDrainServiceMutex.RUnlock()
	return len(fake.createAndBindSyslogDrainServiceArgsForCall)
}

func (fake *FakeCfWorkflow) CreateAndBindSyslogDrainServiceArgsForCall(i int) (cfCmdGenerator.CfCmdGenerator, string) {
	fake.createAndBindSyslogDrainServiceMutex.RLock()
	defer fake.createAndBindSyslogDrainServiceMutex.RUnlock()
	return fake.createAndBindSyslogDrainServiceArgsForCall[i].cfCmdGenerator, fake.createAndBindSyslogDrainServiceArgsForCall[i].serviceName
}

func (fake *FakeCfWorkflow) CreateAndBindSyslogDrainServiceReturns(result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.CreateAndBindSyslogDrainServiceStub = nil
	fake.createAndBindSyslogDrainServiceReturns = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) CreateAndBindSyslogDrainServiceReturnsOnCall(i int, result1 []cmdStartWaiter.CmdStartWaiter) {
	fake.CreateAndBindSyslogDrainServiceStub = nil
	if fake.createAndBindSyslogDrainServiceReturnsOnCall == nil {
		fake.createAndBindSyslogDrainServiceReturnsOnCall = make(map[int]struct {
			result1 []cmdStartWaiter.CmdStartWaiter
		})
	}
	fake.createAndBindSyslogDrainServiceReturnsOnCall[i] = struct {
		result1 []cmdStartWaiter.CmdStartWaiter
	}{result1}
}

func (fake *FakeCfWorkflow) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.appUrlMutex.RLock()
	defer fake.appUrlMutex.RUnlock()
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.tearDownMutex.RLock()
	defer fake.tearDownMutex.RUnlock()
	fake.recentLogsMutex.RLock()
	defer fake.recentLogsMutex.RUnlock()
	fake.streamLogsMutex.RLock()
	defer fake.streamLogsMutex.RUnlock()
	fake.mapRouteMutex.RLock()
	defer fake.mapRouteMutex.RUnlock()
	fake.createAndBindSyslogDrainServiceMutex.RLock()
	defer fake.createAndBindSyslogDrainServiceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCfWorkflow) recordInvocation(key string, args []interface{}) {
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

var _ cfWorkflow.CfWorkflow = new(FakeCfWorkflow)
