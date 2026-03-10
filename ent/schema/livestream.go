package schema

import "entgo.io/ent"

// LiveStream holds the schema definition for the LiveStream entity.
type LiveStream struct {
	ent.Schema
}

// Fields of the LiveStream.
func (LiveStream) Fields() []ent.Field {
	return nil
}

// Edges of the LiveStream.
func (LiveStream) Edges() []ent.Edge {
	return nil
}
