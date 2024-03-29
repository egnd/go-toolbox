// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	xmlparse "github.com/egnd/go-toolbox/xmlparse"
	mock "github.com/stretchr/testify/mock"
)

// Rule is an autogenerated mock type for the Rule type
type Rule struct {
	mock.Mock
}

// Execute provides a mock function with given fields: next
func (_m *Rule) Execute(next xmlparse.TokenHandler) xmlparse.TokenHandler {
	ret := _m.Called(next)

	var r0 xmlparse.TokenHandler
	if rf, ok := ret.Get(0).(func(xmlparse.TokenHandler) xmlparse.TokenHandler); ok {
		r0 = rf(next)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(xmlparse.TokenHandler)
		}
	}

	return r0
}

type mockConstructorTestingTNewRule interface {
	mock.TestingT
	Cleanup(func())
}

// NewRule creates a new instance of Rule. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRule(t mockConstructorTestingTNewRule) *Rule {
	mock := &Rule{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
