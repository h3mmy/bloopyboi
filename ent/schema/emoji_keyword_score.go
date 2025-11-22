package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

type EmojiKeywordScore struct {
	ent.Schema
}

func (EmojiKeywordScore) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("emoji_id", "keyword_id"),
	}
}

func (EmojiKeywordScore) Fields() []ent.Field {
	return []ent.Field{
		field.Float32("score"),
		field.Float32("topicality"),
		field.UUID("emoji_id", uuid.UUID{}),
		field.UUID("keyword_id", uuid.UUID{}),
	}
}

func (EmojiKeywordScore) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("keyword", Keyword.Type).
			Required().
			Unique().
			Field("keyword_id"),
		edge.To("emoji", Emoji.Type).
			Required().
			Unique().
			Field("emoji_id"),
	}
}

func (EmojiKeywordScore) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("emoji_id", "keyword_id"),
		index.Fields("score"),
	}
}
