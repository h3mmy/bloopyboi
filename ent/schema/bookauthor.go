package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BookAuthor holds the schema definition for the BookAuthor entity.
type BookAuthor struct {
	ent.Schema
}

// Fields of the BookAuthor.
func (BookAuthor) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("full_name"),
	}
}

// Edges of the BookAuthor.
func (BookAuthor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("books", Book.Type),
	}
}
