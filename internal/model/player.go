package model

import (
	"github.com/tellmeac/go-template/pkg/ulid"
	"net/http"
	"time"
)

var timeNow = time.Now

type PollPlayerPost struct {
	PollID                 ulid.ULID  `json:"pollId"`
	CloseAt                *time.Time `json:"closeAt"`
	MultipleAnswersAllowed bool       `json:"multipleAnswersAllowed"`
}

func StartPlayer(poll Poll, closeAt *time.Time, multipleAllowed bool) (*PollPlayer, error) {
	now := timeNow()
	if closeAt != nil && now.After(*closeAt) {
		return nil, NewError("start date after close date").WithCode(http.StatusBadRequest)
	}

	counters := make([]AnswerCounter, 0, len(poll.Answers))
	for _, answer := range poll.Answers {
		counters = append(counters, AnswerCounter{
			AnswerID: answer.ID,
			Counter:  0,
		})
	}

	return &PollPlayer{
		ID:              ulid.New(timeNow()),
		PollID:          poll.ID,
		StartAt:         now,
		CloseAt:         closeAt,
		Counters:        counters,
		MultipleAllowed: multipleAllowed,
	}, nil
}

type PollPlayer struct {
	ID              ulid.ULID       `json:"id"`
	PollID          ulid.ULID       `json:"pollID"`
	StartAt         time.Time       `json:"startAt"`
	CloseAt         *time.Time      `json:"closeAt"`
	MultipleAllowed bool            `json:"multipleAllowed"`
	Counters        []AnswerCounter `json:"counters"`
}

func (player *PollPlayer) Vote(answerID int64) error {
	if !player.isAvailable() {
		return NewError("poll is unavailable").WithCode(http.StatusForbidden)
	}

	for _, c := range player.Counters {
		if c.AnswerID == answerID {
			c.Inc()
			return nil
		}
	}
	return NewError("answer id not exists").WithCode(http.StatusBadRequest)
}

func (player *PollPlayer) isAvailable() bool {
	if player.CloseAt == nil {
		return true
	}
	return timeNow().Before(*player.CloseAt)
}

type AnswerCounter struct {
	AnswerID int64 `json:"answer"`
	Counter  int64 `json:"counter"`
}

func (c *AnswerCounter) Inc() {
	c.Counter++
}
