package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// StoryView holds the schema definition for the StoryView entity.
type StoryView struct {
	ent.Schema
}

// Fields of the StoryView.
func (StoryView) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("story_id"),
		field.Uint64("profile_id"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the StoryView.
func (StoryView) Edges() []ent.Edge {
	return nil
}

// Annotations of the StoryView.
func (StoryView) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "story_views"},
	}
}
