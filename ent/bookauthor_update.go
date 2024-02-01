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
	"github.com/h3mmy/bloopyboi/ent/book"
	"github.com/h3mmy/bloopyboi/ent/bookauthor"
	"github.com/h3mmy/bloopyboi/ent/predicate"
)

// BookAuthorUpdate is the builder for updating BookAuthor entities.
type BookAuthorUpdate struct {
	config
	hooks    []Hook
	mutation *BookAuthorMutation
}

// Where appends a list predicates to the BookAuthorUpdate builder.
func (bau *BookAuthorUpdate) Where(ps ...predicate.BookAuthor) *BookAuthorUpdate {
	bau.mutation.Where(ps...)
	return bau
}

// SetFullName sets the "full_name" field.
func (bau *BookAuthorUpdate) SetFullName(s string) *BookAuthorUpdate {
	bau.mutation.SetFullName(s)
	return bau
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (bau *BookAuthorUpdate) AddBookIDs(ids ...uuid.UUID) *BookAuthorUpdate {
	bau.mutation.AddBookIDs(ids...)
	return bau
}

// AddBooks adds the "books" edges to the Book entity.
func (bau *BookAuthorUpdate) AddBooks(b ...*Book) *BookAuthorUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bau.AddBookIDs(ids...)
}

// Mutation returns the BookAuthorMutation object of the builder.
func (bau *BookAuthorUpdate) Mutation() *BookAuthorMutation {
	return bau.mutation
}

// ClearBooks clears all "books" edges to the Book entity.
func (bau *BookAuthorUpdate) ClearBooks() *BookAuthorUpdate {
	bau.mutation.ClearBooks()
	return bau
}

// RemoveBookIDs removes the "books" edge to Book entities by IDs.
func (bau *BookAuthorUpdate) RemoveBookIDs(ids ...uuid.UUID) *BookAuthorUpdate {
	bau.mutation.RemoveBookIDs(ids...)
	return bau
}

// RemoveBooks removes "books" edges to Book entities.
func (bau *BookAuthorUpdate) RemoveBooks(b ...*Book) *BookAuthorUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bau.RemoveBookIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bau *BookAuthorUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bau.sqlSave, bau.mutation, bau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bau *BookAuthorUpdate) SaveX(ctx context.Context) int {
	affected, err := bau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bau *BookAuthorUpdate) Exec(ctx context.Context) error {
	_, err := bau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bau *BookAuthorUpdate) ExecX(ctx context.Context) {
	if err := bau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bau *BookAuthorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(bookauthor.Table, bookauthor.Columns, sqlgraph.NewFieldSpec(bookauthor.FieldID, field.TypeUUID))
	if ps := bau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bau.mutation.FullName(); ok {
		_spec.SetField(bookauthor.FieldFullName, field.TypeString, value)
	}
	if bau.mutation.BooksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bau.mutation.RemovedBooksIDs(); len(nodes) > 0 && !bau.mutation.BooksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bau.mutation.BooksIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookauthor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bau.mutation.done = true
	return n, nil
}

// BookAuthorUpdateOne is the builder for updating a single BookAuthor entity.
type BookAuthorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookAuthorMutation
}

// SetFullName sets the "full_name" field.
func (bauo *BookAuthorUpdateOne) SetFullName(s string) *BookAuthorUpdateOne {
	bauo.mutation.SetFullName(s)
	return bauo
}

// AddBookIDs adds the "books" edge to the Book entity by IDs.
func (bauo *BookAuthorUpdateOne) AddBookIDs(ids ...uuid.UUID) *BookAuthorUpdateOne {
	bauo.mutation.AddBookIDs(ids...)
	return bauo
}

// AddBooks adds the "books" edges to the Book entity.
func (bauo *BookAuthorUpdateOne) AddBooks(b ...*Book) *BookAuthorUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bauo.AddBookIDs(ids...)
}

// Mutation returns the BookAuthorMutation object of the builder.
func (bauo *BookAuthorUpdateOne) Mutation() *BookAuthorMutation {
	return bauo.mutation
}

// ClearBooks clears all "books" edges to the Book entity.
func (bauo *BookAuthorUpdateOne) ClearBooks() *BookAuthorUpdateOne {
	bauo.mutation.ClearBooks()
	return bauo
}

// RemoveBookIDs removes the "books" edge to Book entities by IDs.
func (bauo *BookAuthorUpdateOne) RemoveBookIDs(ids ...uuid.UUID) *BookAuthorUpdateOne {
	bauo.mutation.RemoveBookIDs(ids...)
	return bauo
}

// RemoveBooks removes "books" edges to Book entities.
func (bauo *BookAuthorUpdateOne) RemoveBooks(b ...*Book) *BookAuthorUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bauo.RemoveBookIDs(ids...)
}

// Where appends a list predicates to the BookAuthorUpdate builder.
func (bauo *BookAuthorUpdateOne) Where(ps ...predicate.BookAuthor) *BookAuthorUpdateOne {
	bauo.mutation.Where(ps...)
	return bauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bauo *BookAuthorUpdateOne) Select(field string, fields ...string) *BookAuthorUpdateOne {
	bauo.fields = append([]string{field}, fields...)
	return bauo
}

// Save executes the query and returns the updated BookAuthor entity.
func (bauo *BookAuthorUpdateOne) Save(ctx context.Context) (*BookAuthor, error) {
	return withHooks(ctx, bauo.sqlSave, bauo.mutation, bauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bauo *BookAuthorUpdateOne) SaveX(ctx context.Context) *BookAuthor {
	node, err := bauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bauo *BookAuthorUpdateOne) Exec(ctx context.Context) error {
	_, err := bauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bauo *BookAuthorUpdateOne) ExecX(ctx context.Context) {
	if err := bauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bauo *BookAuthorUpdateOne) sqlSave(ctx context.Context) (_node *BookAuthor, err error) {
	_spec := sqlgraph.NewUpdateSpec(bookauthor.Table, bookauthor.Columns, sqlgraph.NewFieldSpec(bookauthor.FieldID, field.TypeUUID))
	id, ok := bauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BookAuthor.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bookauthor.FieldID)
		for _, f := range fields {
			if !bookauthor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != bookauthor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bauo.mutation.FullName(); ok {
		_spec.SetField(bookauthor.FieldFullName, field.TypeString, value)
	}
	if bauo.mutation.BooksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bauo.mutation.RemovedBooksIDs(); len(nodes) > 0 && !bauo.mutation.BooksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bauo.mutation.BooksIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BookAuthor{config: bauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bookauthor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bauo.mutation.done = true
	return _node, nil
}