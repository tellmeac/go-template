package actions

import (
	"context"
	"github.com/tellmeac/go-template/internal/storage/db"
	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/tellmeac/go-template/internal/core/entities"
)

func GetUser(ctx context.Context, uid uuid.UUID, db *db.Storage) (*entities.User, error) {
	zap.L().Info("GetUser", zap.Any("uid", uid))

	return db.User.Get(ctx, uid)
}
