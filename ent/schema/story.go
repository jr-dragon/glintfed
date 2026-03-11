package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Story holds the schema definition for the Story entity.
type Story struct {
	ent.Schema
}

// Fields of the Story.
func (Story) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("profile_id"),
		field.String("type").Optional(),
		field.Uint32("size").Optional(),
		field.String("mime").Optional(),
		field.Uint16("duration"),
		field.String("path").Optional(),
		field.String("remote_url").Unique().Optional(),
		field.String("media_url").Unique().Optional(),
		field.String("cdn_url").Optional(),
		field.Bool("public").Default(false),
		field.Bool("local").Default(false),
		field.Uint32("view_count").Default(0),
		field.Uint32("comment_count").Optional(),
		field.JSON("story", map[string]any{}).Optional(),
		field.Time("expires_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Bool("is_archived").Default(false).Optional(),
		field.String("name").Optional(),
		field.Bool("active").Optional(),
		field.Bool("can_reply").Default(true),
		field.Bool("can_react").Default(true),
		field.String("object_id").Unique().Optional(),
		field.String("object_uri").Unique().Optional(),
		field.String("bearcap_token").Optional(),
	}
}

// Edges of the Story.
func (Story) Edges() []ent.Edge {
	return nil
}

// Annotations of the Story.
func (Story) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "stories"},
	}
}
