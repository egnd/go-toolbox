// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	pipelines "github.com/egnd/go-toolbox/pipelines"
	mock "github.com/stretchr/testify/mock"
)

// Doer is an autogenerated mock type for the Doer type
type Doer struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Doer) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Do provides a mock function with given fields: _a0
func (_m *Doer) Do(_a0 pipelines.Task) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(pipelines.Task) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDoer interface {
	mock.TestingT
	Cleanup(func())
}

// NewDoer creates a new instance of Doer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDoer(t mockConstructorTestingTNewDoer) *Doer {
	mock := &Doer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
