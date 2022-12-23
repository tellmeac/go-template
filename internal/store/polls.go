package store

import (
	"context"
	"github.com/tellmeac/go-template/internal/model"
)

type PollStore interface {
	Create(ctx context.Context, poll model.Poll) error
	All(ctx context.Context) ([]model.Poll, error)
}
