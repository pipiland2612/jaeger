// Copyright (c) The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0
//
// Run 'make generate-mocks' to regenerate.

// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/jaegertracing/jaeger-idl/model/v1"
	mock "github.com/stretchr/testify/mock"
)

// Writer is an autogenerated mock type for the Writer type
type Writer struct {
	mock.Mock
}

// WriteSpan provides a mock function with given fields: ctx, span
func (_m *Writer) WriteSpan(ctx context.Context, span *model.Span) error {
	ret := _m.Called(ctx, span)

	if len(ret) == 0 {
		panic("no return value specified for WriteSpan")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Span) error); ok {
		r0 = rf(ctx, span)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewWriter creates a new instance of Writer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWriter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Writer {
	mock := &Writer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
