// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (gr *Group) Tenant(ctx context.Context) (*Tenant, error) {
	result, err := gr.Edges.TenantOrErr()
	if IsNotLoaded(err) {
		result, err = gr.QueryTenant().Only(ctx)
	}
	return result, err
}

func (gr *Group) Users(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = gr.NamedUsers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = gr.Edges.UsersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = gr.QueryUsers().All(ctx)
	}
	return result, err
}

func (u *User) Tenant(ctx context.Context) (*Tenant, error) {
	result, err := u.Edges.TenantOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryTenant().Only(ctx)
	}
	return result, err
}

func (u *User) Groups(ctx context.Context) (result []*Group, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedGroups(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.GroupsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryGroups().All(ctx)
	}
	return result, err
}
