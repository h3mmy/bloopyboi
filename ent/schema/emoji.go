package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Emoji holds the schema definition for the Emoji entity.
type Emoji struct {
	ent.Schema
}

// Fields of the Emoji.
func (Emoji) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("emoji_id").
			Unique(),
		field.String("name").
			NotEmpty(),
		field.Bool("animated").
			Default(false),
		field.String("image_uri").
			Nillable().
			Optional(),
		field.Int("adult_likelihood").
			Default(0),
		field.Int("spoof_likelihood").
			Default(0),
		field.Int("medical_likelihood").
			Default(0),
		field.Int("violence_likelihood").
			Default(0),
		field.Int("racy_likelihood").
			Default(0),
	}
}

// Edges of the Emoji.
func (Emoji) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("guild", DiscordGuild.Type).
			Ref("guild_emojis").
			Unique(),
		edge.To("keywords", Keyword.Type).
			Through("emoji_keyword_scores", EmojiKeywordScore.Type),
	}
}
