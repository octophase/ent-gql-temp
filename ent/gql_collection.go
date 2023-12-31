// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"myapp/ent/group"
	"myapp/ent/tenant"
	"myapp/ent/user"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (gr *GroupQuery) CollectFields(ctx context.Context, satisfies ...string) (*GroupQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return gr, nil
	}
	if err := gr.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return gr, nil
}

func (gr *GroupQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(group.Columns))
		selectedFields = []string{group.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "tenant":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TenantClient{config: gr.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, tenantImplementors)...); err != nil {
				return err
			}
			gr.withTenant = query
			if _, ok := fieldSeen[group.FieldTenantID]; !ok {
				selectedFields = append(selectedFields, group.FieldTenantID)
				fieldSeen[group.FieldTenantID] = struct{}{}
			}
		case "users":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: gr.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			gr.WithNamedUsers(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "tenantID":
			if _, ok := fieldSeen[group.FieldTenantID]; !ok {
				selectedFields = append(selectedFields, group.FieldTenantID)
				fieldSeen[group.FieldTenantID] = struct{}{}
			}
		case "username":
			if _, ok := fieldSeen[group.FieldUsername]; !ok {
				selectedFields = append(selectedFields, group.FieldUsername)
				fieldSeen[group.FieldUsername] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		gr.Select(selectedFields...)
	}
	return nil
}

type groupPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []GroupPaginateOption
}

func newGroupPaginateArgs(rv map[string]any) *groupPaginateArgs {
	args := &groupPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*GroupWhereInput); ok {
		args.opts = append(args.opts, WithGroupFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TenantQuery) CollectFields(ctx context.Context, satisfies ...string) (*TenantQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return t, nil
	}
	if err := t.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TenantQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(tenant.Columns))
		selectedFields = []string{tenant.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "name":
			if _, ok := fieldSeen[tenant.FieldName]; !ok {
				selectedFields = append(selectedFields, tenant.FieldName)
				fieldSeen[tenant.FieldName] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		t.Select(selectedFields...)
	}
	return nil
}

type tenantPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TenantPaginateOption
}

func newTenantPaginateArgs(rv map[string]any) *tenantPaginateArgs {
	args := &tenantPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*TenantWhereInput); ok {
		args.opts = append(args.opts, WithTenantFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(user.Columns))
		selectedFields = []string{user.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "tenant":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TenantClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, tenantImplementors)...); err != nil {
				return err
			}
			u.withTenant = query
			if _, ok := fieldSeen[user.FieldTenantID]; !ok {
				selectedFields = append(selectedFields, user.FieldTenantID)
				fieldSeen[user.FieldTenantID] = struct{}{}
			}
		case "groups":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&GroupClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, mayAddCondition(satisfies, groupImplementors)...); err != nil {
				return err
			}
			u.WithNamedGroups(alias, func(wq *GroupQuery) {
				*wq = *query
			})
		case "tenantID":
			if _, ok := fieldSeen[user.FieldTenantID]; !ok {
				selectedFields = append(selectedFields, user.FieldTenantID)
				fieldSeen[user.FieldTenantID] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[user.FieldName]; !ok {
				selectedFields = append(selectedFields, user.FieldName)
				fieldSeen[user.FieldName] = struct{}{}
			}
		case "username":
			if _, ok := fieldSeen[user.FieldUsername]; !ok {
				selectedFields = append(selectedFields, user.FieldUsername)
				fieldSeen[user.FieldUsername] = struct{}{}
			}
		case "email":
			if _, ok := fieldSeen[user.FieldEmail]; !ok {
				selectedFields = append(selectedFields, user.FieldEmail)
				fieldSeen[user.FieldEmail] = struct{}{}
			}
		case "description":
			if _, ok := fieldSeen[user.FieldDescription]; !ok {
				selectedFields = append(selectedFields, user.FieldDescription)
				fieldSeen[user.FieldDescription] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[user.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, user.FieldCreatedAt)
				fieldSeen[user.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[user.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, user.FieldUpdatedAt)
				fieldSeen[user.FieldUpdatedAt] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		u.Select(selectedFields...)
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]any) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*UserWhereInput); ok {
		args.opts = append(args.opts, WithUserFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}
