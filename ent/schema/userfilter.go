package schema

import "entgo.io/ent"

// UserFilter holds the schema definition for the UserFilter entity.
type UserFilter struct {
	ent.Schema
}

// Fields of the UserFilter.
func (UserFilter) Fields() []ent.Field {
	return nil
}

// Edges of the UserFilter.
func (UserFilter) Edges() []ent.Edge {
	return nil
}
