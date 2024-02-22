// Code generated by ent, DO NOT EDIT.

package book

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the book type in the database.
	Label = "book"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldGoodreadsID holds the string denoting the goodreads_id field in the database.
	FieldGoodreadsID = "goodreads_id"
	// FieldGoogleVolumeID holds the string denoting the google_volume_id field in the database.
	FieldGoogleVolumeID = "google_volume_id"
	// FieldIsbn10 holds the string denoting the isbn_10 field in the database.
	FieldIsbn10 = "isbn_10"
	// FieldIsbn13 holds the string denoting the isbn_13 field in the database.
	FieldIsbn13 = "isbn_13"
	// FieldPublisher holds the string denoting the publisher field in the database.
	FieldPublisher = "publisher"
	// FieldImageURL holds the string denoting the image_url field in the database.
	FieldImageURL = "image_url"
	// EdgeBookAuthor holds the string denoting the book_author edge name in mutations.
	EdgeBookAuthor = "book_author"
	// EdgeMediaRequest holds the string denoting the media_request edge name in mutations.
	EdgeMediaRequest = "media_request"
	// Table holds the table name of the book in the database.
	Table = "books"
	// BookAuthorTable is the table that holds the book_author relation/edge. The primary key declared below.
	BookAuthorTable = "book_author_books"
	// BookAuthorInverseTable is the table name for the BookAuthor entity.
	// It exists in this package in order to avoid circular dependency with the "bookauthor" package.
	BookAuthorInverseTable = "book_authors"
	// MediaRequestTable is the table that holds the media_request relation/edge.
	MediaRequestTable = "books"
	// MediaRequestInverseTable is the table name for the MediaRequest entity.
	// It exists in this package in order to avoid circular dependency with the "mediarequest" package.
	MediaRequestInverseTable = "media_requests"
	// MediaRequestColumn is the table column denoting the media_request relation/edge.
	MediaRequestColumn = "book_media_request"
)

// Columns holds all SQL columns for book fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDescription,
	FieldGoodreadsID,
	FieldGoogleVolumeID,
	FieldIsbn10,
	FieldIsbn13,
	FieldPublisher,
	FieldImageURL,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "books"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"book_media_request",
}

var (
	// BookAuthorPrimaryKey and BookAuthorColumn2 are the table columns denoting the
	// primary key for the book_author relation (M2M).
	BookAuthorPrimaryKey = []string{"book_author_id", "book_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Book queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByGoodreadsID orders the results by the goodreads_id field.
func ByGoodreadsID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGoodreadsID, opts...).ToFunc()
}

// ByGoogleVolumeID orders the results by the google_volume_id field.
func ByGoogleVolumeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGoogleVolumeID, opts...).ToFunc()
}

// ByIsbn10 orders the results by the isbn_10 field.
func ByIsbn10(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsbn10, opts...).ToFunc()
}

// ByIsbn13 orders the results by the isbn_13 field.
func ByIsbn13(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsbn13, opts...).ToFunc()
}

// ByPublisher orders the results by the publisher field.
func ByPublisher(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublisher, opts...).ToFunc()
}

// ByImageURL orders the results by the image_url field.
func ByImageURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImageURL, opts...).ToFunc()
}

// ByBookAuthorCount orders the results by book_author count.
func ByBookAuthorCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBookAuthorStep(), opts...)
	}
}

// ByBookAuthor orders the results by book_author terms.
func ByBookAuthor(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBookAuthorStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMediaRequestField orders the results by media_request field.
func ByMediaRequestField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMediaRequestStep(), sql.OrderByField(field, opts...))
	}
}
func newBookAuthorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BookAuthorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, BookAuthorTable, BookAuthorPrimaryKey...),
	)
}
func newMediaRequestStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MediaRequestInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, MediaRequestTable, MediaRequestColumn),
	)
}
