// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"database/sql"
	"sync"

	"github.com/confluentinc/go-health/v2/checkers"
)

type FakeSQLExecer struct {
	ExecContextStub        func(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	execContextMutex       sync.RWMutex
	execContextArgsForCall []struct {
		ctx   context.Context
		query string
		args  []interface{}
	}
	execContextReturns struct {
		result1 sql.Result
		result2 error
	}
	execContextReturnsOnCall map[int]struct {
		result1 sql.Result
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSQLExecer) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	fake.execContextMutex.Lock()
	ret, specificReturn := fake.execContextReturnsOnCall[len(fake.execContextArgsForCall)]
	fake.execContextArgsForCall = append(fake.execContextArgsForCall, struct {
		ctx   context.Context
		query string
		args  []interface{}
	}{ctx, query, args})
	fake.recordInvocation("ExecContext", []interface{}{ctx, query, args})
	fake.execContextMutex.Unlock()
	if fake.ExecContextStub != nil {
		return fake.ExecContextStub(ctx, query, args...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.execContextReturns.result1, fake.execContextReturns.result2
}

func (fake *FakeSQLExecer) ExecContextCallCount() int {
	fake.execContextMutex.RLock()
	defer fake.execContextMutex.RUnlock()
	return len(fake.execContextArgsForCall)
}

func (fake *FakeSQLExecer) ExecContextArgsForCall(i int) (context.Context, string, []interface{}) {
	fake.execContextMutex.RLock()
	defer fake.execContextMutex.RUnlock()
	return fake.execContextArgsForCall[i].ctx, fake.execContextArgsForCall[i].query, fake.execContextArgsForCall[i].args
}

func (fake *FakeSQLExecer) ExecContextReturns(result1 sql.Result, result2 error) {
	fake.ExecContextStub = nil
	fake.execContextReturns = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeSQLExecer) ExecContextReturnsOnCall(i int, result1 sql.Result, result2 error) {
	fake.ExecContextStub = nil
	if fake.execContextReturnsOnCall == nil {
		fake.execContextReturnsOnCall = make(map[int]struct {
			result1 sql.Result
			result2 error
		})
	}
	fake.execContextReturnsOnCall[i] = struct {
		result1 sql.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeSQLExecer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.execContextMutex.RLock()
	defer fake.execContextMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSQLExecer) recordInvocation(key string, args []interface{}) {
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

var _ checkers.SQLExecer = new(FakeSQLExecer)
