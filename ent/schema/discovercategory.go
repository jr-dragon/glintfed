package schema

import "entgo.io/ent"

// DiscoverCategory holds the schema definition for the DiscoverCategory entity.
type DiscoverCategory struct {
	ent.Schema
}

// Fields of the DiscoverCategory.
func (DiscoverCategory) Fields() []ent.Field {
	return nil
}

// Edges of the DiscoverCategory.
func (DiscoverCategory) Edges() []ent.Edge {
	return nil
}
