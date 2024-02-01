package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/bwmarrin/discordgo"
)

// DiscordMessage holds the schema definition for the DiscordMessage entity.
type DiscordMessage struct {
	ent.Schema
}

// Fields of the DiscordMessage.
func (DiscordMessage) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.JSON("raw", discordgo.Message{}),
	}
}

// Edges of the DiscordMessage.
func (DiscordMessage) Edges() []ent.Edge {
	return nil
}

func (DiscordMessage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
