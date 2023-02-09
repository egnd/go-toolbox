// Package tg contains telegram bot tools.
package tg

import (
	"github.com/spf13/viper"
	"gopkg.in/telebot.v3"
)

// NewTeleBotCfg returns telebot config.
func NewTeleBotCfg(cfg *viper.Viper) telebot.Settings {
	return telebot.Settings{ //nolint:exhaustruct
		Poller: &telebot.LongPoller{ //nolint:exhaustruct
			Timeout:        cfg.GetDuration("timeout"),
			AllowedUpdates: cfg.GetStringSlice("events"),
			Limit:          cfg.GetInt("limit"),
		},
		Token:       cfg.GetString("token"),
		Verbose:     cfg.GetBool("verbose"),
		Offline:     cfg.GetBool("offline"),
		Synchronous: cfg.GetBool("synchronous"),
		Updates:     cfg.GetInt("updates_cap"),
		URL:         cfg.GetString("url"),
		ParseMode:   cfg.GetString("parse_mode"),
	}
}

// NewTeleBot return telegram bot instance.
func NewTeleBot(cfg telebot.Settings) *telebot.Bot {
	res, err := telebot.NewBot(cfg)
	if err != nil {
		panic(err)
	}

	return res
}
