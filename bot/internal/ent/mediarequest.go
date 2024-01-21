// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/h3mmy/bloopyboi/bot/internal/ent/mediarequest"
	"github.com/h3mmy/bloopyboi/bot/internal/ent/user"
)

// MediaRequest is the model entity for the MediaRequest schema.
type MediaRequest struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// MediaType holds the value of the "mediaType" field.
	MediaType mediarequest.MediaType `json:"mediaType,omitempty"`
	// RequestId holds the value of the "requestId" field.
	RequestId string `json:"requestId,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MediaRequestQuery when eager-loading is set.
	Edges              MediaRequestEdges `json:"edges"`
	user_media_request *int
}

// MediaRequestEdges holds the relations/edges for other nodes in the graph.
type MediaRequestEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MediaRequestEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MediaRequest) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case mediarequest.FieldID:
			values[i] = new(sql.NullInt64)
		case mediarequest.FieldStatus, mediarequest.FieldMediaType, mediarequest.FieldRequestId:
			values[i] = new(sql.NullString)
		case mediarequest.ForeignKeys[0]: // user_media_request
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MediaRequest", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MediaRequest fields.
func (mr *MediaRequest) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mediarequest.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mr.ID = int(value.Int64)
		case mediarequest.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				mr.Status = value.String
			}
		case mediarequest.FieldMediaType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mediaType", values[i])
			} else if value.Valid {
				mr.MediaType = mediarequest.MediaType(value.String)
			}
		case mediarequest.FieldRequestId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field requestId", values[i])
			} else if value.Valid {
				mr.RequestId = value.String
			}
		case mediarequest.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_media_request", value)
			} else if value.Valid {
				mr.user_media_request = new(int)
				*mr.user_media_request = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the MediaRequest entity.
func (mr *MediaRequest) QueryUser() *UserQuery {
	return (&MediaRequestClient{config: mr.config}).QueryUser(mr)
}

// Update returns a builder for updating this MediaRequest.
// Note that you need to call MediaRequest.Unwrap() before calling this method if this MediaRequest
// was returned from a transaction, and the transaction was committed or rolled back.
func (mr *MediaRequest) Update() *MediaRequestUpdateOne {
	return (&MediaRequestClient{config: mr.config}).UpdateOne(mr)
}

// Unwrap unwraps the MediaRequest entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mr *MediaRequest) Unwrap() *MediaRequest {
	_tx, ok := mr.config.driver.(*txDriver)
	if !ok {
		panic("ent: MediaRequest is not a transactional entity")
	}
	mr.config.driver = _tx.drv
	return mr
}

// String implements the fmt.Stringer.
func (mr *MediaRequest) String() string {
	var builder strings.Builder
	builder.WriteString("MediaRequest(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mr.ID))
	builder.WriteString("status=")
	builder.WriteString(mr.Status)
	builder.WriteString(", ")
	builder.WriteString("mediaType=")
	builder.WriteString(fmt.Sprintf("%v", mr.MediaType))
	builder.WriteString(", ")
	builder.WriteString("requestId=")
	builder.WriteString(mr.RequestId)
	builder.WriteByte(')')
	return builder.String()
}

// MediaRequests is a parsable slice of MediaRequest.
type MediaRequests []*MediaRequest

func (mr MediaRequests) config(cfg config) {
	for _i := range mr {
		mr[_i].config = cfg
	}
}
