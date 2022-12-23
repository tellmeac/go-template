package app

import (
	"context"
	"github.com/tellmeac/go-template/internal/model"
	"github.com/tellmeac/go-template/internal/store"
	"github.com/tellmeac/go-template/pkg/ulid"
)

type App struct {
	store store.Store
}

func (a App) AllPolls(ctx context.Context) ([]model.Poll, error) {
	//TODO implement me
	panic("implement me")
}

func (a App) GetPlayer(ctx context.Context, id ulid.ULID) (model.PollPlayer, error) {
	//TODO implement me
	panic("implement me")
}

func (a App) GetPlayersByPoll(ctx context.Context, id ulid.ULID) ([]model.PollPlayer, error) {
	//TODO implement me
	panic("implement me")
}

func (a App) CreatePoll(ctx context.Context, poll model.PollPost) error {
	//TODO implement me
	panic("implement me")
}

func (a App) StartPoll(ctx context.Context, post model.PollPlayerPost) error {
	//TODO implement me
	panic("implement me")
}

func (a App) Vote(ctx context.Context, request model.VoteRecord) error {
	//TODO implement me
	panic("implement me")
}
