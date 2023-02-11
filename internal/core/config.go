package core

import (
	"github.com/spf13/viper"
	"github.com/tellmeac/go-template/internal/storage/db"
	"github.com/tellmeac/go-template/pkg/web"
)

type Config struct {
	Debug        bool             `yaml:"debug"`
	ServerConfig web.ServerConfig `yaml:"server"`
	Storage      db.StorageConfig `yaml:"storage"`
}

func ParseConfig(loader *viper.Viper) (*Config, error) {
	conf := &Config{}

	if err := loader.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := loader.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
