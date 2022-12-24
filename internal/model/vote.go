package model

import "github.com/oklog/ulid/v2"

type VoteRecord struct {
	UserID       string
	PollPlayerID ulid.ULID `json:"pollPlayerId"`
	AnswerIDs    []int64   `json:"answerIds"`
}
