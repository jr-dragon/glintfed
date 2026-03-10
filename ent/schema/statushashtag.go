package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// StatusHashtag holds the schema definition for the StatusHashtag entity.
type StatusHashtag struct {
	ent.Schema
}

// Fields of the StatusHashtag.
func (StatusHashtag) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("status_id"),
		field.Uint64("hashtag_id"),
		field.Uint64("profile_id").Optional(),
		field.String("status_visibility").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the StatusHashtag.
func (StatusHashtag) Edges() []ent.Edge {
	return nil
}

// Annotations of the StatusHashtag.
func (StatusHashtag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "status_hashtags"},
	}
}
