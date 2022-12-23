package store

import (
	"context"
	"github.com/tellmeac/go-template/internal/model"
	"github.com/tellmeac/go-template/pkg/ulid"
)

type PlayerStore interface {
	Get(ctx context.Context, id ulid.ULID) (model.PollPlayer, error)
	GetByPoll(ctx context.Context, id ulid.ULID) ([]model.PollPlayer, error)
	Create(ctx context.Context, player model.PollPlayer) error
	Update(ctx context.Context, player model.PollPlayer) error
}
