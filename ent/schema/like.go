package schema

import "entgo.io/ent"

// Like holds the schema definition for the Like entity.
type Like struct {
	ent.Schema
}

// Fields of the Like.
func (Like) Fields() []ent.Field {
	return nil
}

// Edges of the Like.
func (Like) Edges() []ent.Edge {
	return nil
}
