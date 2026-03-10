package schema

import "entgo.io/ent"

// MediaTag holds the schema definition for the MediaTag entity.
type MediaTag struct {
	ent.Schema
}

// Fields of the MediaTag.
func (MediaTag) Fields() []ent.Field {
	return nil
}

// Edges of the MediaTag.
func (MediaTag) Edges() []ent.Edge {
	return nil
}
