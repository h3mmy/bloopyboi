// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/h3mmy/bloopyboi/ent/discordmessage"
	"github.com/h3mmy/bloopyboi/ent/discorduser"
	"github.com/h3mmy/bloopyboi/ent/mediarequest"
)

// DiscordUserCreate is the builder for creating a DiscordUser entity.
type DiscordUserCreate struct {
	config
	mutation *DiscordUserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDiscordid sets the "discordid" field.
func (duc *DiscordUserCreate) SetDiscordid(s string) *DiscordUserCreate {
	duc.mutation.SetDiscordid(s)
	return duc
}

// SetUsername sets the "username" field.
func (duc *DiscordUserCreate) SetUsername(s string) *DiscordUserCreate {
	duc.mutation.SetUsername(s)
	return duc
}

// SetEmail sets the "email" field.
func (duc *DiscordUserCreate) SetEmail(s string) *DiscordUserCreate {
	duc.mutation.SetEmail(s)
	return duc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (duc *DiscordUserCreate) SetNillableEmail(s *string) *DiscordUserCreate {
	if s != nil {
		duc.SetEmail(*s)
	}
	return duc
}

// SetDiscriminator sets the "discriminator" field.
func (duc *DiscordUserCreate) SetDiscriminator(s string) *DiscordUserCreate {
	duc.mutation.SetDiscriminator(s)
	return duc
}

// SetNillableDiscriminator sets the "discriminator" field if the given value is not nil.
func (duc *DiscordUserCreate) SetNillableDiscriminator(s *string) *DiscordUserCreate {
	if s != nil {
		duc.SetDiscriminator(*s)
	}
	return duc
}

// SetID sets the "id" field.
func (duc *DiscordUserCreate) SetID(u uuid.UUID) *DiscordUserCreate {
	duc.mutation.SetID(u)
	return duc
}

// AddDiscordMessageIDs adds the "discord_messages" edge to the DiscordMessage entity by IDs.
func (duc *DiscordUserCreate) AddDiscordMessageIDs(ids ...string) *DiscordUserCreate {
	duc.mutation.AddDiscordMessageIDs(ids...)
	return duc
}

// AddDiscordMessages adds the "discord_messages" edges to the DiscordMessage entity.
func (duc *DiscordUserCreate) AddDiscordMessages(d ...*DiscordMessage) *DiscordUserCreate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duc.AddDiscordMessageIDs(ids...)
}

// AddMediaRequestIDs adds the "media_requests" edge to the MediaRequest entity by IDs.
func (duc *DiscordUserCreate) AddMediaRequestIDs(ids ...uuid.UUID) *DiscordUserCreate {
	duc.mutation.AddMediaRequestIDs(ids...)
	return duc
}

// AddMediaRequests adds the "media_requests" edges to the MediaRequest entity.
func (duc *DiscordUserCreate) AddMediaRequests(m ...*MediaRequest) *DiscordUserCreate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return duc.AddMediaRequestIDs(ids...)
}

// Mutation returns the DiscordUserMutation object of the builder.
func (duc *DiscordUserCreate) Mutation() *DiscordUserMutation {
	return duc.mutation
}

// Save creates the DiscordUser in the database.
func (duc *DiscordUserCreate) Save(ctx context.Context) (*DiscordUser, error) {
	return withHooks(ctx, duc.sqlSave, duc.mutation, duc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (duc *DiscordUserCreate) SaveX(ctx context.Context) *DiscordUser {
	v, err := duc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (duc *DiscordUserCreate) Exec(ctx context.Context) error {
	_, err := duc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duc *DiscordUserCreate) ExecX(ctx context.Context) {
	if err := duc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duc *DiscordUserCreate) check() error {
	if _, ok := duc.mutation.Discordid(); !ok {
		return &ValidationError{Name: "discordid", err: errors.New(`ent: missing required field "DiscordUser.discordid"`)}
	}
	if _, ok := duc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "DiscordUser.username"`)}
	}
	return nil
}

func (duc *DiscordUserCreate) sqlSave(ctx context.Context) (*DiscordUser, error) {
	if err := duc.check(); err != nil {
		return nil, err
	}
	_node, _spec := duc.createSpec()
	if err := sqlgraph.CreateNode(ctx, duc.driver, _spec); err != nil {
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
	duc.mutation.id = &_node.ID
	duc.mutation.done = true
	return _node, nil
}

func (duc *DiscordUserCreate) createSpec() (*DiscordUser, *sqlgraph.CreateSpec) {
	var (
		_node = &DiscordUser{config: duc.config}
		_spec = sqlgraph.NewCreateSpec(discorduser.Table, sqlgraph.NewFieldSpec(discorduser.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = duc.conflict
	if id, ok := duc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := duc.mutation.Discordid(); ok {
		_spec.SetField(discorduser.FieldDiscordid, field.TypeString, value)
		_node.Discordid = value
	}
	if value, ok := duc.mutation.Username(); ok {
		_spec.SetField(discorduser.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := duc.mutation.Email(); ok {
		_spec.SetField(discorduser.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := duc.mutation.Discriminator(); ok {
		_spec.SetField(discorduser.FieldDiscriminator, field.TypeString, value)
		_node.Discriminator = value
	}
	if nodes := duc.mutation.DiscordMessagesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := duc.mutation.MediaRequestsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DiscordUser.Create().
//		SetDiscordid(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DiscordUserUpsert) {
//			SetDiscordid(v+v).
//		}).
//		Exec(ctx)
func (duc *DiscordUserCreate) OnConflict(opts ...sql.ConflictOption) *DiscordUserUpsertOne {
	duc.conflict = opts
	return &DiscordUserUpsertOne{
		create: duc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DiscordUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (duc *DiscordUserCreate) OnConflictColumns(columns ...string) *DiscordUserUpsertOne {
	duc.conflict = append(duc.conflict, sql.ConflictColumns(columns...))
	return &DiscordUserUpsertOne{
		create: duc,
	}
}

type (
	// DiscordUserUpsertOne is the builder for "upsert"-ing
	//  one DiscordUser node.
	DiscordUserUpsertOne struct {
		create *DiscordUserCreate
	}

	// DiscordUserUpsert is the "OnConflict" setter.
	DiscordUserUpsert struct {
		*sql.UpdateSet
	}
)

// SetDiscordid sets the "discordid" field.
func (u *DiscordUserUpsert) SetDiscordid(v string) *DiscordUserUpsert {
	u.Set(discorduser.FieldDiscordid, v)
	return u
}

// UpdateDiscordid sets the "discordid" field to the value that was provided on create.
func (u *DiscordUserUpsert) UpdateDiscordid() *DiscordUserUpsert {
	u.SetExcluded(discorduser.FieldDiscordid)
	return u
}

// SetUsername sets the "username" field.
func (u *DiscordUserUpsert) SetUsername(v string) *DiscordUserUpsert {
	u.Set(discorduser.FieldUsername, v)
	return u
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *DiscordUserUpsert) UpdateUsername() *DiscordUserUpsert {
	u.SetExcluded(discorduser.FieldUsername)
	return u
}

// SetEmail sets the "email" field.
func (u *DiscordUserUpsert) SetEmail(v string) *DiscordUserUpsert {
	u.Set(discorduser.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *DiscordUserUpsert) UpdateEmail() *DiscordUserUpsert {
	u.SetExcluded(discorduser.FieldEmail)
	return u
}

// ClearEmail clears the value of the "email" field.
func (u *DiscordUserUpsert) ClearEmail() *DiscordUserUpsert {
	u.SetNull(discorduser.FieldEmail)
	return u
}

// SetDiscriminator sets the "discriminator" field.
func (u *DiscordUserUpsert) SetDiscriminator(v string) *DiscordUserUpsert {
	u.Set(discorduser.FieldDiscriminator, v)
	return u
}

// UpdateDiscriminator sets the "discriminator" field to the value that was provided on create.
func (u *DiscordUserUpsert) UpdateDiscriminator() *DiscordUserUpsert {
	u.SetExcluded(discorduser.FieldDiscriminator)
	return u
}

// ClearDiscriminator clears the value of the "discriminator" field.
func (u *DiscordUserUpsert) ClearDiscriminator() *DiscordUserUpsert {
	u.SetNull(discorduser.FieldDiscriminator)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DiscordUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(discorduser.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DiscordUserUpsertOne) UpdateNewValues() *DiscordUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(discorduser.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DiscordUser.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DiscordUserUpsertOne) Ignore() *DiscordUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DiscordUserUpsertOne) DoNothing() *DiscordUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DiscordUserCreate.OnConflict
// documentation for more info.
func (u *DiscordUserUpsertOne) Update(set func(*DiscordUserUpsert)) *DiscordUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DiscordUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetDiscordid sets the "discordid" field.
func (u *DiscordUserUpsertOne) SetDiscordid(v string) *DiscordUserUpsertOne {
	return u.Update(func(s *DiscordUserUpsert) {
		s.SetDiscordid(v)
	})
}

// UpdateDiscordid sets the "discordid" field to the value that was provided on create.
func (u *DiscordUserUpsertOne) UpdateDiscordid() *DiscordUserUpsertOne {
	return u.Update(func(s *DiscordUserUpsert) {
		s.UpdateDiscordid()
	})
}

// SetUsername sets the "username" field.
func (u *DiscordUserUpsertOne) SetUsername(v string) *DiscordUserUpsertOne {
	return u.Update(func(s *DiscordUserUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *DiscordUserUpsertOne) UpdateUsername() *DiscordUserUpsertOne {
	return u.Update(func(s *DiscordUserUpsert) {
		s.UpdateUsername()
	})
}

// SetEmail sets the "email" field.
func (u *DiscordUserUpsertOne) SetEmail(v string) *DiscordUserUpsertOne {
	return u.Update(func(s *DiscordUserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *DiscordUserUpsertOne) UpdateEmail() *DiscordUserUpsertOne {
	return u.Update(func(s *DiscordUserUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *DiscordUserUpsertOne) ClearEmail() *DiscordUserUpsertOne {
	return u.Update(func(s *DiscordUserUpsert) {
		s.ClearEmail()
	})
}

// SetDiscriminator sets the "discriminator" field.
func (u *DiscordUserUpsertOne) SetDiscriminator(v string) *DiscordUserUpsertOne {
	return u.Update(func(s *DiscordUserUpsert) {
		s.SetDiscriminator(v)
	})
}

// UpdateDiscriminator sets the "discriminator" field to the value that was provided on create.
func (u *DiscordUserUpsertOne) UpdateDiscriminator() *DiscordUserUpsertOne {
	return u.Update(func(s *DiscordUserUpsert) {
		s.UpdateDiscriminator()
	})
}

// ClearDiscriminator clears the value of the "discriminator" field.
func (u *DiscordUserUpsertOne) ClearDiscriminator() *DiscordUserUpsertOne {
	return u.Update(func(s *DiscordUserUpsert) {
		s.ClearDiscriminator()
	})
}

// Exec executes the query.
func (u *DiscordUserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DiscordUserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DiscordUserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DiscordUserUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: DiscordUserUpsertOne.ID is not supported by MySQL driver. Use DiscordUserUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DiscordUserUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DiscordUserCreateBulk is the builder for creating many DiscordUser entities in bulk.
type DiscordUserCreateBulk struct {
	config
	err      error
	builders []*DiscordUserCreate
	conflict []sql.ConflictOption
}

// Save creates the DiscordUser entities in the database.
func (ducb *DiscordUserCreateBulk) Save(ctx context.Context) ([]*DiscordUser, error) {
	if ducb.err != nil {
		return nil, ducb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ducb.builders))
	nodes := make([]*DiscordUser, len(ducb.builders))
	mutators := make([]Mutator, len(ducb.builders))
	for i := range ducb.builders {
		func(i int, root context.Context) {
			builder := ducb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DiscordUserMutation)
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
					_, err = mutators[i+1].Mutate(root, ducb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ducb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ducb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ducb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ducb *DiscordUserCreateBulk) SaveX(ctx context.Context) []*DiscordUser {
	v, err := ducb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ducb *DiscordUserCreateBulk) Exec(ctx context.Context) error {
	_, err := ducb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ducb *DiscordUserCreateBulk) ExecX(ctx context.Context) {
	if err := ducb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DiscordUser.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DiscordUserUpsert) {
//			SetDiscordid(v+v).
//		}).
//		Exec(ctx)
func (ducb *DiscordUserCreateBulk) OnConflict(opts ...sql.ConflictOption) *DiscordUserUpsertBulk {
	ducb.conflict = opts
	return &DiscordUserUpsertBulk{
		create: ducb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DiscordUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ducb *DiscordUserCreateBulk) OnConflictColumns(columns ...string) *DiscordUserUpsertBulk {
	ducb.conflict = append(ducb.conflict, sql.ConflictColumns(columns...))
	return &DiscordUserUpsertBulk{
		create: ducb,
	}
}

// DiscordUserUpsertBulk is the builder for "upsert"-ing
// a bulk of DiscordUser nodes.
type DiscordUserUpsertBulk struct {
	create *DiscordUserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DiscordUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(discorduser.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DiscordUserUpsertBulk) UpdateNewValues() *DiscordUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(discorduser.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DiscordUser.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DiscordUserUpsertBulk) Ignore() *DiscordUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DiscordUserUpsertBulk) DoNothing() *DiscordUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DiscordUserCreateBulk.OnConflict
// documentation for more info.
func (u *DiscordUserUpsertBulk) Update(set func(*DiscordUserUpsert)) *DiscordUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DiscordUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetDiscordid sets the "discordid" field.
func (u *DiscordUserUpsertBulk) SetDiscordid(v string) *DiscordUserUpsertBulk {
	return u.Update(func(s *DiscordUserUpsert) {
		s.SetDiscordid(v)
	})
}

// UpdateDiscordid sets the "discordid" field to the value that was provided on create.
func (u *DiscordUserUpsertBulk) UpdateDiscordid() *DiscordUserUpsertBulk {
	return u.Update(func(s *DiscordUserUpsert) {
		s.UpdateDiscordid()
	})
}

// SetUsername sets the "username" field.
func (u *DiscordUserUpsertBulk) SetUsername(v string) *DiscordUserUpsertBulk {
	return u.Update(func(s *DiscordUserUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *DiscordUserUpsertBulk) UpdateUsername() *DiscordUserUpsertBulk {
	return u.Update(func(s *DiscordUserUpsert) {
		s.UpdateUsername()
	})
}

// SetEmail sets the "email" field.
func (u *DiscordUserUpsertBulk) SetEmail(v string) *DiscordUserUpsertBulk {
	return u.Update(func(s *DiscordUserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *DiscordUserUpsertBulk) UpdateEmail() *DiscordUserUpsertBulk {
	return u.Update(func(s *DiscordUserUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *DiscordUserUpsertBulk) ClearEmail() *DiscordUserUpsertBulk {
	return u.Update(func(s *DiscordUserUpsert) {
		s.ClearEmail()
	})
}

// SetDiscriminator sets the "discriminator" field.
func (u *DiscordUserUpsertBulk) SetDiscriminator(v string) *DiscordUserUpsertBulk {
	return u.Update(func(s *DiscordUserUpsert) {
		s.SetDiscriminator(v)
	})
}

// UpdateDiscriminator sets the "discriminator" field to the value that was provided on create.
func (u *DiscordUserUpsertBulk) UpdateDiscriminator() *DiscordUserUpsertBulk {
	return u.Update(func(s *DiscordUserUpsert) {
		s.UpdateDiscriminator()
	})
}

// ClearDiscriminator clears the value of the "discriminator" field.
func (u *DiscordUserUpsertBulk) ClearDiscriminator() *DiscordUserUpsertBulk {
	return u.Update(func(s *DiscordUserUpsert) {
		s.ClearDiscriminator()
	})
}

// Exec executes the query.
func (u *DiscordUserUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DiscordUserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DiscordUserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DiscordUserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
