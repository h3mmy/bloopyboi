package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DiscordUser holds the schema definition for the DiscordUser entity.
type DiscordUser struct {
	ent.Schema
}

// Fields of the DiscordUser.
func (DiscordUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("discordid").
			Default("unknown"),
		field.String("username"),
	}
}

// Edges of the DiscordUser.
func (DiscordUser) Edges() []ent.Edge {
	return []ent.Edge{
		// create an inverse-edge called "groups" of type `Group`
		// and reference it to the "users" edge (in Group schema)
		// explicitly using the `Ref` method.
		edge.To("discord_messages", DiscordMessage.Type),
	}
}
