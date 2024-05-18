package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

// DiscordMessage holds the schema definition for the DiscordMessage entity.
type DiscordMessage struct {
	ent.Schema
}

// Fields of the DiscordMessage.
func (DiscordMessage) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).
			Unique(),
		field.String("discordid").
			Unique(),
		field.JSON("raw", discordgo.Message{}),
	}
}

// Edges of the DiscordMessage.
func (DiscordMessage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", DiscordUser.Type).Ref("discord_messages").Unique(),
		edge.To("message_reactions", DiscordMessageReaction.Type),
		// edge.To("channel", DiscordChannel.Type),
		edge.From("guild", DiscordGuild.Type).Ref("discord_messages").Unique(),
	}
}

func (DiscordMessage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
