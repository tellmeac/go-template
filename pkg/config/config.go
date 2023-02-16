package config

import (
	"github.com/spf13/viper"
	"strings"
)

type LoaderOption func(*viper.Viper)

func WithConfigPath(p string) LoaderOption {
	return func(v *viper.Viper) {
		v.SetConfigFile(p)
	}
}

func WithEnvPrefix(prefix string) LoaderOption {
	return func(v *viper.Viper) {
		v.SetEnvPrefix(prefix)
	}
}

func WithEnvKeyReplacement(oldnew ...string) LoaderOption {
	return func(v *viper.Viper) {
		v.SetEnvKeyReplacer(strings.NewReplacer(oldnew...))
	}
}

func PrepareLoader(options ...LoaderOption) *viper.Viper {
	v := viper.New()
	v.AutomaticEnv()
	// Default env key replacing, can be overridden with options.
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	for _, opt := range options {
		opt(v)
	}

	return v
}
