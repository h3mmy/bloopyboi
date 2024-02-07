// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/h3mmy/bloopyboi/ent/discorduser"
)

// DiscordUser is the model entity for the DiscordUser schema.
type DiscordUser struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Discordid holds the value of the "discordid" field.
	Discordid string `json:"discordid,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Discriminator holds the value of the "discriminator" field.
	Discriminator string `json:"discriminator,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DiscordUserQuery when eager-loading is set.
	Edges        DiscordUserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DiscordUserEdges holds the relations/edges for other nodes in the graph.
type DiscordUserEdges struct {
	// DiscordMessages holds the value of the discord_messages edge.
	DiscordMessages []*DiscordMessage `json:"discord_messages,omitempty"`
	// MediaRequests holds the value of the media_requests edge.
	MediaRequests []*MediaRequest `json:"media_requests,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes          [2]bool
	namedDiscordMessages map[string][]*DiscordMessage
	namedMediaRequests   map[string][]*MediaRequest
}

// DiscordMessagesOrErr returns the DiscordMessages value or an error if the edge
// was not loaded in eager-loading.
func (e DiscordUserEdges) DiscordMessagesOrErr() ([]*DiscordMessage, error) {
	if e.loadedTypes[0] {
		return e.DiscordMessages, nil
	}
	return nil, &NotLoadedError{edge: "discord_messages"}
}

// MediaRequestsOrErr returns the MediaRequests value or an error if the edge
// was not loaded in eager-loading.
func (e DiscordUserEdges) MediaRequestsOrErr() ([]*MediaRequest, error) {
	if e.loadedTypes[1] {
		return e.MediaRequests, nil
	}
	return nil, &NotLoadedError{edge: "media_requests"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DiscordUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case discorduser.FieldDiscordid, discorduser.FieldUsername, discorduser.FieldEmail, discorduser.FieldDiscriminator:
			values[i] = new(sql.NullString)
		case discorduser.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DiscordUser fields.
func (du *DiscordUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case discorduser.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				du.ID = *value
			}
		case discorduser.FieldDiscordid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field discordid", values[i])
			} else if value.Valid {
				du.Discordid = value.String
			}
		case discorduser.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				du.Username = value.String
			}
		case discorduser.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				du.Email = value.String
			}
		case discorduser.FieldDiscriminator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field discriminator", values[i])
			} else if value.Valid {
				du.Discriminator = value.String
			}
		default:
			du.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DiscordUser.
// This includes values selected through modifiers, order, etc.
func (du *DiscordUser) Value(name string) (ent.Value, error) {
	return du.selectValues.Get(name)
}

// QueryDiscordMessages queries the "discord_messages" edge of the DiscordUser entity.
func (du *DiscordUser) QueryDiscordMessages() *DiscordMessageQuery {
	return NewDiscordUserClient(du.config).QueryDiscordMessages(du)
}

// QueryMediaRequests queries the "media_requests" edge of the DiscordUser entity.
func (du *DiscordUser) QueryMediaRequests() *MediaRequestQuery {
	return NewDiscordUserClient(du.config).QueryMediaRequests(du)
}

// Update returns a builder for updating this DiscordUser.
// Note that you need to call DiscordUser.Unwrap() before calling this method if this DiscordUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (du *DiscordUser) Update() *DiscordUserUpdateOne {
	return NewDiscordUserClient(du.config).UpdateOne(du)
}

// Unwrap unwraps the DiscordUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (du *DiscordUser) Unwrap() *DiscordUser {
	_tx, ok := du.config.driver.(*txDriver)
	if !ok {
		panic("ent: DiscordUser is not a transactional entity")
	}
	du.config.driver = _tx.drv
	return du
}

// String implements the fmt.Stringer.
func (du *DiscordUser) String() string {
	var builder strings.Builder
	builder.WriteString("DiscordUser(")
	builder.WriteString(fmt.Sprintf("id=%v, ", du.ID))
	builder.WriteString("discordid=")
	builder.WriteString(du.Discordid)
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(du.Username)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(du.Email)
	builder.WriteString(", ")
	builder.WriteString("discriminator=")
	builder.WriteString(du.Discriminator)
	builder.WriteByte(')')
	return builder.String()
}

// NamedDiscordMessages returns the DiscordMessages named value or an error if the edge was not
// loaded in eager-loading with this name.
func (du *DiscordUser) NamedDiscordMessages(name string) ([]*DiscordMessage, error) {
	if du.Edges.namedDiscordMessages == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := du.Edges.namedDiscordMessages[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (du *DiscordUser) appendNamedDiscordMessages(name string, edges ...*DiscordMessage) {
	if du.Edges.namedDiscordMessages == nil {
		du.Edges.namedDiscordMessages = make(map[string][]*DiscordMessage)
	}
	if len(edges) == 0 {
		du.Edges.namedDiscordMessages[name] = []*DiscordMessage{}
	} else {
		du.Edges.namedDiscordMessages[name] = append(du.Edges.namedDiscordMessages[name], edges...)
	}
}

// NamedMediaRequests returns the MediaRequests named value or an error if the edge was not
// loaded in eager-loading with this name.
func (du *DiscordUser) NamedMediaRequests(name string) ([]*MediaRequest, error) {
	if du.Edges.namedMediaRequests == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := du.Edges.namedMediaRequests[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (du *DiscordUser) appendNamedMediaRequests(name string, edges ...*MediaRequest) {
	if du.Edges.namedMediaRequests == nil {
		du.Edges.namedMediaRequests = make(map[string][]*MediaRequest)
	}
	if len(edges) == 0 {
		du.Edges.namedMediaRequests[name] = []*MediaRequest{}
	} else {
		du.Edges.namedMediaRequests[name] = append(du.Edges.namedMediaRequests[name], edges...)
	}
}

// DiscordUsers is a parsable slice of DiscordUser.
type DiscordUsers []*DiscordUser
