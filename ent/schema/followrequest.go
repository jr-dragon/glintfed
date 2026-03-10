package schema

import "entgo.io/ent"

// FollowRequest holds the schema definition for the FollowRequest entity.
type FollowRequest struct {
	ent.Schema
}

// Fields of the FollowRequest.
func (FollowRequest) Fields() []ent.Field {
	return nil
}

// Edges of the FollowRequest.
func (FollowRequest) Edges() []ent.Edge {
	return nil
}
