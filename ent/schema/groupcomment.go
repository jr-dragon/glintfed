package schema

import "entgo.io/ent"

// GroupComment holds the schema definition for the GroupComment entity.
type GroupComment struct {
	ent.Schema
}

// Fields of the GroupComment.
func (GroupComment) Fields() []ent.Field {
	return nil
}

// Edges of the GroupComment.
func (GroupComment) Edges() []ent.Edge {
	return nil
}
