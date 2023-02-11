package actions

import (
	"context"
	"github.com/tellmeac/go-template/internal/storage/db"
	"github.com/tellmeac/go-template/pkg/ctxlog"
	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/tellmeac/go-template/internal/core/entities"
)

func GetUser(ctx context.Context, uid uuid.UUID, db *db.Storage) (*entities.User, error) {
	ctx = ctxlog.WithFields(ctx, zap.Any("uid", uid))
	ctxlog.Ctx(ctx, zap.L()).Info("Get user action")

	return db.User.Get(ctx, uid)
}
