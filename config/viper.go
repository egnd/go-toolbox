// Package config contains tools for working with configs.
package config

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

// ViperParams is a params for viper cfg instance creation.
type ViperParams struct {
	UseEnv         bool
	UseOverride    bool
	EnvPrefix      string
	OverrideSuffix string
	Path           string
	// @TODO: cli flags binding
}

// NewViperCfg creates instance of viper config.
func NewViperCfg(params ViperParams) *viper.Viper {
	if _, err := os.Stat(params.Path); err != nil {
		panic(err)
	}

	cfg := viper.New()
	cfg.SetConfigFile(params.Path)

	viperEnableEnvVars(params, cfg)

	if err := cfg.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viperOverrideWith(params, cfg); err != nil {
		panic(err)
	}

	// workaround because viper does not treat env vars the same as other config
	for _, key := range cfg.AllKeys() {
		val := cfg.Get(key)
		cfg.Set(key, val)
	}

	return cfg
}

func viperEnableEnvVars(params ViperParams, cfg *viper.Viper) {
	if params.UseEnv {
		if params.EnvPrefix != "" {
			cfg.SetEnvPrefix(params.EnvPrefix)
		}

		cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		cfg.AutomaticEnv()
	}
}

func viperOverrideWith(params ViperParams, cfg *viper.Viper) error {
	if params.UseOverride {
		suffix := "override"
		if params.OverrideSuffix != "" {
			suffix = params.OverrideSuffix
		}

		overrideFile := fmt.Sprintf("%s.%s.%s",
			strings.TrimRight(strings.TrimSuffix(params.Path, path.Ext(params.Path)), "."),
			strings.Trim(suffix, "."),
			strings.Trim(path.Ext(params.Path), "."),
		)
		if _, fsErr := os.Stat(overrideFile); fsErr == nil {
			cfg.SetConfigFile(overrideFile)

			if err := cfg.MergeInConfig(); err != nil {
				return err //nolint:wrapcheck
			}
		}
	}

	return nil
}
