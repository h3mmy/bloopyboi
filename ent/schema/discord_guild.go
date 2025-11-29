package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DiscordGuild holds the schema definition for the DiscordGuild entity.
type DiscordGuild struct {
	ent.Schema
}

// Fields of the DiscordGuild.
// https://discord.com/developers/docs/resources/guild
func (DiscordGuild) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("discordid").
			Unique(),
		field.String("name"),
		field.String("description").
			Optional(),
		field.String("rules_channel_id").
			Optional(),
		field.String("public_updates_channel_id").
			Optional(),
		field.Int("nsfw_level").
			Optional(),
	}
}

// Edges of the DiscordGuild.
func (DiscordGuild) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("members", DiscordUser.Type),
		edge.To("discord_messages", DiscordMessage.Type),
		edge.To("guild_channels", DiscordChannel.Type),
		edge.To("guild_emojis", Emoji.Type),
	}
}
