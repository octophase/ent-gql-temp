// Code generated by ent, DO NOT EDIT.

package ent

import (
	"errors"
	"fmt"
	"myapp/ent/group"
	"myapp/ent/predicate"
	"myapp/ent/tenant"
	"myapp/ent/user"
	"time"
)

// GroupWhereInput represents a where input for filtering Group queries.
type GroupWhereInput struct {
	Predicates []predicate.Group  `json:"-"`
	Not        *GroupWhereInput   `json:"not,omitempty"`
	Or         []*GroupWhereInput `json:"or,omitempty"`
	And        []*GroupWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "tenant_id" field predicates.
	TenantID      *int  `json:"tenantID,omitempty"`
	TenantIDNEQ   *int  `json:"tenantIDNEQ,omitempty"`
	TenantIDIn    []int `json:"tenantIDIn,omitempty"`
	TenantIDNotIn []int `json:"tenantIDNotIn,omitempty"`

	// "username" field predicates.
	Username             *string  `json:"username,omitempty"`
	UsernameNEQ          *string  `json:"usernameNEQ,omitempty"`
	UsernameIn           []string `json:"usernameIn,omitempty"`
	UsernameNotIn        []string `json:"usernameNotIn,omitempty"`
	UsernameGT           *string  `json:"usernameGT,omitempty"`
	UsernameGTE          *string  `json:"usernameGTE,omitempty"`
	UsernameLT           *string  `json:"usernameLT,omitempty"`
	UsernameLTE          *string  `json:"usernameLTE,omitempty"`
	UsernameContains     *string  `json:"usernameContains,omitempty"`
	UsernameHasPrefix    *string  `json:"usernameHasPrefix,omitempty"`
	UsernameHasSuffix    *string  `json:"usernameHasSuffix,omitempty"`
	UsernameEqualFold    *string  `json:"usernameEqualFold,omitempty"`
	UsernameContainsFold *string  `json:"usernameContainsFold,omitempty"`

	// "tenant" edge predicates.
	HasTenant     *bool               `json:"hasTenant,omitempty"`
	HasTenantWith []*TenantWhereInput `json:"hasTenantWith,omitempty"`

	// "users" edge predicates.
	HasUsers     *bool             `json:"hasUsers,omitempty"`
	HasUsersWith []*UserWhereInput `json:"hasUsersWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *GroupWhereInput) AddPredicates(predicates ...predicate.Group) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the GroupWhereInput filter on the GroupQuery builder.
func (i *GroupWhereInput) Filter(q *GroupQuery) (*GroupQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyGroupWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyGroupWhereInput is returned in case the GroupWhereInput is empty.
var ErrEmptyGroupWhereInput = errors.New("ent: empty predicate GroupWhereInput")

// P returns a predicate for filtering groups.
// An error is returned if the input is empty or invalid.
func (i *GroupWhereInput) P() (predicate.Group, error) {
	var predicates []predicate.Group
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, group.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Group, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, group.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Group, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, group.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, group.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, group.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, group.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, group.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, group.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, group.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, group.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, group.IDLTE(*i.IDLTE))
	}
	if i.TenantID != nil {
		predicates = append(predicates, group.TenantIDEQ(*i.TenantID))
	}
	if i.TenantIDNEQ != nil {
		predicates = append(predicates, group.TenantIDNEQ(*i.TenantIDNEQ))
	}
	if len(i.TenantIDIn) > 0 {
		predicates = append(predicates, group.TenantIDIn(i.TenantIDIn...))
	}
	if len(i.TenantIDNotIn) > 0 {
		predicates = append(predicates, group.TenantIDNotIn(i.TenantIDNotIn...))
	}
	if i.Username != nil {
		predicates = append(predicates, group.UsernameEQ(*i.Username))
	}
	if i.UsernameNEQ != nil {
		predicates = append(predicates, group.UsernameNEQ(*i.UsernameNEQ))
	}
	if len(i.UsernameIn) > 0 {
		predicates = append(predicates, group.UsernameIn(i.UsernameIn...))
	}
	if len(i.UsernameNotIn) > 0 {
		predicates = append(predicates, group.UsernameNotIn(i.UsernameNotIn...))
	}
	if i.UsernameGT != nil {
		predicates = append(predicates, group.UsernameGT(*i.UsernameGT))
	}
	if i.UsernameGTE != nil {
		predicates = append(predicates, group.UsernameGTE(*i.UsernameGTE))
	}
	if i.UsernameLT != nil {
		predicates = append(predicates, group.UsernameLT(*i.UsernameLT))
	}
	if i.UsernameLTE != nil {
		predicates = append(predicates, group.UsernameLTE(*i.UsernameLTE))
	}
	if i.UsernameContains != nil {
		predicates = append(predicates, group.UsernameContains(*i.UsernameContains))
	}
	if i.UsernameHasPrefix != nil {
		predicates = append(predicates, group.UsernameHasPrefix(*i.UsernameHasPrefix))
	}
	if i.UsernameHasSuffix != nil {
		predicates = append(predicates, group.UsernameHasSuffix(*i.UsernameHasSuffix))
	}
	if i.UsernameEqualFold != nil {
		predicates = append(predicates, group.UsernameEqualFold(*i.UsernameEqualFold))
	}
	if i.UsernameContainsFold != nil {
		predicates = append(predicates, group.UsernameContainsFold(*i.UsernameContainsFold))
	}

	if i.HasTenant != nil {
		p := group.HasTenant()
		if !*i.HasTenant {
			p = group.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasTenantWith) > 0 {
		with := make([]predicate.Tenant, 0, len(i.HasTenantWith))
		for _, w := range i.HasTenantWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasTenantWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, group.HasTenantWith(with...))
	}
	if i.HasUsers != nil {
		p := group.HasUsers()
		if !*i.HasUsers {
			p = group.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasUsersWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasUsersWith))
		for _, w := range i.HasUsersWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasUsersWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, group.HasUsersWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyGroupWhereInput
	case 1:
		return predicates[0], nil
	default:
		return group.And(predicates...), nil
	}
}

// TenantWhereInput represents a where input for filtering Tenant queries.
type TenantWhereInput struct {
	Predicates []predicate.Tenant  `json:"-"`
	Not        *TenantWhereInput   `json:"not,omitempty"`
	Or         []*TenantWhereInput `json:"or,omitempty"`
	And        []*TenantWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *TenantWhereInput) AddPredicates(predicates ...predicate.Tenant) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the TenantWhereInput filter on the TenantQuery builder.
func (i *TenantWhereInput) Filter(q *TenantQuery) (*TenantQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyTenantWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyTenantWhereInput is returned in case the TenantWhereInput is empty.
var ErrEmptyTenantWhereInput = errors.New("ent: empty predicate TenantWhereInput")

// P returns a predicate for filtering tenants.
// An error is returned if the input is empty or invalid.
func (i *TenantWhereInput) P() (predicate.Tenant, error) {
	var predicates []predicate.Tenant
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, tenant.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Tenant, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, tenant.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Tenant, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, tenant.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, tenant.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, tenant.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, tenant.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, tenant.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, tenant.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, tenant.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, tenant.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, tenant.IDLTE(*i.IDLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, tenant.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, tenant.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, tenant.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, tenant.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, tenant.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, tenant.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, tenant.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, tenant.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, tenant.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, tenant.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, tenant.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, tenant.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, tenant.NameContainsFold(*i.NameContainsFold))
	}

	switch len(predicates) {
	case 0:
		return nil, ErrEmptyTenantWhereInput
	case 1:
		return predicates[0], nil
	default:
		return tenant.And(predicates...), nil
	}
}

// UserWhereInput represents a where input for filtering User queries.
type UserWhereInput struct {
	Predicates []predicate.User  `json:"-"`
	Not        *UserWhereInput   `json:"not,omitempty"`
	Or         []*UserWhereInput `json:"or,omitempty"`
	And        []*UserWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "tenant_id" field predicates.
	TenantID      *int  `json:"tenantID,omitempty"`
	TenantIDNEQ   *int  `json:"tenantIDNEQ,omitempty"`
	TenantIDIn    []int `json:"tenantIDIn,omitempty"`
	TenantIDNotIn []int `json:"tenantIDNotIn,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameIsNil        bool     `json:"nameIsNil,omitempty"`
	NameNotNil       bool     `json:"nameNotNil,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "username" field predicates.
	Username             *string  `json:"username,omitempty"`
	UsernameNEQ          *string  `json:"usernameNEQ,omitempty"`
	UsernameIn           []string `json:"usernameIn,omitempty"`
	UsernameNotIn        []string `json:"usernameNotIn,omitempty"`
	UsernameGT           *string  `json:"usernameGT,omitempty"`
	UsernameGTE          *string  `json:"usernameGTE,omitempty"`
	UsernameLT           *string  `json:"usernameLT,omitempty"`
	UsernameLTE          *string  `json:"usernameLTE,omitempty"`
	UsernameContains     *string  `json:"usernameContains,omitempty"`
	UsernameHasPrefix    *string  `json:"usernameHasPrefix,omitempty"`
	UsernameHasSuffix    *string  `json:"usernameHasSuffix,omitempty"`
	UsernameEqualFold    *string  `json:"usernameEqualFold,omitempty"`
	UsernameContainsFold *string  `json:"usernameContainsFold,omitempty"`

	// "email" field predicates.
	Email             *string  `json:"email,omitempty"`
	EmailNEQ          *string  `json:"emailNEQ,omitempty"`
	EmailIn           []string `json:"emailIn,omitempty"`
	EmailNotIn        []string `json:"emailNotIn,omitempty"`
	EmailGT           *string  `json:"emailGT,omitempty"`
	EmailGTE          *string  `json:"emailGTE,omitempty"`
	EmailLT           *string  `json:"emailLT,omitempty"`
	EmailLTE          *string  `json:"emailLTE,omitempty"`
	EmailContains     *string  `json:"emailContains,omitempty"`
	EmailHasPrefix    *string  `json:"emailHasPrefix,omitempty"`
	EmailHasSuffix    *string  `json:"emailHasSuffix,omitempty"`
	EmailEqualFold    *string  `json:"emailEqualFold,omitempty"`
	EmailContainsFold *string  `json:"emailContainsFold,omitempty"`

	// "description" field predicates.
	Description             *string  `json:"description,omitempty"`
	DescriptionNEQ          *string  `json:"descriptionNEQ,omitempty"`
	DescriptionIn           []string `json:"descriptionIn,omitempty"`
	DescriptionNotIn        []string `json:"descriptionNotIn,omitempty"`
	DescriptionGT           *string  `json:"descriptionGT,omitempty"`
	DescriptionGTE          *string  `json:"descriptionGTE,omitempty"`
	DescriptionLT           *string  `json:"descriptionLT,omitempty"`
	DescriptionLTE          *string  `json:"descriptionLTE,omitempty"`
	DescriptionContains     *string  `json:"descriptionContains,omitempty"`
	DescriptionHasPrefix    *string  `json:"descriptionHasPrefix,omitempty"`
	DescriptionHasSuffix    *string  `json:"descriptionHasSuffix,omitempty"`
	DescriptionIsNil        bool     `json:"descriptionIsNil,omitempty"`
	DescriptionNotNil       bool     `json:"descriptionNotNil,omitempty"`
	DescriptionEqualFold    *string  `json:"descriptionEqualFold,omitempty"`
	DescriptionContainsFold *string  `json:"descriptionContainsFold,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`

	// "tenant" edge predicates.
	HasTenant     *bool               `json:"hasTenant,omitempty"`
	HasTenantWith []*TenantWhereInput `json:"hasTenantWith,omitempty"`

	// "groups" edge predicates.
	HasGroups     *bool              `json:"hasGroups,omitempty"`
	HasGroupsWith []*GroupWhereInput `json:"hasGroupsWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *UserWhereInput) AddPredicates(predicates ...predicate.User) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the UserWhereInput filter on the UserQuery builder.
func (i *UserWhereInput) Filter(q *UserQuery) (*UserQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyUserWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyUserWhereInput is returned in case the UserWhereInput is empty.
var ErrEmptyUserWhereInput = errors.New("ent: empty predicate UserWhereInput")

// P returns a predicate for filtering users.
// An error is returned if the input is empty or invalid.
func (i *UserWhereInput) P() (predicate.User, error) {
	var predicates []predicate.User
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, user.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.User, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, user.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.User, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, user.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, user.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, user.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, user.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, user.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, user.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, user.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, user.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, user.IDLTE(*i.IDLTE))
	}
	if i.TenantID != nil {
		predicates = append(predicates, user.TenantIDEQ(*i.TenantID))
	}
	if i.TenantIDNEQ != nil {
		predicates = append(predicates, user.TenantIDNEQ(*i.TenantIDNEQ))
	}
	if len(i.TenantIDIn) > 0 {
		predicates = append(predicates, user.TenantIDIn(i.TenantIDIn...))
	}
	if len(i.TenantIDNotIn) > 0 {
		predicates = append(predicates, user.TenantIDNotIn(i.TenantIDNotIn...))
	}
	if i.Name != nil {
		predicates = append(predicates, user.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, user.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, user.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, user.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, user.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, user.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, user.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, user.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, user.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, user.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, user.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameIsNil {
		predicates = append(predicates, user.NameIsNil())
	}
	if i.NameNotNil {
		predicates = append(predicates, user.NameNotNil())
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, user.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, user.NameContainsFold(*i.NameContainsFold))
	}
	if i.Username != nil {
		predicates = append(predicates, user.UsernameEQ(*i.Username))
	}
	if i.UsernameNEQ != nil {
		predicates = append(predicates, user.UsernameNEQ(*i.UsernameNEQ))
	}
	if len(i.UsernameIn) > 0 {
		predicates = append(predicates, user.UsernameIn(i.UsernameIn...))
	}
	if len(i.UsernameNotIn) > 0 {
		predicates = append(predicates, user.UsernameNotIn(i.UsernameNotIn...))
	}
	if i.UsernameGT != nil {
		predicates = append(predicates, user.UsernameGT(*i.UsernameGT))
	}
	if i.UsernameGTE != nil {
		predicates = append(predicates, user.UsernameGTE(*i.UsernameGTE))
	}
	if i.UsernameLT != nil {
		predicates = append(predicates, user.UsernameLT(*i.UsernameLT))
	}
	if i.UsernameLTE != nil {
		predicates = append(predicates, user.UsernameLTE(*i.UsernameLTE))
	}
	if i.UsernameContains != nil {
		predicates = append(predicates, user.UsernameContains(*i.UsernameContains))
	}
	if i.UsernameHasPrefix != nil {
		predicates = append(predicates, user.UsernameHasPrefix(*i.UsernameHasPrefix))
	}
	if i.UsernameHasSuffix != nil {
		predicates = append(predicates, user.UsernameHasSuffix(*i.UsernameHasSuffix))
	}
	if i.UsernameEqualFold != nil {
		predicates = append(predicates, user.UsernameEqualFold(*i.UsernameEqualFold))
	}
	if i.UsernameContainsFold != nil {
		predicates = append(predicates, user.UsernameContainsFold(*i.UsernameContainsFold))
	}
	if i.Email != nil {
		predicates = append(predicates, user.EmailEQ(*i.Email))
	}
	if i.EmailNEQ != nil {
		predicates = append(predicates, user.EmailNEQ(*i.EmailNEQ))
	}
	if len(i.EmailIn) > 0 {
		predicates = append(predicates, user.EmailIn(i.EmailIn...))
	}
	if len(i.EmailNotIn) > 0 {
		predicates = append(predicates, user.EmailNotIn(i.EmailNotIn...))
	}
	if i.EmailGT != nil {
		predicates = append(predicates, user.EmailGT(*i.EmailGT))
	}
	if i.EmailGTE != nil {
		predicates = append(predicates, user.EmailGTE(*i.EmailGTE))
	}
	if i.EmailLT != nil {
		predicates = append(predicates, user.EmailLT(*i.EmailLT))
	}
	if i.EmailLTE != nil {
		predicates = append(predicates, user.EmailLTE(*i.EmailLTE))
	}
	if i.EmailContains != nil {
		predicates = append(predicates, user.EmailContains(*i.EmailContains))
	}
	if i.EmailHasPrefix != nil {
		predicates = append(predicates, user.EmailHasPrefix(*i.EmailHasPrefix))
	}
	if i.EmailHasSuffix != nil {
		predicates = append(predicates, user.EmailHasSuffix(*i.EmailHasSuffix))
	}
	if i.EmailEqualFold != nil {
		predicates = append(predicates, user.EmailEqualFold(*i.EmailEqualFold))
	}
	if i.EmailContainsFold != nil {
		predicates = append(predicates, user.EmailContainsFold(*i.EmailContainsFold))
	}
	if i.Description != nil {
		predicates = append(predicates, user.DescriptionEQ(*i.Description))
	}
	if i.DescriptionNEQ != nil {
		predicates = append(predicates, user.DescriptionNEQ(*i.DescriptionNEQ))
	}
	if len(i.DescriptionIn) > 0 {
		predicates = append(predicates, user.DescriptionIn(i.DescriptionIn...))
	}
	if len(i.DescriptionNotIn) > 0 {
		predicates = append(predicates, user.DescriptionNotIn(i.DescriptionNotIn...))
	}
	if i.DescriptionGT != nil {
		predicates = append(predicates, user.DescriptionGT(*i.DescriptionGT))
	}
	if i.DescriptionGTE != nil {
		predicates = append(predicates, user.DescriptionGTE(*i.DescriptionGTE))
	}
	if i.DescriptionLT != nil {
		predicates = append(predicates, user.DescriptionLT(*i.DescriptionLT))
	}
	if i.DescriptionLTE != nil {
		predicates = append(predicates, user.DescriptionLTE(*i.DescriptionLTE))
	}
	if i.DescriptionContains != nil {
		predicates = append(predicates, user.DescriptionContains(*i.DescriptionContains))
	}
	if i.DescriptionHasPrefix != nil {
		predicates = append(predicates, user.DescriptionHasPrefix(*i.DescriptionHasPrefix))
	}
	if i.DescriptionHasSuffix != nil {
		predicates = append(predicates, user.DescriptionHasSuffix(*i.DescriptionHasSuffix))
	}
	if i.DescriptionIsNil {
		predicates = append(predicates, user.DescriptionIsNil())
	}
	if i.DescriptionNotNil {
		predicates = append(predicates, user.DescriptionNotNil())
	}
	if i.DescriptionEqualFold != nil {
		predicates = append(predicates, user.DescriptionEqualFold(*i.DescriptionEqualFold))
	}
	if i.DescriptionContainsFold != nil {
		predicates = append(predicates, user.DescriptionContainsFold(*i.DescriptionContainsFold))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, user.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, user.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, user.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, user.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, user.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, user.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, user.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, user.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, user.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, user.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, user.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, user.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, user.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, user.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, user.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, user.UpdatedAtLTE(*i.UpdatedAtLTE))
	}

	if i.HasTenant != nil {
		p := user.HasTenant()
		if !*i.HasTenant {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasTenantWith) > 0 {
		with := make([]predicate.Tenant, 0, len(i.HasTenantWith))
		for _, w := range i.HasTenantWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasTenantWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasTenantWith(with...))
	}
	if i.HasGroups != nil {
		p := user.HasGroups()
		if !*i.HasGroups {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasGroupsWith) > 0 {
		with := make([]predicate.Group, 0, len(i.HasGroupsWith))
		for _, w := range i.HasGroupsWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasGroupsWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasGroupsWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyUserWhereInput
	case 1:
		return predicates[0], nil
	default:
		return user.And(predicates...), nil
	}
}