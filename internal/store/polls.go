package store

import (
	"context"
	"github.com/tellmeac/go-template/internal/model"
	"github.com/tellmeac/go-template/pkg/ulid"
)

type PollStore interface {
	Get(ctx context.Context, id ulid.ULID) (model.Poll, error)
	Create(ctx context.Context, poll model.Poll) error
	All(ctx context.Context) ([]model.Poll, error)
}
