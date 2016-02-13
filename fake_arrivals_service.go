// This file was generated by counterfeiter
package main

import (
	"sync"
)

type FakeArrivalsService struct {
	GetStub        func(stopId string) ([]Arrival, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		stopId string
	}
	getReturns struct {
		result1 []Arrival
		result2 error
	}
}

func (fake *FakeArrivalsService) Get(stopId string) ([]Arrival, error) {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		stopId string
	}{stopId})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(stopId)
	} else {
		return fake.getReturns.result1, fake.getReturns.result2
	}
}

func (fake *FakeArrivalsService) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeArrivalsService) GetArgsForCall(i int) string {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].stopId
}

func (fake *FakeArrivalsService) GetReturns(result1 []Arrival, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 []Arrival
		result2 error
	}{result1, result2}
}