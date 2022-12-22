package models

import "github.com/oklog/ulid/v2"

type VoteRequest struct {
	UserID       string
	PollPlayerID ulid.ULID `json:"pollPlayerId"`
	AnswerIDs    []int     `json:"answerIds"`
}
