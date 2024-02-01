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
	"github.com/h3mmy/bloopyboi/ent/book"
	"github.com/h3mmy/bloopyboi/ent/bookauthor"
)

// BookAuthorCreate is the builder for creating a BookAuthor entity.
type BookAuthorCreate struct {
	config
	mutation *BookAuthorMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetFullName sets the "full_name" field.
func (bac *BookAuthorCreate) SetFullName(s string) *BookAuthorCreate {
	bac.mutation.SetFullName(s)
	return bac
}

// SetID sets the "id" field.
func (bac *BookAuthorCreate) SetID(u uuid.UUID) *BookAuthorCreate {
	bac.mutation.SetID(u)
	return bac
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (bac *BookAuthorCreate) AddBookIDs(ids ...uuid.UUID) *BookAuthorCreate {
	bac.mutation.AddBookIDs(ids...)
	return bac
}

// AddBooks adds the "books" edges to the Book entity.
func (bac *BookAuthorCreate) AddBooks(b ...*Book) *BookAuthorCreate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bac.AddBookIDs(ids...)
}

// Mutation returns the BookAuthorMutation object of the builder.
func (bac *BookAuthorCreate) Mutation() *BookAuthorMutation {
	return bac.mutation
}

// Save creates the BookAuthor in the database.
func (bac *BookAuthorCreate) Save(ctx context.Context) (*BookAuthor, error) {
	return withHooks(ctx, bac.sqlSave, bac.mutation, bac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bac *BookAuthorCreate) SaveX(ctx context.Context) *BookAuthor {
	v, err := bac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bac *BookAuthorCreate) Exec(ctx context.Context) error {
	_, err := bac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bac *BookAuthorCreate) ExecX(ctx context.Context) {
	if err := bac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bac *BookAuthorCreate) check() error {
	if _, ok := bac.mutation.FullName(); !ok {
		return &ValidationError{Name: "full_name", err: errors.New(`ent: missing required field "BookAuthor.full_name"`)}
	}
	return nil
}

func (bac *BookAuthorCreate) sqlSave(ctx context.Context) (*BookAuthor, error) {
	if err := bac.check(); err != nil {
		return nil, err
	}
	_node, _spec := bac.createSpec()
	if err := sqlgraph.CreateNode(ctx, bac.driver, _spec); err != nil {
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
	bac.mutation.id = &_node.ID
	bac.mutation.done = true
	return _node, nil
}

func (bac *BookAuthorCreate) createSpec() (*BookAuthor, *sqlgraph.CreateSpec) {
	var (
		_node = &BookAuthor{config: bac.config}
		_spec = sqlgraph.NewCreateSpec(bookauthor.Table, sqlgraph.NewFieldSpec(bookauthor.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = bac.conflict
	if id, ok := bac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := bac.mutation.FullName(); ok {
		_spec.SetField(bookauthor.FieldFullName, field.TypeString, value)
		_node.FullName = value
	}
	if nodes := bac.mutation.BooksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   bookauthor.BooksTable,
			Columns: bookauthor.BooksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(book.FieldID, field.TypeUUID),
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
//	client.BookAuthor.Create().
//		SetFullName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BookAuthorUpsert) {
//			SetFullName(v+v).
//		}).
//		Exec(ctx)
func (bac *BookAuthorCreate) OnConflict(opts ...sql.ConflictOption) *BookAuthorUpsertOne {
	bac.conflict = opts
	return &BookAuthorUpsertOne{
		create: bac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.BookAuthor.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (bac *BookAuthorCreate) OnConflictColumns(columns ...string) *BookAuthorUpsertOne {
	bac.conflict = append(bac.conflict, sql.ConflictColumns(columns...))
	return &BookAuthorUpsertOne{
		create: bac,
	}
}

type (
	// BookAuthorUpsertOne is the builder for "upsert"-ing
	//  one BookAuthor node.
	BookAuthorUpsertOne struct {
		create *BookAuthorCreate
	}

	// BookAuthorUpsert is the "OnConflict" setter.
	BookAuthorUpsert struct {
		*sql.UpdateSet
	}
)

// SetFullName sets the "full_name" field.
func (u *BookAuthorUpsert) SetFullName(v string) *BookAuthorUpsert {
	u.Set(bookauthor.FieldFullName, v)
	return u
}

// UpdateFullName sets the "full_name" field to the value that was provided on create.
func (u *BookAuthorUpsert) UpdateFullName() *BookAuthorUpsert {
	u.SetExcluded(bookauthor.FieldFullName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.BookAuthor.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(bookauthor.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BookAuthorUpsertOne) UpdateNewValues() *BookAuthorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(bookauthor.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.BookAuthor.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *BookAuthorUpsertOne) Ignore() *BookAuthorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BookAuthorUpsertOne) DoNothing() *BookAuthorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BookAuthorCreate.OnConflict
// documentation for more info.
func (u *BookAuthorUpsertOne) Update(set func(*BookAuthorUpsert)) *BookAuthorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BookAuthorUpsert{UpdateSet: update})
	}))
	return u
}

// SetFullName sets the "full_name" field.
func (u *BookAuthorUpsertOne) SetFullName(v string) *BookAuthorUpsertOne {
	return u.Update(func(s *BookAuthorUpsert) {
		s.SetFullName(v)
	})
}

// UpdateFullName sets the "full_name" field to the value that was provided on create.
func (u *BookAuthorUpsertOne) UpdateFullName() *BookAuthorUpsertOne {
	return u.Update(func(s *BookAuthorUpsert) {
		s.UpdateFullName()
	})
}

// Exec executes the query.
func (u *BookAuthorUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BookAuthorCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BookAuthorUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *BookAuthorUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: BookAuthorUpsertOne.ID is not supported by MySQL driver. Use BookAuthorUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *BookAuthorUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// BookAuthorCreateBulk is the builder for creating many BookAuthor entities in bulk.
type BookAuthorCreateBulk struct {
	config
	builders []*BookAuthorCreate
	conflict []sql.ConflictOption
}

// Save creates the BookAuthor entities in the database.
func (bacb *BookAuthorCreateBulk) Save(ctx context.Context) ([]*BookAuthor, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bacb.builders))
	nodes := make([]*BookAuthor, len(bacb.builders))
	mutators := make([]Mutator, len(bacb.builders))
	for i := range bacb.builders {
		func(i int, root context.Context) {
			builder := bacb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookAuthorMutation)
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
					_, err = mutators[i+1].Mutate(root, bacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = bacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bacb *BookAuthorCreateBulk) SaveX(ctx context.Context) []*BookAuthor {
	v, err := bacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bacb *BookAuthorCreateBulk) Exec(ctx context.Context) error {
	_, err := bacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bacb *BookAuthorCreateBulk) ExecX(ctx context.Context) {
	if err := bacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.BookAuthor.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BookAuthorUpsert) {
//			SetFullName(v+v).
//		}).
//		Exec(ctx)
func (bacb *BookAuthorCreateBulk) OnConflict(opts ...sql.ConflictOption) *BookAuthorUpsertBulk {
	bacb.conflict = opts
	return &BookAuthorUpsertBulk{
		create: bacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.BookAuthor.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (bacb *BookAuthorCreateBulk) OnConflictColumns(columns ...string) *BookAuthorUpsertBulk {
	bacb.conflict = append(bacb.conflict, sql.ConflictColumns(columns...))
	return &BookAuthorUpsertBulk{
		create: bacb,
	}
}

// BookAuthorUpsertBulk is the builder for "upsert"-ing
// a bulk of BookAuthor nodes.
type BookAuthorUpsertBulk struct {
	create *BookAuthorCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.BookAuthor.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(bookauthor.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BookAuthorUpsertBulk) UpdateNewValues() *BookAuthorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(bookauthor.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.BookAuthor.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *BookAuthorUpsertBulk) Ignore() *BookAuthorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BookAuthorUpsertBulk) DoNothing() *BookAuthorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BookAuthorCreateBulk.OnConflict
// documentation for more info.
func (u *BookAuthorUpsertBulk) Update(set func(*BookAuthorUpsert)) *BookAuthorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BookAuthorUpsert{UpdateSet: update})
	}))
	return u
}

// SetFullName sets the "full_name" field.
func (u *BookAuthorUpsertBulk) SetFullName(v string) *BookAuthorUpsertBulk {
	return u.Update(func(s *BookAuthorUpsert) {
		s.SetFullName(v)
	})
}

// UpdateFullName sets the "full_name" field to the value that was provided on create.
func (u *BookAuthorUpsertBulk) UpdateFullName() *BookAuthorUpsertBulk {
	return u.Update(func(s *BookAuthorUpsert) {
		s.UpdateFullName()
	})
}

// Exec executes the query.
func (u *BookAuthorUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the BookAuthorCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BookAuthorCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BookAuthorUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
