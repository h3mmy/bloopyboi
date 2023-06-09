// Code generated by ent, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBloopyId holds the string denoting the bloopyid field in the database.
	FieldBloopyId = "bloopy_id"
	// FieldDiscordid holds the string denoting the discordid field in the database.
	FieldDiscordid = "discordid"
	// FieldPlexid holds the string denoting the plexid field in the database.
	FieldPlexid = "plexid"
	// FieldBloopnetid holds the string denoting the bloopnetid field in the database.
	FieldBloopnetid = "bloopnetid"
	// FieldAuthentikpkid holds the string denoting the authentikpkid field in the database.
	FieldAuthentikpkid = "authentikpkid"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// EdgeMediaRequest holds the string denoting the mediarequest edge name in mutations.
	EdgeMediaRequest = "mediaRequest"
	// Table holds the table name of the user in the database.
	Table = "users"
	// GroupsTable is the table that holds the groups relation/edge. The primary key declared below.
	GroupsTable = "group_users"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
	// MediaRequestTable is the table that holds the mediaRequest relation/edge.
	MediaRequestTable = "media_requests"
	// MediaRequestInverseTable is the table name for the MediaRequest entity.
	// It exists in this package in order to avoid circular dependency with the "mediarequest" package.
	MediaRequestInverseTable = "media_requests"
	// MediaRequestColumn is the table column denoting the mediaRequest relation/edge.
	MediaRequestColumn = "user_media_request"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldBloopyId,
	FieldDiscordid,
	FieldPlexid,
	FieldBloopnetid,
	FieldAuthentikpkid,
}

var (
	// GroupsPrimaryKey and GroupsColumn2 are the table columns denoting the
	// primary key for the groups relation (M2M).
	GroupsPrimaryKey = []string{"group_id", "user_id"}
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
	// DefaultPlexid holds the default value on creation for the "Plexid" field.
	DefaultPlexid string
)
