// Code generated by ent, DO NOT EDIT.

package discorduser

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the discorduser type in the database.
	Label = "discord_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDiscordid holds the string denoting the discordid field in the database.
	FieldDiscordid = "discordid"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// EdgeDiscordMessages holds the string denoting the discord_messages edge name in mutations.
	EdgeDiscordMessages = "discord_messages"
	// Table holds the table name of the discorduser in the database.
	Table = "discord_users"
	// DiscordMessagesTable is the table that holds the discord_messages relation/edge. The primary key declared below.
	DiscordMessagesTable = "discord_user_discord_messages"
	// DiscordMessagesInverseTable is the table name for the DiscordMessage entity.
	// It exists in this package in order to avoid circular dependency with the "discordmessage" package.
	DiscordMessagesInverseTable = "discord_messages"
)

// Columns holds all SQL columns for discorduser fields.
var Columns = []string{
	FieldID,
	FieldDiscordid,
	FieldUsername,
}

var (
	// DiscordMessagesPrimaryKey and DiscordMessagesColumn2 are the table columns denoting the
	// primary key for the discord_messages relation (M2M).
	DiscordMessagesPrimaryKey = []string{"discord_user_id", "discord_message_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultDiscordid holds the default value on creation for the "discordid" field.
	DefaultDiscordid string
)

// OrderOption defines the ordering options for the DiscordUser queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDiscordid orders the results by the discordid field.
func ByDiscordid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDiscordid, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByDiscordMessagesCount orders the results by discord_messages count.
func ByDiscordMessagesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDiscordMessagesStep(), opts...)
	}
}

// ByDiscordMessages orders the results by discord_messages terms.
func ByDiscordMessages(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDiscordMessagesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDiscordMessagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DiscordMessagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, DiscordMessagesTable, DiscordMessagesPrimaryKey...),
	)
}
