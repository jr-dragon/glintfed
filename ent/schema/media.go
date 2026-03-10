package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Media holds the schema definition for the Media entity.
type Media struct {
	ent.Schema
}

// Fields of the Media.
func (Media) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("status_id").Optional(),
		field.Uint64("profile_id").Optional(),
		field.Uint64("user_id").Optional(),
		field.Bool("is_nsfw").Default(false),
		field.Bool("remote_media").Default(false),
		field.String("original_sha256").Optional(),
		field.String("optimized_sha256").Optional(),
		field.String("media_path"),
		field.String("thumbnail_path").Optional(),
		field.Text("cdn_url").Optional(),
		field.String("optimized_url").Optional(),
		field.String("thumbnail_url").Optional(),
		field.String("remote_url").Optional(),
		field.Text("caption").Optional(),
		field.String("hls_path").Optional(),
		field.Uint8("order").Default(1),
		field.String("mime").Optional(),
		field.Uint32("size").Optional(),
		field.String("orientation").Optional(),
		field.String("filter_name").Optional(),
		field.String("filter_class").Optional(),
		field.String("license").Optional(),
		field.Time("processed_at").Optional(),
		field.Time("hls_transcoded_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional(),
		field.String("key").Optional(),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.Int8("version").Default(1),
		field.String("blurhash").Optional(),
		field.JSON("srcset", map[string]any{}).Optional(),
		field.Uint32("width").Optional(),
		field.Uint32("height").Optional(),
		field.Bool("skip_optimize").Optional(),
		field.Time("replicated_at").Optional(),
	}
}

// Edges of the Media.
func (Media) Edges() []ent.Edge {
	return nil
}

// Annotations of the Media.
func (Media) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "media"},
	}
}
