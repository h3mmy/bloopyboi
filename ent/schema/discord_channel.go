package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"github.com/h3mmy/bloopyboi/internal/discord"
)

// DiscordChannel holds the schema definition for the DiscordChannel entity.
type DiscordChannel struct {
	ent.Schema
}

// Fields of the DiscordChannel.
// https://discord.com/developers/docs/resources/channel
func (DiscordChannel) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("discordid").
			Unique(),
		field.String("name"),
		field.Int("type").
			GoType(discord.ChannelType(0)),
		field.Bool("nsfw").
			Default(false),
		field.Int("flags").
			Optional().
			Comment("channel flags combined as a bitfield"),
	}
}

// Edges of the DiscordChannel.
func (DiscordChannel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("discord_guild", DiscordGuild.Type).
			Ref("guild_channels"),
			edge.To("messages", DiscordMessage.Type),
	}
}

func (DiscordChannel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
