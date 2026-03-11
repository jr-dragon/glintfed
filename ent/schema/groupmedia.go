package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupMedia holds the schema definition for the GroupMedia entity.
type GroupMedia struct {
	ent.Schema
}

// Fields of the GroupMedia.
func (GroupMedia) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("group_id"),
		field.Uint64("profile_id"),
		field.Uint64("status_id").Optional(),
		field.String("media_path").Unique(),
		field.Text("thumbnail_url").Optional(),
		field.Text("cdn_url").Optional(),
		field.Text("url").Optional(),
		field.String("mime").Optional(),
		field.Uint32("size").Optional(),
		field.Text("cw_summary").Optional(),
		field.String("license").Optional(),
		field.String("blurhash").Optional(),
		field.Uint8("order").Default(1),
		field.Uint32("width").Optional(),
		field.Uint32("height").Optional(),
		field.Bool("local_user").Default(true),
		field.Bool("is_cached").Default(false),
		field.Bool("is_comment").Default(false),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.String("version").Default("1"),
		field.Bool("skip_optimize").Default(false),
		field.Time("processed_at").Optional(),
		field.Time("thumbnail_generated").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the GroupMedia.
func (GroupMedia) Edges() []ent.Edge {
	return nil
}

// Annotations of the GroupMedia.
func (GroupMedia) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_media"},
	}
}
