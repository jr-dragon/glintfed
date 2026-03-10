package schema

import "entgo.io/ent"

// GroupPost holds the schema definition for the GroupPost entity.
type GroupPost struct {
	ent.Schema
}

// Fields of the GroupPost.
func (GroupPost) Fields() []ent.Field {
	return nil
}

// Edges of the GroupPost.
func (GroupPost) Edges() []ent.Edge {
	return nil
}
