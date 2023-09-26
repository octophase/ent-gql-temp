package schema

import (
	"myapp/ent/privacy"
	"myapp/rule"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition of the Group entity.
type Group struct {
	ent.Schema
}

// Mixin of the User schema.
func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the User
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Default("unknown"),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("groups"),
	}
}

// Policy defines the privacy policy of the Group.
func (Group) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			// Limit DenyMismatchedTenants only for
			// Create operations
			privacy.OnMutationOperation(
				rule.DenyMismatchedTenants(),
				ent.OpCreate,
			),
		},
	}
}
