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
	"github.com/h3mmy/bloopyboi/ent/discordchannel"
	"github.com/h3mmy/bloopyboi/ent/discordguild"
	"github.com/h3mmy/bloopyboi/ent/discordmessage"
	"github.com/h3mmy/bloopyboi/ent/discorduser"
)

// DiscordGuildCreate is the builder for creating a DiscordGuild entity.
type DiscordGuildCreate struct {
	config
	mutation *DiscordGuildMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDiscordid sets the "discordid" field.
func (dgc *DiscordGuildCreate) SetDiscordid(s string) *DiscordGuildCreate {
	dgc.mutation.SetDiscordid(s)
	return dgc
}

// SetName sets the "name" field.
func (dgc *DiscordGuildCreate) SetName(s string) *DiscordGuildCreate {
	dgc.mutation.SetName(s)
	return dgc
}

// SetDescription sets the "description" field.
func (dgc *DiscordGuildCreate) SetDescription(s string) *DiscordGuildCreate {
	dgc.mutation.SetDescription(s)
	return dgc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (dgc *DiscordGuildCreate) SetNillableDescription(s *string) *DiscordGuildCreate {
	if s != nil {
		dgc.SetDescription(*s)
	}
	return dgc
}

// SetRulesChannelID sets the "rules_channel_id" field.
func (dgc *DiscordGuildCreate) SetRulesChannelID(s string) *DiscordGuildCreate {
	dgc.mutation.SetRulesChannelID(s)
	return dgc
}

// SetNillableRulesChannelID sets the "rules_channel_id" field if the given value is not nil.
func (dgc *DiscordGuildCreate) SetNillableRulesChannelID(s *string) *DiscordGuildCreate {
	if s != nil {
		dgc.SetRulesChannelID(*s)
	}
	return dgc
}

// SetPublicUpdatesChannelID sets the "public_updates_channel_id" field.
func (dgc *DiscordGuildCreate) SetPublicUpdatesChannelID(s string) *DiscordGuildCreate {
	dgc.mutation.SetPublicUpdatesChannelID(s)
	return dgc
}

// SetNillablePublicUpdatesChannelID sets the "public_updates_channel_id" field if the given value is not nil.
func (dgc *DiscordGuildCreate) SetNillablePublicUpdatesChannelID(s *string) *DiscordGuildCreate {
	if s != nil {
		dgc.SetPublicUpdatesChannelID(*s)
	}
	return dgc
}

// SetNsfwLevel sets the "nsfw_level" field.
func (dgc *DiscordGuildCreate) SetNsfwLevel(i int) *DiscordGuildCreate {
	dgc.mutation.SetNsfwLevel(i)
	return dgc
}

// SetNillableNsfwLevel sets the "nsfw_level" field if the given value is not nil.
func (dgc *DiscordGuildCreate) SetNillableNsfwLevel(i *int) *DiscordGuildCreate {
	if i != nil {
		dgc.SetNsfwLevel(*i)
	}
	return dgc
}

// SetID sets the "id" field.
func (dgc *DiscordGuildCreate) SetID(u uuid.UUID) *DiscordGuildCreate {
	dgc.mutation.SetID(u)
	return dgc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (dgc *DiscordGuildCreate) SetNillableID(u *uuid.UUID) *DiscordGuildCreate {
	if u != nil {
		dgc.SetID(*u)
	}
	return dgc
}

// AddMemberIDs adds the "members" edge to the DiscordUser entity by IDs.
func (dgc *DiscordGuildCreate) AddMemberIDs(ids ...uuid.UUID) *DiscordGuildCreate {
	dgc.mutation.AddMemberIDs(ids...)
	return dgc
}

// AddMembers adds the "members" edges to the DiscordUser entity.
func (dgc *DiscordGuildCreate) AddMembers(d ...*DiscordUser) *DiscordGuildCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dgc.AddMemberIDs(ids...)
}

// AddDiscordMessageIDs adds the "discord_messages" edge to the DiscordMessage entity by IDs.
func (dgc *DiscordGuildCreate) AddDiscordMessageIDs(ids ...uuid.UUID) *DiscordGuildCreate {
	dgc.mutation.AddDiscordMessageIDs(ids...)
	return dgc
}

// AddDiscordMessages adds the "discord_messages" edges to the DiscordMessage entity.
func (dgc *DiscordGuildCreate) AddDiscordMessages(d ...*DiscordMessage) *DiscordGuildCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dgc.AddDiscordMessageIDs(ids...)
}

// AddGuildChannelIDs adds the "guild_channels" edge to the DiscordChannel entity by IDs.
func (dgc *DiscordGuildCreate) AddGuildChannelIDs(ids ...uuid.UUID) *DiscordGuildCreate {
	dgc.mutation.AddGuildChannelIDs(ids...)
	return dgc
}

// AddGuildChannels adds the "guild_channels" edges to the DiscordChannel entity.
func (dgc *DiscordGuildCreate) AddGuildChannels(d ...*DiscordChannel) *DiscordGuildCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dgc.AddGuildChannelIDs(ids...)
}

// Mutation returns the DiscordGuildMutation object of the builder.
func (dgc *DiscordGuildCreate) Mutation() *DiscordGuildMutation {
	return dgc.mutation
}

// Save creates the DiscordGuild in the database.
func (dgc *DiscordGuildCreate) Save(ctx context.Context) (*DiscordGuild, error) {
	dgc.defaults()
	return withHooks(ctx, dgc.sqlSave, dgc.mutation, dgc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dgc *DiscordGuildCreate) SaveX(ctx context.Context) *DiscordGuild {
	v, err := dgc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dgc *DiscordGuildCreate) Exec(ctx context.Context) error {
	_, err := dgc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dgc *DiscordGuildCreate) ExecX(ctx context.Context) {
	if err := dgc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dgc *DiscordGuildCreate) defaults() {
	if _, ok := dgc.mutation.ID(); !ok {
		v := discordguild.DefaultID()
		dgc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dgc *DiscordGuildCreate) check() error {
	if _, ok := dgc.mutation.Discordid(); !ok {
		return &ValidationError{Name: "discordid", err: errors.New(`ent: missing required field "DiscordGuild.discordid"`)}
	}
	if _, ok := dgc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "DiscordGuild.name"`)}
	}
	return nil
}

func (dgc *DiscordGuildCreate) sqlSave(ctx context.Context) (*DiscordGuild, error) {
	if err := dgc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dgc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dgc.driver, _spec); err != nil {
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
	dgc.mutation.id = &_node.ID
	dgc.mutation.done = true
	return _node, nil
}

func (dgc *DiscordGuildCreate) createSpec() (*DiscordGuild, *sqlgraph.CreateSpec) {
	var (
		_node = &DiscordGuild{config: dgc.config}
		_spec = sqlgraph.NewCreateSpec(discordguild.Table, sqlgraph.NewFieldSpec(discordguild.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = dgc.conflict
	if id, ok := dgc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dgc.mutation.Discordid(); ok {
		_spec.SetField(discordguild.FieldDiscordid, field.TypeString, value)
		_node.Discordid = value
	}
	if value, ok := dgc.mutation.Name(); ok {
		_spec.SetField(discordguild.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dgc.mutation.Description(); ok {
		_spec.SetField(discordguild.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := dgc.mutation.RulesChannelID(); ok {
		_spec.SetField(discordguild.FieldRulesChannelID, field.TypeString, value)
		_node.RulesChannelID = value
	}
	if value, ok := dgc.mutation.PublicUpdatesChannelID(); ok {
		_spec.SetField(discordguild.FieldPublicUpdatesChannelID, field.TypeString, value)
		_node.PublicUpdatesChannelID = value
	}
	if value, ok := dgc.mutation.NsfwLevel(); ok {
		_spec.SetField(discordguild.FieldNsfwLevel, field.TypeInt, value)
		_node.NsfwLevel = value
	}
	if nodes := dgc.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   discordguild.MembersTable,
			Columns: discordguild.MembersPrimaryKey,
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
	if nodes := dgc.mutation.DiscordMessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   discordguild.DiscordMessagesTable,
			Columns: []string{discordguild.DiscordMessagesColumn},
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
	if nodes := dgc.mutation.GuildChannelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   discordguild.GuildChannelsTable,
			Columns: discordguild.GuildChannelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordchannel.FieldID, field.TypeUUID),
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
//	client.DiscordGuild.Create().
//		SetDiscordid(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DiscordGuildUpsert) {
//			SetDiscordid(v+v).
//		}).
//		Exec(ctx)
func (dgc *DiscordGuildCreate) OnConflict(opts ...sql.ConflictOption) *DiscordGuildUpsertOne {
	dgc.conflict = opts
	return &DiscordGuildUpsertOne{
		create: dgc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DiscordGuild.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dgc *DiscordGuildCreate) OnConflictColumns(columns ...string) *DiscordGuildUpsertOne {
	dgc.conflict = append(dgc.conflict, sql.ConflictColumns(columns...))
	return &DiscordGuildUpsertOne{
		create: dgc,
	}
}

type (
	// DiscordGuildUpsertOne is the builder for "upsert"-ing
	//  one DiscordGuild node.
	DiscordGuildUpsertOne struct {
		create *DiscordGuildCreate
	}

	// DiscordGuildUpsert is the "OnConflict" setter.
	DiscordGuildUpsert struct {
		*sql.UpdateSet
	}
)

// SetDiscordid sets the "discordid" field.
func (u *DiscordGuildUpsert) SetDiscordid(v string) *DiscordGuildUpsert {
	u.Set(discordguild.FieldDiscordid, v)
	return u
}

// UpdateDiscordid sets the "discordid" field to the value that was provided on create.
func (u *DiscordGuildUpsert) UpdateDiscordid() *DiscordGuildUpsert {
	u.SetExcluded(discordguild.FieldDiscordid)
	return u
}

// SetName sets the "name" field.
func (u *DiscordGuildUpsert) SetName(v string) *DiscordGuildUpsert {
	u.Set(discordguild.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DiscordGuildUpsert) UpdateName() *DiscordGuildUpsert {
	u.SetExcluded(discordguild.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *DiscordGuildUpsert) SetDescription(v string) *DiscordGuildUpsert {
	u.Set(discordguild.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *DiscordGuildUpsert) UpdateDescription() *DiscordGuildUpsert {
	u.SetExcluded(discordguild.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *DiscordGuildUpsert) ClearDescription() *DiscordGuildUpsert {
	u.SetNull(discordguild.FieldDescription)
	return u
}

// SetRulesChannelID sets the "rules_channel_id" field.
func (u *DiscordGuildUpsert) SetRulesChannelID(v string) *DiscordGuildUpsert {
	u.Set(discordguild.FieldRulesChannelID, v)
	return u
}

// UpdateRulesChannelID sets the "rules_channel_id" field to the value that was provided on create.
func (u *DiscordGuildUpsert) UpdateRulesChannelID() *DiscordGuildUpsert {
	u.SetExcluded(discordguild.FieldRulesChannelID)
	return u
}

// ClearRulesChannelID clears the value of the "rules_channel_id" field.
func (u *DiscordGuildUpsert) ClearRulesChannelID() *DiscordGuildUpsert {
	u.SetNull(discordguild.FieldRulesChannelID)
	return u
}

// SetPublicUpdatesChannelID sets the "public_updates_channel_id" field.
func (u *DiscordGuildUpsert) SetPublicUpdatesChannelID(v string) *DiscordGuildUpsert {
	u.Set(discordguild.FieldPublicUpdatesChannelID, v)
	return u
}

// UpdatePublicUpdatesChannelID sets the "public_updates_channel_id" field to the value that was provided on create.
func (u *DiscordGuildUpsert) UpdatePublicUpdatesChannelID() *DiscordGuildUpsert {
	u.SetExcluded(discordguild.FieldPublicUpdatesChannelID)
	return u
}

// ClearPublicUpdatesChannelID clears the value of the "public_updates_channel_id" field.
func (u *DiscordGuildUpsert) ClearPublicUpdatesChannelID() *DiscordGuildUpsert {
	u.SetNull(discordguild.FieldPublicUpdatesChannelID)
	return u
}

// SetNsfwLevel sets the "nsfw_level" field.
func (u *DiscordGuildUpsert) SetNsfwLevel(v int) *DiscordGuildUpsert {
	u.Set(discordguild.FieldNsfwLevel, v)
	return u
}

// UpdateNsfwLevel sets the "nsfw_level" field to the value that was provided on create.
func (u *DiscordGuildUpsert) UpdateNsfwLevel() *DiscordGuildUpsert {
	u.SetExcluded(discordguild.FieldNsfwLevel)
	return u
}

// AddNsfwLevel adds v to the "nsfw_level" field.
func (u *DiscordGuildUpsert) AddNsfwLevel(v int) *DiscordGuildUpsert {
	u.Add(discordguild.FieldNsfwLevel, v)
	return u
}

// ClearNsfwLevel clears the value of the "nsfw_level" field.
func (u *DiscordGuildUpsert) ClearNsfwLevel() *DiscordGuildUpsert {
	u.SetNull(discordguild.FieldNsfwLevel)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DiscordGuild.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(discordguild.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DiscordGuildUpsertOne) UpdateNewValues() *DiscordGuildUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(discordguild.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DiscordGuild.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DiscordGuildUpsertOne) Ignore() *DiscordGuildUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DiscordGuildUpsertOne) DoNothing() *DiscordGuildUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DiscordGuildCreate.OnConflict
// documentation for more info.
func (u *DiscordGuildUpsertOne) Update(set func(*DiscordGuildUpsert)) *DiscordGuildUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DiscordGuildUpsert{UpdateSet: update})
	}))
	return u
}

// SetDiscordid sets the "discordid" field.
func (u *DiscordGuildUpsertOne) SetDiscordid(v string) *DiscordGuildUpsertOne {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.SetDiscordid(v)
	})
}

// UpdateDiscordid sets the "discordid" field to the value that was provided on create.
func (u *DiscordGuildUpsertOne) UpdateDiscordid() *DiscordGuildUpsertOne {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.UpdateDiscordid()
	})
}

// SetName sets the "name" field.
func (u *DiscordGuildUpsertOne) SetName(v string) *DiscordGuildUpsertOne {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DiscordGuildUpsertOne) UpdateName() *DiscordGuildUpsertOne {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *DiscordGuildUpsertOne) SetDescription(v string) *DiscordGuildUpsertOne {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *DiscordGuildUpsertOne) UpdateDescription() *DiscordGuildUpsertOne {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *DiscordGuildUpsertOne) ClearDescription() *DiscordGuildUpsertOne {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.ClearDescription()
	})
}

// SetRulesChannelID sets the "rules_channel_id" field.
func (u *DiscordGuildUpsertOne) SetRulesChannelID(v string) *DiscordGuildUpsertOne {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.SetRulesChannelID(v)
	})
}

// UpdateRulesChannelID sets the "rules_channel_id" field to the value that was provided on create.
func (u *DiscordGuildUpsertOne) UpdateRulesChannelID() *DiscordGuildUpsertOne {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.UpdateRulesChannelID()
	})
}

// ClearRulesChannelID clears the value of the "rules_channel_id" field.
func (u *DiscordGuildUpsertOne) ClearRulesChannelID() *DiscordGuildUpsertOne {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.ClearRulesChannelID()
	})
}

// SetPublicUpdatesChannelID sets the "public_updates_channel_id" field.
func (u *DiscordGuildUpsertOne) SetPublicUpdatesChannelID(v string) *DiscordGuildUpsertOne {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.SetPublicUpdatesChannelID(v)
	})
}

// UpdatePublicUpdatesChannelID sets the "public_updates_channel_id" field to the value that was provided on create.
func (u *DiscordGuildUpsertOne) UpdatePublicUpdatesChannelID() *DiscordGuildUpsertOne {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.UpdatePublicUpdatesChannelID()
	})
}

// ClearPublicUpdatesChannelID clears the value of the "public_updates_channel_id" field.
func (u *DiscordGuildUpsertOne) ClearPublicUpdatesChannelID() *DiscordGuildUpsertOne {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.ClearPublicUpdatesChannelID()
	})
}

// SetNsfwLevel sets the "nsfw_level" field.
func (u *DiscordGuildUpsertOne) SetNsfwLevel(v int) *DiscordGuildUpsertOne {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.SetNsfwLevel(v)
	})
}

// AddNsfwLevel adds v to the "nsfw_level" field.
func (u *DiscordGuildUpsertOne) AddNsfwLevel(v int) *DiscordGuildUpsertOne {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.AddNsfwLevel(v)
	})
}

// UpdateNsfwLevel sets the "nsfw_level" field to the value that was provided on create.
func (u *DiscordGuildUpsertOne) UpdateNsfwLevel() *DiscordGuildUpsertOne {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.UpdateNsfwLevel()
	})
}

// ClearNsfwLevel clears the value of the "nsfw_level" field.
func (u *DiscordGuildUpsertOne) ClearNsfwLevel() *DiscordGuildUpsertOne {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.ClearNsfwLevel()
	})
}

// Exec executes the query.
func (u *DiscordGuildUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DiscordGuildCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DiscordGuildUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DiscordGuildUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: DiscordGuildUpsertOne.ID is not supported by MySQL driver. Use DiscordGuildUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DiscordGuildUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DiscordGuildCreateBulk is the builder for creating many DiscordGuild entities in bulk.
type DiscordGuildCreateBulk struct {
	config
	err      error
	builders []*DiscordGuildCreate
	conflict []sql.ConflictOption
}

// Save creates the DiscordGuild entities in the database.
func (dgcb *DiscordGuildCreateBulk) Save(ctx context.Context) ([]*DiscordGuild, error) {
	if dgcb.err != nil {
		return nil, dgcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dgcb.builders))
	nodes := make([]*DiscordGuild, len(dgcb.builders))
	mutators := make([]Mutator, len(dgcb.builders))
	for i := range dgcb.builders {
		func(i int, root context.Context) {
			builder := dgcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DiscordGuildMutation)
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
					_, err = mutators[i+1].Mutate(root, dgcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dgcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dgcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dgcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dgcb *DiscordGuildCreateBulk) SaveX(ctx context.Context) []*DiscordGuild {
	v, err := dgcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dgcb *DiscordGuildCreateBulk) Exec(ctx context.Context) error {
	_, err := dgcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dgcb *DiscordGuildCreateBulk) ExecX(ctx context.Context) {
	if err := dgcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DiscordGuild.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DiscordGuildUpsert) {
//			SetDiscordid(v+v).
//		}).
//		Exec(ctx)
func (dgcb *DiscordGuildCreateBulk) OnConflict(opts ...sql.ConflictOption) *DiscordGuildUpsertBulk {
	dgcb.conflict = opts
	return &DiscordGuildUpsertBulk{
		create: dgcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DiscordGuild.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dgcb *DiscordGuildCreateBulk) OnConflictColumns(columns ...string) *DiscordGuildUpsertBulk {
	dgcb.conflict = append(dgcb.conflict, sql.ConflictColumns(columns...))
	return &DiscordGuildUpsertBulk{
		create: dgcb,
	}
}

// DiscordGuildUpsertBulk is the builder for "upsert"-ing
// a bulk of DiscordGuild nodes.
type DiscordGuildUpsertBulk struct {
	create *DiscordGuildCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DiscordGuild.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(discordguild.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DiscordGuildUpsertBulk) UpdateNewValues() *DiscordGuildUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(discordguild.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DiscordGuild.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DiscordGuildUpsertBulk) Ignore() *DiscordGuildUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DiscordGuildUpsertBulk) DoNothing() *DiscordGuildUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DiscordGuildCreateBulk.OnConflict
// documentation for more info.
func (u *DiscordGuildUpsertBulk) Update(set func(*DiscordGuildUpsert)) *DiscordGuildUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DiscordGuildUpsert{UpdateSet: update})
	}))
	return u
}

// SetDiscordid sets the "discordid" field.
func (u *DiscordGuildUpsertBulk) SetDiscordid(v string) *DiscordGuildUpsertBulk {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.SetDiscordid(v)
	})
}

// UpdateDiscordid sets the "discordid" field to the value that was provided on create.
func (u *DiscordGuildUpsertBulk) UpdateDiscordid() *DiscordGuildUpsertBulk {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.UpdateDiscordid()
	})
}

// SetName sets the "name" field.
func (u *DiscordGuildUpsertBulk) SetName(v string) *DiscordGuildUpsertBulk {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DiscordGuildUpsertBulk) UpdateName() *DiscordGuildUpsertBulk {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *DiscordGuildUpsertBulk) SetDescription(v string) *DiscordGuildUpsertBulk {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *DiscordGuildUpsertBulk) UpdateDescription() *DiscordGuildUpsertBulk {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *DiscordGuildUpsertBulk) ClearDescription() *DiscordGuildUpsertBulk {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.ClearDescription()
	})
}

// SetRulesChannelID sets the "rules_channel_id" field.
func (u *DiscordGuildUpsertBulk) SetRulesChannelID(v string) *DiscordGuildUpsertBulk {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.SetRulesChannelID(v)
	})
}

// UpdateRulesChannelID sets the "rules_channel_id" field to the value that was provided on create.
func (u *DiscordGuildUpsertBulk) UpdateRulesChannelID() *DiscordGuildUpsertBulk {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.UpdateRulesChannelID()
	})
}

// ClearRulesChannelID clears the value of the "rules_channel_id" field.
func (u *DiscordGuildUpsertBulk) ClearRulesChannelID() *DiscordGuildUpsertBulk {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.ClearRulesChannelID()
	})
}

// SetPublicUpdatesChannelID sets the "public_updates_channel_id" field.
func (u *DiscordGuildUpsertBulk) SetPublicUpdatesChannelID(v string) *DiscordGuildUpsertBulk {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.SetPublicUpdatesChannelID(v)
	})
}

// UpdatePublicUpdatesChannelID sets the "public_updates_channel_id" field to the value that was provided on create.
func (u *DiscordGuildUpsertBulk) UpdatePublicUpdatesChannelID() *DiscordGuildUpsertBulk {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.UpdatePublicUpdatesChannelID()
	})
}

// ClearPublicUpdatesChannelID clears the value of the "public_updates_channel_id" field.
func (u *DiscordGuildUpsertBulk) ClearPublicUpdatesChannelID() *DiscordGuildUpsertBulk {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.ClearPublicUpdatesChannelID()
	})
}

// SetNsfwLevel sets the "nsfw_level" field.
func (u *DiscordGuildUpsertBulk) SetNsfwLevel(v int) *DiscordGuildUpsertBulk {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.SetNsfwLevel(v)
	})
}

// AddNsfwLevel adds v to the "nsfw_level" field.
func (u *DiscordGuildUpsertBulk) AddNsfwLevel(v int) *DiscordGuildUpsertBulk {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.AddNsfwLevel(v)
	})
}

// UpdateNsfwLevel sets the "nsfw_level" field to the value that was provided on create.
func (u *DiscordGuildUpsertBulk) UpdateNsfwLevel() *DiscordGuildUpsertBulk {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.UpdateNsfwLevel()
	})
}

// ClearNsfwLevel clears the value of the "nsfw_level" field.
func (u *DiscordGuildUpsertBulk) ClearNsfwLevel() *DiscordGuildUpsertBulk {
	return u.Update(func(s *DiscordGuildUpsert) {
		s.ClearNsfwLevel()
	})
}

// Exec executes the query.
func (u *DiscordGuildUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DiscordGuildCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DiscordGuildCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DiscordGuildUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
