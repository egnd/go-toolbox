package telebotmdw_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/egnd/go-toolbox/tg/telebotmdw"
	"github.com/egnd/go-toolbox/tg/telebotmdw/mocks"
	"github.com/stretchr/testify/assert"
	"gopkg.in/telebot.v3"
)

func Test_UsersWhiteList(t *testing.T) {
	cases := []struct {
		users  []string
		sender string
		err    error
	}{
		{
			users:  []string{"aaa", "bbb"},
			sender: "ccc",
			err:    errors.New("UsersWhiteList error: user ccc not found"),
		},
		{
			users:  []string{"aaa", "bbb"},
			sender: "bbb",
		},
	}

	for k, test := range cases {
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			ctx := &mocks.TelebotContext{}
			defer ctx.AssertExpectations(t)

			ctx.On("Sender").Return(&telebot.User{Username: test.sender})

			for _, item := range test.users {
				if item == test.sender {
					ctx.On("Callback").Return(&telebot.Callback{})
					ctx.On("Respond").Return(nil)
				}
			}

			if test.err == nil {
				assert.NoError(t, telebotmdw.UsersWhiteList(test.users)(func(ctx telebot.Context) error { return nil })(ctx))
			} else {
				assert.EqualError(t, telebotmdw.UsersWhiteList(test.users)(func(ctx telebot.Context) error { return nil })(ctx), test.err.Error())
			}
		})
	}
}
