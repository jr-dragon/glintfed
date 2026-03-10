package schema

import "entgo.io/ent"

// StatusView holds the schema definition for the StatusView entity.
type StatusView struct {
	ent.Schema
}

// Fields of the StatusView.
func (StatusView) Fields() []ent.Field {
	return nil
}

// Edges of the StatusView.
func (StatusView) Edges() []ent.Edge {
	return nil
}
