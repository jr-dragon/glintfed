package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// StatusEdit holds the schema definition for the StatusEdit entity.
type StatusEdit struct {
	ent.Schema
}

// Fields of the StatusEdit.
func (StatusEdit) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("status_id"),
		field.Uint64("profile_id"),
		field.Text("caption").Optional(),
		field.Text("spoiler_text").Optional(),
		field.JSON("ordered_media_attachment_ids", []uint64{}).Optional(),
		field.JSON("media_descriptions", map[string]string{}).Optional(),
		field.JSON("poll_options", map[string]any{}).Optional(),
		field.Bool("is_nsfw").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the StatusEdit.
func (StatusEdit) Edges() []ent.Edge {
	return nil
}

// Annotations of the StatusEdit.
func (StatusEdit) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "status_edits"},
	}
}
