// Code generated by ent, DO NOT EDIT.

package mediarequest

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the mediarequest type in the database.
	Label = "media_request"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPriority holds the string denoting the priority field in the database.
	FieldPriority = "priority"
	// EdgeDiscordUser holds the string denoting the discord_user edge name in mutations.
	EdgeDiscordUser = "discord_user"
	// EdgeBooks holds the string denoting the books edge name in mutations.
	EdgeBooks = "books"
	// Table holds the table name of the mediarequest in the database.
	Table = "media_requests"
	// DiscordUserTable is the table that holds the discord_user relation/edge.
	DiscordUserTable = "media_requests"
	// DiscordUserInverseTable is the table name for the DiscordUser entity.
	// It exists in this package in order to avoid circular dependency with the "discorduser" package.
	DiscordUserInverseTable = "discord_users"
	// DiscordUserColumn is the table column denoting the discord_user relation/edge.
	DiscordUserColumn = "discord_user_media_requests"
	// BooksTable is the table that holds the books relation/edge. The primary key declared below.
	BooksTable = "media_request_books"
	// BooksInverseTable is the table name for the Book entity.
	// It exists in this package in order to avoid circular dependency with the "book" package.
	BooksInverseTable = "books"
)

// Columns holds all SQL columns for mediarequest fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStatus,
	FieldPriority,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "media_requests"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"discord_user_media_requests",
}

var (
	// BooksPrimaryKey and BooksColumn2 are the table columns denoting the
	// primary key for the books relation (M2M).
	BooksPrimaryKey = []string{"media_request_id", "book_id"}
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

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultPriority holds the default value on creation for the "priority" field.
	DefaultPriority int
)

// OrderOption defines the ordering options for the MediaRequest queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByPriority orders the results by the priority field.
func ByPriority(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPriority, opts...).ToFunc()
}

// ByDiscordUserField orders the results by discord_user field.
func ByDiscordUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDiscordUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByBooksCount orders the results by books count.
func ByBooksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBooksStep(), opts...)
	}
}

// ByBooks orders the results by books terms.
func ByBooks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBooksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDiscordUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DiscordUserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DiscordUserTable, DiscordUserColumn),
	)
}
func newBooksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BooksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, BooksTable, BooksPrimaryKey...),
	)
}