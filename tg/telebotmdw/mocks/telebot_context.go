// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	telebot "gopkg.in/telebot.v3"

	time "time"
)

// TelebotContext is an autogenerated mock type for the TelebotContext type
type TelebotContext struct {
	mock.Mock
}

// Accept provides a mock function with given fields: errorMessage
func (_m *TelebotContext) Accept(errorMessage ...string) error {
	_va := make([]interface{}, len(errorMessage))
	for _i := range errorMessage {
		_va[_i] = errorMessage[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...string) error); ok {
		r0 = rf(errorMessage...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Answer provides a mock function with given fields: resp
func (_m *TelebotContext) Answer(resp *telebot.QueryResponse) error {
	ret := _m.Called(resp)

	var r0 error
	if rf, ok := ret.Get(0).(func(*telebot.QueryResponse) error); ok {
		r0 = rf(resp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Args provides a mock function with given fields:
func (_m *TelebotContext) Args() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// Bot provides a mock function with given fields:
func (_m *TelebotContext) Bot() *telebot.Bot {
	ret := _m.Called()

	var r0 *telebot.Bot
	if rf, ok := ret.Get(0).(func() *telebot.Bot); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telebot.Bot)
		}
	}

	return r0
}

// Callback provides a mock function with given fields:
func (_m *TelebotContext) Callback() *telebot.Callback {
	ret := _m.Called()

	var r0 *telebot.Callback
	if rf, ok := ret.Get(0).(func() *telebot.Callback); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telebot.Callback)
		}
	}

	return r0
}

// Chat provides a mock function with given fields:
func (_m *TelebotContext) Chat() *telebot.Chat {
	ret := _m.Called()

	var r0 *telebot.Chat
	if rf, ok := ret.Get(0).(func() *telebot.Chat); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telebot.Chat)
		}
	}

	return r0
}

// ChatJoinRequest provides a mock function with given fields:
func (_m *TelebotContext) ChatJoinRequest() *telebot.ChatJoinRequest {
	ret := _m.Called()

	var r0 *telebot.ChatJoinRequest
	if rf, ok := ret.Get(0).(func() *telebot.ChatJoinRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telebot.ChatJoinRequest)
		}
	}

	return r0
}

// ChatMember provides a mock function with given fields:
func (_m *TelebotContext) ChatMember() *telebot.ChatMemberUpdate {
	ret := _m.Called()

	var r0 *telebot.ChatMemberUpdate
	if rf, ok := ret.Get(0).(func() *telebot.ChatMemberUpdate); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telebot.ChatMemberUpdate)
		}
	}

	return r0
}

// Data provides a mock function with given fields:
func (_m *TelebotContext) Data() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Delete provides a mock function with given fields:
func (_m *TelebotContext) Delete() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAfter provides a mock function with given fields: d
func (_m *TelebotContext) DeleteAfter(d time.Duration) *time.Timer {
	ret := _m.Called(d)

	var r0 *time.Timer
	if rf, ok := ret.Get(0).(func(time.Duration) *time.Timer); ok {
		r0 = rf(d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*time.Timer)
		}
	}

	return r0
}

// Edit provides a mock function with given fields: what, opts
func (_m *TelebotContext) Edit(what interface{}, opts ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, what)
	_ca = append(_ca, opts...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) error); ok {
		r0 = rf(what, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditCaption provides a mock function with given fields: caption, opts
func (_m *TelebotContext) EditCaption(caption string, opts ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, caption)
	_ca = append(_ca, opts...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...interface{}) error); ok {
		r0 = rf(caption, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditOrReply provides a mock function with given fields: what, opts
func (_m *TelebotContext) EditOrReply(what interface{}, opts ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, what)
	_ca = append(_ca, opts...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) error); ok {
		r0 = rf(what, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditOrSend provides a mock function with given fields: what, opts
func (_m *TelebotContext) EditOrSend(what interface{}, opts ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, what)
	_ca = append(_ca, opts...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) error); ok {
		r0 = rf(what, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Forward provides a mock function with given fields: msg, opts
func (_m *TelebotContext) Forward(msg telebot.Editable, opts ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, opts...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(telebot.Editable, ...interface{}) error); ok {
		r0 = rf(msg, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ForwardTo provides a mock function with given fields: to, opts
func (_m *TelebotContext) ForwardTo(to telebot.Recipient, opts ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, to)
	_ca = append(_ca, opts...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(telebot.Recipient, ...interface{}) error); ok {
		r0 = rf(to, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: key
func (_m *TelebotContext) Get(key string) interface{} {
	ret := _m.Called(key)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// InlineResult provides a mock function with given fields:
func (_m *TelebotContext) InlineResult() *telebot.InlineResult {
	ret := _m.Called()

	var r0 *telebot.InlineResult
	if rf, ok := ret.Get(0).(func() *telebot.InlineResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telebot.InlineResult)
		}
	}

	return r0
}

// Message provides a mock function with given fields:
func (_m *TelebotContext) Message() *telebot.Message {
	ret := _m.Called()

	var r0 *telebot.Message
	if rf, ok := ret.Get(0).(func() *telebot.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telebot.Message)
		}
	}

	return r0
}

// Migration provides a mock function with given fields:
func (_m *TelebotContext) Migration() (int64, int64) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func() int64); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(int64)
	}

	return r0, r1
}

// Notify provides a mock function with given fields: action
func (_m *TelebotContext) Notify(action telebot.ChatAction) error {
	ret := _m.Called(action)

	var r0 error
	if rf, ok := ret.Get(0).(func(telebot.ChatAction) error); ok {
		r0 = rf(action)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Poll provides a mock function with given fields:
func (_m *TelebotContext) Poll() *telebot.Poll {
	ret := _m.Called()

	var r0 *telebot.Poll
	if rf, ok := ret.Get(0).(func() *telebot.Poll); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telebot.Poll)
		}
	}

	return r0
}

// PollAnswer provides a mock function with given fields:
func (_m *TelebotContext) PollAnswer() *telebot.PollAnswer {
	ret := _m.Called()

	var r0 *telebot.PollAnswer
	if rf, ok := ret.Get(0).(func() *telebot.PollAnswer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telebot.PollAnswer)
		}
	}

	return r0
}

// PreCheckoutQuery provides a mock function with given fields:
func (_m *TelebotContext) PreCheckoutQuery() *telebot.PreCheckoutQuery {
	ret := _m.Called()

	var r0 *telebot.PreCheckoutQuery
	if rf, ok := ret.Get(0).(func() *telebot.PreCheckoutQuery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telebot.PreCheckoutQuery)
		}
	}

	return r0
}

// Query provides a mock function with given fields:
func (_m *TelebotContext) Query() *telebot.Query {
	ret := _m.Called()

	var r0 *telebot.Query
	if rf, ok := ret.Get(0).(func() *telebot.Query); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telebot.Query)
		}
	}

	return r0
}

// Recipient provides a mock function with given fields:
func (_m *TelebotContext) Recipient() telebot.Recipient {
	ret := _m.Called()

	var r0 telebot.Recipient
	if rf, ok := ret.Get(0).(func() telebot.Recipient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(telebot.Recipient)
		}
	}

	return r0
}

// Reply provides a mock function with given fields: what, opts
func (_m *TelebotContext) Reply(what interface{}, opts ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, what)
	_ca = append(_ca, opts...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) error); ok {
		r0 = rf(what, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Respond provides a mock function with given fields: resp
func (_m *TelebotContext) Respond(resp ...*telebot.CallbackResponse) error {
	_va := make([]interface{}, len(resp))
	for _i := range resp {
		_va[_i] = resp[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...*telebot.CallbackResponse) error); ok {
		r0 = rf(resp...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: what, opts
func (_m *TelebotContext) Send(what interface{}, opts ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, what)
	_ca = append(_ca, opts...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) error); ok {
		r0 = rf(what, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendAlbum provides a mock function with given fields: a, opts
func (_m *TelebotContext) SendAlbum(a telebot.Album, opts ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, a)
	_ca = append(_ca, opts...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(telebot.Album, ...interface{}) error); ok {
		r0 = rf(a, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Sender provides a mock function with given fields:
func (_m *TelebotContext) Sender() *telebot.User {
	ret := _m.Called()

	var r0 *telebot.User
	if rf, ok := ret.Get(0).(func() *telebot.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telebot.User)
		}
	}

	return r0
}

// Set provides a mock function with given fields: key, val
func (_m *TelebotContext) Set(key string, val interface{}) {
	_m.Called(key, val)
}

// Ship provides a mock function with given fields: what
func (_m *TelebotContext) Ship(what ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, what...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(what...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ShippingQuery provides a mock function with given fields:
func (_m *TelebotContext) ShippingQuery() *telebot.ShippingQuery {
	ret := _m.Called()

	var r0 *telebot.ShippingQuery
	if rf, ok := ret.Get(0).(func() *telebot.ShippingQuery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telebot.ShippingQuery)
		}
	}

	return r0
}

// Text provides a mock function with given fields:
func (_m *TelebotContext) Text() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Update provides a mock function with given fields:
func (_m *TelebotContext) Update() telebot.Update {
	ret := _m.Called()

	var r0 telebot.Update
	if rf, ok := ret.Get(0).(func() telebot.Update); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(telebot.Update)
	}

	return r0
}

type mockConstructorTestingTNewTelebotContext interface {
	mock.TestingT
	Cleanup(func())
}

// NewTelebotContext creates a new instance of TelebotContext. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTelebotContext(t mockConstructorTestingTNewTelebotContext) *TelebotContext {
	mock := &TelebotContext{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}