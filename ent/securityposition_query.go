// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/bug/ent/predicate"
	"entgo.io/bug/ent/securityposition"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SecurityPositionQuery is the builder for querying SecurityPosition entities.
type SecurityPositionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SecurityPosition
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SecurityPositionQuery builder.
func (spq *SecurityPositionQuery) Where(ps ...predicate.SecurityPosition) *SecurityPositionQuery {
	spq.predicates = append(spq.predicates, ps...)
	return spq
}

// Limit adds a limit step to the query.
func (spq *SecurityPositionQuery) Limit(limit int) *SecurityPositionQuery {
	spq.limit = &limit
	return spq
}

// Offset adds an offset step to the query.
func (spq *SecurityPositionQuery) Offset(offset int) *SecurityPositionQuery {
	spq.offset = &offset
	return spq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (spq *SecurityPositionQuery) Unique(unique bool) *SecurityPositionQuery {
	spq.unique = &unique
	return spq
}

// Order adds an order step to the query.
func (spq *SecurityPositionQuery) Order(o ...OrderFunc) *SecurityPositionQuery {
	spq.order = append(spq.order, o...)
	return spq
}

// First returns the first SecurityPosition entity from the query.
// Returns a *NotFoundError when no SecurityPosition was found.
func (spq *SecurityPositionQuery) First(ctx context.Context) (*SecurityPosition, error) {
	nodes, err := spq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{securityposition.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (spq *SecurityPositionQuery) FirstX(ctx context.Context) *SecurityPosition {
	node, err := spq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SecurityPosition ID from the query.
// Returns a *NotFoundError when no SecurityPosition ID was found.
func (spq *SecurityPositionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = spq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{securityposition.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (spq *SecurityPositionQuery) FirstIDX(ctx context.Context) int {
	id, err := spq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SecurityPosition entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SecurityPosition entity is found.
// Returns a *NotFoundError when no SecurityPosition entities are found.
func (spq *SecurityPositionQuery) Only(ctx context.Context) (*SecurityPosition, error) {
	nodes, err := spq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{securityposition.Label}
	default:
		return nil, &NotSingularError{securityposition.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (spq *SecurityPositionQuery) OnlyX(ctx context.Context) *SecurityPosition {
	node, err := spq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SecurityPosition ID in the query.
// Returns a *NotSingularError when more than one SecurityPosition ID is found.
// Returns a *NotFoundError when no entities are found.
func (spq *SecurityPositionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = spq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{securityposition.Label}
	default:
		err = &NotSingularError{securityposition.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (spq *SecurityPositionQuery) OnlyIDX(ctx context.Context) int {
	id, err := spq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SecurityPositions.
func (spq *SecurityPositionQuery) All(ctx context.Context) ([]*SecurityPosition, error) {
	if err := spq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return spq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (spq *SecurityPositionQuery) AllX(ctx context.Context) []*SecurityPosition {
	nodes, err := spq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SecurityPosition IDs.
func (spq *SecurityPositionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := spq.Select(securityposition.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (spq *SecurityPositionQuery) IDsX(ctx context.Context) []int {
	ids, err := spq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (spq *SecurityPositionQuery) Count(ctx context.Context) (int, error) {
	if err := spq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return spq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (spq *SecurityPositionQuery) CountX(ctx context.Context) int {
	count, err := spq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (spq *SecurityPositionQuery) Exist(ctx context.Context) (bool, error) {
	if err := spq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return spq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (spq *SecurityPositionQuery) ExistX(ctx context.Context) bool {
	exist, err := spq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SecurityPositionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (spq *SecurityPositionQuery) Clone() *SecurityPositionQuery {
	if spq == nil {
		return nil
	}
	return &SecurityPositionQuery{
		config:     spq.config,
		limit:      spq.limit,
		offset:     spq.offset,
		order:      append([]OrderFunc{}, spq.order...),
		predicates: append([]predicate.SecurityPosition{}, spq.predicates...),
		// clone intermediate query.
		sql:    spq.sql.Clone(),
		path:   spq.path,
		unique: spq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ParentID int `json:"parent_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SecurityPosition.Query().
//		GroupBy(securityposition.FieldParentID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (spq *SecurityPositionQuery) GroupBy(field string, fields ...string) *SecurityPositionGroupBy {
	grbuild := &SecurityPositionGroupBy{config: spq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return spq.sqlQuery(ctx), nil
	}
	grbuild.label = securityposition.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ParentID int `json:"parent_id,omitempty"`
//	}
//
//	client.SecurityPosition.Query().
//		Select(securityposition.FieldParentID).
//		Scan(ctx, &v)
//
func (spq *SecurityPositionQuery) Select(fields ...string) *SecurityPositionSelect {
	spq.fields = append(spq.fields, fields...)
	selbuild := &SecurityPositionSelect{SecurityPositionQuery: spq}
	selbuild.label = securityposition.Label
	selbuild.flds, selbuild.scan = &spq.fields, selbuild.Scan
	return selbuild
}

func (spq *SecurityPositionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range spq.fields {
		if !securityposition.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if spq.path != nil {
		prev, err := spq.path(ctx)
		if err != nil {
			return err
		}
		spq.sql = prev
	}
	return nil
}

func (spq *SecurityPositionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SecurityPosition, error) {
	var (
		nodes = []*SecurityPosition{}
		_spec = spq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*SecurityPosition).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &SecurityPosition{config: spq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, spq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (spq *SecurityPositionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := spq.querySpec()
	_spec.Node.Columns = spq.fields
	if len(spq.fields) > 0 {
		_spec.Unique = spq.unique != nil && *spq.unique
	}
	return sqlgraph.CountNodes(ctx, spq.driver, _spec)
}

func (spq *SecurityPositionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := spq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (spq *SecurityPositionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   securityposition.Table,
			Columns: securityposition.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: securityposition.FieldID,
			},
		},
		From:   spq.sql,
		Unique: true,
	}
	if unique := spq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := spq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, securityposition.FieldID)
		for i := range fields {
			if fields[i] != securityposition.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := spq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := spq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := spq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := spq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (spq *SecurityPositionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(spq.driver.Dialect())
	t1 := builder.Table(securityposition.Table)
	columns := spq.fields
	if len(columns) == 0 {
		columns = securityposition.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if spq.sql != nil {
		selector = spq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if spq.unique != nil && *spq.unique {
		selector.Distinct()
	}
	for _, p := range spq.predicates {
		p(selector)
	}
	for _, p := range spq.order {
		p(selector)
	}
	if offset := spq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := spq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SecurityPositionGroupBy is the group-by builder for SecurityPosition entities.
type SecurityPositionGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (spgb *SecurityPositionGroupBy) Aggregate(fns ...AggregateFunc) *SecurityPositionGroupBy {
	spgb.fns = append(spgb.fns, fns...)
	return spgb
}

// Scan applies the group-by query and scans the result into the given value.
func (spgb *SecurityPositionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := spgb.path(ctx)
	if err != nil {
		return err
	}
	spgb.sql = query
	return spgb.sqlScan(ctx, v)
}

func (spgb *SecurityPositionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range spgb.fields {
		if !securityposition.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := spgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := spgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (spgb *SecurityPositionGroupBy) sqlQuery() *sql.Selector {
	selector := spgb.sql.Select()
	aggregation := make([]string, 0, len(spgb.fns))
	for _, fn := range spgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(spgb.fields)+len(spgb.fns))
		for _, f := range spgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(spgb.fields...)...)
}

// SecurityPositionSelect is the builder for selecting fields of SecurityPosition entities.
type SecurityPositionSelect struct {
	*SecurityPositionQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sps *SecurityPositionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sps.prepareQuery(ctx); err != nil {
		return err
	}
	sps.sql = sps.SecurityPositionQuery.sqlQuery(ctx)
	return sps.sqlScan(ctx, v)
}

func (sps *SecurityPositionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sps.sql.Query()
	if err := sps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}