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
	"github.com/h3mmy/bloopyboi/ent/discordchannel"
	"github.com/h3mmy/bloopyboi/ent/discordguild"
	"github.com/h3mmy/bloopyboi/ent/discordmessage"
	"github.com/h3mmy/bloopyboi/internal/discord"
)

// DiscordChannelCreate is the builder for creating a DiscordChannel entity.
type DiscordChannelCreate struct {
	config
	mutation *DiscordChannelMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (dcc *DiscordChannelCreate) SetCreateTime(t time.Time) *DiscordChannelCreate {
	dcc.mutation.SetCreateTime(t)
	return dcc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (dcc *DiscordChannelCreate) SetNillableCreateTime(t *time.Time) *DiscordChannelCreate {
	if t != nil {
		dcc.SetCreateTime(*t)
	}
	return dcc
}

// SetUpdateTime sets the "update_time" field.
func (dcc *DiscordChannelCreate) SetUpdateTime(t time.Time) *DiscordChannelCreate {
	dcc.mutation.SetUpdateTime(t)
	return dcc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (dcc *DiscordChannelCreate) SetNillableUpdateTime(t *time.Time) *DiscordChannelCreate {
	if t != nil {
		dcc.SetUpdateTime(*t)
	}
	return dcc
}

// SetDiscordid sets the "discordid" field.
func (dcc *DiscordChannelCreate) SetDiscordid(s string) *DiscordChannelCreate {
	dcc.mutation.SetDiscordid(s)
	return dcc
}

// SetName sets the "name" field.
func (dcc *DiscordChannelCreate) SetName(s string) *DiscordChannelCreate {
	dcc.mutation.SetName(s)
	return dcc
}

// SetType sets the "type" field.
func (dcc *DiscordChannelCreate) SetType(dt discord.ChannelType) *DiscordChannelCreate {
	dcc.mutation.SetType(dt)
	return dcc
}

// SetNsfw sets the "nsfw" field.
func (dcc *DiscordChannelCreate) SetNsfw(b bool) *DiscordChannelCreate {
	dcc.mutation.SetNsfw(b)
	return dcc
}

// SetNillableNsfw sets the "nsfw" field if the given value is not nil.
func (dcc *DiscordChannelCreate) SetNillableNsfw(b *bool) *DiscordChannelCreate {
	if b != nil {
		dcc.SetNsfw(*b)
	}
	return dcc
}

// SetFlags sets the "flags" field.
func (dcc *DiscordChannelCreate) SetFlags(i int) *DiscordChannelCreate {
	dcc.mutation.SetFlags(i)
	return dcc
}

// SetNillableFlags sets the "flags" field if the given value is not nil.
func (dcc *DiscordChannelCreate) SetNillableFlags(i *int) *DiscordChannelCreate {
	if i != nil {
		dcc.SetFlags(*i)
	}
	return dcc
}

// SetID sets the "id" field.
func (dcc *DiscordChannelCreate) SetID(u uuid.UUID) *DiscordChannelCreate {
	dcc.mutation.SetID(u)
	return dcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dcc *DiscordChannelCreate) SetNillableID(u *uuid.UUID) *DiscordChannelCreate {
	if u != nil {
		dcc.SetID(*u)
	}
	return dcc
}

// AddDiscordGuildIDs adds the "discord_guild" edge to the DiscordGuild entity by IDs.
func (dcc *DiscordChannelCreate) AddDiscordGuildIDs(ids ...uuid.UUID) *DiscordChannelCreate {
	dcc.mutation.AddDiscordGuildIDs(ids...)
	return dcc
}

// AddDiscordGuild adds the "discord_guild" edges to the DiscordGuild entity.
func (dcc *DiscordChannelCreate) AddDiscordGuild(d ...*DiscordGuild) *DiscordChannelCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dcc.AddDiscordGuildIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the DiscordMessage entity by IDs.
func (dcc *DiscordChannelCreate) AddMessageIDs(ids ...uuid.UUID) *DiscordChannelCreate {
	dcc.mutation.AddMessageIDs(ids...)
	return dcc
}

// AddMessages adds the "messages" edges to the DiscordMessage entity.
func (dcc *DiscordChannelCreate) AddMessages(d ...*DiscordMessage) *DiscordChannelCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dcc.AddMessageIDs(ids...)
}

// Mutation returns the DiscordChannelMutation object of the builder.
func (dcc *DiscordChannelCreate) Mutation() *DiscordChannelMutation {
	return dcc.mutation
}

// Save creates the DiscordChannel in the database.
func (dcc *DiscordChannelCreate) Save(ctx context.Context) (*DiscordChannel, error) {
	dcc.defaults()
	return withHooks(ctx, dcc.sqlSave, dcc.mutation, dcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dcc *DiscordChannelCreate) SaveX(ctx context.Context) *DiscordChannel {
	v, err := dcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcc *DiscordChannelCreate) Exec(ctx context.Context) error {
	_, err := dcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcc *DiscordChannelCreate) ExecX(ctx context.Context) {
	if err := dcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dcc *DiscordChannelCreate) defaults() {
	if _, ok := dcc.mutation.CreateTime(); !ok {
		v := discordchannel.DefaultCreateTime()
		dcc.mutation.SetCreateTime(v)
	}
	if _, ok := dcc.mutation.UpdateTime(); !ok {
		v := discordchannel.DefaultUpdateTime()
		dcc.mutation.SetUpdateTime(v)
	}
	if _, ok := dcc.mutation.Nsfw(); !ok {
		v := discordchannel.DefaultNsfw
		dcc.mutation.SetNsfw(v)
	}
	if _, ok := dcc.mutation.ID(); !ok {
		v := discordchannel.DefaultID()
		dcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dcc *DiscordChannelCreate) check() error {
	if _, ok := dcc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "DiscordChannel.create_time"`)}
	}
	if _, ok := dcc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "DiscordChannel.update_time"`)}
	}
	if _, ok := dcc.mutation.Discordid(); !ok {
		return &ValidationError{Name: "discordid", err: errors.New(`ent: missing required field "DiscordChannel.discordid"`)}
	}
	if _, ok := dcc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "DiscordChannel.name"`)}
	}
	if _, ok := dcc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "DiscordChannel.type"`)}
	}
	if _, ok := dcc.mutation.Nsfw(); !ok {
		return &ValidationError{Name: "nsfw", err: errors.New(`ent: missing required field "DiscordChannel.nsfw"`)}
	}
	return nil
}

func (dcc *DiscordChannelCreate) sqlSave(ctx context.Context) (*DiscordChannel, error) {
	if err := dcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dcc.driver, _spec); err != nil {
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
	dcc.mutation.id = &_node.ID
	dcc.mutation.done = true
	return _node, nil
}

func (dcc *DiscordChannelCreate) createSpec() (*DiscordChannel, *sqlgraph.CreateSpec) {
	var (
		_node = &DiscordChannel{config: dcc.config}
		_spec = sqlgraph.NewCreateSpec(discordchannel.Table, sqlgraph.NewFieldSpec(discordchannel.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = dcc.conflict
	if id, ok := dcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dcc.mutation.CreateTime(); ok {
		_spec.SetField(discordchannel.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := dcc.mutation.UpdateTime(); ok {
		_spec.SetField(discordchannel.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := dcc.mutation.Discordid(); ok {
		_spec.SetField(discordchannel.FieldDiscordid, field.TypeString, value)
		_node.Discordid = value
	}
	if value, ok := dcc.mutation.Name(); ok {
		_spec.SetField(discordchannel.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dcc.mutation.GetType(); ok {
		_spec.SetField(discordchannel.FieldType, field.TypeInt, value)
		_node.Type = value
	}
	if value, ok := dcc.mutation.Nsfw(); ok {
		_spec.SetField(discordchannel.FieldNsfw, field.TypeBool, value)
		_node.Nsfw = value
	}
	if value, ok := dcc.mutation.Flags(); ok {
		_spec.SetField(discordchannel.FieldFlags, field.TypeInt, value)
		_node.Flags = value
	}
	if nodes := dcc.mutation.DiscordGuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   discordchannel.DiscordGuildTable,
			Columns: discordchannel.DiscordGuildPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordguild.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dcc.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   discordchannel.MessagesTable,
			Columns: []string{discordchannel.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordmessage.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DiscordChannel.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DiscordChannelUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (dcc *DiscordChannelCreate) OnConflict(opts ...sql.ConflictOption) *DiscordChannelUpsertOne {
	dcc.conflict = opts
	return &DiscordChannelUpsertOne{
		create: dcc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DiscordChannel.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dcc *DiscordChannelCreate) OnConflictColumns(columns ...string) *DiscordChannelUpsertOne {
	dcc.conflict = append(dcc.conflict, sql.ConflictColumns(columns...))
	return &DiscordChannelUpsertOne{
		create: dcc,
	}
}

type (
	// DiscordChannelUpsertOne is the builder for "upsert"-ing
	//  one DiscordChannel node.
	DiscordChannelUpsertOne struct {
		create *DiscordChannelCreate
	}

	// DiscordChannelUpsert is the "OnConflict" setter.
	DiscordChannelUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *DiscordChannelUpsert) SetUpdateTime(v time.Time) *DiscordChannelUpsert {
	u.Set(discordchannel.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *DiscordChannelUpsert) UpdateUpdateTime() *DiscordChannelUpsert {
	u.SetExcluded(discordchannel.FieldUpdateTime)
	return u
}

// SetDiscordid sets the "discordid" field.
func (u *DiscordChannelUpsert) SetDiscordid(v string) *DiscordChannelUpsert {
	u.Set(discordchannel.FieldDiscordid, v)
	return u
}

// UpdateDiscordid sets the "discordid" field to the value that was provided on create.
func (u *DiscordChannelUpsert) UpdateDiscordid() *DiscordChannelUpsert {
	u.SetExcluded(discordchannel.FieldDiscordid)
	return u
}

// SetName sets the "name" field.
func (u *DiscordChannelUpsert) SetName(v string) *DiscordChannelUpsert {
	u.Set(discordchannel.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DiscordChannelUpsert) UpdateName() *DiscordChannelUpsert {
	u.SetExcluded(discordchannel.FieldName)
	return u
}

// SetType sets the "type" field.
func (u *DiscordChannelUpsert) SetType(v discord.ChannelType) *DiscordChannelUpsert {
	u.Set(discordchannel.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *DiscordChannelUpsert) UpdateType() *DiscordChannelUpsert {
	u.SetExcluded(discordchannel.FieldType)
	return u
}

// AddType adds v to the "type" field.
func (u *DiscordChannelUpsert) AddType(v discord.ChannelType) *DiscordChannelUpsert {
	u.Add(discordchannel.FieldType, v)
	return u
}

// SetNsfw sets the "nsfw" field.
func (u *DiscordChannelUpsert) SetNsfw(v bool) *DiscordChannelUpsert {
	u.Set(discordchannel.FieldNsfw, v)
	return u
}

// UpdateNsfw sets the "nsfw" field to the value that was provided on create.
func (u *DiscordChannelUpsert) UpdateNsfw() *DiscordChannelUpsert {
	u.SetExcluded(discordchannel.FieldNsfw)
	return u
}

// SetFlags sets the "flags" field.
func (u *DiscordChannelUpsert) SetFlags(v int) *DiscordChannelUpsert {
	u.Set(discordchannel.FieldFlags, v)
	return u
}

// UpdateFlags sets the "flags" field to the value that was provided on create.
func (u *DiscordChannelUpsert) UpdateFlags() *DiscordChannelUpsert {
	u.SetExcluded(discordchannel.FieldFlags)
	return u
}

// AddFlags adds v to the "flags" field.
func (u *DiscordChannelUpsert) AddFlags(v int) *DiscordChannelUpsert {
	u.Add(discordchannel.FieldFlags, v)
	return u
}

// ClearFlags clears the value of the "flags" field.
func (u *DiscordChannelUpsert) ClearFlags() *DiscordChannelUpsert {
	u.SetNull(discordchannel.FieldFlags)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DiscordChannel.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(discordchannel.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DiscordChannelUpsertOne) UpdateNewValues() *DiscordChannelUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(discordchannel.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(discordchannel.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DiscordChannel.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DiscordChannelUpsertOne) Ignore() *DiscordChannelUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DiscordChannelUpsertOne) DoNothing() *DiscordChannelUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DiscordChannelCreate.OnConflict
// documentation for more info.
func (u *DiscordChannelUpsertOne) Update(set func(*DiscordChannelUpsert)) *DiscordChannelUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DiscordChannelUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *DiscordChannelUpsertOne) SetUpdateTime(v time.Time) *DiscordChannelUpsertOne {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *DiscordChannelUpsertOne) UpdateUpdateTime() *DiscordChannelUpsertOne {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDiscordid sets the "discordid" field.
func (u *DiscordChannelUpsertOne) SetDiscordid(v string) *DiscordChannelUpsertOne {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.SetDiscordid(v)
	})
}

// UpdateDiscordid sets the "discordid" field to the value that was provided on create.
func (u *DiscordChannelUpsertOne) UpdateDiscordid() *DiscordChannelUpsertOne {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.UpdateDiscordid()
	})
}

// SetName sets the "name" field.
func (u *DiscordChannelUpsertOne) SetName(v string) *DiscordChannelUpsertOne {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DiscordChannelUpsertOne) UpdateName() *DiscordChannelUpsertOne {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.UpdateName()
	})
}

// SetType sets the "type" field.
func (u *DiscordChannelUpsertOne) SetType(v discord.ChannelType) *DiscordChannelUpsertOne {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.SetType(v)
	})
}

// AddType adds v to the "type" field.
func (u *DiscordChannelUpsertOne) AddType(v discord.ChannelType) *DiscordChannelUpsertOne {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.AddType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *DiscordChannelUpsertOne) UpdateType() *DiscordChannelUpsertOne {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.UpdateType()
	})
}

// SetNsfw sets the "nsfw" field.
func (u *DiscordChannelUpsertOne) SetNsfw(v bool) *DiscordChannelUpsertOne {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.SetNsfw(v)
	})
}

// UpdateNsfw sets the "nsfw" field to the value that was provided on create.
func (u *DiscordChannelUpsertOne) UpdateNsfw() *DiscordChannelUpsertOne {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.UpdateNsfw()
	})
}

// SetFlags sets the "flags" field.
func (u *DiscordChannelUpsertOne) SetFlags(v int) *DiscordChannelUpsertOne {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.SetFlags(v)
	})
}

// AddFlags adds v to the "flags" field.
func (u *DiscordChannelUpsertOne) AddFlags(v int) *DiscordChannelUpsertOne {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.AddFlags(v)
	})
}

// UpdateFlags sets the "flags" field to the value that was provided on create.
func (u *DiscordChannelUpsertOne) UpdateFlags() *DiscordChannelUpsertOne {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.UpdateFlags()
	})
}

// ClearFlags clears the value of the "flags" field.
func (u *DiscordChannelUpsertOne) ClearFlags() *DiscordChannelUpsertOne {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.ClearFlags()
	})
}

// Exec executes the query.
func (u *DiscordChannelUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DiscordChannelCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DiscordChannelUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DiscordChannelUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: DiscordChannelUpsertOne.ID is not supported by MySQL driver. Use DiscordChannelUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DiscordChannelUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DiscordChannelCreateBulk is the builder for creating many DiscordChannel entities in bulk.
type DiscordChannelCreateBulk struct {
	config
	err      error
	builders []*DiscordChannelCreate
	conflict []sql.ConflictOption
}

// Save creates the DiscordChannel entities in the database.
func (dccb *DiscordChannelCreateBulk) Save(ctx context.Context) ([]*DiscordChannel, error) {
	if dccb.err != nil {
		return nil, dccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dccb.builders))
	nodes := make([]*DiscordChannel, len(dccb.builders))
	mutators := make([]Mutator, len(dccb.builders))
	for i := range dccb.builders {
		func(i int, root context.Context) {
			builder := dccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DiscordChannelMutation)
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
					_, err = mutators[i+1].Mutate(root, dccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dccb *DiscordChannelCreateBulk) SaveX(ctx context.Context) []*DiscordChannel {
	v, err := dccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dccb *DiscordChannelCreateBulk) Exec(ctx context.Context) error {
	_, err := dccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dccb *DiscordChannelCreateBulk) ExecX(ctx context.Context) {
	if err := dccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DiscordChannel.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DiscordChannelUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (dccb *DiscordChannelCreateBulk) OnConflict(opts ...sql.ConflictOption) *DiscordChannelUpsertBulk {
	dccb.conflict = opts
	return &DiscordChannelUpsertBulk{
		create: dccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DiscordChannel.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dccb *DiscordChannelCreateBulk) OnConflictColumns(columns ...string) *DiscordChannelUpsertBulk {
	dccb.conflict = append(dccb.conflict, sql.ConflictColumns(columns...))
	return &DiscordChannelUpsertBulk{
		create: dccb,
	}
}

// DiscordChannelUpsertBulk is the builder for "upsert"-ing
// a bulk of DiscordChannel nodes.
type DiscordChannelUpsertBulk struct {
	create *DiscordChannelCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DiscordChannel.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(discordchannel.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DiscordChannelUpsertBulk) UpdateNewValues() *DiscordChannelUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(discordchannel.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(discordchannel.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DiscordChannel.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DiscordChannelUpsertBulk) Ignore() *DiscordChannelUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DiscordChannelUpsertBulk) DoNothing() *DiscordChannelUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DiscordChannelCreateBulk.OnConflict
// documentation for more info.
func (u *DiscordChannelUpsertBulk) Update(set func(*DiscordChannelUpsert)) *DiscordChannelUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DiscordChannelUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *DiscordChannelUpsertBulk) SetUpdateTime(v time.Time) *DiscordChannelUpsertBulk {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *DiscordChannelUpsertBulk) UpdateUpdateTime() *DiscordChannelUpsertBulk {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDiscordid sets the "discordid" field.
func (u *DiscordChannelUpsertBulk) SetDiscordid(v string) *DiscordChannelUpsertBulk {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.SetDiscordid(v)
	})
}

// UpdateDiscordid sets the "discordid" field to the value that was provided on create.
func (u *DiscordChannelUpsertBulk) UpdateDiscordid() *DiscordChannelUpsertBulk {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.UpdateDiscordid()
	})
}

// SetName sets the "name" field.
func (u *DiscordChannelUpsertBulk) SetName(v string) *DiscordChannelUpsertBulk {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DiscordChannelUpsertBulk) UpdateName() *DiscordChannelUpsertBulk {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.UpdateName()
	})
}

// SetType sets the "type" field.
func (u *DiscordChannelUpsertBulk) SetType(v discord.ChannelType) *DiscordChannelUpsertBulk {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.SetType(v)
	})
}

// AddType adds v to the "type" field.
func (u *DiscordChannelUpsertBulk) AddType(v discord.ChannelType) *DiscordChannelUpsertBulk {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.AddType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *DiscordChannelUpsertBulk) UpdateType() *DiscordChannelUpsertBulk {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.UpdateType()
	})
}

// SetNsfw sets the "nsfw" field.
func (u *DiscordChannelUpsertBulk) SetNsfw(v bool) *DiscordChannelUpsertBulk {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.SetNsfw(v)
	})
}

// UpdateNsfw sets the "nsfw" field to the value that was provided on create.
func (u *DiscordChannelUpsertBulk) UpdateNsfw() *DiscordChannelUpsertBulk {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.UpdateNsfw()
	})
}

// SetFlags sets the "flags" field.
func (u *DiscordChannelUpsertBulk) SetFlags(v int) *DiscordChannelUpsertBulk {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.SetFlags(v)
	})
}

// AddFlags adds v to the "flags" field.
func (u *DiscordChannelUpsertBulk) AddFlags(v int) *DiscordChannelUpsertBulk {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.AddFlags(v)
	})
}

// UpdateFlags sets the "flags" field to the value that was provided on create.
func (u *DiscordChannelUpsertBulk) UpdateFlags() *DiscordChannelUpsertBulk {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.UpdateFlags()
	})
}

// ClearFlags clears the value of the "flags" field.
func (u *DiscordChannelUpsertBulk) ClearFlags() *DiscordChannelUpsertBulk {
	return u.Update(func(s *DiscordChannelUpsert) {
		s.ClearFlags()
	})
}

// Exec executes the query.
func (u *DiscordChannelUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DiscordChannelCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DiscordChannelCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DiscordChannelUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
