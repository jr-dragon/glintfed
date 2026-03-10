package schema

import "entgo.io/ent"

// Follower holds the schema definition for the Follower entity.
type Follower struct {
	ent.Schema
}

// Fields of the Follower.
func (Follower) Fields() []ent.Field {
	return nil
}

// Edges of the Follower.
func (Follower) Edges() []ent.Edge {
	return nil
}
