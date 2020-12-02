// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/HaHadaxigua/melancholy/ent/exitlog"
	"github.com/HaHadaxigua/melancholy/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// ExitLogQuery is the builder for querying ExitLog entities.
type ExitLogQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.ExitLog
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (elq *ExitLogQuery) Where(ps ...predicate.ExitLog) *ExitLogQuery {
	elq.predicates = append(elq.predicates, ps...)
	return elq
}

// Limit adds a limit step to the query.
func (elq *ExitLogQuery) Limit(limit int) *ExitLogQuery {
	elq.limit = &limit
	return elq
}

// Offset adds an offset step to the query.
func (elq *ExitLogQuery) Offset(offset int) *ExitLogQuery {
	elq.offset = &offset
	return elq
}

// Order adds an order step to the query.
func (elq *ExitLogQuery) Order(o ...OrderFunc) *ExitLogQuery {
	elq.order = append(elq.order, o...)
	return elq
}

// First returns the first ExitLog entity in the query. Returns *NotFoundError when no exitlog was found.
func (elq *ExitLogQuery) First(ctx context.Context) (*ExitLog, error) {
	nodes, err := elq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{exitlog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (elq *ExitLogQuery) FirstX(ctx context.Context) *ExitLog {
	node, err := elq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ExitLog id in the query. Returns *NotFoundError when no id was found.
func (elq *ExitLogQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = elq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{exitlog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (elq *ExitLogQuery) FirstIDX(ctx context.Context) int {
	id, err := elq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only ExitLog entity in the query, returns an error if not exactly one entity was returned.
func (elq *ExitLogQuery) Only(ctx context.Context) (*ExitLog, error) {
	nodes, err := elq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{exitlog.Label}
	default:
		return nil, &NotSingularError{exitlog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (elq *ExitLogQuery) OnlyX(ctx context.Context) *ExitLog {
	node, err := elq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only ExitLog id in the query, returns an error if not exactly one id was returned.
func (elq *ExitLogQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = elq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{exitlog.Label}
	default:
		err = &NotSingularError{exitlog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (elq *ExitLogQuery) OnlyIDX(ctx context.Context) int {
	id, err := elq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ExitLogs.
func (elq *ExitLogQuery) All(ctx context.Context) ([]*ExitLog, error) {
	if err := elq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return elq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (elq *ExitLogQuery) AllX(ctx context.Context) []*ExitLog {
	nodes, err := elq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ExitLog ids.
func (elq *ExitLogQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := elq.Select(exitlog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (elq *ExitLogQuery) IDsX(ctx context.Context) []int {
	ids, err := elq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (elq *ExitLogQuery) Count(ctx context.Context) (int, error) {
	if err := elq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return elq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (elq *ExitLogQuery) CountX(ctx context.Context) int {
	count, err := elq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (elq *ExitLogQuery) Exist(ctx context.Context) (bool, error) {
	if err := elq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return elq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (elq *ExitLogQuery) ExistX(ctx context.Context) bool {
	exist, err := elq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (elq *ExitLogQuery) Clone() *ExitLogQuery {
	if elq == nil {
		return nil
	}
	return &ExitLogQuery{
		config:     elq.config,
		limit:      elq.limit,
		offset:     elq.offset,
		order:      append([]OrderFunc{}, elq.order...),
		unique:     append([]string{}, elq.unique...),
		predicates: append([]predicate.ExitLog{}, elq.predicates...),
		// clone intermediate query.
		sql:  elq.sql.Clone(),
		path: elq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID int `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ExitLog.Query().
//		GroupBy(exitlog.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (elq *ExitLogQuery) GroupBy(field string, fields ...string) *ExitLogGroupBy {
	group := &ExitLogGroupBy{config: elq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := elq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return elq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		UserID int `json:"user_id,omitempty"`
//	}
//
//	client.ExitLog.Query().
//		Select(exitlog.FieldUserID).
//		Scan(ctx, &v)
//
func (elq *ExitLogQuery) Select(field string, fields ...string) *ExitLogSelect {
	selector := &ExitLogSelect{config: elq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := elq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return elq.sqlQuery(), nil
	}
	return selector
}

func (elq *ExitLogQuery) prepareQuery(ctx context.Context) error {
	if elq.path != nil {
		prev, err := elq.path(ctx)
		if err != nil {
			return err
		}
		elq.sql = prev
	}
	return nil
}

func (elq *ExitLogQuery) sqlAll(ctx context.Context) ([]*ExitLog, error) {
	var (
		nodes   = []*ExitLog{}
		withFKs = elq.withFKs
		_spec   = elq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, exitlog.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &ExitLog{config: elq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, elq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (elq *ExitLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := elq.querySpec()
	return sqlgraph.CountNodes(ctx, elq.driver, _spec)
}

func (elq *ExitLogQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := elq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (elq *ExitLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   exitlog.Table,
			Columns: exitlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: exitlog.FieldID,
			},
		},
		From:   elq.sql,
		Unique: true,
	}
	if ps := elq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := elq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := elq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := elq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, exitlog.ValidColumn)
			}
		}
	}
	return _spec
}

func (elq *ExitLogQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(elq.driver.Dialect())
	t1 := builder.Table(exitlog.Table)
	selector := builder.Select(t1.Columns(exitlog.Columns...)...).From(t1)
	if elq.sql != nil {
		selector = elq.sql
		selector.Select(selector.Columns(exitlog.Columns...)...)
	}
	for _, p := range elq.predicates {
		p(selector)
	}
	for _, p := range elq.order {
		p(selector, exitlog.ValidColumn)
	}
	if offset := elq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := elq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ExitLogGroupBy is the builder for group-by ExitLog entities.
type ExitLogGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (elgb *ExitLogGroupBy) Aggregate(fns ...AggregateFunc) *ExitLogGroupBy {
	elgb.fns = append(elgb.fns, fns...)
	return elgb
}

// Scan applies the group-by query and scan the result into the given value.
func (elgb *ExitLogGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := elgb.path(ctx)
	if err != nil {
		return err
	}
	elgb.sql = query
	return elgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (elgb *ExitLogGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := elgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (elgb *ExitLogGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(elgb.fields) > 1 {
		return nil, errors.New("ent: ExitLogGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := elgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (elgb *ExitLogGroupBy) StringsX(ctx context.Context) []string {
	v, err := elgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (elgb *ExitLogGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = elgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{exitlog.Label}
	default:
		err = fmt.Errorf("ent: ExitLogGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (elgb *ExitLogGroupBy) StringX(ctx context.Context) string {
	v, err := elgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (elgb *ExitLogGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(elgb.fields) > 1 {
		return nil, errors.New("ent: ExitLogGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := elgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (elgb *ExitLogGroupBy) IntsX(ctx context.Context) []int {
	v, err := elgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (elgb *ExitLogGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = elgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{exitlog.Label}
	default:
		err = fmt.Errorf("ent: ExitLogGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (elgb *ExitLogGroupBy) IntX(ctx context.Context) int {
	v, err := elgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (elgb *ExitLogGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(elgb.fields) > 1 {
		return nil, errors.New("ent: ExitLogGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := elgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (elgb *ExitLogGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := elgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (elgb *ExitLogGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = elgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{exitlog.Label}
	default:
		err = fmt.Errorf("ent: ExitLogGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (elgb *ExitLogGroupBy) Float64X(ctx context.Context) float64 {
	v, err := elgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (elgb *ExitLogGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(elgb.fields) > 1 {
		return nil, errors.New("ent: ExitLogGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := elgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (elgb *ExitLogGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := elgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (elgb *ExitLogGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = elgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{exitlog.Label}
	default:
		err = fmt.Errorf("ent: ExitLogGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (elgb *ExitLogGroupBy) BoolX(ctx context.Context) bool {
	v, err := elgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (elgb *ExitLogGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range elgb.fields {
		if !exitlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := elgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := elgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (elgb *ExitLogGroupBy) sqlQuery() *sql.Selector {
	selector := elgb.sql
	columns := make([]string, 0, len(elgb.fields)+len(elgb.fns))
	columns = append(columns, elgb.fields...)
	for _, fn := range elgb.fns {
		columns = append(columns, fn(selector, exitlog.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(elgb.fields...)
}

// ExitLogSelect is the builder for select fields of ExitLog entities.
type ExitLogSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (els *ExitLogSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := els.path(ctx)
	if err != nil {
		return err
	}
	els.sql = query
	return els.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (els *ExitLogSelect) ScanX(ctx context.Context, v interface{}) {
	if err := els.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (els *ExitLogSelect) Strings(ctx context.Context) ([]string, error) {
	if len(els.fields) > 1 {
		return nil, errors.New("ent: ExitLogSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := els.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (els *ExitLogSelect) StringsX(ctx context.Context) []string {
	v, err := els.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (els *ExitLogSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = els.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{exitlog.Label}
	default:
		err = fmt.Errorf("ent: ExitLogSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (els *ExitLogSelect) StringX(ctx context.Context) string {
	v, err := els.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (els *ExitLogSelect) Ints(ctx context.Context) ([]int, error) {
	if len(els.fields) > 1 {
		return nil, errors.New("ent: ExitLogSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := els.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (els *ExitLogSelect) IntsX(ctx context.Context) []int {
	v, err := els.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (els *ExitLogSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = els.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{exitlog.Label}
	default:
		err = fmt.Errorf("ent: ExitLogSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (els *ExitLogSelect) IntX(ctx context.Context) int {
	v, err := els.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (els *ExitLogSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(els.fields) > 1 {
		return nil, errors.New("ent: ExitLogSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := els.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (els *ExitLogSelect) Float64sX(ctx context.Context) []float64 {
	v, err := els.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (els *ExitLogSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = els.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{exitlog.Label}
	default:
		err = fmt.Errorf("ent: ExitLogSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (els *ExitLogSelect) Float64X(ctx context.Context) float64 {
	v, err := els.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (els *ExitLogSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(els.fields) > 1 {
		return nil, errors.New("ent: ExitLogSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := els.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (els *ExitLogSelect) BoolsX(ctx context.Context) []bool {
	v, err := els.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (els *ExitLogSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = els.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{exitlog.Label}
	default:
		err = fmt.Errorf("ent: ExitLogSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (els *ExitLogSelect) BoolX(ctx context.Context) bool {
	v, err := els.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (els *ExitLogSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range els.fields {
		if !exitlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := els.sqlQuery().Query()
	if err := els.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (els *ExitLogSelect) sqlQuery() sql.Querier {
	selector := els.sql
	selector.Select(selector.Columns(els.fields...)...)
	return selector
}
