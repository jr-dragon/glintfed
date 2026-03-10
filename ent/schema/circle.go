package schema

import "entgo.io/ent"

// Circle holds the schema definition for the Circle entity.
type Circle struct {
	ent.Schema
}

// Fields of the Circle.
func (Circle) Fields() []ent.Field {
	return nil
}

// Edges of the Circle.
func (Circle) Edges() []ent.Edge {
	return nil
}
