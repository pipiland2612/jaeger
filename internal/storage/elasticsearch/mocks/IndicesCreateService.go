// Copyright (c) The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0
//
// Run 'make generate-mocks' to regenerate.

// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	es "github.com/jaegertracing/jaeger/internal/storage/elasticsearch"
	elastic "github.com/olivere/elastic"

	mock "github.com/stretchr/testify/mock"
)

// IndicesCreateService is an autogenerated mock type for the IndicesCreateService type
type IndicesCreateService struct {
	mock.Mock
}

// Body provides a mock function with given fields: mapping
func (_m *IndicesCreateService) Body(mapping string) es.IndicesCreateService {
	ret := _m.Called(mapping)

	if len(ret) == 0 {
		panic("no return value specified for Body")
	}

	var r0 es.IndicesCreateService
	if rf, ok := ret.Get(0).(func(string) es.IndicesCreateService); ok {
		r0 = rf(mapping)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(es.IndicesCreateService)
		}
	}

	return r0
}

// Do provides a mock function with given fields: ctx
func (_m *IndicesCreateService) Do(ctx context.Context) (*elastic.IndicesCreateResult, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Do")
	}

	var r0 *elastic.IndicesCreateResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*elastic.IndicesCreateResult, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *elastic.IndicesCreateResult); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*elastic.IndicesCreateResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIndicesCreateService creates a new instance of IndicesCreateService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIndicesCreateService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IndicesCreateService {
	mock := &IndicesCreateService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
