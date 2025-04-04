// Code generated by ent, DO NOT EDIT.

package ent

import (
	"baseProject/ent/predicate"
	"baseProject/ent/price"
	"baseProject/ent/product"
	"baseProject/ent/review"
	"baseProject/ent/store"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StoreQuery is the builder for querying Store entities.
type StoreQuery struct {
	config
	ctx          *QueryContext
	order        []store.OrderOption
	inters       []Interceptor
	predicates   []predicate.Store
	withProducts *ProductQuery
	withPrices   *PriceQuery
	withReviews  *ReviewQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StoreQuery builder.
func (sq *StoreQuery) Where(ps ...predicate.Store) *StoreQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *StoreQuery) Limit(limit int) *StoreQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *StoreQuery) Offset(offset int) *StoreQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *StoreQuery) Unique(unique bool) *StoreQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *StoreQuery) Order(o ...store.OrderOption) *StoreQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryProducts chains the current query on the "products" edge.
func (sq *StoreQuery) QueryProducts() *ProductQuery {
	query := (&ProductClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(store.Table, store.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, store.ProductsTable, store.ProductsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPrices chains the current query on the "prices" edge.
func (sq *StoreQuery) QueryPrices() *PriceQuery {
	query := (&PriceClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(store.Table, store.FieldID, selector),
			sqlgraph.To(price.Table, price.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, store.PricesTable, store.PricesColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReviews chains the current query on the "reviews" edge.
func (sq *StoreQuery) QueryReviews() *ReviewQuery {
	query := (&ReviewClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(store.Table, store.FieldID, selector),
			sqlgraph.To(review.Table, review.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, store.ReviewsTable, store.ReviewsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Store entity from the query.
// Returns a *NotFoundError when no Store was found.
func (sq *StoreQuery) First(ctx context.Context) (*Store, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{store.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *StoreQuery) FirstX(ctx context.Context) *Store {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Store ID from the query.
// Returns a *NotFoundError when no Store ID was found.
func (sq *StoreQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{store.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *StoreQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Store entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Store entity is found.
// Returns a *NotFoundError when no Store entities are found.
func (sq *StoreQuery) Only(ctx context.Context) (*Store, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{store.Label}
	default:
		return nil, &NotSingularError{store.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *StoreQuery) OnlyX(ctx context.Context) *Store {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Store ID in the query.
// Returns a *NotSingularError when more than one Store ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *StoreQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{store.Label}
	default:
		err = &NotSingularError{store.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *StoreQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Stores.
func (sq *StoreQuery) All(ctx context.Context) ([]*Store, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryAll)
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Store, *StoreQuery]()
	return withInterceptors[[]*Store](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *StoreQuery) AllX(ctx context.Context) []*Store {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Store IDs.
func (sq *StoreQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryIDs)
	if err = sq.Select(store.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *StoreQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *StoreQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryCount)
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*StoreQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *StoreQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *StoreQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryExist)
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *StoreQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StoreQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *StoreQuery) Clone() *StoreQuery {
	if sq == nil {
		return nil
	}
	return &StoreQuery{
		config:       sq.config,
		ctx:          sq.ctx.Clone(),
		order:        append([]store.OrderOption{}, sq.order...),
		inters:       append([]Interceptor{}, sq.inters...),
		predicates:   append([]predicate.Store{}, sq.predicates...),
		withProducts: sq.withProducts.Clone(),
		withPrices:   sq.withPrices.Clone(),
		withReviews:  sq.withReviews.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithProducts tells the query-builder to eager-load the nodes that are connected to
// the "products" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *StoreQuery) WithProducts(opts ...func(*ProductQuery)) *StoreQuery {
	query := (&ProductClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withProducts = query
	return sq
}

// WithPrices tells the query-builder to eager-load the nodes that are connected to
// the "prices" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *StoreQuery) WithPrices(opts ...func(*PriceQuery)) *StoreQuery {
	query := (&PriceClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withPrices = query
	return sq
}

// WithReviews tells the query-builder to eager-load the nodes that are connected to
// the "reviews" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *StoreQuery) WithReviews(opts ...func(*ReviewQuery)) *StoreQuery {
	query := (&ReviewClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withReviews = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Store.Query().
//		GroupBy(store.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *StoreQuery) GroupBy(field string, fields ...string) *StoreGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &StoreGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = store.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Store.Query().
//		Select(store.FieldName).
//		Scan(ctx, &v)
func (sq *StoreQuery) Select(fields ...string) *StoreSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &StoreSelect{StoreQuery: sq}
	sbuild.label = store.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a StoreSelect configured with the given aggregations.
func (sq *StoreQuery) Aggregate(fns ...AggregateFunc) *StoreSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *StoreQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !store.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *StoreQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Store, error) {
	var (
		nodes       = []*Store{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [3]bool{
			sq.withProducts != nil,
			sq.withPrices != nil,
			sq.withReviews != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, store.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Store).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Store{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withProducts; query != nil {
		if err := sq.loadProducts(ctx, query, nodes,
			func(n *Store) { n.Edges.Products = []*Product{} },
			func(n *Store, e *Product) { n.Edges.Products = append(n.Edges.Products, e) }); err != nil {
			return nil, err
		}
	}
	if query := sq.withPrices; query != nil {
		if err := sq.loadPrices(ctx, query, nodes,
			func(n *Store) { n.Edges.Prices = []*Price{} },
			func(n *Store, e *Price) { n.Edges.Prices = append(n.Edges.Prices, e) }); err != nil {
			return nil, err
		}
	}
	if query := sq.withReviews; query != nil {
		if err := sq.loadReviews(ctx, query, nodes,
			func(n *Store) { n.Edges.Reviews = []*Review{} },
			func(n *Store, e *Review) { n.Edges.Reviews = append(n.Edges.Reviews, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *StoreQuery) loadProducts(ctx context.Context, query *ProductQuery, nodes []*Store, init func(*Store), assign func(*Store, *Product)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Store)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Product(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(store.ProductsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.store_products
		if fk == nil {
			return fmt.Errorf(`foreign-key "store_products" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "store_products" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (sq *StoreQuery) loadPrices(ctx context.Context, query *PriceQuery, nodes []*Store, init func(*Store), assign func(*Store, *Price)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Store)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(price.FieldStoreID)
	}
	query.Where(predicate.Price(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(store.PricesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.StoreID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "store_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (sq *StoreQuery) loadReviews(ctx context.Context, query *ReviewQuery, nodes []*Store, init func(*Store), assign func(*Store, *Review)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Store)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(review.FieldStoreID)
	}
	query.Where(predicate.Review(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(store.ReviewsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.StoreID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "store_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (sq *StoreQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *StoreQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(store.Table, store.Columns, sqlgraph.NewFieldSpec(store.FieldID, field.TypeInt))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, store.FieldID)
		for i := range fields {
			if fields[i] != store.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *StoreQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(store.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = store.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StoreGroupBy is the group-by builder for Store entities.
type StoreGroupBy struct {
	selector
	build *StoreQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *StoreGroupBy) Aggregate(fns ...AggregateFunc) *StoreGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *StoreGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, ent.OpQueryGroupBy)
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StoreQuery, *StoreGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *StoreGroupBy) sqlScan(ctx context.Context, root *StoreQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// StoreSelect is the builder for selecting fields of Store entities.
type StoreSelect struct {
	*StoreQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *StoreSelect) Aggregate(fns ...AggregateFunc) *StoreSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *StoreSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, ent.OpQuerySelect)
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StoreQuery, *StoreSelect](ctx, ss.StoreQuery, ss, ss.inters, v)
}

func (ss *StoreSelect) sqlScan(ctx context.Context, root *StoreQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
