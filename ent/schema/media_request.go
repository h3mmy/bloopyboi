package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// MediaRequest holds the schema definition for the MediaRequest entity.
type MediaRequest struct {
	ent.Schema
}

// Fields of the MediaRequest.
func (MediaRequest) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).
			Unique(),
		field.String("status"),
		field.Int("priority").
			Default(50),
	}
}

// Edges of the MediaRequest.
func (MediaRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("discord_user", DiscordUser.Type).
			Ref("media_requests").
			Unique(),
		edge.To("books", Book.Type),
	}
}

func (MediaRequest) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
