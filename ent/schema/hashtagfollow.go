package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// HashtagFollow holds the schema definition for the HashtagFollow entity.
type HashtagFollow struct {
	ent.Schema
}

// Fields of the HashtagFollow.
func (HashtagFollow) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("user_id"),
		field.Uint64("profile_id"),
		field.Uint64("hashtag_id"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the HashtagFollow.
func (HashtagFollow) Edges() []ent.Edge {
	return nil
}

// Annotations of the HashtagFollow.
func (HashtagFollow) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "hashtag_follows"},
	}
}
