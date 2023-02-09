package tg_test

import (
	"testing"
	"time"

	"github.com/egnd/go-toolbox/tg"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gopkg.in/telebot.v3"
)

func Test_Telebot(t *testing.T) {
	cfg := viper.New()

	cfg.Set("timeout", "10s")
	cfg.Set("events", []string{"aaa", "bbb"})
	cfg.Set("limit", 101)
	cfg.Set("token", "sometoken")
	cfg.Set("verbose", true)
	cfg.Set("offline", true)
	cfg.Set("synchronous", true)
	cfg.Set("updates_cap", 100)
	cfg.Set("url", "someurl")
	cfg.Set("parse_mode", "HTML")

	botCfg := tg.NewTeleBotCfg(cfg)

	assert.Empty(t, botCfg.Client)
	assert.Empty(t, botCfg.OnError)

	assert.EqualValues(t, 10*time.Second, botCfg.Poller.(*telebot.LongPoller).Timeout)
	assert.EqualValues(t, []string{"aaa", "bbb"}, botCfg.Poller.(*telebot.LongPoller).AllowedUpdates)
	assert.EqualValues(t, 101, botCfg.Poller.(*telebot.LongPoller).Limit)

	assert.EqualValues(t, "sometoken", botCfg.Token)
	assert.EqualValues(t, true, botCfg.Verbose)
	assert.EqualValues(t, true, botCfg.Offline)
	assert.EqualValues(t, true, botCfg.Synchronous)
	assert.EqualValues(t, 100, botCfg.Updates)
	assert.EqualValues(t, "someurl", botCfg.URL)
	assert.EqualValues(t, telebot.ModeHTML, botCfg.ParseMode)

	assert.NotPanics(t, func() {
		assert.NotEmpty(t, tg.NewTeleBot(botCfg))
	})
}

func Test_TelebotPanic(t *testing.T) {
	assert.Panics(t, func() {
		tg.NewTeleBot(tg.NewTeleBotCfg(viper.New()))
	})
}
