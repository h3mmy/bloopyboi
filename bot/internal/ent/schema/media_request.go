package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MediaRequest holds the schema definition for the MediaRequest entity.
type MediaRequest struct {
	ent.Schema
}

// Fields of the MediaRequest.
func (MediaRequest) Fields() []ent.Field {
	return []ent.Field{
		field.String("status"),
		field.Enum("mediaType").
			Values("movie", "tv", "book", "music"),
		field.String("requestId"),
	}
}

// Edges of the MediaRequest.
func (MediaRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("mediaRequest").
			Unique(),
	}
}
