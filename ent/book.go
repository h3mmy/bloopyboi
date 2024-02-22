// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/h3mmy/bloopyboi/ent/book"
	"github.com/h3mmy/bloopyboi/ent/mediarequest"
)

// Book is the model entity for the Book schema.
type Book struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// GoodreadsID holds the value of the "goodreads_id" field.
	GoodreadsID string `json:"goodreads_id,omitempty"`
	// GoogleVolumeID holds the value of the "google_volume_id" field.
	GoogleVolumeID string `json:"google_volume_id,omitempty"`
	// Isbn10 holds the value of the "isbn_10" field.
	Isbn10 string `json:"isbn_10,omitempty"`
	// Isbn13 holds the value of the "isbn_13" field.
	Isbn13 string `json:"isbn_13,omitempty"`
	// Publisher holds the value of the "publisher" field.
	Publisher string `json:"publisher,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BookQuery when eager-loading is set.
	Edges              BookEdges `json:"edges"`
	book_media_request *uuid.UUID
	selectValues       sql.SelectValues
}

// BookEdges holds the relations/edges for other nodes in the graph.
type BookEdges struct {
	// BookAuthor holds the value of the book_author edge.
	BookAuthor []*BookAuthor `json:"book_author,omitempty"`
	// MediaRequest holds the value of the media_request edge.
	MediaRequest *MediaRequest `json:"media_request,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes     [2]bool
	namedBookAuthor map[string][]*BookAuthor
}

// BookAuthorOrErr returns the BookAuthor value or an error if the edge
// was not loaded in eager-loading.
func (e BookEdges) BookAuthorOrErr() ([]*BookAuthor, error) {
	if e.loadedTypes[0] {
		return e.BookAuthor, nil
	}
	return nil, &NotLoadedError{edge: "book_author"}
}

// MediaRequestOrErr returns the MediaRequest value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookEdges) MediaRequestOrErr() (*MediaRequest, error) {
	if e.loadedTypes[1] {
		if e.MediaRequest == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: mediarequest.Label}
		}
		return e.MediaRequest, nil
	}
	return nil, &NotLoadedError{edge: "media_request"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Book) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case book.FieldTitle, book.FieldDescription, book.FieldGoodreadsID, book.FieldGoogleVolumeID, book.FieldIsbn10, book.FieldIsbn13, book.FieldPublisher, book.FieldImageURL:
			values[i] = new(sql.NullString)
		case book.FieldID:
			values[i] = new(uuid.UUID)
		case book.ForeignKeys[0]: // book_media_request
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Book fields.
func (b *Book) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case book.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				b.ID = *value
			}
		case book.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				b.Title = value.String
			}
		case book.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				b.Description = value.String
			}
		case book.FieldGoodreadsID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field goodreads_id", values[i])
			} else if value.Valid {
				b.GoodreadsID = value.String
			}
		case book.FieldGoogleVolumeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field google_volume_id", values[i])
			} else if value.Valid {
				b.GoogleVolumeID = value.String
			}
		case book.FieldIsbn10:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field isbn_10", values[i])
			} else if value.Valid {
				b.Isbn10 = value.String
			}
		case book.FieldIsbn13:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field isbn_13", values[i])
			} else if value.Valid {
				b.Isbn13 = value.String
			}
		case book.FieldPublisher:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field publisher", values[i])
			} else if value.Valid {
				b.Publisher = value.String
			}
		case book.FieldImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
				b.ImageURL = value.String
			}
		case book.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field book_media_request", values[i])
			} else if value.Valid {
				b.book_media_request = new(uuid.UUID)
				*b.book_media_request = *value.S.(*uuid.UUID)
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Book.
// This includes values selected through modifiers, order, etc.
func (b *Book) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryBookAuthor queries the "book_author" edge of the Book entity.
func (b *Book) QueryBookAuthor() *BookAuthorQuery {
	return NewBookClient(b.config).QueryBookAuthor(b)
}

// QueryMediaRequest queries the "media_request" edge of the Book entity.
func (b *Book) QueryMediaRequest() *MediaRequestQuery {
	return NewBookClient(b.config).QueryMediaRequest(b)
}

// Update returns a builder for updating this Book.
// Note that you need to call Book.Unwrap() before calling this method if this Book
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Book) Update() *BookUpdateOne {
	return NewBookClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Book entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Book) Unwrap() *Book {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Book is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Book) String() string {
	var builder strings.Builder
	builder.WriteString("Book(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("title=")
	builder.WriteString(b.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(b.Description)
	builder.WriteString(", ")
	builder.WriteString("goodreads_id=")
	builder.WriteString(b.GoodreadsID)
	builder.WriteString(", ")
	builder.WriteString("google_volume_id=")
	builder.WriteString(b.GoogleVolumeID)
	builder.WriteString(", ")
	builder.WriteString("isbn_10=")
	builder.WriteString(b.Isbn10)
	builder.WriteString(", ")
	builder.WriteString("isbn_13=")
	builder.WriteString(b.Isbn13)
	builder.WriteString(", ")
	builder.WriteString("publisher=")
	builder.WriteString(b.Publisher)
	builder.WriteString(", ")
	builder.WriteString("image_url=")
	builder.WriteString(b.ImageURL)
	builder.WriteByte(')')
	return builder.String()
}

// NamedBookAuthor returns the BookAuthor named value or an error if the edge was not
// loaded in eager-loading with this name.
func (b *Book) NamedBookAuthor(name string) ([]*BookAuthor, error) {
	if b.Edges.namedBookAuthor == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := b.Edges.namedBookAuthor[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (b *Book) appendNamedBookAuthor(name string, edges ...*BookAuthor) {
	if b.Edges.namedBookAuthor == nil {
		b.Edges.namedBookAuthor = make(map[string][]*BookAuthor)
	}
	if len(edges) == 0 {
		b.Edges.namedBookAuthor[name] = []*BookAuthor{}
	} else {
		b.Edges.namedBookAuthor[name] = append(b.Edges.namedBookAuthor[name], edges...)
	}
}

// Books is a parsable slice of Book.
type Books []*Book
