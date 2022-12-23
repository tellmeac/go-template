package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/tellmeac/go-template/internal/model"
	"github.com/tellmeac/go-template/pkg/ulid"
)

type PollPlayer struct {
	ent.Schema
}

func (PollPlayer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", ulid.ULID{}),
		field.UUID("PollID", ulid.ULID{}),
		field.Time("StartAt"),
		field.Time("CloseAt").Nillable(),
		field.Bool("MultipleAllowed"),
		field.JSON("Counters", []model.AnswerCounter{}).
			Comment("Contains list of counters for answers"),
	}
}

func (PollPlayer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("records", VoteRecord.Type),
		edge.From("poll", Poll.Type).
			Ref("players").
			Field("PollID").
			Unique().
			Required(),
	}
}
