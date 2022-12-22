package models

import "github.com/tellmeac/go-template/pkg/ulid"

type Poll struct {
	ID           ulid.ULID `json:"id"`
	Question     string    `json:"question"`
	Answers      []Answer  `json:"answers"`
	answersCount int64
}

func NewPoll(question string) *Poll {
	return &Poll{
		ID:       ulid.New(timeNow()),
		Question: question,
	}
}

func (p *Poll) WithAnswers(contents ...string) *Poll {
	for _, content := range contents {
		p.answersCount++

		p.Answers = append(p.Answers, Answer{
			ID:      p.answersCount,
			Content: content,
		})
	}

	return p
}

type Answer struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}
