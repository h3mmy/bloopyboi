// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/h3mmy/bloopyboi/ent/discordmessage"
	"github.com/h3mmy/bloopyboi/ent/discorduser"
	"github.com/h3mmy/bloopyboi/ent/mediarequest"
	"github.com/h3mmy/bloopyboi/ent/predicate"
)

// DiscordUserUpdate is the builder for updating DiscordUser entities.
type DiscordUserUpdate struct {
	config
	hooks    []Hook
	mutation *DiscordUserMutation
}

// Where appends a list predicates to the DiscordUserUpdate builder.
func (duu *DiscordUserUpdate) Where(ps ...predicate.DiscordUser) *DiscordUserUpdate {
	duu.mutation.Where(ps...)
	return duu
}

// SetDiscordid sets the "discordid" field.
func (duu *DiscordUserUpdate) SetDiscordid(s string) *DiscordUserUpdate {
	duu.mutation.SetDiscordid(s)
	return duu
}

// SetNillableDiscordid sets the "discordid" field if the given value is not nil.
func (duu *DiscordUserUpdate) SetNillableDiscordid(s *string) *DiscordUserUpdate {
	if s != nil {
		duu.SetDiscordid(*s)
	}
	return duu
}

// SetUsername sets the "username" field.
func (duu *DiscordUserUpdate) SetUsername(s string) *DiscordUserUpdate {
	duu.mutation.SetUsername(s)
	return duu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (duu *DiscordUserUpdate) SetNillableUsername(s *string) *DiscordUserUpdate {
	if s != nil {
		duu.SetUsername(*s)
	}
	return duu
}

// SetEmail sets the "email" field.
func (duu *DiscordUserUpdate) SetEmail(s string) *DiscordUserUpdate {
	duu.mutation.SetEmail(s)
	return duu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (duu *DiscordUserUpdate) SetNillableEmail(s *string) *DiscordUserUpdate {
	if s != nil {
		duu.SetEmail(*s)
	}
	return duu
}

// ClearEmail clears the value of the "email" field.
func (duu *DiscordUserUpdate) ClearEmail() *DiscordUserUpdate {
	duu.mutation.ClearEmail()
	return duu
}

// SetDiscriminator sets the "discriminator" field.
func (duu *DiscordUserUpdate) SetDiscriminator(s string) *DiscordUserUpdate {
	duu.mutation.SetDiscriminator(s)
	return duu
}

// SetNillableDiscriminator sets the "discriminator" field if the given value is not nil.
func (duu *DiscordUserUpdate) SetNillableDiscriminator(s *string) *DiscordUserUpdate {
	if s != nil {
		duu.SetDiscriminator(*s)
	}
	return duu
}

// ClearDiscriminator clears the value of the "discriminator" field.
func (duu *DiscordUserUpdate) ClearDiscriminator() *DiscordUserUpdate {
	duu.mutation.ClearDiscriminator()
	return duu
}

// AddDiscordMessageIDs adds the "discord_messages" edge to the DiscordMessage entity by IDs.
func (duu *DiscordUserUpdate) AddDiscordMessageIDs(ids ...string) *DiscordUserUpdate {
	duu.mutation.AddDiscordMessageIDs(ids...)
	return duu
}

// AddDiscordMessages adds the "discord_messages" edges to the DiscordMessage entity.
func (duu *DiscordUserUpdate) AddDiscordMessages(d ...*DiscordMessage) *DiscordUserUpdate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duu.AddDiscordMessageIDs(ids...)
}

// AddMediaRequestIDs adds the "media_requests" edge to the MediaRequest entity by IDs.
func (duu *DiscordUserUpdate) AddMediaRequestIDs(ids ...uuid.UUID) *DiscordUserUpdate {
	duu.mutation.AddMediaRequestIDs(ids...)
	return duu
}

// AddMediaRequests adds the "media_requests" edges to the MediaRequest entity.
func (duu *DiscordUserUpdate) AddMediaRequests(m ...*MediaRequest) *DiscordUserUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return duu.AddMediaRequestIDs(ids...)
}

// Mutation returns the DiscordUserMutation object of the builder.
func (duu *DiscordUserUpdate) Mutation() *DiscordUserMutation {
	return duu.mutation
}

// ClearDiscordMessages clears all "discord_messages" edges to the DiscordMessage entity.
func (duu *DiscordUserUpdate) ClearDiscordMessages() *DiscordUserUpdate {
	duu.mutation.ClearDiscordMessages()
	return duu
}

// RemoveDiscordMessageIDs removes the "discord_messages" edge to DiscordMessage entities by IDs.
func (duu *DiscordUserUpdate) RemoveDiscordMessageIDs(ids ...string) *DiscordUserUpdate {
	duu.mutation.RemoveDiscordMessageIDs(ids...)
	return duu
}

// RemoveDiscordMessages removes "discord_messages" edges to DiscordMessage entities.
func (duu *DiscordUserUpdate) RemoveDiscordMessages(d ...*DiscordMessage) *DiscordUserUpdate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duu.RemoveDiscordMessageIDs(ids...)
}

// ClearMediaRequests clears all "media_requests" edges to the MediaRequest entity.
func (duu *DiscordUserUpdate) ClearMediaRequests() *DiscordUserUpdate {
	duu.mutation.ClearMediaRequests()
	return duu
}

// RemoveMediaRequestIDs removes the "media_requests" edge to MediaRequest entities by IDs.
func (duu *DiscordUserUpdate) RemoveMediaRequestIDs(ids ...uuid.UUID) *DiscordUserUpdate {
	duu.mutation.RemoveMediaRequestIDs(ids...)
	return duu
}

// RemoveMediaRequests removes "media_requests" edges to MediaRequest entities.
func (duu *DiscordUserUpdate) RemoveMediaRequests(m ...*MediaRequest) *DiscordUserUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return duu.RemoveMediaRequestIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (duu *DiscordUserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, duu.sqlSave, duu.mutation, duu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duu *DiscordUserUpdate) SaveX(ctx context.Context) int {
	affected, err := duu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (duu *DiscordUserUpdate) Exec(ctx context.Context) error {
	_, err := duu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duu *DiscordUserUpdate) ExecX(ctx context.Context) {
	if err := duu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duu *DiscordUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(discorduser.Table, discorduser.Columns, sqlgraph.NewFieldSpec(discorduser.FieldID, field.TypeUUID))
	if ps := duu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duu.mutation.Discordid(); ok {
		_spec.SetField(discorduser.FieldDiscordid, field.TypeString, value)
	}
	if value, ok := duu.mutation.Username(); ok {
		_spec.SetField(discorduser.FieldUsername, field.TypeString, value)
	}
	if value, ok := duu.mutation.Email(); ok {
		_spec.SetField(discorduser.FieldEmail, field.TypeString, value)
	}
	if duu.mutation.EmailCleared() {
		_spec.ClearField(discorduser.FieldEmail, field.TypeString)
	}
	if value, ok := duu.mutation.Discriminator(); ok {
		_spec.SetField(discorduser.FieldDiscriminator, field.TypeString, value)
	}
	if duu.mutation.DiscriminatorCleared() {
		_spec.ClearField(discorduser.FieldDiscriminator, field.TypeString)
	}
	if duu.mutation.DiscordMessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   discorduser.DiscordMessagesTable,
			Columns: discorduser.DiscordMessagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordmessage.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duu.mutation.RemovedDiscordMessagesIDs(); len(nodes) > 0 && !duu.mutation.DiscordMessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   discorduser.DiscordMessagesTable,
			Columns: discorduser.DiscordMessagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordmessage.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duu.mutation.DiscordMessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   discorduser.DiscordMessagesTable,
			Columns: discorduser.DiscordMessagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordmessage.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duu.mutation.MediaRequestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   discorduser.MediaRequestsTable,
			Columns: []string{discorduser.MediaRequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mediarequest.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duu.mutation.RemovedMediaRequestsIDs(); len(nodes) > 0 && !duu.mutation.MediaRequestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   discorduser.MediaRequestsTable,
			Columns: []string{discorduser.MediaRequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mediarequest.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duu.mutation.MediaRequestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   discorduser.MediaRequestsTable,
			Columns: []string{discorduser.MediaRequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mediarequest.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, duu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{discorduser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	duu.mutation.done = true
	return n, nil
}

// DiscordUserUpdateOne is the builder for updating a single DiscordUser entity.
type DiscordUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DiscordUserMutation
}

// SetDiscordid sets the "discordid" field.
func (duuo *DiscordUserUpdateOne) SetDiscordid(s string) *DiscordUserUpdateOne {
	duuo.mutation.SetDiscordid(s)
	return duuo
}

// SetNillableDiscordid sets the "discordid" field if the given value is not nil.
func (duuo *DiscordUserUpdateOne) SetNillableDiscordid(s *string) *DiscordUserUpdateOne {
	if s != nil {
		duuo.SetDiscordid(*s)
	}
	return duuo
}

// SetUsername sets the "username" field.
func (duuo *DiscordUserUpdateOne) SetUsername(s string) *DiscordUserUpdateOne {
	duuo.mutation.SetUsername(s)
	return duuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (duuo *DiscordUserUpdateOne) SetNillableUsername(s *string) *DiscordUserUpdateOne {
	if s != nil {
		duuo.SetUsername(*s)
	}
	return duuo
}

// SetEmail sets the "email" field.
func (duuo *DiscordUserUpdateOne) SetEmail(s string) *DiscordUserUpdateOne {
	duuo.mutation.SetEmail(s)
	return duuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (duuo *DiscordUserUpdateOne) SetNillableEmail(s *string) *DiscordUserUpdateOne {
	if s != nil {
		duuo.SetEmail(*s)
	}
	return duuo
}

// ClearEmail clears the value of the "email" field.
func (duuo *DiscordUserUpdateOne) ClearEmail() *DiscordUserUpdateOne {
	duuo.mutation.ClearEmail()
	return duuo
}

// SetDiscriminator sets the "discriminator" field.
func (duuo *DiscordUserUpdateOne) SetDiscriminator(s string) *DiscordUserUpdateOne {
	duuo.mutation.SetDiscriminator(s)
	return duuo
}

// SetNillableDiscriminator sets the "discriminator" field if the given value is not nil.
func (duuo *DiscordUserUpdateOne) SetNillableDiscriminator(s *string) *DiscordUserUpdateOne {
	if s != nil {
		duuo.SetDiscriminator(*s)
	}
	return duuo
}

// ClearDiscriminator clears the value of the "discriminator" field.
func (duuo *DiscordUserUpdateOne) ClearDiscriminator() *DiscordUserUpdateOne {
	duuo.mutation.ClearDiscriminator()
	return duuo
}

// AddDiscordMessageIDs adds the "discord_messages" edge to the DiscordMessage entity by IDs.
func (duuo *DiscordUserUpdateOne) AddDiscordMessageIDs(ids ...string) *DiscordUserUpdateOne {
	duuo.mutation.AddDiscordMessageIDs(ids...)
	return duuo
}

// AddDiscordMessages adds the "discord_messages" edges to the DiscordMessage entity.
func (duuo *DiscordUserUpdateOne) AddDiscordMessages(d ...*DiscordMessage) *DiscordUserUpdateOne {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duuo.AddDiscordMessageIDs(ids...)
}

// AddMediaRequestIDs adds the "media_requests" edge to the MediaRequest entity by IDs.
func (duuo *DiscordUserUpdateOne) AddMediaRequestIDs(ids ...uuid.UUID) *DiscordUserUpdateOne {
	duuo.mutation.AddMediaRequestIDs(ids...)
	return duuo
}

// AddMediaRequests adds the "media_requests" edges to the MediaRequest entity.
func (duuo *DiscordUserUpdateOne) AddMediaRequests(m ...*MediaRequest) *DiscordUserUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return duuo.AddMediaRequestIDs(ids...)
}

// Mutation returns the DiscordUserMutation object of the builder.
func (duuo *DiscordUserUpdateOne) Mutation() *DiscordUserMutation {
	return duuo.mutation
}

// ClearDiscordMessages clears all "discord_messages" edges to the DiscordMessage entity.
func (duuo *DiscordUserUpdateOne) ClearDiscordMessages() *DiscordUserUpdateOne {
	duuo.mutation.ClearDiscordMessages()
	return duuo
}

// RemoveDiscordMessageIDs removes the "discord_messages" edge to DiscordMessage entities by IDs.
func (duuo *DiscordUserUpdateOne) RemoveDiscordMessageIDs(ids ...string) *DiscordUserUpdateOne {
	duuo.mutation.RemoveDiscordMessageIDs(ids...)
	return duuo
}

// RemoveDiscordMessages removes "discord_messages" edges to DiscordMessage entities.
func (duuo *DiscordUserUpdateOne) RemoveDiscordMessages(d ...*DiscordMessage) *DiscordUserUpdateOne {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duuo.RemoveDiscordMessageIDs(ids...)
}

// ClearMediaRequests clears all "media_requests" edges to the MediaRequest entity.
func (duuo *DiscordUserUpdateOne) ClearMediaRequests() *DiscordUserUpdateOne {
	duuo.mutation.ClearMediaRequests()
	return duuo
}

// RemoveMediaRequestIDs removes the "media_requests" edge to MediaRequest entities by IDs.
func (duuo *DiscordUserUpdateOne) RemoveMediaRequestIDs(ids ...uuid.UUID) *DiscordUserUpdateOne {
	duuo.mutation.RemoveMediaRequestIDs(ids...)
	return duuo
}

// RemoveMediaRequests removes "media_requests" edges to MediaRequest entities.
func (duuo *DiscordUserUpdateOne) RemoveMediaRequests(m ...*MediaRequest) *DiscordUserUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return duuo.RemoveMediaRequestIDs(ids...)
}

// Where appends a list predicates to the DiscordUserUpdate builder.
func (duuo *DiscordUserUpdateOne) Where(ps ...predicate.DiscordUser) *DiscordUserUpdateOne {
	duuo.mutation.Where(ps...)
	return duuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duuo *DiscordUserUpdateOne) Select(field string, fields ...string) *DiscordUserUpdateOne {
	duuo.fields = append([]string{field}, fields...)
	return duuo
}

// Save executes the query and returns the updated DiscordUser entity.
func (duuo *DiscordUserUpdateOne) Save(ctx context.Context) (*DiscordUser, error) {
	return withHooks(ctx, duuo.sqlSave, duuo.mutation, duuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duuo *DiscordUserUpdateOne) SaveX(ctx context.Context) *DiscordUser {
	node, err := duuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duuo *DiscordUserUpdateOne) Exec(ctx context.Context) error {
	_, err := duuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duuo *DiscordUserUpdateOne) ExecX(ctx context.Context) {
	if err := duuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duuo *DiscordUserUpdateOne) sqlSave(ctx context.Context) (_node *DiscordUser, err error) {
	_spec := sqlgraph.NewUpdateSpec(discorduser.Table, discorduser.Columns, sqlgraph.NewFieldSpec(discorduser.FieldID, field.TypeUUID))
	id, ok := duuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DiscordUser.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, discorduser.FieldID)
		for _, f := range fields {
			if !discorduser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != discorduser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duuo.mutation.Discordid(); ok {
		_spec.SetField(discorduser.FieldDiscordid, field.TypeString, value)
	}
	if value, ok := duuo.mutation.Username(); ok {
		_spec.SetField(discorduser.FieldUsername, field.TypeString, value)
	}
	if value, ok := duuo.mutation.Email(); ok {
		_spec.SetField(discorduser.FieldEmail, field.TypeString, value)
	}
	if duuo.mutation.EmailCleared() {
		_spec.ClearField(discorduser.FieldEmail, field.TypeString)
	}
	if value, ok := duuo.mutation.Discriminator(); ok {
		_spec.SetField(discorduser.FieldDiscriminator, field.TypeString, value)
	}
	if duuo.mutation.DiscriminatorCleared() {
		_spec.ClearField(discorduser.FieldDiscriminator, field.TypeString)
	}
	if duuo.mutation.DiscordMessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   discorduser.DiscordMessagesTable,
			Columns: discorduser.DiscordMessagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordmessage.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duuo.mutation.RemovedDiscordMessagesIDs(); len(nodes) > 0 && !duuo.mutation.DiscordMessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   discorduser.DiscordMessagesTable,
			Columns: discorduser.DiscordMessagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordmessage.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duuo.mutation.DiscordMessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   discorduser.DiscordMessagesTable,
			Columns: discorduser.DiscordMessagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordmessage.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duuo.mutation.MediaRequestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   discorduser.MediaRequestsTable,
			Columns: []string{discorduser.MediaRequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mediarequest.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duuo.mutation.RemovedMediaRequestsIDs(); len(nodes) > 0 && !duuo.mutation.MediaRequestsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   discorduser.MediaRequestsTable,
			Columns: []string{discorduser.MediaRequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mediarequest.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duuo.mutation.MediaRequestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   discorduser.MediaRequestsTable,
			Columns: []string{discorduser.MediaRequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mediarequest.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DiscordUser{config: duuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{discorduser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duuo.mutation.done = true
	return _node, nil
}
