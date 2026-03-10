package schema

import "entgo.io/ent"

// Avatar holds the schema definition for the Avatar entity.
type Avatar struct {
	ent.Schema
}

// Fields of the Avatar.
func (Avatar) Fields() []ent.Field {
	return nil
}

// Edges of the Avatar.
func (Avatar) Edges() []ent.Edge {
	return nil
}
