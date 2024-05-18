// Code generated by ent, DO NOT EDIT.

package discordmessagereaction

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/h3mmy/bloopyboi/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldEQ(FieldUpdateTime, v))
}

// Removed applies equality check predicate on the "removed" field. It's identical to RemovedEQ.
func Removed(v bool) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldEQ(FieldRemoved, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldLTE(FieldUpdateTime, v))
}

// RemovedEQ applies the EQ predicate on the "removed" field.
func RemovedEQ(v bool) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldEQ(FieldRemoved, v))
}

// RemovedNEQ applies the NEQ predicate on the "removed" field.
func RemovedNEQ(v bool) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.FieldNEQ(FieldRemoved, v))
}

// HasDiscordMessage applies the HasEdge predicate on the "discord_message" edge.
func HasDiscordMessage() predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DiscordMessageTable, DiscordMessageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDiscordMessageWith applies the HasEdge predicate on the "discord_message" edge with a given conditions (other predicates).
func HasDiscordMessageWith(preds ...predicate.DiscordMessage) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(func(s *sql.Selector) {
		step := newDiscordMessageStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.DiscordUser) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(func(s *sql.Selector) {
		step := newAuthorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DiscordMessageReaction) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DiscordMessageReaction) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DiscordMessageReaction) predicate.DiscordMessageReaction {
	return predicate.DiscordMessageReaction(sql.NotPredicates(p))
}
