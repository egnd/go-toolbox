// Package logging contains tools for working with logging.
package logging

import (
	"io"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

// ZerologCfg is a params for zerolog logger instance creation.
type ZerologCfg struct {
	Colors           bool
	Pretty           bool
	Caller           bool
	LevelFieldName   string
	MessageFieldName string
	TimeFieldName    string
	TimeFormat       string
	Level            zerolog.Level
	DurationUnit     time.Duration
}

// NewZerologCfgViper creates zerolog config from viper config instance.
func NewZerologCfgViper(v *viper.Viper) ZerologCfg { //nolint:varnamelen
	if !v.InConfig("level") {
		panic("zerolog: invalid viper config")
	}

	level, err := zerolog.ParseLevel(v.GetString("level"))
	if err != nil {
		panic(err)
	}

	return ZerologCfg{
		Level:            level,
		LevelFieldName:   v.GetString("fields_names.level"),
		MessageFieldName: v.GetString("fields_names.msg"),
		TimeFieldName:    v.GetString("fields_names.time"),
		TimeFormat:       v.GetString("time_format"),
		DurationUnit:     v.GetDuration("duration_unit"),
		Caller:           v.GetBool("caller"),
		Pretty:           v.GetBool("pretty"),
		Colors:           v.GetBool("colors"),
	}
}

// NewZerolog creates instance of zerolog logger.
func NewZerolog(cfg ZerologCfg, writer io.Writer) zerolog.Logger {
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().In(time.Local)
	}

	if cfg.DurationUnit > 0 {
		zerolog.DurationFieldUnit = cfg.DurationUnit
	}

	if cfg.LevelFieldName != "" {
		zerolog.LevelFieldName = cfg.LevelFieldName
	}

	if cfg.MessageFieldName != "" {
		zerolog.MessageFieldName = cfg.MessageFieldName
	}

	if cfg.TimeFieldName != "" {
		zerolog.TimestampFieldName = cfg.TimeFieldName
	}

	if cfg.TimeFormat != "" {
		zerolog.TimeFieldFormat = cfg.TimeFormat
	}

	if cfg.Pretty {
		writer = zerolog.ConsoleWriter{Out: writer, NoColor: !cfg.Colors} //nolint:exhaustruct,exhaustivestruct
	}

	ctx := zerolog.New(writer).Level(cfg.Level).With().Timestamp()

	if cfg.Caller {
		ctx = ctx.Caller()
	}

	return ctx.Logger()
}
