// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/confluentinc/go-health/v2/checkers"
)

type FakeSQLPinger struct {
	PingContextStub        func(ctx context.Context) error
	pingContextMutex       sync.RWMutex
	pingContextArgsForCall []struct {
		ctx context.Context
	}
	pingContextReturns struct {
		result1 error
	}
	pingContextReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSQLPinger) PingContext(ctx context.Context) error {
	fake.pingContextMutex.Lock()
	ret, specificReturn := fake.pingContextReturnsOnCall[len(fake.pingContextArgsForCall)]
	fake.pingContextArgsForCall = append(fake.pingContextArgsForCall, struct {
		ctx context.Context
	}{ctx})
	fake.recordInvocation("PingContext", []interface{}{ctx})
	fake.pingContextMutex.Unlock()
	if fake.PingContextStub != nil {
		return fake.PingContextStub(ctx)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pingContextReturns.result1
}

func (fake *FakeSQLPinger) PingContextCallCount() int {
	fake.pingContextMutex.RLock()
	defer fake.pingContextMutex.RUnlock()
	return len(fake.pingContextArgsForCall)
}

func (fake *FakeSQLPinger) PingContextArgsForCall(i int) context.Context {
	fake.pingContextMutex.RLock()
	defer fake.pingContextMutex.RUnlock()
	return fake.pingContextArgsForCall[i].ctx
}

func (fake *FakeSQLPinger) PingContextReturns(result1 error) {
	fake.PingContextStub = nil
	fake.pingContextReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSQLPinger) PingContextReturnsOnCall(i int, result1 error) {
	fake.PingContextStub = nil
	if fake.pingContextReturnsOnCall == nil {
		fake.pingContextReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pingContextReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSQLPinger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pingContextMutex.RLock()
	defer fake.pingContextMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSQLPinger) recordInvocation(key string, args []interface{}) {
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

var _ checkers.SQLPinger = new(FakeSQLPinger)
