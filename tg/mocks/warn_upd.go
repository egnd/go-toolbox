// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	mock "github.com/stretchr/testify/mock"
)

// WarnUpd is an autogenerated mock type for the WarnUpd type
type WarnUpd struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *WarnUpd) Execute(_a0 string, _a1 tgbotapi.Update) {
	_m.Called(_a0, _a1)
}

type mockConstructorTestingTNewWarnUpd interface {
	mock.TestingT
	Cleanup(func())
}

// NewWarnUpd creates a new instance of WarnUpd. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWarnUpd(t mockConstructorTestingTNewWarnUpd) *WarnUpd {
	mock := &WarnUpd{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
