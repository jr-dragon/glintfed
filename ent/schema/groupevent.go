package schema

import "entgo.io/ent"

// GroupEvent holds the schema definition for the GroupEvent entity.
type GroupEvent struct {
	ent.Schema
}

// Fields of the GroupEvent.
func (GroupEvent) Fields() []ent.Field {
	return nil
}

// Edges of the GroupEvent.
func (GroupEvent) Edges() []ent.Edge {
	return nil
}
