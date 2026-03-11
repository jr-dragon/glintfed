package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CustomEmoji holds the schema definition for the CustomEmoji entity.
type CustomEmoji struct {
	ent.Schema
}

// Fields of the CustomEmoji.
func (CustomEmoji) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("shortcode"),
		field.String("media_path").Optional(),
		field.String("domain").Optional(),
		field.Bool("disabled").Default(false),
		field.String("uri").Optional(),
		field.String("image_remote_url").Optional(),
		field.Uint32("category_id").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CustomEmoji.
func (CustomEmoji) Edges() []ent.Edge {
	return nil
}

// Annotations of the CustomEmoji.
func (CustomEmoji) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "custom_emoji"},
	}
}
