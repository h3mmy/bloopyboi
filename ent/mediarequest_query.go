// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/h3mmy/bloopyboi/ent/book"
	"github.com/h3mmy/bloopyboi/ent/discorduser"
	"github.com/h3mmy/bloopyboi/ent/mediarequest"
	"github.com/h3mmy/bloopyboi/ent/predicate"
)

// MediaRequestQuery is the builder for querying MediaRequest entities.
type MediaRequestQuery struct {
	config
	ctx             *QueryContext
	order           []mediarequest.OrderOption
	inters          []Interceptor
	predicates      []predicate.MediaRequest
	withDiscordUser *DiscordUserQuery
	withBooks       *BookQuery
	withFKs         bool
	modifiers       []func(*sql.Selector)
	withNamedBooks  map[string]*BookQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MediaRequestQuery builder.
func (mrq *MediaRequestQuery) Where(ps ...predicate.MediaRequest) *MediaRequestQuery {
	mrq.predicates = append(mrq.predicates, ps...)
	return mrq
}

// Limit the number of records to be returned by this query.
func (mrq *MediaRequestQuery) Limit(limit int) *MediaRequestQuery {
	mrq.ctx.Limit = &limit
	return mrq
}

// Offset to start from.
func (mrq *MediaRequestQuery) Offset(offset int) *MediaRequestQuery {
	mrq.ctx.Offset = &offset
	return mrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mrq *MediaRequestQuery) Unique(unique bool) *MediaRequestQuery {
	mrq.ctx.Unique = &unique
	return mrq
}

// Order specifies how the records should be ordered.
func (mrq *MediaRequestQuery) Order(o ...mediarequest.OrderOption) *MediaRequestQuery {
	mrq.order = append(mrq.order, o...)
	return mrq
}

// QueryDiscordUser chains the current query on the "discord_user" edge.
func (mrq *MediaRequestQuery) QueryDiscordUser() *DiscordUserQuery {
	query := (&DiscordUserClient{config: mrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(mediarequest.Table, mediarequest.FieldID, selector),
			sqlgraph.To(discorduser.Table, discorduser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, mediarequest.DiscordUserTable, mediarequest.DiscordUserColumn),
		)
		fromU = sqlgraph.SetNeighbors(mrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBooks chains the current query on the "books" edge.
func (mrq *MediaRequestQuery) QueryBooks() *BookQuery {
	query := (&BookClient{config: mrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(mediarequest.Table, mediarequest.FieldID, selector),
			sqlgraph.To(book.Table, book.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, mediarequest.BooksTable, mediarequest.BooksPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(mrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MediaRequest entity from the query.
// Returns a *NotFoundError when no MediaRequest was found.
func (mrq *MediaRequestQuery) First(ctx context.Context) (*MediaRequest, error) {
	nodes, err := mrq.Limit(1).All(setContextOp(ctx, mrq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{mediarequest.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mrq *MediaRequestQuery) FirstX(ctx context.Context) *MediaRequest {
	node, err := mrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MediaRequest ID from the query.
// Returns a *NotFoundError when no MediaRequest ID was found.
func (mrq *MediaRequestQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mrq.Limit(1).IDs(setContextOp(ctx, mrq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{mediarequest.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mrq *MediaRequestQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := mrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MediaRequest entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MediaRequest entity is found.
// Returns a *NotFoundError when no MediaRequest entities are found.
func (mrq *MediaRequestQuery) Only(ctx context.Context) (*MediaRequest, error) {
	nodes, err := mrq.Limit(2).All(setContextOp(ctx, mrq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{mediarequest.Label}
	default:
		return nil, &NotSingularError{mediarequest.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mrq *MediaRequestQuery) OnlyX(ctx context.Context) *MediaRequest {
	node, err := mrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MediaRequest ID in the query.
// Returns a *NotSingularError when more than one MediaRequest ID is found.
// Returns a *NotFoundError when no entities are found.
func (mrq *MediaRequestQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mrq.Limit(2).IDs(setContextOp(ctx, mrq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{mediarequest.Label}
	default:
		err = &NotSingularError{mediarequest.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mrq *MediaRequestQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := mrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MediaRequests.
func (mrq *MediaRequestQuery) All(ctx context.Context) ([]*MediaRequest, error) {
	ctx = setContextOp(ctx, mrq.ctx, "All")
	if err := mrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MediaRequest, *MediaRequestQuery]()
	return withInterceptors[[]*MediaRequest](ctx, mrq, qr, mrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mrq *MediaRequestQuery) AllX(ctx context.Context) []*MediaRequest {
	nodes, err := mrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MediaRequest IDs.
func (mrq *MediaRequestQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if mrq.ctx.Unique == nil && mrq.path != nil {
		mrq.Unique(true)
	}
	ctx = setContextOp(ctx, mrq.ctx, "IDs")
	if err = mrq.Select(mediarequest.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mrq *MediaRequestQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := mrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mrq *MediaRequestQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mrq.ctx, "Count")
	if err := mrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mrq, querierCount[*MediaRequestQuery](), mrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mrq *MediaRequestQuery) CountX(ctx context.Context) int {
	count, err := mrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mrq *MediaRequestQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mrq.ctx, "Exist")
	switch _, err := mrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mrq *MediaRequestQuery) ExistX(ctx context.Context) bool {
	exist, err := mrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MediaRequestQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mrq *MediaRequestQuery) Clone() *MediaRequestQuery {
	if mrq == nil {
		return nil
	}
	return &MediaRequestQuery{
		config:          mrq.config,
		ctx:             mrq.ctx.Clone(),
		order:           append([]mediarequest.OrderOption{}, mrq.order...),
		inters:          append([]Interceptor{}, mrq.inters...),
		predicates:      append([]predicate.MediaRequest{}, mrq.predicates...),
		withDiscordUser: mrq.withDiscordUser.Clone(),
		withBooks:       mrq.withBooks.Clone(),
		// clone intermediate query.
		sql:  mrq.sql.Clone(),
		path: mrq.path,
	}
}

// WithDiscordUser tells the query-builder to eager-load the nodes that are connected to
// the "discord_user" edge. The optional arguments are used to configure the query builder of the edge.
func (mrq *MediaRequestQuery) WithDiscordUser(opts ...func(*DiscordUserQuery)) *MediaRequestQuery {
	query := (&DiscordUserClient{config: mrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mrq.withDiscordUser = query
	return mrq
}

// WithBooks tells the query-builder to eager-load the nodes that are connected to
// the "books" edge. The optional arguments are used to configure the query builder of the edge.
func (mrq *MediaRequestQuery) WithBooks(opts ...func(*BookQuery)) *MediaRequestQuery {
	query := (&BookClient{config: mrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mrq.withBooks = query
	return mrq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MediaRequest.Query().
//		GroupBy(mediarequest.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mrq *MediaRequestQuery) GroupBy(field string, fields ...string) *MediaRequestGroupBy {
	mrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MediaRequestGroupBy{build: mrq}
	grbuild.flds = &mrq.ctx.Fields
	grbuild.label = mediarequest.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.MediaRequest.Query().
//		Select(mediarequest.FieldCreateTime).
//		Scan(ctx, &v)
func (mrq *MediaRequestQuery) Select(fields ...string) *MediaRequestSelect {
	mrq.ctx.Fields = append(mrq.ctx.Fields, fields...)
	sbuild := &MediaRequestSelect{MediaRequestQuery: mrq}
	sbuild.label = mediarequest.Label
	sbuild.flds, sbuild.scan = &mrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MediaRequestSelect configured with the given aggregations.
func (mrq *MediaRequestQuery) Aggregate(fns ...AggregateFunc) *MediaRequestSelect {
	return mrq.Select().Aggregate(fns...)
}

func (mrq *MediaRequestQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mrq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mrq); err != nil {
				return err
			}
		}
	}
	for _, f := range mrq.ctx.Fields {
		if !mediarequest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mrq.path != nil {
		prev, err := mrq.path(ctx)
		if err != nil {
			return err
		}
		mrq.sql = prev
	}
	return nil
}

func (mrq *MediaRequestQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MediaRequest, error) {
	var (
		nodes       = []*MediaRequest{}
		withFKs     = mrq.withFKs
		_spec       = mrq.querySpec()
		loadedTypes = [2]bool{
			mrq.withDiscordUser != nil,
			mrq.withBooks != nil,
		}
	)
	if mrq.withDiscordUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, mediarequest.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MediaRequest).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MediaRequest{config: mrq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(mrq.modifiers) > 0 {
		_spec.Modifiers = mrq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mrq.withDiscordUser; query != nil {
		if err := mrq.loadDiscordUser(ctx, query, nodes, nil,
			func(n *MediaRequest, e *DiscordUser) { n.Edges.DiscordUser = e }); err != nil {
			return nil, err
		}
	}
	if query := mrq.withBooks; query != nil {
		if err := mrq.loadBooks(ctx, query, nodes,
			func(n *MediaRequest) { n.Edges.Books = []*Book{} },
			func(n *MediaRequest, e *Book) { n.Edges.Books = append(n.Edges.Books, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range mrq.withNamedBooks {
		if err := mrq.loadBooks(ctx, query, nodes,
			func(n *MediaRequest) { n.appendNamedBooks(name) },
			func(n *MediaRequest, e *Book) { n.appendNamedBooks(name, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mrq *MediaRequestQuery) loadDiscordUser(ctx context.Context, query *DiscordUserQuery, nodes []*MediaRequest, init func(*MediaRequest), assign func(*MediaRequest, *DiscordUser)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*MediaRequest)
	for i := range nodes {
		if nodes[i].discord_user_media_requests == nil {
			continue
		}
		fk := *nodes[i].discord_user_media_requests
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(discorduser.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "discord_user_media_requests" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mrq *MediaRequestQuery) loadBooks(ctx context.Context, query *BookQuery, nodes []*MediaRequest, init func(*MediaRequest), assign func(*MediaRequest, *Book)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*MediaRequest)
	nids := make(map[uuid.UUID]map[*MediaRequest]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(mediarequest.BooksTable)
		s.Join(joinT).On(s.C(book.FieldID), joinT.C(mediarequest.BooksPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(mediarequest.BooksPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(mediarequest.BooksPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*MediaRequest]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Book](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "books" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (mrq *MediaRequestQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mrq.querySpec()
	if len(mrq.modifiers) > 0 {
		_spec.Modifiers = mrq.modifiers
	}
	_spec.Node.Columns = mrq.ctx.Fields
	if len(mrq.ctx.Fields) > 0 {
		_spec.Unique = mrq.ctx.Unique != nil && *mrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mrq.driver, _spec)
}

func (mrq *MediaRequestQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(mediarequest.Table, mediarequest.Columns, sqlgraph.NewFieldSpec(mediarequest.FieldID, field.TypeUUID))
	_spec.From = mrq.sql
	if unique := mrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mrq.path != nil {
		_spec.Unique = true
	}
	if fields := mrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mediarequest.FieldID)
		for i := range fields {
			if fields[i] != mediarequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mrq *MediaRequestQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mrq.driver.Dialect())
	t1 := builder.Table(mediarequest.Table)
	columns := mrq.ctx.Fields
	if len(columns) == 0 {
		columns = mediarequest.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mrq.sql != nil {
		selector = mrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mrq.ctx.Unique != nil && *mrq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range mrq.modifiers {
		m(selector)
	}
	for _, p := range mrq.predicates {
		p(selector)
	}
	for _, p := range mrq.order {
		p(selector)
	}
	if offset := mrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (mrq *MediaRequestQuery) ForUpdate(opts ...sql.LockOption) *MediaRequestQuery {
	if mrq.driver.Dialect() == dialect.Postgres {
		mrq.Unique(false)
	}
	mrq.modifiers = append(mrq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return mrq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (mrq *MediaRequestQuery) ForShare(opts ...sql.LockOption) *MediaRequestQuery {
	if mrq.driver.Dialect() == dialect.Postgres {
		mrq.Unique(false)
	}
	mrq.modifiers = append(mrq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return mrq
}

// WithNamedBooks tells the query-builder to eager-load the nodes that are connected to the "books"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (mrq *MediaRequestQuery) WithNamedBooks(name string, opts ...func(*BookQuery)) *MediaRequestQuery {
	query := (&BookClient{config: mrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if mrq.withNamedBooks == nil {
		mrq.withNamedBooks = make(map[string]*BookQuery)
	}
	mrq.withNamedBooks[name] = query
	return mrq
}

// MediaRequestGroupBy is the group-by builder for MediaRequest entities.
type MediaRequestGroupBy struct {
	selector
	build *MediaRequestQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mrgb *MediaRequestGroupBy) Aggregate(fns ...AggregateFunc) *MediaRequestGroupBy {
	mrgb.fns = append(mrgb.fns, fns...)
	return mrgb
}

// Scan applies the selector query and scans the result into the given value.
func (mrgb *MediaRequestGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mrgb.build.ctx, "GroupBy")
	if err := mrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MediaRequestQuery, *MediaRequestGroupBy](ctx, mrgb.build, mrgb, mrgb.build.inters, v)
}

func (mrgb *MediaRequestGroupBy) sqlScan(ctx context.Context, root *MediaRequestQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mrgb.fns))
	for _, fn := range mrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mrgb.flds)+len(mrgb.fns))
		for _, f := range *mrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MediaRequestSelect is the builder for selecting fields of MediaRequest entities.
type MediaRequestSelect struct {
	*MediaRequestQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mrs *MediaRequestSelect) Aggregate(fns ...AggregateFunc) *MediaRequestSelect {
	mrs.fns = append(mrs.fns, fns...)
	return mrs
}

// Scan applies the selector query and scans the result into the given value.
func (mrs *MediaRequestSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mrs.ctx, "Select")
	if err := mrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MediaRequestQuery, *MediaRequestSelect](ctx, mrs.MediaRequestQuery, mrs, mrs.inters, v)
}

func (mrs *MediaRequestSelect) sqlScan(ctx context.Context, root *MediaRequestQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mrs.fns))
	for _, fn := range mrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}