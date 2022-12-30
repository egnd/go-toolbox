package logging_test

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/egnd/go-toolbox/logging"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func Test_ZerologCfgViper(t *testing.T) {
	cases := []struct {
		cfg    *viper.Viper
		res    logging.ZerologCfg
		panics string
	}{
		{
			cfg: func() *viper.Viper {
				cfg := viper.New()
				cfg.SetConfigType("yaml")
				cfg.ReadConfig(bytes.NewBuffer([]byte(`
logs:
  level: info
  time_format: UNIXMS
  duration_unit: 1s
  caller: true
  pretty: true
  colors: true
  fields_names:
    level: l
    msg: m
    time: t`)))
				return cfg.Sub("logs")
			}(),
			res: logging.ZerologCfg{
				Level:            zerolog.InfoLevel,
				LevelFieldName:   "l",
				MessageFieldName: "m",
				TimeFieldName:    "t",
				TimeFormat:       zerolog.TimeFormatUnixMs,
				DurationUnit:     time.Second,
				Caller:           true,
				Pretty:           true,
				Colors:           true,
			},
		},
		{
			cfg: func() *viper.Viper {
				cfg := viper.New()
				cfg.SetConfigType("yaml")
				cfg.ReadConfig(bytes.NewBuffer([]byte("level: info222")))
				return cfg
			}(),
			panics: "Unknown Level String: 'info222', defaulting to NoLevel",
		},
	}

	for k, test := range cases {
		k, test := k, test
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			if test.panics != "" {
				assert.PanicsWithError(t, test.panics, func() { logging.NewZerologCfgViper(test.cfg) })
				return
			}

			assert.EqualValues(t, test.res, logging.NewZerologCfgViper(test.cfg))
		})
	}
}

func Test_Zerolog(t *testing.T) {
	cases := []struct {
		cfg   logging.ZerologCfg
		log   func(*zerolog.Logger)
		entry *regexp.Regexp
	}{
		{
			cfg: logging.ZerologCfg{
				Level:            zerolog.InfoLevel,
				LevelFieldName:   "l",
				MessageFieldName: "m",
				TimeFieldName:    "t",
				TimeFormat:       zerolog.TimeFormatUnixMs,
				DurationUnit:     time.Second,
				Caller:           true,
			},
			log:   func(logger *zerolog.Logger) { logger.Info().Dur("dur", time.Millisecond*1500).Msg("info msg") },
			entry: regexp.MustCompile(`{"l":"info","dur":1.5,"t":[0-9]+,"caller":".+/go-toolbox/logging/zerolog_test.go:[0-9]+","m":"info msg"}`),
		},
		{
			cfg: logging.ZerologCfg{
				Level:  zerolog.InfoLevel,
				Pretty: true,
			},
			log:   func(logger *zerolog.Logger) { logger.Info().Msg("info msg") },
			entry: regexp.MustCompile(`[0-9]+:[0-9]+[AP]M INF info msg`),
		},
	}

	for k, test := range cases {
		k, test := k, test
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			var buffer bytes.Buffer
			logger := logging.NewZerolog(test.cfg, &buffer)
			test.log(&logger)
			if test.entry != nil {
				assert.EqualValuesf(t, true, test.entry.Match(buffer.Bytes()), "%s != %s", test.entry.String(), buffer.String())
			}
		})
	}
}
