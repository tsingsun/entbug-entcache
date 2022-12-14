// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/bug/ent/predicate"
	"entgo.io/bug/ent/securityjournal"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SecurityJournalQuery is the builder for querying SecurityJournal entities.
type SecurityJournalQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SecurityJournal
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SecurityJournalQuery builder.
func (sjq *SecurityJournalQuery) Where(ps ...predicate.SecurityJournal) *SecurityJournalQuery {
	sjq.predicates = append(sjq.predicates, ps...)
	return sjq
}

// Limit adds a limit step to the query.
func (sjq *SecurityJournalQuery) Limit(limit int) *SecurityJournalQuery {
	sjq.limit = &limit
	return sjq
}

// Offset adds an offset step to the query.
func (sjq *SecurityJournalQuery) Offset(offset int) *SecurityJournalQuery {
	sjq.offset = &offset
	return sjq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sjq *SecurityJournalQuery) Unique(unique bool) *SecurityJournalQuery {
	sjq.unique = &unique
	return sjq
}

// Order adds an order step to the query.
func (sjq *SecurityJournalQuery) Order(o ...OrderFunc) *SecurityJournalQuery {
	sjq.order = append(sjq.order, o...)
	return sjq
}

// First returns the first SecurityJournal entity from the query.
// Returns a *NotFoundError when no SecurityJournal was found.
func (sjq *SecurityJournalQuery) First(ctx context.Context) (*SecurityJournal, error) {
	nodes, err := sjq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{securityjournal.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sjq *SecurityJournalQuery) FirstX(ctx context.Context) *SecurityJournal {
	node, err := sjq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SecurityJournal ID from the query.
// Returns a *NotFoundError when no SecurityJournal ID was found.
func (sjq *SecurityJournalQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sjq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{securityjournal.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sjq *SecurityJournalQuery) FirstIDX(ctx context.Context) int {
	id, err := sjq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SecurityJournal entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SecurityJournal entity is found.
// Returns a *NotFoundError when no SecurityJournal entities are found.
func (sjq *SecurityJournalQuery) Only(ctx context.Context) (*SecurityJournal, error) {
	nodes, err := sjq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{securityjournal.Label}
	default:
		return nil, &NotSingularError{securityjournal.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sjq *SecurityJournalQuery) OnlyX(ctx context.Context) *SecurityJournal {
	node, err := sjq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SecurityJournal ID in the query.
// Returns a *NotSingularError when more than one SecurityJournal ID is found.
// Returns a *NotFoundError when no entities are found.
func (sjq *SecurityJournalQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sjq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{securityjournal.Label}
	default:
		err = &NotSingularError{securityjournal.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sjq *SecurityJournalQuery) OnlyIDX(ctx context.Context) int {
	id, err := sjq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SecurityJournals.
func (sjq *SecurityJournalQuery) All(ctx context.Context) ([]*SecurityJournal, error) {
	if err := sjq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sjq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sjq *SecurityJournalQuery) AllX(ctx context.Context) []*SecurityJournal {
	nodes, err := sjq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SecurityJournal IDs.
func (sjq *SecurityJournalQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := sjq.Select(securityjournal.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sjq *SecurityJournalQuery) IDsX(ctx context.Context) []int {
	ids, err := sjq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sjq *SecurityJournalQuery) Count(ctx context.Context) (int, error) {
	if err := sjq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sjq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sjq *SecurityJournalQuery) CountX(ctx context.Context) int {
	count, err := sjq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sjq *SecurityJournalQuery) Exist(ctx context.Context) (bool, error) {
	if err := sjq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sjq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sjq *SecurityJournalQuery) ExistX(ctx context.Context) bool {
	exist, err := sjq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SecurityJournalQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sjq *SecurityJournalQuery) Clone() *SecurityJournalQuery {
	if sjq == nil {
		return nil
	}
	return &SecurityJournalQuery{
		config:     sjq.config,
		limit:      sjq.limit,
		offset:     sjq.offset,
		order:      append([]OrderFunc{}, sjq.order...),
		predicates: append([]predicate.SecurityJournal{}, sjq.predicates...),
		// clone intermediate query.
		sql:    sjq.sql.Clone(),
		path:   sjq.path,
		unique: sjq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SecurityAccountID int `json:"security_account_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SecurityJournal.Query().
//		GroupBy(securityjournal.FieldSecurityAccountID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sjq *SecurityJournalQuery) GroupBy(field string, fields ...string) *SecurityJournalGroupBy {
	grbuild := &SecurityJournalGroupBy{config: sjq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sjq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sjq.sqlQuery(ctx), nil
	}
	grbuild.label = securityjournal.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SecurityAccountID int `json:"security_account_id,omitempty"`
//	}
//
//	client.SecurityJournal.Query().
//		Select(securityjournal.FieldSecurityAccountID).
//		Scan(ctx, &v)
//
func (sjq *SecurityJournalQuery) Select(fields ...string) *SecurityJournalSelect {
	sjq.fields = append(sjq.fields, fields...)
	selbuild := &SecurityJournalSelect{SecurityJournalQuery: sjq}
	selbuild.label = securityjournal.Label
	selbuild.flds, selbuild.scan = &sjq.fields, selbuild.Scan
	return selbuild
}

func (sjq *SecurityJournalQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sjq.fields {
		if !securityjournal.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sjq.path != nil {
		prev, err := sjq.path(ctx)
		if err != nil {
			return err
		}
		sjq.sql = prev
	}
	return nil
}

func (sjq *SecurityJournalQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SecurityJournal, error) {
	var (
		nodes = []*SecurityJournal{}
		_spec = sjq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*SecurityJournal).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &SecurityJournal{config: sjq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sjq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sjq *SecurityJournalQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sjq.querySpec()
	_spec.Node.Columns = sjq.fields
	if len(sjq.fields) > 0 {
		_spec.Unique = sjq.unique != nil && *sjq.unique
	}
	return sqlgraph.CountNodes(ctx, sjq.driver, _spec)
}

func (sjq *SecurityJournalQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sjq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (sjq *SecurityJournalQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   securityjournal.Table,
			Columns: securityjournal.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: securityjournal.FieldID,
			},
		},
		From:   sjq.sql,
		Unique: true,
	}
	if unique := sjq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sjq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, securityjournal.FieldID)
		for i := range fields {
			if fields[i] != securityjournal.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sjq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sjq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sjq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sjq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sjq *SecurityJournalQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sjq.driver.Dialect())
	t1 := builder.Table(securityjournal.Table)
	columns := sjq.fields
	if len(columns) == 0 {
		columns = securityjournal.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sjq.sql != nil {
		selector = sjq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sjq.unique != nil && *sjq.unique {
		selector.Distinct()
	}
	for _, p := range sjq.predicates {
		p(selector)
	}
	for _, p := range sjq.order {
		p(selector)
	}
	if offset := sjq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sjq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SecurityJournalGroupBy is the group-by builder for SecurityJournal entities.
type SecurityJournalGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sjgb *SecurityJournalGroupBy) Aggregate(fns ...AggregateFunc) *SecurityJournalGroupBy {
	sjgb.fns = append(sjgb.fns, fns...)
	return sjgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sjgb *SecurityJournalGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sjgb.path(ctx)
	if err != nil {
		return err
	}
	sjgb.sql = query
	return sjgb.sqlScan(ctx, v)
}

func (sjgb *SecurityJournalGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sjgb.fields {
		if !securityjournal.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sjgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sjgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sjgb *SecurityJournalGroupBy) sqlQuery() *sql.Selector {
	selector := sjgb.sql.Select()
	aggregation := make([]string, 0, len(sjgb.fns))
	for _, fn := range sjgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sjgb.fields)+len(sjgb.fns))
		for _, f := range sjgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sjgb.fields...)...)
}

// SecurityJournalSelect is the builder for selecting fields of SecurityJournal entities.
type SecurityJournalSelect struct {
	*SecurityJournalQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sjs *SecurityJournalSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sjs.prepareQuery(ctx); err != nil {
		return err
	}
	sjs.sql = sjs.SecurityJournalQuery.sqlQuery(ctx)
	return sjs.sqlScan(ctx, v)
}

func (sjs *SecurityJournalSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sjs.sql.Query()
	if err := sjs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
