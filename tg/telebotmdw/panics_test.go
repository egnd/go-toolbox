package telebotmdw_test

import (
	"errors"
	"testing"

	"github.com/egnd/go-toolbox/tg/telebotmdw"
	"github.com/egnd/go-toolbox/tg/telebotmdw/mocks"
	"github.com/stretchr/testify/assert"
	"gopkg.in/telebot.v3"
)

func Test_CatchPanics_Success(t *testing.T) {

	assert.NoError(t, telebotmdw.CatchPanics(false)(func(_ telebot.Context) error { return nil })(&mocks.TelebotContext{}))
}

func Test_CatchPanics_Error(t *testing.T) {
	assert.EqualError(t, telebotmdw.CatchPanics(false)(func(_ telebot.Context) error { return errors.New("some error") })(&mocks.TelebotContext{}), "some error")
}

func Test_CatchPanics_Catch_Panic_Err(t *testing.T) {
	assert.EqualError(t, telebotmdw.CatchPanics(true)(func(_ telebot.Context) error { panic(errors.New("some panic")) })(&mocks.TelebotContext{}), "panic: some panic")
}

func Test_CatchPanics_Catch_Panic_Str(t *testing.T) {
	assert.EqualError(t, telebotmdw.CatchPanics(true)(func(_ telebot.Context) error { panic("some panic") })(&mocks.TelebotContext{}), "panic: some panic")
}

func Test_CatchPanics_Catch_Panic_Val(t *testing.T) {
	assert.EqualError(t, telebotmdw.CatchPanics(true)(func(_ telebot.Context) error { panic(123) })(&mocks.TelebotContext{}), "panic: 123")
}

func Test_CatchPanics_NoCatch_Panic(t *testing.T) {
	assert.PanicsWithValue(t, "some panic", func() {
		telebotmdw.CatchPanics(false)(func(_ telebot.Context) error { panic("some panic") })(&mocks.TelebotContext{})
	})
}
