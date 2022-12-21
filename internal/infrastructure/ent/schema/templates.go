package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Templates struct {
	ent.Schema
}

func (Templates) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("template"),
	}
}

func (Templates) Index() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique(),
	}
}
