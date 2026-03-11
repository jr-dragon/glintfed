package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Follower holds the schema definition for the Follower entity.
type Follower struct {
	ent.Schema
}

// Fields of the Follower.
func (Follower) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("profile_id"),
		field.Uint64("following_id"),
		field.Bool("local_profile").Default(true),
		field.Bool("local_following").Default(true),
		field.Bool("show_reblogs").Default(true),
		field.Bool("notify").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Follower.
func (Follower) Edges() []ent.Edge {
	return nil
}

// Annotations of the Follower.
func (Follower) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "followers"},
	}
}
