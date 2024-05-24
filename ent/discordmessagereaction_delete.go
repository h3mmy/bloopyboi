// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/h3mmy/bloopyboi/ent/discordmessagereaction"
	"github.com/h3mmy/bloopyboi/ent/predicate"
)

// DiscordMessageReactionDelete is the builder for deleting a DiscordMessageReaction entity.
type DiscordMessageReactionDelete struct {
	config
	hooks    []Hook
	mutation *DiscordMessageReactionMutation
}

// Where appends a list predicates to the DiscordMessageReactionDelete builder.
func (dmrd *DiscordMessageReactionDelete) Where(ps ...predicate.DiscordMessageReaction) *DiscordMessageReactionDelete {
	dmrd.mutation.Where(ps...)
	return dmrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dmrd *DiscordMessageReactionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dmrd.sqlExec, dmrd.mutation, dmrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dmrd *DiscordMessageReactionDelete) ExecX(ctx context.Context) int {
	n, err := dmrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dmrd *DiscordMessageReactionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(discordmessagereaction.Table, sqlgraph.NewFieldSpec(discordmessagereaction.FieldID, field.TypeUUID))
	if ps := dmrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dmrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dmrd.mutation.done = true
	return affected, err
}

// DiscordMessageReactionDeleteOne is the builder for deleting a single DiscordMessageReaction entity.
type DiscordMessageReactionDeleteOne struct {
	dmrd *DiscordMessageReactionDelete
}

// Where appends a list predicates to the DiscordMessageReactionDelete builder.
func (dmrdo *DiscordMessageReactionDeleteOne) Where(ps ...predicate.DiscordMessageReaction) *DiscordMessageReactionDeleteOne {
	dmrdo.dmrd.mutation.Where(ps...)
	return dmrdo
}

// Exec executes the deletion query.
func (dmrdo *DiscordMessageReactionDeleteOne) Exec(ctx context.Context) error {
	n, err := dmrdo.dmrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{discordmessagereaction.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dmrdo *DiscordMessageReactionDeleteOne) ExecX(ctx context.Context) {
	if err := dmrdo.Exec(ctx); err != nil {
		panic(err)
	}
}