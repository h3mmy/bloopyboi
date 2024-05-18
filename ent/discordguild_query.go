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
	"github.com/h3mmy/bloopyboi/ent/discordguild"
	"github.com/h3mmy/bloopyboi/ent/discordmessage"
	"github.com/h3mmy/bloopyboi/ent/discorduser"
	"github.com/h3mmy/bloopyboi/ent/predicate"
)

// DiscordGuildQuery is the builder for querying DiscordGuild entities.
type DiscordGuildQuery struct {
	config
	ctx                      *QueryContext
	order                    []discordguild.OrderOption
	inters                   []Interceptor
	predicates               []predicate.DiscordGuild
	withMembers              *DiscordUserQuery
	withDiscordMessages      *DiscordMessageQuery
	modifiers                []func(*sql.Selector)
	withNamedMembers         map[string]*DiscordUserQuery
	withNamedDiscordMessages map[string]*DiscordMessageQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DiscordGuildQuery builder.
func (dgq *DiscordGuildQuery) Where(ps ...predicate.DiscordGuild) *DiscordGuildQuery {
	dgq.predicates = append(dgq.predicates, ps...)
	return dgq
}

// Limit the number of records to be returned by this query.
func (dgq *DiscordGuildQuery) Limit(limit int) *DiscordGuildQuery {
	dgq.ctx.Limit = &limit
	return dgq
}

// Offset to start from.
func (dgq *DiscordGuildQuery) Offset(offset int) *DiscordGuildQuery {
	dgq.ctx.Offset = &offset
	return dgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dgq *DiscordGuildQuery) Unique(unique bool) *DiscordGuildQuery {
	dgq.ctx.Unique = &unique
	return dgq
}

// Order specifies how the records should be ordered.
func (dgq *DiscordGuildQuery) Order(o ...discordguild.OrderOption) *DiscordGuildQuery {
	dgq.order = append(dgq.order, o...)
	return dgq
}

// QueryMembers chains the current query on the "members" edge.
func (dgq *DiscordGuildQuery) QueryMembers() *DiscordUserQuery {
	query := (&DiscordUserClient{config: dgq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(discordguild.Table, discordguild.FieldID, selector),
			sqlgraph.To(discorduser.Table, discorduser.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, discordguild.MembersTable, discordguild.MembersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(dgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDiscordMessages chains the current query on the "discord_messages" edge.
func (dgq *DiscordGuildQuery) QueryDiscordMessages() *DiscordMessageQuery {
	query := (&DiscordMessageClient{config: dgq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(discordguild.Table, discordguild.FieldID, selector),
			sqlgraph.To(discordmessage.Table, discordmessage.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, discordguild.DiscordMessagesTable, discordguild.DiscordMessagesColumn),
		)
		fromU = sqlgraph.SetNeighbors(dgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DiscordGuild entity from the query.
// Returns a *NotFoundError when no DiscordGuild was found.
func (dgq *DiscordGuildQuery) First(ctx context.Context) (*DiscordGuild, error) {
	nodes, err := dgq.Limit(1).All(setContextOp(ctx, dgq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{discordguild.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dgq *DiscordGuildQuery) FirstX(ctx context.Context) *DiscordGuild {
	node, err := dgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DiscordGuild ID from the query.
// Returns a *NotFoundError when no DiscordGuild ID was found.
func (dgq *DiscordGuildQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dgq.Limit(1).IDs(setContextOp(ctx, dgq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{discordguild.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dgq *DiscordGuildQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DiscordGuild entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DiscordGuild entity is found.
// Returns a *NotFoundError when no DiscordGuild entities are found.
func (dgq *DiscordGuildQuery) Only(ctx context.Context) (*DiscordGuild, error) {
	nodes, err := dgq.Limit(2).All(setContextOp(ctx, dgq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{discordguild.Label}
	default:
		return nil, &NotSingularError{discordguild.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dgq *DiscordGuildQuery) OnlyX(ctx context.Context) *DiscordGuild {
	node, err := dgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DiscordGuild ID in the query.
// Returns a *NotSingularError when more than one DiscordGuild ID is found.
// Returns a *NotFoundError when no entities are found.
func (dgq *DiscordGuildQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dgq.Limit(2).IDs(setContextOp(ctx, dgq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{discordguild.Label}
	default:
		err = &NotSingularError{discordguild.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dgq *DiscordGuildQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DiscordGuilds.
func (dgq *DiscordGuildQuery) All(ctx context.Context) ([]*DiscordGuild, error) {
	ctx = setContextOp(ctx, dgq.ctx, "All")
	if err := dgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DiscordGuild, *DiscordGuildQuery]()
	return withInterceptors[[]*DiscordGuild](ctx, dgq, qr, dgq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dgq *DiscordGuildQuery) AllX(ctx context.Context) []*DiscordGuild {
	nodes, err := dgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DiscordGuild IDs.
func (dgq *DiscordGuildQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if dgq.ctx.Unique == nil && dgq.path != nil {
		dgq.Unique(true)
	}
	ctx = setContextOp(ctx, dgq.ctx, "IDs")
	if err = dgq.Select(discordguild.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dgq *DiscordGuildQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dgq *DiscordGuildQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dgq.ctx, "Count")
	if err := dgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dgq, querierCount[*DiscordGuildQuery](), dgq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dgq *DiscordGuildQuery) CountX(ctx context.Context) int {
	count, err := dgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dgq *DiscordGuildQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dgq.ctx, "Exist")
	switch _, err := dgq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dgq *DiscordGuildQuery) ExistX(ctx context.Context) bool {
	exist, err := dgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DiscordGuildQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dgq *DiscordGuildQuery) Clone() *DiscordGuildQuery {
	if dgq == nil {
		return nil
	}
	return &DiscordGuildQuery{
		config:              dgq.config,
		ctx:                 dgq.ctx.Clone(),
		order:               append([]discordguild.OrderOption{}, dgq.order...),
		inters:              append([]Interceptor{}, dgq.inters...),
		predicates:          append([]predicate.DiscordGuild{}, dgq.predicates...),
		withMembers:         dgq.withMembers.Clone(),
		withDiscordMessages: dgq.withDiscordMessages.Clone(),
		// clone intermediate query.
		sql:  dgq.sql.Clone(),
		path: dgq.path,
	}
}

// WithMembers tells the query-builder to eager-load the nodes that are connected to
// the "members" edge. The optional arguments are used to configure the query builder of the edge.
func (dgq *DiscordGuildQuery) WithMembers(opts ...func(*DiscordUserQuery)) *DiscordGuildQuery {
	query := (&DiscordUserClient{config: dgq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dgq.withMembers = query
	return dgq
}

// WithDiscordMessages tells the query-builder to eager-load the nodes that are connected to
// the "discord_messages" edge. The optional arguments are used to configure the query builder of the edge.
func (dgq *DiscordGuildQuery) WithDiscordMessages(opts ...func(*DiscordMessageQuery)) *DiscordGuildQuery {
	query := (&DiscordMessageClient{config: dgq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dgq.withDiscordMessages = query
	return dgq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Discordid string `json:"discordid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DiscordGuild.Query().
//		GroupBy(discordguild.FieldDiscordid).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dgq *DiscordGuildQuery) GroupBy(field string, fields ...string) *DiscordGuildGroupBy {
	dgq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DiscordGuildGroupBy{build: dgq}
	grbuild.flds = &dgq.ctx.Fields
	grbuild.label = discordguild.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Discordid string `json:"discordid,omitempty"`
//	}
//
//	client.DiscordGuild.Query().
//		Select(discordguild.FieldDiscordid).
//		Scan(ctx, &v)
func (dgq *DiscordGuildQuery) Select(fields ...string) *DiscordGuildSelect {
	dgq.ctx.Fields = append(dgq.ctx.Fields, fields...)
	sbuild := &DiscordGuildSelect{DiscordGuildQuery: dgq}
	sbuild.label = discordguild.Label
	sbuild.flds, sbuild.scan = &dgq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DiscordGuildSelect configured with the given aggregations.
func (dgq *DiscordGuildQuery) Aggregate(fns ...AggregateFunc) *DiscordGuildSelect {
	return dgq.Select().Aggregate(fns...)
}

func (dgq *DiscordGuildQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dgq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dgq); err != nil {
				return err
			}
		}
	}
	for _, f := range dgq.ctx.Fields {
		if !discordguild.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dgq.path != nil {
		prev, err := dgq.path(ctx)
		if err != nil {
			return err
		}
		dgq.sql = prev
	}
	return nil
}

func (dgq *DiscordGuildQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DiscordGuild, error) {
	var (
		nodes       = []*DiscordGuild{}
		_spec       = dgq.querySpec()
		loadedTypes = [2]bool{
			dgq.withMembers != nil,
			dgq.withDiscordMessages != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DiscordGuild).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DiscordGuild{config: dgq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(dgq.modifiers) > 0 {
		_spec.Modifiers = dgq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dgq.withMembers; query != nil {
		if err := dgq.loadMembers(ctx, query, nodes,
			func(n *DiscordGuild) { n.Edges.Members = []*DiscordUser{} },
			func(n *DiscordGuild, e *DiscordUser) { n.Edges.Members = append(n.Edges.Members, e) }); err != nil {
			return nil, err
		}
	}
	if query := dgq.withDiscordMessages; query != nil {
		if err := dgq.loadDiscordMessages(ctx, query, nodes,
			func(n *DiscordGuild) { n.Edges.DiscordMessages = []*DiscordMessage{} },
			func(n *DiscordGuild, e *DiscordMessage) { n.Edges.DiscordMessages = append(n.Edges.DiscordMessages, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range dgq.withNamedMembers {
		if err := dgq.loadMembers(ctx, query, nodes,
			func(n *DiscordGuild) { n.appendNamedMembers(name) },
			func(n *DiscordGuild, e *DiscordUser) { n.appendNamedMembers(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range dgq.withNamedDiscordMessages {
		if err := dgq.loadDiscordMessages(ctx, query, nodes,
			func(n *DiscordGuild) { n.appendNamedDiscordMessages(name) },
			func(n *DiscordGuild, e *DiscordMessage) { n.appendNamedDiscordMessages(name, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dgq *DiscordGuildQuery) loadMembers(ctx context.Context, query *DiscordUserQuery, nodes []*DiscordGuild, init func(*DiscordGuild), assign func(*DiscordGuild, *DiscordUser)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*DiscordGuild)
	nids := make(map[uuid.UUID]map[*DiscordGuild]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(discordguild.MembersTable)
		s.Join(joinT).On(s.C(discorduser.FieldID), joinT.C(discordguild.MembersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(discordguild.MembersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(discordguild.MembersPrimaryKey[0]))
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
					nids[inValue] = map[*DiscordGuild]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*DiscordUser](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "members" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (dgq *DiscordGuildQuery) loadDiscordMessages(ctx context.Context, query *DiscordMessageQuery, nodes []*DiscordGuild, init func(*DiscordGuild), assign func(*DiscordGuild, *DiscordMessage)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*DiscordGuild)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.DiscordMessage(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(discordguild.DiscordMessagesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.discord_guild_discord_messages
		if fk == nil {
			return fmt.Errorf(`foreign-key "discord_guild_discord_messages" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "discord_guild_discord_messages" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (dgq *DiscordGuildQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dgq.querySpec()
	if len(dgq.modifiers) > 0 {
		_spec.Modifiers = dgq.modifiers
	}
	_spec.Node.Columns = dgq.ctx.Fields
	if len(dgq.ctx.Fields) > 0 {
		_spec.Unique = dgq.ctx.Unique != nil && *dgq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dgq.driver, _spec)
}

func (dgq *DiscordGuildQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(discordguild.Table, discordguild.Columns, sqlgraph.NewFieldSpec(discordguild.FieldID, field.TypeUUID))
	_spec.From = dgq.sql
	if unique := dgq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dgq.path != nil {
		_spec.Unique = true
	}
	if fields := dgq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, discordguild.FieldID)
		for i := range fields {
			if fields[i] != discordguild.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dgq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dgq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dgq *DiscordGuildQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dgq.driver.Dialect())
	t1 := builder.Table(discordguild.Table)
	columns := dgq.ctx.Fields
	if len(columns) == 0 {
		columns = discordguild.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dgq.sql != nil {
		selector = dgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dgq.ctx.Unique != nil && *dgq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range dgq.modifiers {
		m(selector)
	}
	for _, p := range dgq.predicates {
		p(selector)
	}
	for _, p := range dgq.order {
		p(selector)
	}
	if offset := dgq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dgq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (dgq *DiscordGuildQuery) ForUpdate(opts ...sql.LockOption) *DiscordGuildQuery {
	if dgq.driver.Dialect() == dialect.Postgres {
		dgq.Unique(false)
	}
	dgq.modifiers = append(dgq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return dgq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (dgq *DiscordGuildQuery) ForShare(opts ...sql.LockOption) *DiscordGuildQuery {
	if dgq.driver.Dialect() == dialect.Postgres {
		dgq.Unique(false)
	}
	dgq.modifiers = append(dgq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return dgq
}

// WithNamedMembers tells the query-builder to eager-load the nodes that are connected to the "members"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (dgq *DiscordGuildQuery) WithNamedMembers(name string, opts ...func(*DiscordUserQuery)) *DiscordGuildQuery {
	query := (&DiscordUserClient{config: dgq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if dgq.withNamedMembers == nil {
		dgq.withNamedMembers = make(map[string]*DiscordUserQuery)
	}
	dgq.withNamedMembers[name] = query
	return dgq
}

// WithNamedDiscordMessages tells the query-builder to eager-load the nodes that are connected to the "discord_messages"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (dgq *DiscordGuildQuery) WithNamedDiscordMessages(name string, opts ...func(*DiscordMessageQuery)) *DiscordGuildQuery {
	query := (&DiscordMessageClient{config: dgq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if dgq.withNamedDiscordMessages == nil {
		dgq.withNamedDiscordMessages = make(map[string]*DiscordMessageQuery)
	}
	dgq.withNamedDiscordMessages[name] = query
	return dgq
}

// DiscordGuildGroupBy is the group-by builder for DiscordGuild entities.
type DiscordGuildGroupBy struct {
	selector
	build *DiscordGuildQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dggb *DiscordGuildGroupBy) Aggregate(fns ...AggregateFunc) *DiscordGuildGroupBy {
	dggb.fns = append(dggb.fns, fns...)
	return dggb
}

// Scan applies the selector query and scans the result into the given value.
func (dggb *DiscordGuildGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dggb.build.ctx, "GroupBy")
	if err := dggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DiscordGuildQuery, *DiscordGuildGroupBy](ctx, dggb.build, dggb, dggb.build.inters, v)
}

func (dggb *DiscordGuildGroupBy) sqlScan(ctx context.Context, root *DiscordGuildQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dggb.fns))
	for _, fn := range dggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dggb.flds)+len(dggb.fns))
		for _, f := range *dggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DiscordGuildSelect is the builder for selecting fields of DiscordGuild entities.
type DiscordGuildSelect struct {
	*DiscordGuildQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dgs *DiscordGuildSelect) Aggregate(fns ...AggregateFunc) *DiscordGuildSelect {
	dgs.fns = append(dgs.fns, fns...)
	return dgs
}

// Scan applies the selector query and scans the result into the given value.
func (dgs *DiscordGuildSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgs.ctx, "Select")
	if err := dgs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DiscordGuildQuery, *DiscordGuildSelect](ctx, dgs.DiscordGuildQuery, dgs, dgs.inters, v)
}

func (dgs *DiscordGuildSelect) sqlScan(ctx context.Context, root *DiscordGuildQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dgs.fns))
	for _, fn := range dgs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dgs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
