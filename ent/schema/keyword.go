package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Keyword struct {
	ent.Schema
}

// Fields of the Keyword.
func (Keyword) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("keyword").Unique(),
	}
}

// Edges of the Keyword.
func (Keyword) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("emojis", Emoji.Type).
			Ref("keywords").
			Through("emoji_keyword_scores", EmojiKeywordScore.Type),
	}

}
