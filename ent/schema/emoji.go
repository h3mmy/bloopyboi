package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Emoji holds the schema definition for the Emoji entity.
type Emoji struct {
	ent.Schema
}

// Fields of the Emoji.
func (Emoji) Fields() []ent.Field {
	return []ent.Field{
		field.String("emoji_id").
			Unique(),
		field.String("name").
			NotEmpty(),
		field.Bool("animated").
			Default(false),
		field.JSON("keywords", []string{}).
			Optional(),
	}
}

// Edges of the Emoji.
func (Emoji) Edges() []ent.Edge {
	return nil
}
