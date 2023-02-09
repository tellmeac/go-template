package actions

import (
	"context"
	"github.com/tellmeac/go-template/internal/storage/db"

	"github.com/google/uuid"
	"github.com/tellmeac/go-template/internal/core/entities"
)

func GetUserAction(ctx context.Context, uid uuid.UUID, db *db.Storage) (*entities.User, error) {
	return db.User.Get(ctx, uid)
}
