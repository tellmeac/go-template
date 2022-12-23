package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/tellmeac/go-template/pkg/ulid"
)

type VoteRecord struct {
	ent.Schema
}

func (VoteRecord) Fields() []ent.Field {
	return []ent.Field{
		field.String("UserID"),
		field.UUID("PollPlayerID", ulid.ULID{}),
		field.JSON("Answers", []int{}).
			Comment("Contains number of requested answers in poll"),
	}
}

func (VoteRecord) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("UserID", "PollPlayerID").Unique(),
	}
}

func (VoteRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("player", PollPlayer.Type).
			Ref("records").
			Field("PollPlayerID").
			Unique().
			Required(),
	}
}
