package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/tellmeac/go-template/internal/model"
	"github.com/tellmeac/go-template/pkg/ulid"
)

type Poll struct {
	ent.Schema
}

func (Poll) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", ulid.ULID{}),
		field.String("Question"),
		field.JSON("Answers", []model.Answer{}).
			Comment("Contains list of possible answers in poll"),
	}
}

func (Poll) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("players", PollPlayer.Type),
	}
}
