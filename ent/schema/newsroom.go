package schema

import "entgo.io/ent"

// Newsroom holds the schema definition for the Newsroom entity.
type Newsroom struct {
	ent.Schema
}

// Fields of the Newsroom.
func (Newsroom) Fields() []ent.Field {
	return nil
}

// Edges of the Newsroom.
func (Newsroom) Edges() []ent.Edge {
	return nil
}
