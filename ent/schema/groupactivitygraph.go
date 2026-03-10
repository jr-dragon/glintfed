package schema

import "entgo.io/ent"

// GroupActivityGraph holds the schema definition for the GroupActivityGraph entity.
type GroupActivityGraph struct {
	ent.Schema
}

// Fields of the GroupActivityGraph.
func (GroupActivityGraph) Fields() []ent.Field {
	return nil
}

// Edges of the GroupActivityGraph.
func (GroupActivityGraph) Edges() []ent.Edge {
	return nil
}
