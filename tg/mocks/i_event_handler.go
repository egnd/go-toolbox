// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	mock "github.com/stretchr/testify/mock"

	tgchain "github.com/egnd/go-toolbox/tg/tgchain"
)

// IEventHandler is an autogenerated mock type for the IEventHandler type
type IEventHandler struct {
	mock.Mock
}

// Decorate provides a mock function with given fields: _a0
func (_m *IEventHandler) Decorate(_a0 tgchain.IEventHandler) {
	_m.Called(_a0)
}

// GetDescription provides a mock function with given fields:
func (_m *IEventHandler) GetDescription() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetName provides a mock function with given fields:
func (_m *IEventHandler) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Handle provides a mock function with given fields: _a0, _a1
func (_m *IEventHandler) Handle(_a0 context.Context, _a1 tgbotapi.Update) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, tgbotapi.Update) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Next provides a mock function with given fields: _a0, _a1
func (_m *IEventHandler) Next(_a0 context.Context, _a1 tgbotapi.Update) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, tgbotapi.Update) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIEventHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewIEventHandler creates a new instance of IEventHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIEventHandler(t mockConstructorTestingTNewIEventHandler) *IEventHandler {
	mock := &IEventHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
