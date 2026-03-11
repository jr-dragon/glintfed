package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// MediaTag holds the schema definition for the MediaTag entity.
type MediaTag struct {
	ent.Schema
}

// Fields of the MediaTag.
func (MediaTag) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("status_id").Optional(),
		field.Uint64("media_id"),
		field.Uint64("profile_id"),
		field.String("tagged_username").Optional(),
		field.Bool("is_public").Default(true),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the MediaTag.
func (MediaTag) Edges() []ent.Edge {
	return nil
}

// Annotations of the MediaTag.
func (MediaTag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "media_tags"},
	}
}
