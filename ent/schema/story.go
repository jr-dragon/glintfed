package schema

import "entgo.io/ent"

// Story holds the schema definition for the Story entity.
type Story struct {
	ent.Schema
}

// Fields of the Story.
func (Story) Fields() []ent.Field {
	return nil
}

// Edges of the Story.
func (Story) Edges() []ent.Edge {
	return nil
}
