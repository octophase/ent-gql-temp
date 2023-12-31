// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"myapp/ent/group"
	"myapp/ent/tenant"
	"myapp/ent/user"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[int]
	PageInfo       = entgql.PageInfo[int]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// GroupEdge is the edge representation of Group.
type GroupEdge struct {
	Node   *Group `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// GroupConnection is the connection containing edges to Group.
type GroupConnection struct {
	Edges      []*GroupEdge `json:"edges"`
	PageInfo   PageInfo     `json:"pageInfo"`
	TotalCount int          `json:"totalCount"`
}

func (c *GroupConnection) build(nodes []*Group, pager *groupPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Group
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Group {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Group {
			return nodes[i]
		}
	}
	c.Edges = make([]*GroupEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &GroupEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// GroupPaginateOption enables pagination customization.
type GroupPaginateOption func(*groupPager) error

// WithGroupOrder configures pagination ordering.
func WithGroupOrder(order *GroupOrder) GroupPaginateOption {
	if order == nil {
		order = DefaultGroupOrder
	}
	o := *order
	return func(pager *groupPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultGroupOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithGroupFilter configures pagination filter.
func WithGroupFilter(filter func(*GroupQuery) (*GroupQuery, error)) GroupPaginateOption {
	return func(pager *groupPager) error {
		if filter == nil {
			return errors.New("GroupQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type groupPager struct {
	reverse bool
	order   *GroupOrder
	filter  func(*GroupQuery) (*GroupQuery, error)
}

func newGroupPager(opts []GroupPaginateOption, reverse bool) (*groupPager, error) {
	pager := &groupPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultGroupOrder
	}
	return pager, nil
}

func (p *groupPager) applyFilter(query *GroupQuery) (*GroupQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *groupPager) toCursor(gr *Group) Cursor {
	return p.order.Field.toCursor(gr)
}

func (p *groupPager) applyCursors(query *GroupQuery, after, before *Cursor) (*GroupQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultGroupOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *groupPager) applyOrder(query *GroupQuery) *GroupQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultGroupOrder.Field {
		query = query.Order(DefaultGroupOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *groupPager) orderExpr(query *GroupQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultGroupOrder.Field {
			b.Comma().Ident(DefaultGroupOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Group.
func (gr *GroupQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...GroupPaginateOption,
) (*GroupConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newGroupPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if gr, err = pager.applyFilter(gr); err != nil {
		return nil, err
	}
	conn := &GroupConnection{Edges: []*GroupEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := gr.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if gr, err = pager.applyCursors(gr, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		gr.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := gr.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	gr = pager.applyOrder(gr)
	nodes, err := gr.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// GroupOrderField defines the ordering field of Group.
type GroupOrderField struct {
	// Value extracts the ordering value from the given Group.
	Value    func(*Group) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) group.OrderOption
	toCursor func(*Group) Cursor
}

// GroupOrder defines the ordering of Group.
type GroupOrder struct {
	Direction OrderDirection   `json:"direction"`
	Field     *GroupOrderField `json:"field"`
}

// DefaultGroupOrder is the default ordering of Group.
var DefaultGroupOrder = &GroupOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &GroupOrderField{
		Value: func(gr *Group) (ent.Value, error) {
			return gr.ID, nil
		},
		column: group.FieldID,
		toTerm: group.ByID,
		toCursor: func(gr *Group) Cursor {
			return Cursor{ID: gr.ID}
		},
	},
}

// ToEdge converts Group into GroupEdge.
func (gr *Group) ToEdge(order *GroupOrder) *GroupEdge {
	if order == nil {
		order = DefaultGroupOrder
	}
	return &GroupEdge{
		Node:   gr,
		Cursor: order.Field.toCursor(gr),
	}
}

// TenantEdge is the edge representation of Tenant.
type TenantEdge struct {
	Node   *Tenant `json:"node"`
	Cursor Cursor  `json:"cursor"`
}

// TenantConnection is the connection containing edges to Tenant.
type TenantConnection struct {
	Edges      []*TenantEdge `json:"edges"`
	PageInfo   PageInfo      `json:"pageInfo"`
	TotalCount int           `json:"totalCount"`
}

func (c *TenantConnection) build(nodes []*Tenant, pager *tenantPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Tenant
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Tenant {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Tenant {
			return nodes[i]
		}
	}
	c.Edges = make([]*TenantEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &TenantEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// TenantPaginateOption enables pagination customization.
type TenantPaginateOption func(*tenantPager) error

// WithTenantOrder configures pagination ordering.
func WithTenantOrder(order *TenantOrder) TenantPaginateOption {
	if order == nil {
		order = DefaultTenantOrder
	}
	o := *order
	return func(pager *tenantPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultTenantOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithTenantFilter configures pagination filter.
func WithTenantFilter(filter func(*TenantQuery) (*TenantQuery, error)) TenantPaginateOption {
	return func(pager *tenantPager) error {
		if filter == nil {
			return errors.New("TenantQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type tenantPager struct {
	reverse bool
	order   *TenantOrder
	filter  func(*TenantQuery) (*TenantQuery, error)
}

func newTenantPager(opts []TenantPaginateOption, reverse bool) (*tenantPager, error) {
	pager := &tenantPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultTenantOrder
	}
	return pager, nil
}

func (p *tenantPager) applyFilter(query *TenantQuery) (*TenantQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *tenantPager) toCursor(t *Tenant) Cursor {
	return p.order.Field.toCursor(t)
}

func (p *tenantPager) applyCursors(query *TenantQuery, after, before *Cursor) (*TenantQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultTenantOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *tenantPager) applyOrder(query *TenantQuery) *TenantQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultTenantOrder.Field {
		query = query.Order(DefaultTenantOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *tenantPager) orderExpr(query *TenantQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultTenantOrder.Field {
			b.Comma().Ident(DefaultTenantOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to Tenant.
func (t *TenantQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...TenantPaginateOption,
) (*TenantConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newTenantPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if t, err = pager.applyFilter(t); err != nil {
		return nil, err
	}
	conn := &TenantConnection{Edges: []*TenantEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := t.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if t, err = pager.applyCursors(t, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		t.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := t.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	t = pager.applyOrder(t)
	nodes, err := t.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// TenantOrderField defines the ordering field of Tenant.
type TenantOrderField struct {
	// Value extracts the ordering value from the given Tenant.
	Value    func(*Tenant) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) tenant.OrderOption
	toCursor func(*Tenant) Cursor
}

// TenantOrder defines the ordering of Tenant.
type TenantOrder struct {
	Direction OrderDirection    `json:"direction"`
	Field     *TenantOrderField `json:"field"`
}

// DefaultTenantOrder is the default ordering of Tenant.
var DefaultTenantOrder = &TenantOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &TenantOrderField{
		Value: func(t *Tenant) (ent.Value, error) {
			return t.ID, nil
		},
		column: tenant.FieldID,
		toTerm: tenant.ByID,
		toCursor: func(t *Tenant) Cursor {
			return Cursor{ID: t.ID}
		},
	},
}

// ToEdge converts Tenant into TenantEdge.
func (t *Tenant) ToEdge(order *TenantOrder) *TenantEdge {
	if order == nil {
		order = DefaultTenantOrder
	}
	return &TenantEdge{
		Node:   t,
		Cursor: order.Field.toCursor(t),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *UserConnection) build(nodes []*User, pager *userPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}
	c.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order *UserOrder) UserPaginateOption {
	if order == nil {
		order = DefaultUserOrder
	}
	o := *order
	return func(pager *userPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultUserOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	reverse bool
	order   *UserOrder
	filter  func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption, reverse bool) (*userPager, error) {
	pager := &userPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultUserOrder
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	return p.order.Field.toCursor(u)
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) (*UserQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultUserOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *userPager) applyOrder(query *UserQuery) *UserQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultUserOrder.Field {
		query = query.Order(DefaultUserOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *userPager) orderExpr(query *UserQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultUserOrder.Field {
			b.Comma().Ident(DefaultUserOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}
	conn := &UserConnection{Edges: []*UserEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := u.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if u, err = pager.applyCursors(u, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		u.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := u.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	u = pager.applyOrder(u)
	nodes, err := u.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	// Value extracts the ordering value from the given User.
	Value    func(*User) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) user.OrderOption
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.ID, nil
		},
		column: user.FieldID,
		toTerm: user.ByID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}
