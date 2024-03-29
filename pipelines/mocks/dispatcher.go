// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	pipelines "github.com/egnd/go-toolbox/pipelines"
	mock "github.com/stretchr/testify/mock"
)

// Dispatcher is an autogenerated mock type for the Dispatcher type
type Dispatcher struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Dispatcher) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Push provides a mock function with given fields: _a0
func (_m *Dispatcher) Push(_a0 pipelines.Task) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(pipelines.Task) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Wait provides a mock function with given fields:
func (_m *Dispatcher) Wait() {
	_m.Called()
}

type mockConstructorTestingTNewDispatcher interface {
	mock.TestingT
	Cleanup(func())
}

// NewDispatcher creates a new instance of Dispatcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDispatcher(t mockConstructorTestingTNewDispatcher) *Dispatcher {
	mock := &Dispatcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
