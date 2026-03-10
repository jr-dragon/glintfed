package schema

import "entgo.io/ent"

// Place holds the schema definition for the Place entity.
type Place struct {
	ent.Schema
}

// Fields of the Place.
func (Place) Fields() []ent.Field {
	return nil
}

// Edges of the Place.
func (Place) Edges() []ent.Edge {
	return nil
}
