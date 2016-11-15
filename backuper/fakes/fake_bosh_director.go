// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotal-cf/pcf-backup-and-restore/backuper"
)

type FakeBoshDirector struct {
	FindInstancesStub        func(deploymentName string) ([]backuper.Instance, error)
	findInstancesMutex       sync.RWMutex
	findInstancesArgsForCall []struct {
		deploymentName string
	}
	findInstancesReturns struct {
		result1 []backuper.Instance
		result2 error
	}
	GetManifestStub        func(deploymentName string) (string, error)
	getManifestMutex       sync.RWMutex
	getManifestArgsForCall []struct {
		deploymentName string
	}
	getManifestReturns struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBoshDirector) FindInstances(deploymentName string) ([]backuper.Instance, error) {
	fake.findInstancesMutex.Lock()
	fake.findInstancesArgsForCall = append(fake.findInstancesArgsForCall, struct {
		deploymentName string
	}{deploymentName})
	fake.recordInvocation("FindInstances", []interface{}{deploymentName})
	fake.findInstancesMutex.Unlock()
	if fake.FindInstancesStub != nil {
		return fake.FindInstancesStub(deploymentName)
	} else {
		return fake.findInstancesReturns.result1, fake.findInstancesReturns.result2
	}
}

func (fake *FakeBoshDirector) FindInstancesCallCount() int {
	fake.findInstancesMutex.RLock()
	defer fake.findInstancesMutex.RUnlock()
	return len(fake.findInstancesArgsForCall)
}

func (fake *FakeBoshDirector) FindInstancesArgsForCall(i int) string {
	fake.findInstancesMutex.RLock()
	defer fake.findInstancesMutex.RUnlock()
	return fake.findInstancesArgsForCall[i].deploymentName
}

func (fake *FakeBoshDirector) FindInstancesReturns(result1 []backuper.Instance, result2 error) {
	fake.FindInstancesStub = nil
	fake.findInstancesReturns = struct {
		result1 []backuper.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshDirector) GetManifest(deploymentName string) (string, error) {
	fake.getManifestMutex.Lock()
	fake.getManifestArgsForCall = append(fake.getManifestArgsForCall, struct {
		deploymentName string
	}{deploymentName})
	fake.recordInvocation("GetManifest", []interface{}{deploymentName})
	fake.getManifestMutex.Unlock()
	if fake.GetManifestStub != nil {
		return fake.GetManifestStub(deploymentName)
	} else {
		return fake.getManifestReturns.result1, fake.getManifestReturns.result2
	}
}

func (fake *FakeBoshDirector) GetManifestCallCount() int {
	fake.getManifestMutex.RLock()
	defer fake.getManifestMutex.RUnlock()
	return len(fake.getManifestArgsForCall)
}

func (fake *FakeBoshDirector) GetManifestArgsForCall(i int) string {
	fake.getManifestMutex.RLock()
	defer fake.getManifestMutex.RUnlock()
	return fake.getManifestArgsForCall[i].deploymentName
}

func (fake *FakeBoshDirector) GetManifestReturns(result1 string, result2 error) {
	fake.GetManifestStub = nil
	fake.getManifestReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshDirector) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findInstancesMutex.RLock()
	defer fake.findInstancesMutex.RUnlock()
	fake.getManifestMutex.RLock()
	defer fake.getManifestMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBoshDirector) recordInvocation(key string, args []interface{}) {
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

var _ backuper.BoshDirector = new(FakeBoshDirector)
