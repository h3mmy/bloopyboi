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
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/h3mmy/bloopyboi/ent/discordguild"
	"github.com/h3mmy/bloopyboi/ent/discordmessage"
	"github.com/h3mmy/bloopyboi/ent/discordmessagereaction"
	"github.com/h3mmy/bloopyboi/ent/discorduser"
)

// DiscordMessageCreate is the builder for creating a DiscordMessage entity.
type DiscordMessageCreate struct {
	config
	mutation *DiscordMessageMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (dmc *DiscordMessageCreate) SetCreateTime(t time.Time) *DiscordMessageCreate {
	dmc.mutation.SetCreateTime(t)
	return dmc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (dmc *DiscordMessageCreate) SetNillableCreateTime(t *time.Time) *DiscordMessageCreate {
	if t != nil {
		dmc.SetCreateTime(*t)
	}
	return dmc
}

// SetUpdateTime sets the "update_time" field.
func (dmc *DiscordMessageCreate) SetUpdateTime(t time.Time) *DiscordMessageCreate {
	dmc.mutation.SetUpdateTime(t)
	return dmc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (dmc *DiscordMessageCreate) SetNillableUpdateTime(t *time.Time) *DiscordMessageCreate {
	if t != nil {
		dmc.SetUpdateTime(*t)
	}
	return dmc
}

// SetDiscordid sets the "discordid" field.
func (dmc *DiscordMessageCreate) SetDiscordid(s string) *DiscordMessageCreate {
	dmc.mutation.SetDiscordid(s)
	return dmc
}

// SetRaw sets the "raw" field.
func (dmc *DiscordMessageCreate) SetRaw(d discordgo.Message) *DiscordMessageCreate {
	dmc.mutation.SetRaw(d)
	return dmc
}

// SetID sets the "id" field.
func (dmc *DiscordMessageCreate) SetID(u uuid.UUID) *DiscordMessageCreate {
	dmc.mutation.SetID(u)
	return dmc
}

// SetAuthorID sets the "author" edge to the DiscordUser entity by ID.
func (dmc *DiscordMessageCreate) SetAuthorID(id uuid.UUID) *DiscordMessageCreate {
	dmc.mutation.SetAuthorID(id)
	return dmc
}

// SetNillableAuthorID sets the "author" edge to the DiscordUser entity by ID if the given value is not nil.
func (dmc *DiscordMessageCreate) SetNillableAuthorID(id *uuid.UUID) *DiscordMessageCreate {
	if id != nil {
		dmc = dmc.SetAuthorID(*id)
	}
	return dmc
}

// SetAuthor sets the "author" edge to the DiscordUser entity.
func (dmc *DiscordMessageCreate) SetAuthor(d *DiscordUser) *DiscordMessageCreate {
	return dmc.SetAuthorID(d.ID)
}

// AddMessageReactionIDs adds the "message_reactions" edge to the DiscordMessageReaction entity by IDs.
func (dmc *DiscordMessageCreate) AddMessageReactionIDs(ids ...uuid.UUID) *DiscordMessageCreate {
	dmc.mutation.AddMessageReactionIDs(ids...)
	return dmc
}

// AddMessageReactions adds the "message_reactions" edges to the DiscordMessageReaction entity.
func (dmc *DiscordMessageCreate) AddMessageReactions(d ...*DiscordMessageReaction) *DiscordMessageCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dmc.AddMessageReactionIDs(ids...)
}

// SetGuildID sets the "guild" edge to the DiscordGuild entity by ID.
func (dmc *DiscordMessageCreate) SetGuildID(id uuid.UUID) *DiscordMessageCreate {
	dmc.mutation.SetGuildID(id)
	return dmc
}

// SetNillableGuildID sets the "guild" edge to the DiscordGuild entity by ID if the given value is not nil.
func (dmc *DiscordMessageCreate) SetNillableGuildID(id *uuid.UUID) *DiscordMessageCreate {
	if id != nil {
		dmc = dmc.SetGuildID(*id)
	}
	return dmc
}

// SetGuild sets the "guild" edge to the DiscordGuild entity.
func (dmc *DiscordMessageCreate) SetGuild(d *DiscordGuild) *DiscordMessageCreate {
	return dmc.SetGuildID(d.ID)
}

// Mutation returns the DiscordMessageMutation object of the builder.
func (dmc *DiscordMessageCreate) Mutation() *DiscordMessageMutation {
	return dmc.mutation
}

// Save creates the DiscordMessage in the database.
func (dmc *DiscordMessageCreate) Save(ctx context.Context) (*DiscordMessage, error) {
	dmc.defaults()
	return withHooks(ctx, dmc.sqlSave, dmc.mutation, dmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dmc *DiscordMessageCreate) SaveX(ctx context.Context) *DiscordMessage {
	v, err := dmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dmc *DiscordMessageCreate) Exec(ctx context.Context) error {
	_, err := dmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dmc *DiscordMessageCreate) ExecX(ctx context.Context) {
	if err := dmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dmc *DiscordMessageCreate) defaults() {
	if _, ok := dmc.mutation.CreateTime(); !ok {
		v := discordmessage.DefaultCreateTime()
		dmc.mutation.SetCreateTime(v)
	}
	if _, ok := dmc.mutation.UpdateTime(); !ok {
		v := discordmessage.DefaultUpdateTime()
		dmc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dmc *DiscordMessageCreate) check() error {
	if _, ok := dmc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "DiscordMessage.create_time"`)}
	}
	if _, ok := dmc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "DiscordMessage.update_time"`)}
	}
	if _, ok := dmc.mutation.Discordid(); !ok {
		return &ValidationError{Name: "discordid", err: errors.New(`ent: missing required field "DiscordMessage.discordid"`)}
	}
	if _, ok := dmc.mutation.Raw(); !ok {
		return &ValidationError{Name: "raw", err: errors.New(`ent: missing required field "DiscordMessage.raw"`)}
	}
	return nil
}

func (dmc *DiscordMessageCreate) sqlSave(ctx context.Context) (*DiscordMessage, error) {
	if err := dmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dmc.driver, _spec); err != nil {
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
	dmc.mutation.id = &_node.ID
	dmc.mutation.done = true
	return _node, nil
}

func (dmc *DiscordMessageCreate) createSpec() (*DiscordMessage, *sqlgraph.CreateSpec) {
	var (
		_node = &DiscordMessage{config: dmc.config}
		_spec = sqlgraph.NewCreateSpec(discordmessage.Table, sqlgraph.NewFieldSpec(discordmessage.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = dmc.conflict
	if id, ok := dmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dmc.mutation.CreateTime(); ok {
		_spec.SetField(discordmessage.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := dmc.mutation.UpdateTime(); ok {
		_spec.SetField(discordmessage.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := dmc.mutation.Discordid(); ok {
		_spec.SetField(discordmessage.FieldDiscordid, field.TypeString, value)
		_node.Discordid = value
	}
	if value, ok := dmc.mutation.Raw(); ok {
		_spec.SetField(discordmessage.FieldRaw, field.TypeJSON, value)
		_node.Raw = value
	}
	if nodes := dmc.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   discordmessage.AuthorTable,
			Columns: []string{discordmessage.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discorduser.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.discord_user_discord_messages = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dmc.mutation.MessageReactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   discordmessage.MessageReactionsTable,
			Columns: []string{discordmessage.MessageReactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordmessagereaction.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dmc.mutation.GuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   discordmessage.GuildTable,
			Columns: []string{discordmessage.GuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordguild.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.discord_guild_discord_messages = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DiscordMessage.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DiscordMessageUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (dmc *DiscordMessageCreate) OnConflict(opts ...sql.ConflictOption) *DiscordMessageUpsertOne {
	dmc.conflict = opts
	return &DiscordMessageUpsertOne{
		create: dmc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DiscordMessage.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dmc *DiscordMessageCreate) OnConflictColumns(columns ...string) *DiscordMessageUpsertOne {
	dmc.conflict = append(dmc.conflict, sql.ConflictColumns(columns...))
	return &DiscordMessageUpsertOne{
		create: dmc,
	}
}

type (
	// DiscordMessageUpsertOne is the builder for "upsert"-ing
	//  one DiscordMessage node.
	DiscordMessageUpsertOne struct {
		create *DiscordMessageCreate
	}

	// DiscordMessageUpsert is the "OnConflict" setter.
	DiscordMessageUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *DiscordMessageUpsert) SetUpdateTime(v time.Time) *DiscordMessageUpsert {
	u.Set(discordmessage.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *DiscordMessageUpsert) UpdateUpdateTime() *DiscordMessageUpsert {
	u.SetExcluded(discordmessage.FieldUpdateTime)
	return u
}

// SetDiscordid sets the "discordid" field.
func (u *DiscordMessageUpsert) SetDiscordid(v string) *DiscordMessageUpsert {
	u.Set(discordmessage.FieldDiscordid, v)
	return u
}

// UpdateDiscordid sets the "discordid" field to the value that was provided on create.
func (u *DiscordMessageUpsert) UpdateDiscordid() *DiscordMessageUpsert {
	u.SetExcluded(discordmessage.FieldDiscordid)
	return u
}

// SetRaw sets the "raw" field.
func (u *DiscordMessageUpsert) SetRaw(v discordgo.Message) *DiscordMessageUpsert {
	u.Set(discordmessage.FieldRaw, v)
	return u
}

// UpdateRaw sets the "raw" field to the value that was provided on create.
func (u *DiscordMessageUpsert) UpdateRaw() *DiscordMessageUpsert {
	u.SetExcluded(discordmessage.FieldRaw)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DiscordMessage.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(discordmessage.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DiscordMessageUpsertOne) UpdateNewValues() *DiscordMessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(discordmessage.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(discordmessage.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DiscordMessage.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DiscordMessageUpsertOne) Ignore() *DiscordMessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DiscordMessageUpsertOne) DoNothing() *DiscordMessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DiscordMessageCreate.OnConflict
// documentation for more info.
func (u *DiscordMessageUpsertOne) Update(set func(*DiscordMessageUpsert)) *DiscordMessageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DiscordMessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *DiscordMessageUpsertOne) SetUpdateTime(v time.Time) *DiscordMessageUpsertOne {
	return u.Update(func(s *DiscordMessageUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *DiscordMessageUpsertOne) UpdateUpdateTime() *DiscordMessageUpsertOne {
	return u.Update(func(s *DiscordMessageUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDiscordid sets the "discordid" field.
func (u *DiscordMessageUpsertOne) SetDiscordid(v string) *DiscordMessageUpsertOne {
	return u.Update(func(s *DiscordMessageUpsert) {
		s.SetDiscordid(v)
	})
}

// UpdateDiscordid sets the "discordid" field to the value that was provided on create.
func (u *DiscordMessageUpsertOne) UpdateDiscordid() *DiscordMessageUpsertOne {
	return u.Update(func(s *DiscordMessageUpsert) {
		s.UpdateDiscordid()
	})
}

// SetRaw sets the "raw" field.
func (u *DiscordMessageUpsertOne) SetRaw(v discordgo.Message) *DiscordMessageUpsertOne {
	return u.Update(func(s *DiscordMessageUpsert) {
		s.SetRaw(v)
	})
}

// UpdateRaw sets the "raw" field to the value that was provided on create.
func (u *DiscordMessageUpsertOne) UpdateRaw() *DiscordMessageUpsertOne {
	return u.Update(func(s *DiscordMessageUpsert) {
		s.UpdateRaw()
	})
}

// Exec executes the query.
func (u *DiscordMessageUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DiscordMessageCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DiscordMessageUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DiscordMessageUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: DiscordMessageUpsertOne.ID is not supported by MySQL driver. Use DiscordMessageUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DiscordMessageUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DiscordMessageCreateBulk is the builder for creating many DiscordMessage entities in bulk.
type DiscordMessageCreateBulk struct {
	config
	err      error
	builders []*DiscordMessageCreate
	conflict []sql.ConflictOption
}

// Save creates the DiscordMessage entities in the database.
func (dmcb *DiscordMessageCreateBulk) Save(ctx context.Context) ([]*DiscordMessage, error) {
	if dmcb.err != nil {
		return nil, dmcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dmcb.builders))
	nodes := make([]*DiscordMessage, len(dmcb.builders))
	mutators := make([]Mutator, len(dmcb.builders))
	for i := range dmcb.builders {
		func(i int, root context.Context) {
			builder := dmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DiscordMessageMutation)
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
					_, err = mutators[i+1].Mutate(root, dmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dmcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dmcb *DiscordMessageCreateBulk) SaveX(ctx context.Context) []*DiscordMessage {
	v, err := dmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dmcb *DiscordMessageCreateBulk) Exec(ctx context.Context) error {
	_, err := dmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dmcb *DiscordMessageCreateBulk) ExecX(ctx context.Context) {
	if err := dmcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DiscordMessage.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DiscordMessageUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (dmcb *DiscordMessageCreateBulk) OnConflict(opts ...sql.ConflictOption) *DiscordMessageUpsertBulk {
	dmcb.conflict = opts
	return &DiscordMessageUpsertBulk{
		create: dmcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DiscordMessage.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dmcb *DiscordMessageCreateBulk) OnConflictColumns(columns ...string) *DiscordMessageUpsertBulk {
	dmcb.conflict = append(dmcb.conflict, sql.ConflictColumns(columns...))
	return &DiscordMessageUpsertBulk{
		create: dmcb,
	}
}

// DiscordMessageUpsertBulk is the builder for "upsert"-ing
// a bulk of DiscordMessage nodes.
type DiscordMessageUpsertBulk struct {
	create *DiscordMessageCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DiscordMessage.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(discordmessage.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DiscordMessageUpsertBulk) UpdateNewValues() *DiscordMessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(discordmessage.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(discordmessage.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DiscordMessage.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DiscordMessageUpsertBulk) Ignore() *DiscordMessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DiscordMessageUpsertBulk) DoNothing() *DiscordMessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DiscordMessageCreateBulk.OnConflict
// documentation for more info.
func (u *DiscordMessageUpsertBulk) Update(set func(*DiscordMessageUpsert)) *DiscordMessageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DiscordMessageUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *DiscordMessageUpsertBulk) SetUpdateTime(v time.Time) *DiscordMessageUpsertBulk {
	return u.Update(func(s *DiscordMessageUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *DiscordMessageUpsertBulk) UpdateUpdateTime() *DiscordMessageUpsertBulk {
	return u.Update(func(s *DiscordMessageUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetDiscordid sets the "discordid" field.
func (u *DiscordMessageUpsertBulk) SetDiscordid(v string) *DiscordMessageUpsertBulk {
	return u.Update(func(s *DiscordMessageUpsert) {
		s.SetDiscordid(v)
	})
}

// UpdateDiscordid sets the "discordid" field to the value that was provided on create.
func (u *DiscordMessageUpsertBulk) UpdateDiscordid() *DiscordMessageUpsertBulk {
	return u.Update(func(s *DiscordMessageUpsert) {
		s.UpdateDiscordid()
	})
}

// SetRaw sets the "raw" field.
func (u *DiscordMessageUpsertBulk) SetRaw(v discordgo.Message) *DiscordMessageUpsertBulk {
	return u.Update(func(s *DiscordMessageUpsert) {
		s.SetRaw(v)
	})
}

// UpdateRaw sets the "raw" field to the value that was provided on create.
func (u *DiscordMessageUpsertBulk) UpdateRaw() *DiscordMessageUpsertBulk {
	return u.Update(func(s *DiscordMessageUpsert) {
		s.UpdateRaw()
	})
}

// Exec executes the query.
func (u *DiscordMessageUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DiscordMessageCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DiscordMessageCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DiscordMessageUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
