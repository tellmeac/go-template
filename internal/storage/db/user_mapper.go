package db

import (
	"context"
	"github.com/Masterminds/squirrel"

	"github.com/google/uuid"
	"github.com/tellmeac/go-template/internal/core/entities"
)

type UserMapper struct {
	Storage *Storage
}

func (m *UserMapper) Get(ctx context.Context, uid uuid.UUID) (*entities.User, error) {
	query, args, err := squirrel.Select("uid", "name").From("users").
		Where(squirrel.Eq{
			"uid": uid,
		}).ToSql()
	if err != nil {
		return nil, err
	}

	row := m.Storage.Pool.QueryRow(ctx, query, args)

	user := &entities.User{}
	err = row.Scan(&user.UID, &user.Name)
	if err != nil {
		return nil, err
	}

	return user, nil
}
