// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
)

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Name        *string
	Username    string
	Email       string
	Description *string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	TenantID    int
	GroupIDs    []int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	m.SetUsername(i.Username)
	m.SetEmail(i.Email)
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetTenantID(i.TenantID)
	if v := i.GroupIDs; len(v) > 0 {
		m.AddGroupIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}