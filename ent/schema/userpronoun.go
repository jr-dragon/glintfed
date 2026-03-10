package schema

import "entgo.io/ent"

// UserPronoun holds the schema definition for the UserPronoun entity.
type UserPronoun struct {
	ent.Schema
}

// Fields of the UserPronoun.
func (UserPronoun) Fields() []ent.Field {
	return nil
}

// Edges of the UserPronoun.
func (UserPronoun) Edges() []ent.Edge {
	return nil
}
