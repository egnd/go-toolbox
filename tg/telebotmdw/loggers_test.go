package telebotmdw_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/tg/telebotmdw"
	"github.com/egnd/go-toolbox/tg/telebotmdw/mocks"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/telebot.v3"
)

func Test_CtxGetLogger(t *testing.T) {
	ctx := &mocks.Context{}
	defer ctx.AssertExpectations(t)

	ctx.On("Get", "logger").Return(nil, false)
	_, ok := telebotmdw.CtxGetLogger(ctx)
	assert.False(t, ok)
}

func Test_CtxWithLogger(t *testing.T) {
	logger := zerolog.Nop()

	ctx := &mocks.Context{}
	defer ctx.AssertExpectations(t)

	ctx.On("Set", "logger", logger)

	assert.NotEmpty(t, telebotmdw.CtxWithLogger(ctx, logger))
}

func Test_LogUpdate(t *testing.T) {
	cases := []struct {
		details bool
		err     error
	}{
		{
			details: true,
		},
		{
			err: errors.New("some error"),
		},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			ctx := &mocks.Context{}
			defer ctx.AssertExpectations(t)

			if test.err == nil {
				ctx.On("Get", "logger").Return(zerolog.Nop(), true)
			}

			if test.details {
				ctx.On("Update").Return(telebot.Update{})
			}

			if test.err == nil {
				assert.NoError(t, telebotmdw.LogUpdate(test.details)(func(ctx telebot.Context) error { return test.err })(ctx))
			} else {
				assert.EqualError(t, telebotmdw.LogUpdate(test.details)(func(ctx telebot.Context) error { return test.err })(ctx), test.err.Error())
			}
		})
	}
}

func Test_AddCtxLogger(t *testing.T) {
	cases := []struct {
		upd  telebot.Update
		user *telebot.User
		data string
		text string
		args []string
	}{
		{
			upd:  telebot.Update{Message: &telebot.Message{}},
			user: &telebot.User{Username: "login", LastName: "lname", FirstName: "fname"},
			data: "111",
			text: "222",
			args: []string{"333", "444"},
		},
		{
			upd:  telebot.Update{EditedMessage: &telebot.Message{}},
			user: &telebot.User{Username: "login"},
		},
		{
			upd:  telebot.Update{Callback: &telebot.Callback{Unique: "asdfg"}},
			user: &telebot.User{Username: "login"},
		},
		{
			upd:  telebot.Update{Query: &telebot.Query{}},
			user: &telebot.User{Username: "login"},
		},
		{
			upd:  telebot.Update{InlineResult: &telebot.InlineResult{}},
			user: &telebot.User{Username: "login"},
		},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			logger := zerolog.Nop()
			ctx := &mocks.Context{}
			defer ctx.AssertExpectations(t)

			ctx.On("Update").Return(test.upd)
			ctx.On("Sender").Return(test.user)
			ctx.On("Args").Return(test.args)
			ctx.On("Data").Return(test.data)
			ctx.On("Text").Return(test.text)
			ctx.On("Set", "logger", mock.AnythingOfType(fmt.Sprintf("%T", logger))).Return(ctx)

			assert.NoError(t, telebotmdw.AddCtxLogger(&logger)(func(_ telebot.Context) error { return nil })(ctx))
		})
	}
}
