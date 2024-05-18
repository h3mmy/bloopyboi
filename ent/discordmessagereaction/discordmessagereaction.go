// Code generated by ent, DO NOT EDIT.

package discordmessagereaction

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the discordmessagereaction type in the database.
	Label = "discord_message_reaction"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldRemoved holds the string denoting the removed field in the database.
	FieldRemoved = "removed"
	// FieldRaw holds the string denoting the raw field in the database.
	FieldRaw = "raw"
	// EdgeDiscordMessage holds the string denoting the discord_message edge name in mutations.
	EdgeDiscordMessage = "discord_message"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// Table holds the table name of the discordmessagereaction in the database.
	Table = "discord_message_reactions"
	// DiscordMessageTable is the table that holds the discord_message relation/edge.
	DiscordMessageTable = "discord_message_reactions"
	// DiscordMessageInverseTable is the table name for the DiscordMessage entity.
	// It exists in this package in order to avoid circular dependency with the "discordmessage" package.
	DiscordMessageInverseTable = "discord_messages"
	// DiscordMessageColumn is the table column denoting the discord_message relation/edge.
	DiscordMessageColumn = "discord_message_message_reactions"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "discord_message_reactions"
	// AuthorInverseTable is the table name for the DiscordUser entity.
	// It exists in this package in order to avoid circular dependency with the "discorduser" package.
	AuthorInverseTable = "discord_users"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "discord_user_message_reactions"
)

// Columns holds all SQL columns for discordmessagereaction fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldRemoved,
	FieldRaw,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "discord_message_reactions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"discord_message_message_reactions",
	"discord_user_message_reactions",
}

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
	// DefaultRemoved holds the default value on creation for the "removed" field.
	DefaultRemoved bool
)

// OrderOption defines the ordering options for the DiscordMessageReaction queries.
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

// ByRemoved orders the results by the removed field.
func ByRemoved(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemoved, opts...).ToFunc()
}

// ByDiscordMessageField orders the results by discord_message field.
func ByDiscordMessageField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDiscordMessageStep(), sql.OrderByField(field, opts...))
	}
}

// ByAuthorField orders the results by author field.
func ByAuthorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorStep(), sql.OrderByField(field, opts...))
	}
}
func newDiscordMessageStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DiscordMessageInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DiscordMessageTable, DiscordMessageColumn),
	)
}
func newAuthorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
	)
}
