package app

import (
	"context"
	"errors"
	"github.com/tellmeac/go-template/internal/model"
	"github.com/tellmeac/go-template/internal/store"
	"github.com/tellmeac/go-template/pkg/ulid"
	"net/http"
)

func (a App) GetPlayer(ctx context.Context, id ulid.ULID) (model.PollPlayer, error) {
	player, err := a.store.Players().Get(ctx, id)
	if err != nil {
		return model.PollPlayer{}, model.NewError("failed to get player").Wrap(err).WithCode(http.StatusInternalServerError)
	}
	return player, nil
}

func (a App) GetPlayersByPoll(ctx context.Context, id ulid.ULID) ([]model.PollPlayer, error) {
	players, err := a.store.Players().GetByPoll(ctx, id)
	if err != nil {
		return nil, model.NewError("failed to get players").Wrap(err).WithCode(http.StatusInternalServerError)
	}
	return players, nil
}

func (a App) StartPoll(ctx context.Context, request model.PollPlayerPost) error {
	poll, err := a.store.Polls().Get(ctx, request.PollID)
	if err != nil {
		var nfErr *store.ErrNotFound
		switch {
		case errors.As(err, &nfErr):
			return model.NewError("failed to get poll").Wrap(err).WithCode(http.StatusNotFound)
		default:
			return model.NewError("failed to get poll").Wrap(err).WithCode(http.StatusInternalServerError)
		}
	}

	player, err := model.StartPlayer(poll, request.CloseAt, request.MultipleAnswersAllowed)
	if err != nil {
		return err
	}

	err = a.store.Players().Create(ctx, player)
	if err != nil {
		return model.NewError("failed to save player").Wrap(err).WithCode(http.StatusInternalServerError)
	}
	return nil
}

func (a App) Vote(ctx context.Context, request model.VoteRecord) error {
	player, err := a.store.Players().Get(ctx, request.PollPlayerID)
	if err != nil {
		var nfErr *store.ErrNotFound
		switch {
		case errors.As(err, &nfErr):
			return model.NewError("failed to get player").Wrap(err).WithCode(http.StatusNotFound)
		default:
			return model.NewError("failed to get player").Wrap(err).WithCode(http.StatusInternalServerError)
		}
	}

	mPlayer := &player
	for _, answerID := range request.AnswerIDs {
		err := mPlayer.Vote(answerID)
		if err != nil {
			return err
		}
	}

	err = a.store.VoteRecords().Create(ctx, request)
	if err != nil {
		var cErr *store.ErrConflict
		switch {
		case errors.As(err, &cErr):
			return model.NewError("already voted").Wrap(err).WithCode(http.StatusBadRequest)
		default:
			return model.NewError("failed to save vote record").Wrap(err).WithCode(http.StatusInternalServerError)
		}
	}

	err = a.store.Players().Update(ctx, *mPlayer)
	if err != nil {
		return model.NewError("failed to update player").Wrap(err).WithCode(http.StatusInternalServerError)
	}
	return nil
}
