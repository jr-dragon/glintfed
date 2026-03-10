package schema

import "entgo.io/ent"

// Portfolio holds the schema definition for the Portfolio entity.
type Portfolio struct {
	ent.Schema
}

// Fields of the Portfolio.
func (Portfolio) Fields() []ent.Field {
	return nil
}

// Edges of the Portfolio.
func (Portfolio) Edges() []ent.Edge {
	return nil
}
