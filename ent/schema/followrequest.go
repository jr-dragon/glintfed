package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// FollowRequest holds the schema definition for the FollowRequest entity.
type FollowRequest struct {
	ent.Schema
}

// Fields of the FollowRequest.
func (FollowRequest) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("follower_id"),
		field.Uint64("following_id"),
		field.JSON("activity", map[string]any{}).Optional(),
		field.Bool("is_rejected").Default(false),
		field.Bool("is_local").Default(false),
		field.Time("handled_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the FollowRequest.
func (FollowRequest) Edges() []ent.Edge {
	return nil
}

// Annotations of the FollowRequest.
func (FollowRequest) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "follow_requests"},
	}
}
