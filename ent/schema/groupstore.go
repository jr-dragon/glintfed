package schema

import "entgo.io/ent"

// GroupStore holds the schema definition for the GroupStore entity.
type GroupStore struct {
	ent.Schema
}

// Fields of the GroupStore.
func (GroupStore) Fields() []ent.Field {
	return nil
}

// Edges of the GroupStore.
func (GroupStore) Edges() []ent.Edge {
	return nil
}
