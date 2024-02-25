// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/h3mmy/bloopyboi/ent/book"
	"github.com/h3mmy/bloopyboi/ent/discorduser"
	"github.com/h3mmy/bloopyboi/ent/mediarequest"
	"github.com/h3mmy/bloopyboi/internal/models"
)

// MediaRequestCreate is the builder for creating a MediaRequest entity.
type MediaRequestCreate struct {
	config
	mutation *MediaRequestMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (mrc *MediaRequestCreate) SetCreateTime(t time.Time) *MediaRequestCreate {
	mrc.mutation.SetCreateTime(t)
	return mrc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (mrc *MediaRequestCreate) SetNillableCreateTime(t *time.Time) *MediaRequestCreate {
	if t != nil {
		mrc.SetCreateTime(*t)
	}
	return mrc
}

// SetUpdateTime sets the "update_time" field.
func (mrc *MediaRequestCreate) SetUpdateTime(t time.Time) *MediaRequestCreate {
	mrc.mutation.SetUpdateTime(t)
	return mrc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (mrc *MediaRequestCreate) SetNillableUpdateTime(t *time.Time) *MediaRequestCreate {
	if t != nil {
		mrc.SetUpdateTime(*t)
	}
	return mrc
}

// SetStatus sets the "status" field.
func (mrc *MediaRequestCreate) SetStatus(mrs models.MediaRequestStatus) *MediaRequestCreate {
	mrc.mutation.SetStatus(mrs)
	return mrc
}

// SetPriority sets the "priority" field.
func (mrc *MediaRequestCreate) SetPriority(i int) *MediaRequestCreate {
	mrc.mutation.SetPriority(i)
	return mrc
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (mrc *MediaRequestCreate) SetNillablePriority(i *int) *MediaRequestCreate {
	if i != nil {
		mrc.SetPriority(*i)
	}
	return mrc
}

// SetID sets the "id" field.
func (mrc *MediaRequestCreate) SetID(u uuid.UUID) *MediaRequestCreate {
	mrc.mutation.SetID(u)
	return mrc
}

// AddDiscordUserIDs adds the "discord_users" edge to the DiscordUser entity by IDs.
func (mrc *MediaRequestCreate) AddDiscordUserIDs(ids ...uuid.UUID) *MediaRequestCreate {
	mrc.mutation.AddDiscordUserIDs(ids...)
	return mrc
}

// AddDiscordUsers adds the "discord_users" edges to the DiscordUser entity.
func (mrc *MediaRequestCreate) AddDiscordUsers(d ...*DiscordUser) *MediaRequestCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return mrc.AddDiscordUserIDs(ids...)
}

// SetBookID sets the "book" edge to the Book entity by ID.
func (mrc *MediaRequestCreate) SetBookID(id uuid.UUID) *MediaRequestCreate {
	mrc.mutation.SetBookID(id)
	return mrc
}

// SetNillableBookID sets the "book" edge to the Book entity by ID if the given value is not nil.
func (mrc *MediaRequestCreate) SetNillableBookID(id *uuid.UUID) *MediaRequestCreate {
	if id != nil {
		mrc = mrc.SetBookID(*id)
	}
	return mrc
}

// SetBook sets the "book" edge to the Book entity.
func (mrc *MediaRequestCreate) SetBook(b *Book) *MediaRequestCreate {
	return mrc.SetBookID(b.ID)
}

// Mutation returns the MediaRequestMutation object of the builder.
func (mrc *MediaRequestCreate) Mutation() *MediaRequestMutation {
	return mrc.mutation
}

// Save creates the MediaRequest in the database.
func (mrc *MediaRequestCreate) Save(ctx context.Context) (*MediaRequest, error) {
	mrc.defaults()
	return withHooks(ctx, mrc.sqlSave, mrc.mutation, mrc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mrc *MediaRequestCreate) SaveX(ctx context.Context) *MediaRequest {
	v, err := mrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mrc *MediaRequestCreate) Exec(ctx context.Context) error {
	_, err := mrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mrc *MediaRequestCreate) ExecX(ctx context.Context) {
	if err := mrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mrc *MediaRequestCreate) defaults() {
	if _, ok := mrc.mutation.CreateTime(); !ok {
		v := mediarequest.DefaultCreateTime()
		mrc.mutation.SetCreateTime(v)
	}
	if _, ok := mrc.mutation.UpdateTime(); !ok {
		v := mediarequest.DefaultUpdateTime()
		mrc.mutation.SetUpdateTime(v)
	}
	if _, ok := mrc.mutation.Priority(); !ok {
		v := mediarequest.DefaultPriority
		mrc.mutation.SetPriority(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mrc *MediaRequestCreate) check() error {
	if _, ok := mrc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "MediaRequest.create_time"`)}
	}
	if _, ok := mrc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "MediaRequest.update_time"`)}
	}
	if _, ok := mrc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "MediaRequest.status"`)}
	}
	if v, ok := mrc.mutation.Status(); ok {
		if err := mediarequest.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "MediaRequest.status": %w`, err)}
		}
	}
	if _, ok := mrc.mutation.Priority(); !ok {
		return &ValidationError{Name: "priority", err: errors.New(`ent: missing required field "MediaRequest.priority"`)}
	}
	return nil
}

func (mrc *MediaRequestCreate) sqlSave(ctx context.Context) (*MediaRequest, error) {
	if err := mrc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	mrc.mutation.id = &_node.ID
	mrc.mutation.done = true
	return _node, nil
}

func (mrc *MediaRequestCreate) createSpec() (*MediaRequest, *sqlgraph.CreateSpec) {
	var (
		_node = &MediaRequest{config: mrc.config}
		_spec = sqlgraph.NewCreateSpec(mediarequest.Table, sqlgraph.NewFieldSpec(mediarequest.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = mrc.conflict
	if id, ok := mrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mrc.mutation.CreateTime(); ok {
		_spec.SetField(mediarequest.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := mrc.mutation.UpdateTime(); ok {
		_spec.SetField(mediarequest.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := mrc.mutation.Status(); ok {
		_spec.SetField(mediarequest.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := mrc.mutation.Priority(); ok {
		_spec.SetField(mediarequest.FieldPriority, field.TypeInt, value)
		_node.Priority = value
	}
	if nodes := mrc.mutation.DiscordUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   mediarequest.DiscordUsersTable,
			Columns: mediarequest.DiscordUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discorduser.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mrc.mutation.BookIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   mediarequest.BookTable,
			Columns: []string{mediarequest.BookColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.book_media_request = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MediaRequest.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MediaRequestUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (mrc *MediaRequestCreate) OnConflict(opts ...sql.ConflictOption) *MediaRequestUpsertOne {
	mrc.conflict = opts
	return &MediaRequestUpsertOne{
		create: mrc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MediaRequest.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mrc *MediaRequestCreate) OnConflictColumns(columns ...string) *MediaRequestUpsertOne {
	mrc.conflict = append(mrc.conflict, sql.ConflictColumns(columns...))
	return &MediaRequestUpsertOne{
		create: mrc,
	}
}

type (
	// MediaRequestUpsertOne is the builder for "upsert"-ing
	//  one MediaRequest node.
	MediaRequestUpsertOne struct {
		create *MediaRequestCreate
	}

	// MediaRequestUpsert is the "OnConflict" setter.
	MediaRequestUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *MediaRequestUpsert) SetUpdateTime(v time.Time) *MediaRequestUpsert {
	u.Set(mediarequest.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *MediaRequestUpsert) UpdateUpdateTime() *MediaRequestUpsert {
	u.SetExcluded(mediarequest.FieldUpdateTime)
	return u
}

// SetStatus sets the "status" field.
func (u *MediaRequestUpsert) SetStatus(v models.MediaRequestStatus) *MediaRequestUpsert {
	u.Set(mediarequest.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *MediaRequestUpsert) UpdateStatus() *MediaRequestUpsert {
	u.SetExcluded(mediarequest.FieldStatus)
	return u
}

// SetPriority sets the "priority" field.
func (u *MediaRequestUpsert) SetPriority(v int) *MediaRequestUpsert {
	u.Set(mediarequest.FieldPriority, v)
	return u
}

// UpdatePriority sets the "priority" field to the value that was provided on create.
func (u *MediaRequestUpsert) UpdatePriority() *MediaRequestUpsert {
	u.SetExcluded(mediarequest.FieldPriority)
	return u
}

// AddPriority adds v to the "priority" field.
func (u *MediaRequestUpsert) AddPriority(v int) *MediaRequestUpsert {
	u.Add(mediarequest.FieldPriority, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.MediaRequest.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mediarequest.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MediaRequestUpsertOne) UpdateNewValues() *MediaRequestUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(mediarequest.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(mediarequest.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MediaRequest.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MediaRequestUpsertOne) Ignore() *MediaRequestUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MediaRequestUpsertOne) DoNothing() *MediaRequestUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MediaRequestCreate.OnConflict
// documentation for more info.
func (u *MediaRequestUpsertOne) Update(set func(*MediaRequestUpsert)) *MediaRequestUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MediaRequestUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *MediaRequestUpsertOne) SetUpdateTime(v time.Time) *MediaRequestUpsertOne {
	return u.Update(func(s *MediaRequestUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *MediaRequestUpsertOne) UpdateUpdateTime() *MediaRequestUpsertOne {
	return u.Update(func(s *MediaRequestUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetStatus sets the "status" field.
func (u *MediaRequestUpsertOne) SetStatus(v models.MediaRequestStatus) *MediaRequestUpsertOne {
	return u.Update(func(s *MediaRequestUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *MediaRequestUpsertOne) UpdateStatus() *MediaRequestUpsertOne {
	return u.Update(func(s *MediaRequestUpsert) {
		s.UpdateStatus()
	})
}

// SetPriority sets the "priority" field.
func (u *MediaRequestUpsertOne) SetPriority(v int) *MediaRequestUpsertOne {
	return u.Update(func(s *MediaRequestUpsert) {
		s.SetPriority(v)
	})
}

// AddPriority adds v to the "priority" field.
func (u *MediaRequestUpsertOne) AddPriority(v int) *MediaRequestUpsertOne {
	return u.Update(func(s *MediaRequestUpsert) {
		s.AddPriority(v)
	})
}

// UpdatePriority sets the "priority" field to the value that was provided on create.
func (u *MediaRequestUpsertOne) UpdatePriority() *MediaRequestUpsertOne {
	return u.Update(func(s *MediaRequestUpsert) {
		s.UpdatePriority()
	})
}

// Exec executes the query.
func (u *MediaRequestUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MediaRequestCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MediaRequestUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MediaRequestUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: MediaRequestUpsertOne.ID is not supported by MySQL driver. Use MediaRequestUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MediaRequestUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MediaRequestCreateBulk is the builder for creating many MediaRequest entities in bulk.
type MediaRequestCreateBulk struct {
	config
	err      error
	builders []*MediaRequestCreate
	conflict []sql.ConflictOption
}

// Save creates the MediaRequest entities in the database.
func (mrcb *MediaRequestCreateBulk) Save(ctx context.Context) ([]*MediaRequest, error) {
	if mrcb.err != nil {
		return nil, mrcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mrcb.builders))
	nodes := make([]*MediaRequest, len(mrcb.builders))
	mutators := make([]Mutator, len(mrcb.builders))
	for i := range mrcb.builders {
		func(i int, root context.Context) {
			builder := mrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MediaRequestMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mrcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mrcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mrcb *MediaRequestCreateBulk) SaveX(ctx context.Context) []*MediaRequest {
	v, err := mrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mrcb *MediaRequestCreateBulk) Exec(ctx context.Context) error {
	_, err := mrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mrcb *MediaRequestCreateBulk) ExecX(ctx context.Context) {
	if err := mrcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MediaRequest.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MediaRequestUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (mrcb *MediaRequestCreateBulk) OnConflict(opts ...sql.ConflictOption) *MediaRequestUpsertBulk {
	mrcb.conflict = opts
	return &MediaRequestUpsertBulk{
		create: mrcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MediaRequest.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mrcb *MediaRequestCreateBulk) OnConflictColumns(columns ...string) *MediaRequestUpsertBulk {
	mrcb.conflict = append(mrcb.conflict, sql.ConflictColumns(columns...))
	return &MediaRequestUpsertBulk{
		create: mrcb,
	}
}

// MediaRequestUpsertBulk is the builder for "upsert"-ing
// a bulk of MediaRequest nodes.
type MediaRequestUpsertBulk struct {
	create *MediaRequestCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.MediaRequest.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(mediarequest.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MediaRequestUpsertBulk) UpdateNewValues() *MediaRequestUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(mediarequest.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(mediarequest.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MediaRequest.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MediaRequestUpsertBulk) Ignore() *MediaRequestUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MediaRequestUpsertBulk) DoNothing() *MediaRequestUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MediaRequestCreateBulk.OnConflict
// documentation for more info.
func (u *MediaRequestUpsertBulk) Update(set func(*MediaRequestUpsert)) *MediaRequestUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MediaRequestUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *MediaRequestUpsertBulk) SetUpdateTime(v time.Time) *MediaRequestUpsertBulk {
	return u.Update(func(s *MediaRequestUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *MediaRequestUpsertBulk) UpdateUpdateTime() *MediaRequestUpsertBulk {
	return u.Update(func(s *MediaRequestUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetStatus sets the "status" field.
func (u *MediaRequestUpsertBulk) SetStatus(v models.MediaRequestStatus) *MediaRequestUpsertBulk {
	return u.Update(func(s *MediaRequestUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *MediaRequestUpsertBulk) UpdateStatus() *MediaRequestUpsertBulk {
	return u.Update(func(s *MediaRequestUpsert) {
		s.UpdateStatus()
	})
}

// SetPriority sets the "priority" field.
func (u *MediaRequestUpsertBulk) SetPriority(v int) *MediaRequestUpsertBulk {
	return u.Update(func(s *MediaRequestUpsert) {
		s.SetPriority(v)
	})
}

// AddPriority adds v to the "priority" field.
func (u *MediaRequestUpsertBulk) AddPriority(v int) *MediaRequestUpsertBulk {
	return u.Update(func(s *MediaRequestUpsert) {
		s.AddPriority(v)
	})
}

// UpdatePriority sets the "priority" field to the value that was provided on create.
func (u *MediaRequestUpsertBulk) UpdatePriority() *MediaRequestUpsertBulk {
	return u.Update(func(s *MediaRequestUpsert) {
		s.UpdatePriority()
	})
}

// Exec executes the query.
func (u *MediaRequestUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MediaRequestCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MediaRequestCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MediaRequestUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
