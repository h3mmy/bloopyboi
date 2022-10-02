// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/ent/mediarequest"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/ent/predicate"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/ent/user"
)

// MediaRequestUpdate is the builder for updating MediaRequest entities.
type MediaRequestUpdate struct {
	config
	hooks    []Hook
	mutation *MediaRequestMutation
}

// Where appends a list predicates to the MediaRequestUpdate builder.
func (mru *MediaRequestUpdate) Where(ps ...predicate.MediaRequest) *MediaRequestUpdate {
	mru.mutation.Where(ps...)
	return mru
}

// SetStatus sets the "status" field.
func (mru *MediaRequestUpdate) SetStatus(s string) *MediaRequestUpdate {
	mru.mutation.SetStatus(s)
	return mru
}

// SetMediaType sets the "mediaType" field.
func (mru *MediaRequestUpdate) SetMediaType(mt mediarequest.MediaType) *MediaRequestUpdate {
	mru.mutation.SetMediaType(mt)
	return mru
}

// SetRequestId sets the "requestId" field.
func (mru *MediaRequestUpdate) SetRequestId(s string) *MediaRequestUpdate {
	mru.mutation.SetRequestId(s)
	return mru
}

// SetUserID sets the "user" edge to the User entity by ID.
func (mru *MediaRequestUpdate) SetUserID(id int) *MediaRequestUpdate {
	mru.mutation.SetUserID(id)
	return mru
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (mru *MediaRequestUpdate) SetNillableUserID(id *int) *MediaRequestUpdate {
	if id != nil {
		mru = mru.SetUserID(*id)
	}
	return mru
}

// SetUser sets the "user" edge to the User entity.
func (mru *MediaRequestUpdate) SetUser(u *User) *MediaRequestUpdate {
	return mru.SetUserID(u.ID)
}

// Mutation returns the MediaRequestMutation object of the builder.
func (mru *MediaRequestUpdate) Mutation() *MediaRequestMutation {
	return mru.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (mru *MediaRequestUpdate) ClearUser() *MediaRequestUpdate {
	mru.mutation.ClearUser()
	return mru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mru *MediaRequestUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mru.hooks) == 0 {
		if err = mru.check(); err != nil {
			return 0, err
		}
		affected, err = mru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MediaRequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mru.check(); err != nil {
				return 0, err
			}
			mru.mutation = mutation
			affected, err = mru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mru.hooks) - 1; i >= 0; i-- {
			if mru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mru *MediaRequestUpdate) SaveX(ctx context.Context) int {
	affected, err := mru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mru *MediaRequestUpdate) Exec(ctx context.Context) error {
	_, err := mru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mru *MediaRequestUpdate) ExecX(ctx context.Context) {
	if err := mru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mru *MediaRequestUpdate) check() error {
	if v, ok := mru.mutation.MediaType(); ok {
		if err := mediarequest.MediaTypeValidator(v); err != nil {
			return &ValidationError{Name: "mediaType", err: fmt.Errorf(`ent: validator failed for field "MediaRequest.mediaType": %w`, err)}
		}
	}
	return nil
}

func (mru *MediaRequestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mediarequest.Table,
			Columns: mediarequest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: mediarequest.FieldID,
			},
		},
	}
	if ps := mru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mru.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mediarequest.FieldStatus,
		})
	}
	if value, ok := mru.mutation.MediaType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: mediarequest.FieldMediaType,
		})
	}
	if value, ok := mru.mutation.RequestId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mediarequest.FieldRequestId,
		})
	}
	if mru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mediarequest.UserTable,
			Columns: []string{mediarequest.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mediarequest.UserTable,
			Columns: []string{mediarequest.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mediarequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MediaRequestUpdateOne is the builder for updating a single MediaRequest entity.
type MediaRequestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MediaRequestMutation
}

// SetStatus sets the "status" field.
func (mruo *MediaRequestUpdateOne) SetStatus(s string) *MediaRequestUpdateOne {
	mruo.mutation.SetStatus(s)
	return mruo
}

// SetMediaType sets the "mediaType" field.
func (mruo *MediaRequestUpdateOne) SetMediaType(mt mediarequest.MediaType) *MediaRequestUpdateOne {
	mruo.mutation.SetMediaType(mt)
	return mruo
}

// SetRequestId sets the "requestId" field.
func (mruo *MediaRequestUpdateOne) SetRequestId(s string) *MediaRequestUpdateOne {
	mruo.mutation.SetRequestId(s)
	return mruo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (mruo *MediaRequestUpdateOne) SetUserID(id int) *MediaRequestUpdateOne {
	mruo.mutation.SetUserID(id)
	return mruo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (mruo *MediaRequestUpdateOne) SetNillableUserID(id *int) *MediaRequestUpdateOne {
	if id != nil {
		mruo = mruo.SetUserID(*id)
	}
	return mruo
}

// SetUser sets the "user" edge to the User entity.
func (mruo *MediaRequestUpdateOne) SetUser(u *User) *MediaRequestUpdateOne {
	return mruo.SetUserID(u.ID)
}

// Mutation returns the MediaRequestMutation object of the builder.
func (mruo *MediaRequestUpdateOne) Mutation() *MediaRequestMutation {
	return mruo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (mruo *MediaRequestUpdateOne) ClearUser() *MediaRequestUpdateOne {
	mruo.mutation.ClearUser()
	return mruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mruo *MediaRequestUpdateOne) Select(field string, fields ...string) *MediaRequestUpdateOne {
	mruo.fields = append([]string{field}, fields...)
	return mruo
}

// Save executes the query and returns the updated MediaRequest entity.
func (mruo *MediaRequestUpdateOne) Save(ctx context.Context) (*MediaRequest, error) {
	var (
		err  error
		node *MediaRequest
	)
	if len(mruo.hooks) == 0 {
		if err = mruo.check(); err != nil {
			return nil, err
		}
		node, err = mruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MediaRequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mruo.check(); err != nil {
				return nil, err
			}
			mruo.mutation = mutation
			node, err = mruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mruo.hooks) - 1; i >= 0; i-- {
			if mruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*MediaRequest)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MediaRequestMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mruo *MediaRequestUpdateOne) SaveX(ctx context.Context) *MediaRequest {
	node, err := mruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mruo *MediaRequestUpdateOne) Exec(ctx context.Context) error {
	_, err := mruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mruo *MediaRequestUpdateOne) ExecX(ctx context.Context) {
	if err := mruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mruo *MediaRequestUpdateOne) check() error {
	if v, ok := mruo.mutation.MediaType(); ok {
		if err := mediarequest.MediaTypeValidator(v); err != nil {
			return &ValidationError{Name: "mediaType", err: fmt.Errorf(`ent: validator failed for field "MediaRequest.mediaType": %w`, err)}
		}
	}
	return nil
}

func (mruo *MediaRequestUpdateOne) sqlSave(ctx context.Context) (_node *MediaRequest, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mediarequest.Table,
			Columns: mediarequest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: mediarequest.FieldID,
			},
		},
	}
	id, ok := mruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MediaRequest.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mediarequest.FieldID)
		for _, f := range fields {
			if !mediarequest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mediarequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mruo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mediarequest.FieldStatus,
		})
	}
	if value, ok := mruo.mutation.MediaType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: mediarequest.FieldMediaType,
		})
	}
	if value, ok := mruo.mutation.RequestId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mediarequest.FieldRequestId,
		})
	}
	if mruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mediarequest.UserTable,
			Columns: []string{mediarequest.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   mediarequest.UserTable,
			Columns: []string{mediarequest.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MediaRequest{config: mruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mediarequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
