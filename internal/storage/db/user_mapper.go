package db

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/tellmeac/go-template/internal/core/entities"
)

type UserMapper struct {
	Storage *Storage
}

func (m *UserMapper) Get(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	query, args, err := squirrel.Select("*").From("users").
		PlaceholderFormat(squirrel.Dollar).
		Where("uid=?", id).ToSql()
	if err != nil {
		return nil, fmt.Errorf("query: %s", err)
	}

	row := m.Storage.Pool.QueryRow(ctx, query, args...)

	user := &entities.User{}
	err = row.Scan(&user.UID, &user.Name)
	if err != nil {
		return nil, fmt.Errorf("scan result: %s", err)
	}

	return user, nil
}
