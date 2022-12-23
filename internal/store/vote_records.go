package store

import (
	"context"
	"github.com/tellmeac/go-template/internal/model"
	"github.com/tellmeac/go-template/pkg/ulid"
)

type VoteRecordStore interface {
	GetForPlayer(ctx context.Context, playerID ulid.ULID, userID ulid.ULID) (model.VoteRecord, error)
	Create(ctx context.Context, record model.VoteRecord) error
}
