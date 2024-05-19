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
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("full_name").Unique(), // This may not be true but I need to meet postgres constraints
	}
}

// Edges of the BookAuthor.
func (BookAuthor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("books", Book.Type),
	}
}
