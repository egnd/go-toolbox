package config_test

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/egnd/go-toolbox/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func Test_ViperCfg(t *testing.T) {
	cases := []struct {
		params          config.ViperParams
		panics          string
		cfgFile         string
		cfgOverrideFile string
		envVars         map[string]string
		check           func(*testing.T, *viper.Viper)
	}{
		{
			params: config.ViperParams{Path: "/tmp/cfg1.yml",
				UseOverride: true, OverrideSuffix: "upd",
				UseEnv: true, EnvPrefix: "TC1",
			},
			cfgFile:         `var1: orig-val`,
			cfgOverrideFile: `var1: overrided-val`,
			envVars:         map[string]string{"VAR1": "env-val", "TC1_VAR1": "prefixed-env-val"},
			check: func(t *testing.T, cfg *viper.Viper) {
				assert.EqualValues(t, "prefixed-env-val", cfg.GetString("var1"))
			},
		},
		{
			params: config.ViperParams{Path: "/tmp/cfg1.yml",
				UseOverride: true,
				UseEnv:      true, EnvPrefix: "TC1",
			},
			cfgFile:         `var1: orig-val`,
			cfgOverrideFile: `var1: overrided-val`,
			check: func(t *testing.T, cfg *viper.Viper) {
				assert.EqualValues(t, "overrided-val", cfg.GetString("var1"))
			},
		},
		{
			params: config.ViperParams{Path: "/tmp/cfg1.yml",
				UseOverride: true, OverrideSuffix: "upd",
			},
			cfgFile:         `var1: orig-val`,
			cfgOverrideFile: `var1: overrided-val`,
			check: func(t *testing.T, cfg *viper.Viper) {
				assert.EqualValues(t, "overrided-val", cfg.GetString("var1"))
			},
		},
		{
			params:  config.ViperParams{Path: "/tmp/cfg1.yml"},
			cfgFile: `var1: orig-val`,
			check: func(t *testing.T, cfg *viper.Viper) {
				assert.EqualValues(t, "orig-val", cfg.GetString("var1"))
			},
		},
		{
			panics: "stat : no such file or directory",
		},
		{
			params:  config.ViperParams{Path: "/tmp/cfg1.yml"},
			cfgFile: "-",
			panics:  "While parsing config: yaml: unmarshal errors:\n  line 1: cannot unmarshal !!seq into map[string]interface {}",
		},
		{
			params:          config.ViperParams{Path: "/tmp/cfg1.yml", UseOverride: true},
			cfgFile:         `var1: orig-val`,
			cfgOverrideFile: `-`,
			panics:          "While parsing config: yaml: unmarshal errors:\n  line 1: cannot unmarshal !!seq into map[string]interface {}",
		},
	}

	for k, test := range cases {
		k, test := k, test
		t.Run(fmt.Sprint(k+1), func(t *testing.T) {
			if test.params.Path != "" {
				if err := os.WriteFile(test.params.Path, []byte(test.cfgFile), 0644); err == nil {
					defer os.Remove(test.params.Path)

					if test.params.UseOverride {
						suffix := "override"
						if test.params.OverrideSuffix != "" {
							suffix = test.params.OverrideSuffix
						}
						overrideFile := fmt.Sprintf("%s.%s.%s",
							strings.TrimRight(strings.TrimSuffix(test.params.Path, path.Ext(test.params.Path)), "."),
							strings.Trim(suffix, "."),
							strings.Trim(path.Ext(test.params.Path), "."),
						)
						if err := os.WriteFile(overrideFile, []byte(test.cfgOverrideFile), 0644); err == nil {
							defer os.Remove(overrideFile)
						}
					}

					if test.params.UseEnv {
						for key, val := range test.envVars {
							os.Setenv(key, val)
						}
						defer func() {
							for key := range test.envVars {
								os.Unsetenv(key)
							}
						}()
					}
				}
			}

			if test.panics != "" {
				assert.PanicsWithError(t, test.panics, func() { config.NewViperCfg(test.params) })
				return
			}

			test.check(t, config.NewViperCfg(test.params))
		})
	}
}
