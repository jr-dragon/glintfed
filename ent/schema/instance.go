package schema

import "entgo.io/ent"

// Instance holds the schema definition for the Instance entity.
type Instance struct {
	ent.Schema
}

// Fields of the Instance.
func (Instance) Fields() []ent.Field {
	return nil
}

// Edges of the Instance.
func (Instance) Edges() []ent.Edge {
	return nil
}
