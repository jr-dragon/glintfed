package schema

import "entgo.io/ent"

// GroupBlock holds the schema definition for the GroupBlock entity.
type GroupBlock struct {
	ent.Schema
}

// Fields of the GroupBlock.
func (GroupBlock) Fields() []ent.Field {
	return nil
}

// Edges of the GroupBlock.
func (GroupBlock) Edges() []ent.Edge {
	return nil
}
