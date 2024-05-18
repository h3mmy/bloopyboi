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
	"github.com/h3mmy/bloopyboi/ent/discordmessagereaction"
	"github.com/h3mmy/bloopyboi/ent/discorduser"
	"github.com/h3mmy/bloopyboi/ent/predicate"
)

// DiscordMessageQuery is the builder for querying DiscordMessage entities.
type DiscordMessageQuery struct {
	config
	ctx                       *QueryContext
	order                     []discordmessage.OrderOption
	inters                    []Interceptor
	predicates                []predicate.DiscordMessage
	withAuthor                *DiscordUserQuery
	withMessageReactions      *DiscordMessageReactionQuery
	withGuild                 *DiscordGuildQuery
	withFKs                   bool
	modifiers                 []func(*sql.Selector)
	withNamedMessageReactions map[string]*DiscordMessageReactionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DiscordMessageQuery builder.
func (dmq *DiscordMessageQuery) Where(ps ...predicate.DiscordMessage) *DiscordMessageQuery {
	dmq.predicates = append(dmq.predicates, ps...)
	return dmq
}

// Limit the number of records to be returned by this query.
func (dmq *DiscordMessageQuery) Limit(limit int) *DiscordMessageQuery {
	dmq.ctx.Limit = &limit
	return dmq
}

// Offset to start from.
func (dmq *DiscordMessageQuery) Offset(offset int) *DiscordMessageQuery {
	dmq.ctx.Offset = &offset
	return dmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dmq *DiscordMessageQuery) Unique(unique bool) *DiscordMessageQuery {
	dmq.ctx.Unique = &unique
	return dmq
}

// Order specifies how the records should be ordered.
func (dmq *DiscordMessageQuery) Order(o ...discordmessage.OrderOption) *DiscordMessageQuery {
	dmq.order = append(dmq.order, o...)
	return dmq
}

// QueryAuthor chains the current query on the "author" edge.
func (dmq *DiscordMessageQuery) QueryAuthor() *DiscordUserQuery {
	query := (&DiscordUserClient{config: dmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(discordmessage.Table, discordmessage.FieldID, selector),
			sqlgraph.To(discorduser.Table, discorduser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, discordmessage.AuthorTable, discordmessage.AuthorColumn),
		)
		fromU = sqlgraph.SetNeighbors(dmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMessageReactions chains the current query on the "message_reactions" edge.
func (dmq *DiscordMessageQuery) QueryMessageReactions() *DiscordMessageReactionQuery {
	query := (&DiscordMessageReactionClient{config: dmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(discordmessage.Table, discordmessage.FieldID, selector),
			sqlgraph.To(discordmessagereaction.Table, discordmessagereaction.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, discordmessage.MessageReactionsTable, discordmessage.MessageReactionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(dmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGuild chains the current query on the "guild" edge.
func (dmq *DiscordMessageQuery) QueryGuild() *DiscordGuildQuery {
	query := (&DiscordGuildClient{config: dmq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(discordmessage.Table, discordmessage.FieldID, selector),
			sqlgraph.To(discordguild.Table, discordguild.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, discordmessage.GuildTable, discordmessage.GuildColumn),
		)
		fromU = sqlgraph.SetNeighbors(dmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DiscordMessage entity from the query.
// Returns a *NotFoundError when no DiscordMessage was found.
func (dmq *DiscordMessageQuery) First(ctx context.Context) (*DiscordMessage, error) {
	nodes, err := dmq.Limit(1).All(setContextOp(ctx, dmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{discordmessage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dmq *DiscordMessageQuery) FirstX(ctx context.Context) *DiscordMessage {
	node, err := dmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DiscordMessage ID from the query.
// Returns a *NotFoundError when no DiscordMessage ID was found.
func (dmq *DiscordMessageQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dmq.Limit(1).IDs(setContextOp(ctx, dmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{discordmessage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dmq *DiscordMessageQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := dmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DiscordMessage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DiscordMessage entity is found.
// Returns a *NotFoundError when no DiscordMessage entities are found.
func (dmq *DiscordMessageQuery) Only(ctx context.Context) (*DiscordMessage, error) {
	nodes, err := dmq.Limit(2).All(setContextOp(ctx, dmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{discordmessage.Label}
	default:
		return nil, &NotSingularError{discordmessage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dmq *DiscordMessageQuery) OnlyX(ctx context.Context) *DiscordMessage {
	node, err := dmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DiscordMessage ID in the query.
// Returns a *NotSingularError when more than one DiscordMessage ID is found.
// Returns a *NotFoundError when no entities are found.
func (dmq *DiscordMessageQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = dmq.Limit(2).IDs(setContextOp(ctx, dmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{discordmessage.Label}
	default:
		err = &NotSingularError{discordmessage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dmq *DiscordMessageQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := dmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DiscordMessages.
func (dmq *DiscordMessageQuery) All(ctx context.Context) ([]*DiscordMessage, error) {
	ctx = setContextOp(ctx, dmq.ctx, "All")
	if err := dmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DiscordMessage, *DiscordMessageQuery]()
	return withInterceptors[[]*DiscordMessage](ctx, dmq, qr, dmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dmq *DiscordMessageQuery) AllX(ctx context.Context) []*DiscordMessage {
	nodes, err := dmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DiscordMessage IDs.
func (dmq *DiscordMessageQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if dmq.ctx.Unique == nil && dmq.path != nil {
		dmq.Unique(true)
	}
	ctx = setContextOp(ctx, dmq.ctx, "IDs")
	if err = dmq.Select(discordmessage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dmq *DiscordMessageQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := dmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dmq *DiscordMessageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dmq.ctx, "Count")
	if err := dmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dmq, querierCount[*DiscordMessageQuery](), dmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dmq *DiscordMessageQuery) CountX(ctx context.Context) int {
	count, err := dmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dmq *DiscordMessageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dmq.ctx, "Exist")
	switch _, err := dmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dmq *DiscordMessageQuery) ExistX(ctx context.Context) bool {
	exist, err := dmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DiscordMessageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dmq *DiscordMessageQuery) Clone() *DiscordMessageQuery {
	if dmq == nil {
		return nil
	}
	return &DiscordMessageQuery{
		config:               dmq.config,
		ctx:                  dmq.ctx.Clone(),
		order:                append([]discordmessage.OrderOption{}, dmq.order...),
		inters:               append([]Interceptor{}, dmq.inters...),
		predicates:           append([]predicate.DiscordMessage{}, dmq.predicates...),
		withAuthor:           dmq.withAuthor.Clone(),
		withMessageReactions: dmq.withMessageReactions.Clone(),
		withGuild:            dmq.withGuild.Clone(),
		// clone intermediate query.
		sql:  dmq.sql.Clone(),
		path: dmq.path,
	}
}

// WithAuthor tells the query-builder to eager-load the nodes that are connected to
// the "author" edge. The optional arguments are used to configure the query builder of the edge.
func (dmq *DiscordMessageQuery) WithAuthor(opts ...func(*DiscordUserQuery)) *DiscordMessageQuery {
	query := (&DiscordUserClient{config: dmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dmq.withAuthor = query
	return dmq
}

// WithMessageReactions tells the query-builder to eager-load the nodes that are connected to
// the "message_reactions" edge. The optional arguments are used to configure the query builder of the edge.
func (dmq *DiscordMessageQuery) WithMessageReactions(opts ...func(*DiscordMessageReactionQuery)) *DiscordMessageQuery {
	query := (&DiscordMessageReactionClient{config: dmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dmq.withMessageReactions = query
	return dmq
}

// WithGuild tells the query-builder to eager-load the nodes that are connected to
// the "guild" edge. The optional arguments are used to configure the query builder of the edge.
func (dmq *DiscordMessageQuery) WithGuild(opts ...func(*DiscordGuildQuery)) *DiscordMessageQuery {
	query := (&DiscordGuildClient{config: dmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dmq.withGuild = query
	return dmq
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
//	client.DiscordMessage.Query().
//		GroupBy(discordmessage.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dmq *DiscordMessageQuery) GroupBy(field string, fields ...string) *DiscordMessageGroupBy {
	dmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DiscordMessageGroupBy{build: dmq}
	grbuild.flds = &dmq.ctx.Fields
	grbuild.label = discordmessage.Label
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
//	client.DiscordMessage.Query().
//		Select(discordmessage.FieldCreateTime).
//		Scan(ctx, &v)
func (dmq *DiscordMessageQuery) Select(fields ...string) *DiscordMessageSelect {
	dmq.ctx.Fields = append(dmq.ctx.Fields, fields...)
	sbuild := &DiscordMessageSelect{DiscordMessageQuery: dmq}
	sbuild.label = discordmessage.Label
	sbuild.flds, sbuild.scan = &dmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DiscordMessageSelect configured with the given aggregations.
func (dmq *DiscordMessageQuery) Aggregate(fns ...AggregateFunc) *DiscordMessageSelect {
	return dmq.Select().Aggregate(fns...)
}

func (dmq *DiscordMessageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dmq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dmq); err != nil {
				return err
			}
		}
	}
	for _, f := range dmq.ctx.Fields {
		if !discordmessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dmq.path != nil {
		prev, err := dmq.path(ctx)
		if err != nil {
			return err
		}
		dmq.sql = prev
	}
	return nil
}

func (dmq *DiscordMessageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DiscordMessage, error) {
	var (
		nodes       = []*DiscordMessage{}
		withFKs     = dmq.withFKs
		_spec       = dmq.querySpec()
		loadedTypes = [3]bool{
			dmq.withAuthor != nil,
			dmq.withMessageReactions != nil,
			dmq.withGuild != nil,
		}
	)
	if dmq.withAuthor != nil || dmq.withGuild != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, discordmessage.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DiscordMessage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DiscordMessage{config: dmq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(dmq.modifiers) > 0 {
		_spec.Modifiers = dmq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dmq.withAuthor; query != nil {
		if err := dmq.loadAuthor(ctx, query, nodes, nil,
			func(n *DiscordMessage, e *DiscordUser) { n.Edges.Author = e }); err != nil {
			return nil, err
		}
	}
	if query := dmq.withMessageReactions; query != nil {
		if err := dmq.loadMessageReactions(ctx, query, nodes,
			func(n *DiscordMessage) { n.Edges.MessageReactions = []*DiscordMessageReaction{} },
			func(n *DiscordMessage, e *DiscordMessageReaction) {
				n.Edges.MessageReactions = append(n.Edges.MessageReactions, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := dmq.withGuild; query != nil {
		if err := dmq.loadGuild(ctx, query, nodes, nil,
			func(n *DiscordMessage, e *DiscordGuild) { n.Edges.Guild = e }); err != nil {
			return nil, err
		}
	}
	for name, query := range dmq.withNamedMessageReactions {
		if err := dmq.loadMessageReactions(ctx, query, nodes,
			func(n *DiscordMessage) { n.appendNamedMessageReactions(name) },
			func(n *DiscordMessage, e *DiscordMessageReaction) { n.appendNamedMessageReactions(name, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dmq *DiscordMessageQuery) loadAuthor(ctx context.Context, query *DiscordUserQuery, nodes []*DiscordMessage, init func(*DiscordMessage), assign func(*DiscordMessage, *DiscordUser)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*DiscordMessage)
	for i := range nodes {
		if nodes[i].discord_user_discord_messages == nil {
			continue
		}
		fk := *nodes[i].discord_user_discord_messages
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
			return fmt.Errorf(`unexpected foreign-key "discord_user_discord_messages" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dmq *DiscordMessageQuery) loadMessageReactions(ctx context.Context, query *DiscordMessageReactionQuery, nodes []*DiscordMessage, init func(*DiscordMessage), assign func(*DiscordMessage, *DiscordMessageReaction)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*DiscordMessage)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.DiscordMessageReaction(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(discordmessage.MessageReactionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.discord_message_message_reactions
		if fk == nil {
			return fmt.Errorf(`foreign-key "discord_message_message_reactions" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "discord_message_message_reactions" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (dmq *DiscordMessageQuery) loadGuild(ctx context.Context, query *DiscordGuildQuery, nodes []*DiscordMessage, init func(*DiscordMessage), assign func(*DiscordMessage, *DiscordGuild)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*DiscordMessage)
	for i := range nodes {
		if nodes[i].discord_guild_discord_messages == nil {
			continue
		}
		fk := *nodes[i].discord_guild_discord_messages
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(discordguild.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "discord_guild_discord_messages" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dmq *DiscordMessageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dmq.querySpec()
	if len(dmq.modifiers) > 0 {
		_spec.Modifiers = dmq.modifiers
	}
	_spec.Node.Columns = dmq.ctx.Fields
	if len(dmq.ctx.Fields) > 0 {
		_spec.Unique = dmq.ctx.Unique != nil && *dmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dmq.driver, _spec)
}

func (dmq *DiscordMessageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(discordmessage.Table, discordmessage.Columns, sqlgraph.NewFieldSpec(discordmessage.FieldID, field.TypeUUID))
	_spec.From = dmq.sql
	if unique := dmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dmq.path != nil {
		_spec.Unique = true
	}
	if fields := dmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, discordmessage.FieldID)
		for i := range fields {
			if fields[i] != discordmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dmq *DiscordMessageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dmq.driver.Dialect())
	t1 := builder.Table(discordmessage.Table)
	columns := dmq.ctx.Fields
	if len(columns) == 0 {
		columns = discordmessage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dmq.sql != nil {
		selector = dmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dmq.ctx.Unique != nil && *dmq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range dmq.modifiers {
		m(selector)
	}
	for _, p := range dmq.predicates {
		p(selector)
	}
	for _, p := range dmq.order {
		p(selector)
	}
	if offset := dmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (dmq *DiscordMessageQuery) ForUpdate(opts ...sql.LockOption) *DiscordMessageQuery {
	if dmq.driver.Dialect() == dialect.Postgres {
		dmq.Unique(false)
	}
	dmq.modifiers = append(dmq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return dmq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (dmq *DiscordMessageQuery) ForShare(opts ...sql.LockOption) *DiscordMessageQuery {
	if dmq.driver.Dialect() == dialect.Postgres {
		dmq.Unique(false)
	}
	dmq.modifiers = append(dmq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return dmq
}

// WithNamedMessageReactions tells the query-builder to eager-load the nodes that are connected to the "message_reactions"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (dmq *DiscordMessageQuery) WithNamedMessageReactions(name string, opts ...func(*DiscordMessageReactionQuery)) *DiscordMessageQuery {
	query := (&DiscordMessageReactionClient{config: dmq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if dmq.withNamedMessageReactions == nil {
		dmq.withNamedMessageReactions = make(map[string]*DiscordMessageReactionQuery)
	}
	dmq.withNamedMessageReactions[name] = query
	return dmq
}

// DiscordMessageGroupBy is the group-by builder for DiscordMessage entities.
type DiscordMessageGroupBy struct {
	selector
	build *DiscordMessageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dmgb *DiscordMessageGroupBy) Aggregate(fns ...AggregateFunc) *DiscordMessageGroupBy {
	dmgb.fns = append(dmgb.fns, fns...)
	return dmgb
}

// Scan applies the selector query and scans the result into the given value.
func (dmgb *DiscordMessageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dmgb.build.ctx, "GroupBy")
	if err := dmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DiscordMessageQuery, *DiscordMessageGroupBy](ctx, dmgb.build, dmgb, dmgb.build.inters, v)
}

func (dmgb *DiscordMessageGroupBy) sqlScan(ctx context.Context, root *DiscordMessageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dmgb.fns))
	for _, fn := range dmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dmgb.flds)+len(dmgb.fns))
		for _, f := range *dmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DiscordMessageSelect is the builder for selecting fields of DiscordMessage entities.
type DiscordMessageSelect struct {
	*DiscordMessageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dms *DiscordMessageSelect) Aggregate(fns ...AggregateFunc) *DiscordMessageSelect {
	dms.fns = append(dms.fns, fns...)
	return dms
}

// Scan applies the selector query and scans the result into the given value.
func (dms *DiscordMessageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dms.ctx, "Select")
	if err := dms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DiscordMessageQuery, *DiscordMessageSelect](ctx, dms.DiscordMessageQuery, dms, dms.inters, v)
}

func (dms *DiscordMessageSelect) sqlScan(ctx context.Context, root *DiscordMessageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dms.fns))
	for _, fn := range dms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
