package schema

import "entgo.io/ent"

// DirectMessage holds the schema definition for the DirectMessage entity.
type DirectMessage struct {
	ent.Schema
}

// Fields of the DirectMessage.
func (DirectMessage) Fields() []ent.Field {
	return nil
}

// Edges of the DirectMessage.
func (DirectMessage) Edges() []ent.Edge {
	return nil
}
