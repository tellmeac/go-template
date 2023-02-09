package core

import (
	"github.com/spf13/viper"
	"github.com/tellmeac/go-template/internal/storage/db"
)

type Config struct {
	Debug  bool `yaml:"debug"`
	Server struct {
		Listen string `yaml:"listen"`
	} `yaml:"server"`
	Storage db.StorageConfig `yaml:"storage"`
}

func ParseConfig(loader *viper.Viper) (*Config, error) {
	conf := &Config{}

	loader.Debug()

	if err := loader.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := loader.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
