package app

import (
	"context"
	"github.com/tellmeac/go-template/internal/model"
	"net/http"
)

func (a App) AllPolls(ctx context.Context) ([]model.Poll, error) {
	polls, err := a.store.Polls().All(ctx)
	if err != nil {
		return nil, model.NewError("failed to get all polls").Wrap(err).WithCode(http.StatusInternalServerError)
	}
	return polls, nil
}

func (a App) CreatePoll(ctx context.Context, request model.PollPost) error {
	poll := model.NewPoll(request.Question).WithAnswers(request.AnswersContent...)

	err := a.store.Polls().Create(ctx, *poll)
	if err != nil {
		return model.NewError("failed to save new poll").Wrap(err).WithCode(http.StatusInternalServerError)
	}
	return nil
}
