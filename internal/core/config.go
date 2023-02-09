package core

import (
	"context"

	"github.com/heetch/confita"
	"github.com/tellmeac/go-template/internal/storage/db"
)

type Config struct {
	Debug  bool `yaml:"debug"`
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Storage db.Storage `yaml:"storage"`
}

func ParseConfig(ctx context.Context, loader *confita.Loader) (*Config, error) {
	conf := &Config{}

	if loader != nil {
		err := loader.Load(ctx, conf)
		if err != nil {
			return nil, err
		}
	}

	return conf, nil
}
