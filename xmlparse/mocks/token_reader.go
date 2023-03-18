// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	xml "encoding/xml"

	mock "github.com/stretchr/testify/mock"
)

// TokenReader is an autogenerated mock type for the TokenReader type
type TokenReader struct {
	mock.Mock
}

// Token provides a mock function with given fields:
func (_m *TokenReader) Token() (xml.Token, error) {
	ret := _m.Called()

	var r0 xml.Token
	var r1 error
	if rf, ok := ret.Get(0).(func() (xml.Token, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() xml.Token); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(xml.Token)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTokenReader interface {
	mock.TestingT
	Cleanup(func())
}

// NewTokenReader creates a new instance of TokenReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTokenReader(t mockConstructorTestingTNewTokenReader) *TokenReader {
	mock := &TokenReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
