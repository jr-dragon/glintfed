package schema

import "entgo.io/ent"

// GroupLike holds the schema definition for the GroupLike entity.
type GroupLike struct {
	ent.Schema
}

// Fields of the GroupLike.
func (GroupLike) Fields() []ent.Field {
	return nil
}

// Edges of the GroupLike.
func (GroupLike) Edges() []ent.Edge {
	return nil
}
