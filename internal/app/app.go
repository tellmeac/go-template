package app

import (
	"context"
	"github.com/tellmeac/go-template/internal/model"
	"github.com/tellmeac/go-template/pkg/ulid"
)

type App interface {
	GetPoll(ctx context.Context) (model.Poll, error)
	GetPollPlayer(ctx context.Context, id ulid.ULID) (model.PollPlayer, error)
	GetPlayersByPoll(ctx context.Context, id ulid.ULID) ([]model.PollPlayer, error)
	CreatePoll(ctx context.Context, poll model.PollPost) error
	StartPollPlayer(ctx context.Context, post model.PollPlayerPost) error
	Vote(ctx context.Context, request model.VoteRequest) error
}
