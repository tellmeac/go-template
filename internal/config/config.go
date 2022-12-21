package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog/log"
	"sync"
)

// Config is a global configuration object.
type Config struct {
	Debug           bool   `env:"DEBUG" env-default:"false"`
	ListenAddress   string `env:"LISTEN_ADDRESS" env-default:"0.0.0.0:80"`
	DatabaseAddress string `env:"DATABASE_ADDRESS" env-default:"postgres://postgres:postgres@localhost:5432/Greetings?sslmode=disable"`
}

var (
	instance *Config
	once     sync.Once
)

// Get reads config from env once.
func Get() Config {
	once.Do(func() {
		if instance != nil {
			return
		}

		var config Config
		if err := cleanenv.ReadEnv(&config); err != nil {
			log.Fatal().Err(err).Msg("Failed to read config from environment")
		}
		instance = &config
	})
	return *instance
}
