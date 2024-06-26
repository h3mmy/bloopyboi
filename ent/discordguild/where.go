// Code generated by ent, DO NOT EDIT.

package discordguild

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/h3mmy/bloopyboi/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldLTE(FieldID, id))
}

// Discordid applies equality check predicate on the "discordid" field. It's identical to DiscordidEQ.
func Discordid(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldEQ(FieldDiscordid, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldEQ(FieldDescription, v))
}

// RulesChannelID applies equality check predicate on the "rules_channel_id" field. It's identical to RulesChannelIDEQ.
func RulesChannelID(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldEQ(FieldRulesChannelID, v))
}

// PublicUpdatesChannelID applies equality check predicate on the "public_updates_channel_id" field. It's identical to PublicUpdatesChannelIDEQ.
func PublicUpdatesChannelID(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldEQ(FieldPublicUpdatesChannelID, v))
}

// NsfwLevel applies equality check predicate on the "nsfw_level" field. It's identical to NsfwLevelEQ.
func NsfwLevel(v int) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldEQ(FieldNsfwLevel, v))
}

// DiscordidEQ applies the EQ predicate on the "discordid" field.
func DiscordidEQ(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldEQ(FieldDiscordid, v))
}

// DiscordidNEQ applies the NEQ predicate on the "discordid" field.
func DiscordidNEQ(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldNEQ(FieldDiscordid, v))
}

// DiscordidIn applies the In predicate on the "discordid" field.
func DiscordidIn(vs ...string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldIn(FieldDiscordid, vs...))
}

// DiscordidNotIn applies the NotIn predicate on the "discordid" field.
func DiscordidNotIn(vs ...string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldNotIn(FieldDiscordid, vs...))
}

// DiscordidGT applies the GT predicate on the "discordid" field.
func DiscordidGT(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldGT(FieldDiscordid, v))
}

// DiscordidGTE applies the GTE predicate on the "discordid" field.
func DiscordidGTE(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldGTE(FieldDiscordid, v))
}

// DiscordidLT applies the LT predicate on the "discordid" field.
func DiscordidLT(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldLT(FieldDiscordid, v))
}

// DiscordidLTE applies the LTE predicate on the "discordid" field.
func DiscordidLTE(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldLTE(FieldDiscordid, v))
}

// DiscordidContains applies the Contains predicate on the "discordid" field.
func DiscordidContains(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldContains(FieldDiscordid, v))
}

// DiscordidHasPrefix applies the HasPrefix predicate on the "discordid" field.
func DiscordidHasPrefix(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldHasPrefix(FieldDiscordid, v))
}

// DiscordidHasSuffix applies the HasSuffix predicate on the "discordid" field.
func DiscordidHasSuffix(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldHasSuffix(FieldDiscordid, v))
}

// DiscordidEqualFold applies the EqualFold predicate on the "discordid" field.
func DiscordidEqualFold(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldEqualFold(FieldDiscordid, v))
}

// DiscordidContainsFold applies the ContainsFold predicate on the "discordid" field.
func DiscordidContainsFold(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldContainsFold(FieldDiscordid, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldContainsFold(FieldDescription, v))
}

// RulesChannelIDEQ applies the EQ predicate on the "rules_channel_id" field.
func RulesChannelIDEQ(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldEQ(FieldRulesChannelID, v))
}

// RulesChannelIDNEQ applies the NEQ predicate on the "rules_channel_id" field.
func RulesChannelIDNEQ(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldNEQ(FieldRulesChannelID, v))
}

// RulesChannelIDIn applies the In predicate on the "rules_channel_id" field.
func RulesChannelIDIn(vs ...string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldIn(FieldRulesChannelID, vs...))
}

// RulesChannelIDNotIn applies the NotIn predicate on the "rules_channel_id" field.
func RulesChannelIDNotIn(vs ...string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldNotIn(FieldRulesChannelID, vs...))
}

// RulesChannelIDGT applies the GT predicate on the "rules_channel_id" field.
func RulesChannelIDGT(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldGT(FieldRulesChannelID, v))
}

// RulesChannelIDGTE applies the GTE predicate on the "rules_channel_id" field.
func RulesChannelIDGTE(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldGTE(FieldRulesChannelID, v))
}

// RulesChannelIDLT applies the LT predicate on the "rules_channel_id" field.
func RulesChannelIDLT(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldLT(FieldRulesChannelID, v))
}

// RulesChannelIDLTE applies the LTE predicate on the "rules_channel_id" field.
func RulesChannelIDLTE(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldLTE(FieldRulesChannelID, v))
}

// RulesChannelIDContains applies the Contains predicate on the "rules_channel_id" field.
func RulesChannelIDContains(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldContains(FieldRulesChannelID, v))
}

// RulesChannelIDHasPrefix applies the HasPrefix predicate on the "rules_channel_id" field.
func RulesChannelIDHasPrefix(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldHasPrefix(FieldRulesChannelID, v))
}

// RulesChannelIDHasSuffix applies the HasSuffix predicate on the "rules_channel_id" field.
func RulesChannelIDHasSuffix(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldHasSuffix(FieldRulesChannelID, v))
}

// RulesChannelIDIsNil applies the IsNil predicate on the "rules_channel_id" field.
func RulesChannelIDIsNil() predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldIsNull(FieldRulesChannelID))
}

// RulesChannelIDNotNil applies the NotNil predicate on the "rules_channel_id" field.
func RulesChannelIDNotNil() predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldNotNull(FieldRulesChannelID))
}

// RulesChannelIDEqualFold applies the EqualFold predicate on the "rules_channel_id" field.
func RulesChannelIDEqualFold(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldEqualFold(FieldRulesChannelID, v))
}

// RulesChannelIDContainsFold applies the ContainsFold predicate on the "rules_channel_id" field.
func RulesChannelIDContainsFold(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldContainsFold(FieldRulesChannelID, v))
}

// PublicUpdatesChannelIDEQ applies the EQ predicate on the "public_updates_channel_id" field.
func PublicUpdatesChannelIDEQ(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldEQ(FieldPublicUpdatesChannelID, v))
}

// PublicUpdatesChannelIDNEQ applies the NEQ predicate on the "public_updates_channel_id" field.
func PublicUpdatesChannelIDNEQ(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldNEQ(FieldPublicUpdatesChannelID, v))
}

// PublicUpdatesChannelIDIn applies the In predicate on the "public_updates_channel_id" field.
func PublicUpdatesChannelIDIn(vs ...string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldIn(FieldPublicUpdatesChannelID, vs...))
}

// PublicUpdatesChannelIDNotIn applies the NotIn predicate on the "public_updates_channel_id" field.
func PublicUpdatesChannelIDNotIn(vs ...string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldNotIn(FieldPublicUpdatesChannelID, vs...))
}

// PublicUpdatesChannelIDGT applies the GT predicate on the "public_updates_channel_id" field.
func PublicUpdatesChannelIDGT(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldGT(FieldPublicUpdatesChannelID, v))
}

// PublicUpdatesChannelIDGTE applies the GTE predicate on the "public_updates_channel_id" field.
func PublicUpdatesChannelIDGTE(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldGTE(FieldPublicUpdatesChannelID, v))
}

// PublicUpdatesChannelIDLT applies the LT predicate on the "public_updates_channel_id" field.
func PublicUpdatesChannelIDLT(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldLT(FieldPublicUpdatesChannelID, v))
}

// PublicUpdatesChannelIDLTE applies the LTE predicate on the "public_updates_channel_id" field.
func PublicUpdatesChannelIDLTE(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldLTE(FieldPublicUpdatesChannelID, v))
}

// PublicUpdatesChannelIDContains applies the Contains predicate on the "public_updates_channel_id" field.
func PublicUpdatesChannelIDContains(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldContains(FieldPublicUpdatesChannelID, v))
}

// PublicUpdatesChannelIDHasPrefix applies the HasPrefix predicate on the "public_updates_channel_id" field.
func PublicUpdatesChannelIDHasPrefix(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldHasPrefix(FieldPublicUpdatesChannelID, v))
}

// PublicUpdatesChannelIDHasSuffix applies the HasSuffix predicate on the "public_updates_channel_id" field.
func PublicUpdatesChannelIDHasSuffix(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldHasSuffix(FieldPublicUpdatesChannelID, v))
}

// PublicUpdatesChannelIDIsNil applies the IsNil predicate on the "public_updates_channel_id" field.
func PublicUpdatesChannelIDIsNil() predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldIsNull(FieldPublicUpdatesChannelID))
}

// PublicUpdatesChannelIDNotNil applies the NotNil predicate on the "public_updates_channel_id" field.
func PublicUpdatesChannelIDNotNil() predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldNotNull(FieldPublicUpdatesChannelID))
}

// PublicUpdatesChannelIDEqualFold applies the EqualFold predicate on the "public_updates_channel_id" field.
func PublicUpdatesChannelIDEqualFold(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldEqualFold(FieldPublicUpdatesChannelID, v))
}

// PublicUpdatesChannelIDContainsFold applies the ContainsFold predicate on the "public_updates_channel_id" field.
func PublicUpdatesChannelIDContainsFold(v string) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldContainsFold(FieldPublicUpdatesChannelID, v))
}

// NsfwLevelEQ applies the EQ predicate on the "nsfw_level" field.
func NsfwLevelEQ(v int) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldEQ(FieldNsfwLevel, v))
}

// NsfwLevelNEQ applies the NEQ predicate on the "nsfw_level" field.
func NsfwLevelNEQ(v int) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldNEQ(FieldNsfwLevel, v))
}

// NsfwLevelIn applies the In predicate on the "nsfw_level" field.
func NsfwLevelIn(vs ...int) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldIn(FieldNsfwLevel, vs...))
}

// NsfwLevelNotIn applies the NotIn predicate on the "nsfw_level" field.
func NsfwLevelNotIn(vs ...int) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldNotIn(FieldNsfwLevel, vs...))
}

// NsfwLevelGT applies the GT predicate on the "nsfw_level" field.
func NsfwLevelGT(v int) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldGT(FieldNsfwLevel, v))
}

// NsfwLevelGTE applies the GTE predicate on the "nsfw_level" field.
func NsfwLevelGTE(v int) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldGTE(FieldNsfwLevel, v))
}

// NsfwLevelLT applies the LT predicate on the "nsfw_level" field.
func NsfwLevelLT(v int) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldLT(FieldNsfwLevel, v))
}

// NsfwLevelLTE applies the LTE predicate on the "nsfw_level" field.
func NsfwLevelLTE(v int) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldLTE(FieldNsfwLevel, v))
}

// NsfwLevelIsNil applies the IsNil predicate on the "nsfw_level" field.
func NsfwLevelIsNil() predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldIsNull(FieldNsfwLevel))
}

// NsfwLevelNotNil applies the NotNil predicate on the "nsfw_level" field.
func NsfwLevelNotNil() predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.FieldNotNull(FieldNsfwLevel))
}

// HasMembers applies the HasEdge predicate on the "members" edge.
func HasMembers() predicate.DiscordGuild {
	return predicate.DiscordGuild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, MembersTable, MembersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMembersWith applies the HasEdge predicate on the "members" edge with a given conditions (other predicates).
func HasMembersWith(preds ...predicate.DiscordUser) predicate.DiscordGuild {
	return predicate.DiscordGuild(func(s *sql.Selector) {
		step := newMembersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDiscordMessages applies the HasEdge predicate on the "discord_messages" edge.
func HasDiscordMessages() predicate.DiscordGuild {
	return predicate.DiscordGuild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DiscordMessagesTable, DiscordMessagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDiscordMessagesWith applies the HasEdge predicate on the "discord_messages" edge with a given conditions (other predicates).
func HasDiscordMessagesWith(preds ...predicate.DiscordMessage) predicate.DiscordGuild {
	return predicate.DiscordGuild(func(s *sql.Selector) {
		step := newDiscordMessagesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGuildChannels applies the HasEdge predicate on the "guild_channels" edge.
func HasGuildChannels() predicate.DiscordGuild {
	return predicate.DiscordGuild(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, GuildChannelsTable, GuildChannelsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGuildChannelsWith applies the HasEdge predicate on the "guild_channels" edge with a given conditions (other predicates).
func HasGuildChannelsWith(preds ...predicate.DiscordChannel) predicate.DiscordGuild {
	return predicate.DiscordGuild(func(s *sql.Selector) {
		step := newGuildChannelsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DiscordGuild) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DiscordGuild) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DiscordGuild) predicate.DiscordGuild {
	return predicate.DiscordGuild(sql.NotPredicates(p))
}
