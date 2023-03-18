// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Counter is an autogenerated mock type for the Counter type
type Counter struct {
	mock.Mock
}

// Set provides a mock function with given fields: _a0
func (_m *Counter) Set(_a0 float64) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewCounter interface {
	mock.TestingT
	Cleanup(func())
}

// NewCounter creates a new instance of Counter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCounter(t mockConstructorTestingTNewCounter) *Counter {
	mock := &Counter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
