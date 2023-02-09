package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StorageConfig struct {
	ConnectionString string `yaml:"connectionString"`
}

type Storage struct {
	Config StorageConfig
	Pool   *pgxpool.Pool

	User UserMapper
}

func NewStorage(ctx context.Context, config StorageConfig) (*Storage, error) {
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
