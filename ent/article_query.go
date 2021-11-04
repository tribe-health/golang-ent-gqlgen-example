// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"golang-ent-gqlgen-example/ent/article"
	"golang-ent-gqlgen-example/ent/predicate"
	"golang-ent-gqlgen-example/ent/user"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ArticleQuery is the builder for querying Article entities.
type ArticleQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Article
	// eager-loading edges.
	withUser *UserQuery
	withFKs  bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ArticleQuery builder.
func (aq *ArticleQuery) Where(ps ...predicate.Article) *ArticleQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *ArticleQuery) Limit(limit int) *ArticleQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *ArticleQuery) Offset(offset int) *ArticleQuery {
	aq.offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *ArticleQuery) Unique(unique bool) *ArticleQuery {
	aq.unique = &unique
	return aq
}

// Order adds an order step to the query.
func (aq *ArticleQuery) Order(o ...OrderFunc) *ArticleQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryUser chains the current query on the "user" edge.
func (aq *ArticleQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(article.Table, article.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, article.UserTable, article.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Article entity from the query.
// Returns a *NotFoundError when no Article was found.
func (aq *ArticleQuery) First(ctx context.Context) (*Article, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{article.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *ArticleQuery) FirstX(ctx context.Context) *Article {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Article ID from the query.
// Returns a *NotFoundError when no Article ID was found.
func (aq *ArticleQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{article.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *ArticleQuery) FirstIDX(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Article entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Article entity is not found.
// Returns a *NotFoundError when no Article entities are found.
func (aq *ArticleQuery) Only(ctx context.Context) (*Article, error) {
	nodes, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{article.Label}
	default:
		return nil, &NotSingularError{article.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *ArticleQuery) OnlyX(ctx context.Context) *Article {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Article ID in the query.
// Returns a *NotSingularError when exactly one Article ID is not found.
// Returns a *NotFoundError when no entities are found.
func (aq *ArticleQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{article.Label}
	default:
		err = &NotSingularError{article.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *ArticleQuery) OnlyIDX(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Articles.
func (aq *ArticleQuery) All(ctx context.Context) ([]*Article, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *ArticleQuery) AllX(ctx context.Context) []*Article {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Article IDs.
func (aq *ArticleQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := aq.Select(article.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *ArticleQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *ArticleQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *ArticleQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *ArticleQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *ArticleQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ArticleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *ArticleQuery) Clone() *ArticleQuery {
	if aq == nil {
		return nil
	}
	return &ArticleQuery{
		config:     aq.config,
		limit:      aq.limit,
		offset:     aq.offset,
		order:      append([]OrderFunc{}, aq.order...),
		predicates: append([]predicate.Article{}, aq.predicates...),
		withUser:   aq.withUser.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *ArticleQuery) WithUser(opts ...func(*UserQuery)) *ArticleQuery {
	query := &UserQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withUser = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Article.Query().
//		GroupBy(article.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aq *ArticleQuery) GroupBy(field string, fields ...string) *ArticleGroupBy {
	group := &ArticleGroupBy{config: aq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.Article.Query().
//		Select(article.FieldTitle).
//		Scan(ctx, &v)
//
func (aq *ArticleQuery) Select(fields ...string) *ArticleSelect {
	aq.fields = append(aq.fields, fields...)
	return &ArticleSelect{ArticleQuery: aq}
}

func (aq *ArticleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aq.fields {
		if !article.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *ArticleQuery) sqlAll(ctx context.Context) ([]*Article, error) {
	var (
		nodes       = []*Article{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [1]bool{
			aq.withUser != nil,
		}
	)
	if aq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, article.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Article{config: aq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := aq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Article)
		for i := range nodes {
			if nodes[i].user_articles == nil {
				continue
			}
			fk := *nodes[i].user_articles
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_articles" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (aq *ArticleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *ArticleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (aq *ArticleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   article.Table,
			Columns: article.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: article.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if unique := aq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, article.FieldID)
		for i := range fields {
			if fields[i] != article.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *ArticleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(article.Table)
	columns := aq.fields
	if len(columns) == 0 {
		columns = article.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ArticleGroupBy is the group-by builder for Article entities.
type ArticleGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *ArticleGroupBy) Aggregate(fns ...AggregateFunc) *ArticleGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scans the result into the given value.
func (agb *ArticleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (agb *ArticleGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := agb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *ArticleGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: ArticleGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (agb *ArticleGroupBy) StringsX(ctx context.Context) []string {
	v, err := agb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *ArticleGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = agb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{article.Label}
	default:
		err = fmt.Errorf("ent: ArticleGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (agb *ArticleGroupBy) StringX(ctx context.Context) string {
	v, err := agb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *ArticleGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: ArticleGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (agb *ArticleGroupBy) IntsX(ctx context.Context) []int {
	v, err := agb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *ArticleGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = agb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{article.Label}
	default:
		err = fmt.Errorf("ent: ArticleGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (agb *ArticleGroupBy) IntX(ctx context.Context) int {
	v, err := agb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *ArticleGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: ArticleGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (agb *ArticleGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := agb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *ArticleGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = agb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{article.Label}
	default:
		err = fmt.Errorf("ent: ArticleGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (agb *ArticleGroupBy) Float64X(ctx context.Context) float64 {
	v, err := agb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (agb *ArticleGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(agb.fields) > 1 {
		return nil, errors.New("ent: ArticleGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := agb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (agb *ArticleGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := agb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (agb *ArticleGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = agb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{article.Label}
	default:
		err = fmt.Errorf("ent: ArticleGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (agb *ArticleGroupBy) BoolX(ctx context.Context) bool {
	v, err := agb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (agb *ArticleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range agb.fields {
		if !article.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *ArticleGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql.Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(agb.fields)+len(agb.fns))
		for _, f := range agb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(agb.fields...)...)
}

// ArticleSelect is the builder for selecting fields of Article entities.
type ArticleSelect struct {
	*ArticleQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (as *ArticleSelect) Scan(ctx context.Context, v interface{}) error {
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	as.sql = as.ArticleQuery.sqlQuery(ctx)
	return as.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (as *ArticleSelect) ScanX(ctx context.Context, v interface{}) {
	if err := as.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (as *ArticleSelect) Strings(ctx context.Context) ([]string, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: ArticleSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (as *ArticleSelect) StringsX(ctx context.Context) []string {
	v, err := as.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (as *ArticleSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = as.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{article.Label}
	default:
		err = fmt.Errorf("ent: ArticleSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (as *ArticleSelect) StringX(ctx context.Context) string {
	v, err := as.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (as *ArticleSelect) Ints(ctx context.Context) ([]int, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: ArticleSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (as *ArticleSelect) IntsX(ctx context.Context) []int {
	v, err := as.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (as *ArticleSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = as.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{article.Label}
	default:
		err = fmt.Errorf("ent: ArticleSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (as *ArticleSelect) IntX(ctx context.Context) int {
	v, err := as.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (as *ArticleSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: ArticleSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (as *ArticleSelect) Float64sX(ctx context.Context) []float64 {
	v, err := as.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (as *ArticleSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = as.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{article.Label}
	default:
		err = fmt.Errorf("ent: ArticleSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (as *ArticleSelect) Float64X(ctx context.Context) float64 {
	v, err := as.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (as *ArticleSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(as.fields) > 1 {
		return nil, errors.New("ent: ArticleSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := as.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (as *ArticleSelect) BoolsX(ctx context.Context) []bool {
	v, err := as.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (as *ArticleSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = as.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{article.Label}
	default:
		err = fmt.Errorf("ent: ArticleSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (as *ArticleSelect) BoolX(ctx context.Context) bool {
	v, err := as.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (as *ArticleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := as.sql.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
