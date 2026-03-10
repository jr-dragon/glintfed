package schema

import "entgo.io/ent"

// Poll holds the schema definition for the Poll entity.
type Poll struct {
	ent.Schema
}

// Fields of the Poll.
func (Poll) Fields() []ent.Field {
	return nil
}

// Edges of the Poll.
func (Poll) Edges() []ent.Edge {
	return nil
}
