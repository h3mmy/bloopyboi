// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bwmarrin/discordgo"
	"gitlab.com/h3mmy/bloopyboi/bot/internal/ent/discordmessage"
)

// DiscordMessageCreate is the builder for creating a DiscordMessage entity.
type DiscordMessageCreate struct {
	config
	mutation *DiscordMessageMutation
	hooks    []Hook
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

// SetRaw sets the "raw" field.
func (dmc *DiscordMessageCreate) SetRaw(d discordgo.Message) *DiscordMessageCreate {
	dmc.mutation.SetRaw(d)
	return dmc
}

// SetID sets the "id" field.
func (dmc *DiscordMessageCreate) SetID(s string) *DiscordMessageCreate {
	dmc.mutation.SetID(s)
	return dmc
}

// Mutation returns the DiscordMessageMutation object of the builder.
func (dmc *DiscordMessageCreate) Mutation() *DiscordMessageMutation {
	return dmc.mutation
}

// Save creates the DiscordMessage in the database.
func (dmc *DiscordMessageCreate) Save(ctx context.Context) (*DiscordMessage, error) {
	var (
		err  error
		node *DiscordMessage
	)
	dmc.defaults()
	if len(dmc.hooks) == 0 {
		if err = dmc.check(); err != nil {
			return nil, err
		}
		node, err = dmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DiscordMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dmc.check(); err != nil {
				return nil, err
			}
			dmc.mutation = mutation
			if node, err = dmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dmc.hooks) - 1; i >= 0; i-- {
			if dmc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dmc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dmc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DiscordMessage)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DiscordMessageMutation", v)
		}
		node = nv
	}
	return node, err
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
	if _, ok := dmc.mutation.Raw(); !ok {
		return &ValidationError{Name: "raw", err: errors.New(`ent: missing required field "DiscordMessage.raw"`)}
	}
	return nil
}

func (dmc *DiscordMessageCreate) sqlSave(ctx context.Context) (*DiscordMessage, error) {
	_node, _spec := dmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected DiscordMessage.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (dmc *DiscordMessageCreate) createSpec() (*DiscordMessage, *sqlgraph.CreateSpec) {
	var (
		_node = &DiscordMessage{config: dmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: discordmessage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: discordmessage.FieldID,
			},
		}
	)
	if id, ok := dmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dmc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: discordmessage.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := dmc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: discordmessage.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := dmc.mutation.Raw(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: discordmessage.FieldRaw,
		})
		_node.Raw = value
	}
	return _node, _spec
}

// DiscordMessageCreateBulk is the builder for creating many DiscordMessage entities in bulk.
type DiscordMessageCreateBulk struct {
	config
	builders []*DiscordMessageCreate
}

// Save creates the DiscordMessage entities in the database.
func (dmcb *DiscordMessageCreateBulk) Save(ctx context.Context) ([]*DiscordMessage, error) {
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
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
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
