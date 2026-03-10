package schema

import "entgo.io/ent"

// Status holds the schema definition for the Status entity.
type Status struct {
	ent.Schema
}

// Fields of the Status.
func (Status) Fields() []ent.Field {
	return nil
}

// Edges of the Status.
func (Status) Edges() []ent.Edge {
	return nil
}
