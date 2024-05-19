package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

// DiscordMessageReaction holds the schema definition for the DiscordMessageReaction entity.
type DiscordMessageReaction struct {
	ent.Schema
}

// Fields of the DiscordMessageReaction.
func (DiscordMessageReaction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("emoji_api_name"),
		field.Bool("removed").
			Default(false),
		field.JSON("raw", discordgo.MessageReaction{}),
	}
}

// Edges of the DiscordMessageReaction.
func (DiscordMessageReaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("discord_message", DiscordMessage.Type).
			Ref("message_reactions").
			Required().
			Unique(),
		edge.From("author", DiscordUser.Type).
			Ref("message_reactions").
			Required().
			Unique(),
	}
}

func (DiscordMessageReaction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
