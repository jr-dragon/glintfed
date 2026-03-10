package schema

import "entgo.io/ent"

// UserRoles holds the schema definition for the UserRoles entity.
type UserRoles struct {
	ent.Schema
}

// Fields of the UserRoles.
func (UserRoles) Fields() []ent.Field {
	return nil
}

// Edges of the UserRoles.
func (UserRoles) Edges() []ent.Edge {
	return nil
}
