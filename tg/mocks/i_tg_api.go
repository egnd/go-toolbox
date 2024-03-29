// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	mock "github.com/stretchr/testify/mock"
)

// ITgAPI is an autogenerated mock type for the ITgAPI type
type ITgAPI struct {
	mock.Mock
}

// AnswerInlineQuery provides a mock function with given fields: _a0
func (_m *ITgAPI) AnswerInlineQuery(_a0 tgbotapi.InlineConfig) (tgbotapi.APIResponse, error) {
	ret := _m.Called(_a0)

	var r0 tgbotapi.APIResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(tgbotapi.InlineConfig) (tgbotapi.APIResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(tgbotapi.InlineConfig) tgbotapi.APIResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(tgbotapi.APIResponse)
	}

	if rf, ok := ret.Get(1).(func(tgbotapi.InlineConfig) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChatMember provides a mock function with given fields: _a0
func (_m *ITgAPI) GetChatMember(_a0 tgbotapi.ChatConfigWithUser) (tgbotapi.ChatMember, error) {
	ret := _m.Called(_a0)

	var r0 tgbotapi.ChatMember
	var r1 error
	if rf, ok := ret.Get(0).(func(tgbotapi.ChatConfigWithUser) (tgbotapi.ChatMember, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(tgbotapi.ChatConfigWithUser) tgbotapi.ChatMember); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(tgbotapi.ChatMember)
	}

	if rf, ok := ret.Get(1).(func(tgbotapi.ChatConfigWithUser) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMe provides a mock function with given fields:
func (_m *ITgAPI) GetMe() (tgbotapi.User, error) {
	ret := _m.Called()

	var r0 tgbotapi.User
	var r1 error
	if rf, ok := ret.Get(0).(func() (tgbotapi.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() tgbotapi.User); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(tgbotapi.User)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUpdatesChan provides a mock function with given fields: _a0
func (_m *ITgAPI) GetUpdatesChan(_a0 tgbotapi.UpdateConfig) (tgbotapi.UpdatesChannel, error) {
	ret := _m.Called(_a0)

	var r0 tgbotapi.UpdatesChannel
	var r1 error
	if rf, ok := ret.Get(0).(func(tgbotapi.UpdateConfig) (tgbotapi.UpdatesChannel, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tgbotapi.UpdatesChannel)
		}
	}

	if rf, ok := ret.Get(1).(func(tgbotapi.UpdateConfig) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LeaveChat provides a mock function with given fields: _a0
func (_m *ITgAPI) LeaveChat(_a0 tgbotapi.ChatConfig) (tgbotapi.APIResponse, error) {
	ret := _m.Called(_a0)

	var r0 tgbotapi.APIResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(tgbotapi.ChatConfig) (tgbotapi.APIResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(tgbotapi.ChatConfig) tgbotapi.APIResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(tgbotapi.APIResponse)
	}

	if rf, ok := ret.Get(1).(func(tgbotapi.ChatConfig) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Send provides a mock function with given fields: _a0
func (_m *ITgAPI) Send(_a0 tgbotapi.Chattable) (tgbotapi.Message, error) {
	ret := _m.Called(_a0)

	var r0 tgbotapi.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(tgbotapi.Chattable) (tgbotapi.Message, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(tgbotapi.Chattable) tgbotapi.Message); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(tgbotapi.Message)
	}

	if rf, ok := ret.Get(1).(func(tgbotapi.Chattable) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewITgAPI interface {
	mock.TestingT
	Cleanup(func())
}

// NewITgAPI creates a new instance of ITgAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewITgAPI(t mockConstructorTestingTNewITgAPI) *ITgAPI {
	mock := &ITgAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
