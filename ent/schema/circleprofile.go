package schema

import "entgo.io/ent"

// CircleProfile holds the schema definition for the CircleProfile entity.
type CircleProfile struct {
	ent.Schema
}

// Fields of the CircleProfile.
func (CircleProfile) Fields() []ent.Field {
	return nil
}

// Edges of the CircleProfile.
func (CircleProfile) Edges() []ent.Edge {
	return nil
}
