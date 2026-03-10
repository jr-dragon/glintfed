package schema

import "entgo.io/ent"

// GroupCategory holds the schema definition for the GroupCategory entity.
type GroupCategory struct {
	ent.Schema
}

// Fields of the GroupCategory.
func (GroupCategory) Fields() []ent.Field {
	return nil
}

// Edges of the GroupCategory.
func (GroupCategory) Edges() []ent.Edge {
	return nil
}
