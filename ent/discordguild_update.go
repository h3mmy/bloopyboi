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
	"github.com/h3mmy/bloopyboi/ent/discordchannel"
	"github.com/h3mmy/bloopyboi/ent/discordguild"
	"github.com/h3mmy/bloopyboi/ent/discordmessage"
	"github.com/h3mmy/bloopyboi/ent/discorduser"
	"github.com/h3mmy/bloopyboi/ent/predicate"
)

// DiscordGuildUpdate is the builder for updating DiscordGuild entities.
type DiscordGuildUpdate struct {
	config
	hooks    []Hook
	mutation *DiscordGuildMutation
}

// Where appends a list predicates to the DiscordGuildUpdate builder.
func (dgu *DiscordGuildUpdate) Where(ps ...predicate.DiscordGuild) *DiscordGuildUpdate {
	dgu.mutation.Where(ps...)
	return dgu
}

// SetDiscordid sets the "discordid" field.
func (dgu *DiscordGuildUpdate) SetDiscordid(s string) *DiscordGuildUpdate {
	dgu.mutation.SetDiscordid(s)
	return dgu
}

// SetNillableDiscordid sets the "discordid" field if the given value is not nil.
func (dgu *DiscordGuildUpdate) SetNillableDiscordid(s *string) *DiscordGuildUpdate {
	if s != nil {
		dgu.SetDiscordid(*s)
	}
	return dgu
}

// SetName sets the "name" field.
func (dgu *DiscordGuildUpdate) SetName(s string) *DiscordGuildUpdate {
	dgu.mutation.SetName(s)
	return dgu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (dgu *DiscordGuildUpdate) SetNillableName(s *string) *DiscordGuildUpdate {
	if s != nil {
		dgu.SetName(*s)
	}
	return dgu
}

// SetDescription sets the "description" field.
func (dgu *DiscordGuildUpdate) SetDescription(s string) *DiscordGuildUpdate {
	dgu.mutation.SetDescription(s)
	return dgu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (dgu *DiscordGuildUpdate) SetNillableDescription(s *string) *DiscordGuildUpdate {
	if s != nil {
		dgu.SetDescription(*s)
	}
	return dgu
}

// ClearDescription clears the value of the "description" field.
func (dgu *DiscordGuildUpdate) ClearDescription() *DiscordGuildUpdate {
	dgu.mutation.ClearDescription()
	return dgu
}

// SetRulesChannelID sets the "rules_channel_id" field.
func (dgu *DiscordGuildUpdate) SetRulesChannelID(s string) *DiscordGuildUpdate {
	dgu.mutation.SetRulesChannelID(s)
	return dgu
}

// SetNillableRulesChannelID sets the "rules_channel_id" field if the given value is not nil.
func (dgu *DiscordGuildUpdate) SetNillableRulesChannelID(s *string) *DiscordGuildUpdate {
	if s != nil {
		dgu.SetRulesChannelID(*s)
	}
	return dgu
}

// ClearRulesChannelID clears the value of the "rules_channel_id" field.
func (dgu *DiscordGuildUpdate) ClearRulesChannelID() *DiscordGuildUpdate {
	dgu.mutation.ClearRulesChannelID()
	return dgu
}

// SetPublicUpdatesChannelID sets the "public_updates_channel_id" field.
func (dgu *DiscordGuildUpdate) SetPublicUpdatesChannelID(s string) *DiscordGuildUpdate {
	dgu.mutation.SetPublicUpdatesChannelID(s)
	return dgu
}

// SetNillablePublicUpdatesChannelID sets the "public_updates_channel_id" field if the given value is not nil.
func (dgu *DiscordGuildUpdate) SetNillablePublicUpdatesChannelID(s *string) *DiscordGuildUpdate {
	if s != nil {
		dgu.SetPublicUpdatesChannelID(*s)
	}
	return dgu
}

// ClearPublicUpdatesChannelID clears the value of the "public_updates_channel_id" field.
func (dgu *DiscordGuildUpdate) ClearPublicUpdatesChannelID() *DiscordGuildUpdate {
	dgu.mutation.ClearPublicUpdatesChannelID()
	return dgu
}

// SetNsfwLevel sets the "nsfw_level" field.
func (dgu *DiscordGuildUpdate) SetNsfwLevel(i int) *DiscordGuildUpdate {
	dgu.mutation.ResetNsfwLevel()
	dgu.mutation.SetNsfwLevel(i)
	return dgu
}

// SetNillableNsfwLevel sets the "nsfw_level" field if the given value is not nil.
func (dgu *DiscordGuildUpdate) SetNillableNsfwLevel(i *int) *DiscordGuildUpdate {
	if i != nil {
		dgu.SetNsfwLevel(*i)
	}
	return dgu
}

// AddNsfwLevel adds i to the "nsfw_level" field.
func (dgu *DiscordGuildUpdate) AddNsfwLevel(i int) *DiscordGuildUpdate {
	dgu.mutation.AddNsfwLevel(i)
	return dgu
}

// ClearNsfwLevel clears the value of the "nsfw_level" field.
func (dgu *DiscordGuildUpdate) ClearNsfwLevel() *DiscordGuildUpdate {
	dgu.mutation.ClearNsfwLevel()
	return dgu
}

// AddMemberIDs adds the "members" edge to the DiscordUser entity by IDs.
func (dgu *DiscordGuildUpdate) AddMemberIDs(ids ...uuid.UUID) *DiscordGuildUpdate {
	dgu.mutation.AddMemberIDs(ids...)
	return dgu
}

// AddMembers adds the "members" edges to the DiscordUser entity.
func (dgu *DiscordGuildUpdate) AddMembers(d ...*DiscordUser) *DiscordGuildUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dgu.AddMemberIDs(ids...)
}

// AddDiscordMessageIDs adds the "discord_messages" edge to the DiscordMessage entity by IDs.
func (dgu *DiscordGuildUpdate) AddDiscordMessageIDs(ids ...uuid.UUID) *DiscordGuildUpdate {
	dgu.mutation.AddDiscordMessageIDs(ids...)
	return dgu
}

// AddDiscordMessages adds the "discord_messages" edges to the DiscordMessage entity.
func (dgu *DiscordGuildUpdate) AddDiscordMessages(d ...*DiscordMessage) *DiscordGuildUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dgu.AddDiscordMessageIDs(ids...)
}

// AddGuildChannelIDs adds the "guild_channels" edge to the DiscordChannel entity by IDs.
func (dgu *DiscordGuildUpdate) AddGuildChannelIDs(ids ...uuid.UUID) *DiscordGuildUpdate {
	dgu.mutation.AddGuildChannelIDs(ids...)
	return dgu
}

// AddGuildChannels adds the "guild_channels" edges to the DiscordChannel entity.
func (dgu *DiscordGuildUpdate) AddGuildChannels(d ...*DiscordChannel) *DiscordGuildUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dgu.AddGuildChannelIDs(ids...)
}

// Mutation returns the DiscordGuildMutation object of the builder.
func (dgu *DiscordGuildUpdate) Mutation() *DiscordGuildMutation {
	return dgu.mutation
}

// ClearMembers clears all "members" edges to the DiscordUser entity.
func (dgu *DiscordGuildUpdate) ClearMembers() *DiscordGuildUpdate {
	dgu.mutation.ClearMembers()
	return dgu
}

// RemoveMemberIDs removes the "members" edge to DiscordUser entities by IDs.
func (dgu *DiscordGuildUpdate) RemoveMemberIDs(ids ...uuid.UUID) *DiscordGuildUpdate {
	dgu.mutation.RemoveMemberIDs(ids...)
	return dgu
}

// RemoveMembers removes "members" edges to DiscordUser entities.
func (dgu *DiscordGuildUpdate) RemoveMembers(d ...*DiscordUser) *DiscordGuildUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dgu.RemoveMemberIDs(ids...)
}

// ClearDiscordMessages clears all "discord_messages" edges to the DiscordMessage entity.
func (dgu *DiscordGuildUpdate) ClearDiscordMessages() *DiscordGuildUpdate {
	dgu.mutation.ClearDiscordMessages()
	return dgu
}

// RemoveDiscordMessageIDs removes the "discord_messages" edge to DiscordMessage entities by IDs.
func (dgu *DiscordGuildUpdate) RemoveDiscordMessageIDs(ids ...uuid.UUID) *DiscordGuildUpdate {
	dgu.mutation.RemoveDiscordMessageIDs(ids...)
	return dgu
}

// RemoveDiscordMessages removes "discord_messages" edges to DiscordMessage entities.
func (dgu *DiscordGuildUpdate) RemoveDiscordMessages(d ...*DiscordMessage) *DiscordGuildUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dgu.RemoveDiscordMessageIDs(ids...)
}

// ClearGuildChannels clears all "guild_channels" edges to the DiscordChannel entity.
func (dgu *DiscordGuildUpdate) ClearGuildChannels() *DiscordGuildUpdate {
	dgu.mutation.ClearGuildChannels()
	return dgu
}

// RemoveGuildChannelIDs removes the "guild_channels" edge to DiscordChannel entities by IDs.
func (dgu *DiscordGuildUpdate) RemoveGuildChannelIDs(ids ...uuid.UUID) *DiscordGuildUpdate {
	dgu.mutation.RemoveGuildChannelIDs(ids...)
	return dgu
}

// RemoveGuildChannels removes "guild_channels" edges to DiscordChannel entities.
func (dgu *DiscordGuildUpdate) RemoveGuildChannels(d ...*DiscordChannel) *DiscordGuildUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dgu.RemoveGuildChannelIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dgu *DiscordGuildUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, dgu.sqlSave, dgu.mutation, dgu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dgu *DiscordGuildUpdate) SaveX(ctx context.Context) int {
	affected, err := dgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dgu *DiscordGuildUpdate) Exec(ctx context.Context) error {
	_, err := dgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dgu *DiscordGuildUpdate) ExecX(ctx context.Context) {
	if err := dgu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dgu *DiscordGuildUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(discordguild.Table, discordguild.Columns, sqlgraph.NewFieldSpec(discordguild.FieldID, field.TypeUUID))
	if ps := dgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dgu.mutation.Discordid(); ok {
		_spec.SetField(discordguild.FieldDiscordid, field.TypeString, value)
	}
	if value, ok := dgu.mutation.Name(); ok {
		_spec.SetField(discordguild.FieldName, field.TypeString, value)
	}
	if value, ok := dgu.mutation.Description(); ok {
		_spec.SetField(discordguild.FieldDescription, field.TypeString, value)
	}
	if dgu.mutation.DescriptionCleared() {
		_spec.ClearField(discordguild.FieldDescription, field.TypeString)
	}
	if value, ok := dgu.mutation.RulesChannelID(); ok {
		_spec.SetField(discordguild.FieldRulesChannelID, field.TypeString, value)
	}
	if dgu.mutation.RulesChannelIDCleared() {
		_spec.ClearField(discordguild.FieldRulesChannelID, field.TypeString)
	}
	if value, ok := dgu.mutation.PublicUpdatesChannelID(); ok {
		_spec.SetField(discordguild.FieldPublicUpdatesChannelID, field.TypeString, value)
	}
	if dgu.mutation.PublicUpdatesChannelIDCleared() {
		_spec.ClearField(discordguild.FieldPublicUpdatesChannelID, field.TypeString)
	}
	if value, ok := dgu.mutation.NsfwLevel(); ok {
		_spec.SetField(discordguild.FieldNsfwLevel, field.TypeInt, value)
	}
	if value, ok := dgu.mutation.AddedNsfwLevel(); ok {
		_spec.AddField(discordguild.FieldNsfwLevel, field.TypeInt, value)
	}
	if dgu.mutation.NsfwLevelCleared() {
		_spec.ClearField(discordguild.FieldNsfwLevel, field.TypeInt)
	}
	if dgu.mutation.MembersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dgu.mutation.RemovedMembersIDs(); len(nodes) > 0 && !dgu.mutation.MembersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dgu.mutation.MembersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dgu.mutation.DiscordMessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dgu.mutation.RemovedDiscordMessagesIDs(); len(nodes) > 0 && !dgu.mutation.DiscordMessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dgu.mutation.DiscordMessagesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dgu.mutation.GuildChannelsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dgu.mutation.RemovedGuildChannelsIDs(); len(nodes) > 0 && !dgu.mutation.GuildChannelsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dgu.mutation.GuildChannelsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{discordguild.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dgu.mutation.done = true
	return n, nil
}

// DiscordGuildUpdateOne is the builder for updating a single DiscordGuild entity.
type DiscordGuildUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DiscordGuildMutation
}

// SetDiscordid sets the "discordid" field.
func (dguo *DiscordGuildUpdateOne) SetDiscordid(s string) *DiscordGuildUpdateOne {
	dguo.mutation.SetDiscordid(s)
	return dguo
}

// SetNillableDiscordid sets the "discordid" field if the given value is not nil.
func (dguo *DiscordGuildUpdateOne) SetNillableDiscordid(s *string) *DiscordGuildUpdateOne {
	if s != nil {
		dguo.SetDiscordid(*s)
	}
	return dguo
}

// SetName sets the "name" field.
func (dguo *DiscordGuildUpdateOne) SetName(s string) *DiscordGuildUpdateOne {
	dguo.mutation.SetName(s)
	return dguo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (dguo *DiscordGuildUpdateOne) SetNillableName(s *string) *DiscordGuildUpdateOne {
	if s != nil {
		dguo.SetName(*s)
	}
	return dguo
}

// SetDescription sets the "description" field.
func (dguo *DiscordGuildUpdateOne) SetDescription(s string) *DiscordGuildUpdateOne {
	dguo.mutation.SetDescription(s)
	return dguo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (dguo *DiscordGuildUpdateOne) SetNillableDescription(s *string) *DiscordGuildUpdateOne {
	if s != nil {
		dguo.SetDescription(*s)
	}
	return dguo
}

// ClearDescription clears the value of the "description" field.
func (dguo *DiscordGuildUpdateOne) ClearDescription() *DiscordGuildUpdateOne {
	dguo.mutation.ClearDescription()
	return dguo
}

// SetRulesChannelID sets the "rules_channel_id" field.
func (dguo *DiscordGuildUpdateOne) SetRulesChannelID(s string) *DiscordGuildUpdateOne {
	dguo.mutation.SetRulesChannelID(s)
	return dguo
}

// SetNillableRulesChannelID sets the "rules_channel_id" field if the given value is not nil.
func (dguo *DiscordGuildUpdateOne) SetNillableRulesChannelID(s *string) *DiscordGuildUpdateOne {
	if s != nil {
		dguo.SetRulesChannelID(*s)
	}
	return dguo
}

// ClearRulesChannelID clears the value of the "rules_channel_id" field.
func (dguo *DiscordGuildUpdateOne) ClearRulesChannelID() *DiscordGuildUpdateOne {
	dguo.mutation.ClearRulesChannelID()
	return dguo
}

// SetPublicUpdatesChannelID sets the "public_updates_channel_id" field.
func (dguo *DiscordGuildUpdateOne) SetPublicUpdatesChannelID(s string) *DiscordGuildUpdateOne {
	dguo.mutation.SetPublicUpdatesChannelID(s)
	return dguo
}

// SetNillablePublicUpdatesChannelID sets the "public_updates_channel_id" field if the given value is not nil.
func (dguo *DiscordGuildUpdateOne) SetNillablePublicUpdatesChannelID(s *string) *DiscordGuildUpdateOne {
	if s != nil {
		dguo.SetPublicUpdatesChannelID(*s)
	}
	return dguo
}

// ClearPublicUpdatesChannelID clears the value of the "public_updates_channel_id" field.
func (dguo *DiscordGuildUpdateOne) ClearPublicUpdatesChannelID() *DiscordGuildUpdateOne {
	dguo.mutation.ClearPublicUpdatesChannelID()
	return dguo
}

// SetNsfwLevel sets the "nsfw_level" field.
func (dguo *DiscordGuildUpdateOne) SetNsfwLevel(i int) *DiscordGuildUpdateOne {
	dguo.mutation.ResetNsfwLevel()
	dguo.mutation.SetNsfwLevel(i)
	return dguo
}

// SetNillableNsfwLevel sets the "nsfw_level" field if the given value is not nil.
func (dguo *DiscordGuildUpdateOne) SetNillableNsfwLevel(i *int) *DiscordGuildUpdateOne {
	if i != nil {
		dguo.SetNsfwLevel(*i)
	}
	return dguo
}

// AddNsfwLevel adds i to the "nsfw_level" field.
func (dguo *DiscordGuildUpdateOne) AddNsfwLevel(i int) *DiscordGuildUpdateOne {
	dguo.mutation.AddNsfwLevel(i)
	return dguo
}

// ClearNsfwLevel clears the value of the "nsfw_level" field.
func (dguo *DiscordGuildUpdateOne) ClearNsfwLevel() *DiscordGuildUpdateOne {
	dguo.mutation.ClearNsfwLevel()
	return dguo
}

// AddMemberIDs adds the "members" edge to the DiscordUser entity by IDs.
func (dguo *DiscordGuildUpdateOne) AddMemberIDs(ids ...uuid.UUID) *DiscordGuildUpdateOne {
	dguo.mutation.AddMemberIDs(ids...)
	return dguo
}

// AddMembers adds the "members" edges to the DiscordUser entity.
func (dguo *DiscordGuildUpdateOne) AddMembers(d ...*DiscordUser) *DiscordGuildUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dguo.AddMemberIDs(ids...)
}

// AddDiscordMessageIDs adds the "discord_messages" edge to the DiscordMessage entity by IDs.
func (dguo *DiscordGuildUpdateOne) AddDiscordMessageIDs(ids ...uuid.UUID) *DiscordGuildUpdateOne {
	dguo.mutation.AddDiscordMessageIDs(ids...)
	return dguo
}

// AddDiscordMessages adds the "discord_messages" edges to the DiscordMessage entity.
func (dguo *DiscordGuildUpdateOne) AddDiscordMessages(d ...*DiscordMessage) *DiscordGuildUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dguo.AddDiscordMessageIDs(ids...)
}

// AddGuildChannelIDs adds the "guild_channels" edge to the DiscordChannel entity by IDs.
func (dguo *DiscordGuildUpdateOne) AddGuildChannelIDs(ids ...uuid.UUID) *DiscordGuildUpdateOne {
	dguo.mutation.AddGuildChannelIDs(ids...)
	return dguo
}

// AddGuildChannels adds the "guild_channels" edges to the DiscordChannel entity.
func (dguo *DiscordGuildUpdateOne) AddGuildChannels(d ...*DiscordChannel) *DiscordGuildUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dguo.AddGuildChannelIDs(ids...)
}

// Mutation returns the DiscordGuildMutation object of the builder.
func (dguo *DiscordGuildUpdateOne) Mutation() *DiscordGuildMutation {
	return dguo.mutation
}

// ClearMembers clears all "members" edges to the DiscordUser entity.
func (dguo *DiscordGuildUpdateOne) ClearMembers() *DiscordGuildUpdateOne {
	dguo.mutation.ClearMembers()
	return dguo
}

// RemoveMemberIDs removes the "members" edge to DiscordUser entities by IDs.
func (dguo *DiscordGuildUpdateOne) RemoveMemberIDs(ids ...uuid.UUID) *DiscordGuildUpdateOne {
	dguo.mutation.RemoveMemberIDs(ids...)
	return dguo
}

// RemoveMembers removes "members" edges to DiscordUser entities.
func (dguo *DiscordGuildUpdateOne) RemoveMembers(d ...*DiscordUser) *DiscordGuildUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dguo.RemoveMemberIDs(ids...)
}

// ClearDiscordMessages clears all "discord_messages" edges to the DiscordMessage entity.
func (dguo *DiscordGuildUpdateOne) ClearDiscordMessages() *DiscordGuildUpdateOne {
	dguo.mutation.ClearDiscordMessages()
	return dguo
}

// RemoveDiscordMessageIDs removes the "discord_messages" edge to DiscordMessage entities by IDs.
func (dguo *DiscordGuildUpdateOne) RemoveDiscordMessageIDs(ids ...uuid.UUID) *DiscordGuildUpdateOne {
	dguo.mutation.RemoveDiscordMessageIDs(ids...)
	return dguo
}

// RemoveDiscordMessages removes "discord_messages" edges to DiscordMessage entities.
func (dguo *DiscordGuildUpdateOne) RemoveDiscordMessages(d ...*DiscordMessage) *DiscordGuildUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dguo.RemoveDiscordMessageIDs(ids...)
}

// ClearGuildChannels clears all "guild_channels" edges to the DiscordChannel entity.
func (dguo *DiscordGuildUpdateOne) ClearGuildChannels() *DiscordGuildUpdateOne {
	dguo.mutation.ClearGuildChannels()
	return dguo
}

// RemoveGuildChannelIDs removes the "guild_channels" edge to DiscordChannel entities by IDs.
func (dguo *DiscordGuildUpdateOne) RemoveGuildChannelIDs(ids ...uuid.UUID) *DiscordGuildUpdateOne {
	dguo.mutation.RemoveGuildChannelIDs(ids...)
	return dguo
}

// RemoveGuildChannels removes "guild_channels" edges to DiscordChannel entities.
func (dguo *DiscordGuildUpdateOne) RemoveGuildChannels(d ...*DiscordChannel) *DiscordGuildUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dguo.RemoveGuildChannelIDs(ids...)
}

// Where appends a list predicates to the DiscordGuildUpdate builder.
func (dguo *DiscordGuildUpdateOne) Where(ps ...predicate.DiscordGuild) *DiscordGuildUpdateOne {
	dguo.mutation.Where(ps...)
	return dguo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dguo *DiscordGuildUpdateOne) Select(field string, fields ...string) *DiscordGuildUpdateOne {
	dguo.fields = append([]string{field}, fields...)
	return dguo
}

// Save executes the query and returns the updated DiscordGuild entity.
func (dguo *DiscordGuildUpdateOne) Save(ctx context.Context) (*DiscordGuild, error) {
	return withHooks(ctx, dguo.sqlSave, dguo.mutation, dguo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dguo *DiscordGuildUpdateOne) SaveX(ctx context.Context) *DiscordGuild {
	node, err := dguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dguo *DiscordGuildUpdateOne) Exec(ctx context.Context) error {
	_, err := dguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dguo *DiscordGuildUpdateOne) ExecX(ctx context.Context) {
	if err := dguo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dguo *DiscordGuildUpdateOne) sqlSave(ctx context.Context) (_node *DiscordGuild, err error) {
	_spec := sqlgraph.NewUpdateSpec(discordguild.Table, discordguild.Columns, sqlgraph.NewFieldSpec(discordguild.FieldID, field.TypeUUID))
	id, ok := dguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DiscordGuild.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, discordguild.FieldID)
		for _, f := range fields {
			if !discordguild.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != discordguild.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dguo.mutation.Discordid(); ok {
		_spec.SetField(discordguild.FieldDiscordid, field.TypeString, value)
	}
	if value, ok := dguo.mutation.Name(); ok {
		_spec.SetField(discordguild.FieldName, field.TypeString, value)
	}
	if value, ok := dguo.mutation.Description(); ok {
		_spec.SetField(discordguild.FieldDescription, field.TypeString, value)
	}
	if dguo.mutation.DescriptionCleared() {
		_spec.ClearField(discordguild.FieldDescription, field.TypeString)
	}
	if value, ok := dguo.mutation.RulesChannelID(); ok {
		_spec.SetField(discordguild.FieldRulesChannelID, field.TypeString, value)
	}
	if dguo.mutation.RulesChannelIDCleared() {
		_spec.ClearField(discordguild.FieldRulesChannelID, field.TypeString)
	}
	if value, ok := dguo.mutation.PublicUpdatesChannelID(); ok {
		_spec.SetField(discordguild.FieldPublicUpdatesChannelID, field.TypeString, value)
	}
	if dguo.mutation.PublicUpdatesChannelIDCleared() {
		_spec.ClearField(discordguild.FieldPublicUpdatesChannelID, field.TypeString)
	}
	if value, ok := dguo.mutation.NsfwLevel(); ok {
		_spec.SetField(discordguild.FieldNsfwLevel, field.TypeInt, value)
	}
	if value, ok := dguo.mutation.AddedNsfwLevel(); ok {
		_spec.AddField(discordguild.FieldNsfwLevel, field.TypeInt, value)
	}
	if dguo.mutation.NsfwLevelCleared() {
		_spec.ClearField(discordguild.FieldNsfwLevel, field.TypeInt)
	}
	if dguo.mutation.MembersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dguo.mutation.RemovedMembersIDs(); len(nodes) > 0 && !dguo.mutation.MembersCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dguo.mutation.MembersIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dguo.mutation.DiscordMessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dguo.mutation.RemovedDiscordMessagesIDs(); len(nodes) > 0 && !dguo.mutation.DiscordMessagesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dguo.mutation.DiscordMessagesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if dguo.mutation.GuildChannelsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dguo.mutation.RemovedGuildChannelsIDs(); len(nodes) > 0 && !dguo.mutation.GuildChannelsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dguo.mutation.GuildChannelsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DiscordGuild{config: dguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{discordguild.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dguo.mutation.done = true
	return _node, nil
}
