package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("title"),
		field.String("description").
			Optional(),
		field.String("goodreads_id").
			Optional().
			Unique(),
		field.String("google_volume_id").
			Unique(),
		field.String("isbn_10").
			Optional(),
		field.String("isbn_13").
			Optional(),
		field.String("publisher").
			Optional(),
		field.String("image_url").
			Optional(),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("book_author", BookAuthor.Type).Ref("books"),
		edge.To("media_request", MediaRequest.Type).Unique(),
	}
}
