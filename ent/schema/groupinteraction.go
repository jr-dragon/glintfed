package schema

import "entgo.io/ent"

// GroupInteraction holds the schema definition for the GroupInteraction entity.
type GroupInteraction struct {
	ent.Schema
}

// Fields of the GroupInteraction.
func (GroupInteraction) Fields() []ent.Field {
	return nil
}

// Edges of the GroupInteraction.
func (GroupInteraction) Edges() []ent.Edge {
	return nil
}
