// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// BloopyId holds the value of the "bloopyId" field.
	BloopyId uuid.UUID `json:"bloopyId,omitempty"`
	// Discordid holds the value of the "discordid" field.
	Discordid string `json:"discordid,omitempty"`
	// Plexid holds the value of the "Plexid" field.
	Plexid string `json:"Plexid,omitempty"`
	// Bloopnetid holds the value of the "bloopnetid" field.
	Bloopnetid string `json:"bloopnetid,omitempty"`
	// Authentikpkid holds the value of the "authentikpkid" field.
	Authentikpkid string `json:"authentikpkid,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Groups holds the value of the groups edge.
	Groups []*Group `json:"groups,omitempty"`
	// MediaRequest holds the value of the mediaRequest edge.
	MediaRequest []*MediaRequest `json:"mediaRequest,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[0] {
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// MediaRequestOrErr returns the MediaRequest value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MediaRequestOrErr() ([]*MediaRequest, error) {
	if e.loadedTypes[1] {
		return e.MediaRequest, nil
	}
	return nil, &NotLoadedError{edge: "mediaRequest"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldDiscordid, user.FieldPlexid, user.FieldBloopnetid, user.FieldAuthentikpkid:
			values[i] = new(sql.NullString)
		case user.FieldBloopyId:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldBloopyId:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field bloopyId", values[i])
			} else if value != nil {
				u.BloopyId = *value
			}
		case user.FieldDiscordid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field discordid", values[i])
			} else if value.Valid {
				u.Discordid = value.String
			}
		case user.FieldPlexid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Plexid", values[i])
			} else if value.Valid {
				u.Plexid = value.String
			}
		case user.FieldBloopnetid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bloopnetid", values[i])
			} else if value.Valid {
				u.Bloopnetid = value.String
			}
		case user.FieldAuthentikpkid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field authentikpkid", values[i])
			} else if value.Valid {
				u.Authentikpkid = value.String
			}
		}
	}
	return nil
}

// QueryGroups queries the "groups" edge of the User entity.
func (u *User) QueryGroups() *GroupQuery {
	return (&UserClient{config: u.config}).QueryGroups(u)
}

// QueryMediaRequest queries the "mediaRequest" edge of the User entity.
func (u *User) QueryMediaRequest() *MediaRequestQuery {
	return (&UserClient{config: u.config}).QueryMediaRequest(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("bloopyId=")
	builder.WriteString(fmt.Sprintf("%v", u.BloopyId))
	builder.WriteString(", ")
	builder.WriteString("discordid=")
	builder.WriteString(u.Discordid)
	builder.WriteString(", ")
	builder.WriteString("Plexid=")
	builder.WriteString(u.Plexid)
	builder.WriteString(", ")
	builder.WriteString("bloopnetid=")
	builder.WriteString(u.Bloopnetid)
	builder.WriteString(", ")
	builder.WriteString("authentikpkid=")
	builder.WriteString(u.Authentikpkid)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
