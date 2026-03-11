package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ImportPost holds the schema definition for the ImportPost entity.
type ImportPost struct {
	ent.Schema
}

// Fields of the ImportPost.
func (ImportPost) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("profile_id"),
		field.Uint64("user_id"),
		field.String("service"),
		field.String("post_hash").Optional(),
		field.String("filename"),
		field.Uint8("media_count"),
		field.String("post_type").Optional(),
		field.Text("caption").Optional(),
		field.JSON("media", map[string]any{}).Optional(),
		field.Uint8("creation_year").Optional(),
		field.Uint8("creation_month").Optional(),
		field.Uint8("creation_day").Optional(),
		field.Uint8("creation_id").Optional(),
		field.Uint64("status_id").Unique().Optional(),
		field.Time("creation_date").Optional(),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.Bool("skip_missing_media").Default(false),
		field.Bool("uploaded_to_s3").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ImportPost.
func (ImportPost) Edges() []ent.Edge {
	return nil
}

// Annotations of the ImportPost.
func (ImportPost) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "import_posts"},
	}
}
