package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	ConnectionString string `yaml:"connectionString"`
}

type Storage struct {
	Config Config
	Pool   *pgxpool.Pool

	User UserMapper
}

func NewStorage(ctx context.Context, config Config) (*Storage, error) {
	var err error

	storage := &Storage{
		Config: config,
	}

	storage.Pool, err = pgxpool.New(ctx, config.ConnectionString)
	if err != nil {
		return nil, err
	}

	storage.User = UserMapper{
		Storage: storage,
	}

	return storage, nil
}
