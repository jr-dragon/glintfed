package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Status holds the schema definition for the Status entity.
type Status struct {
	ent.Schema
}

// Fields of the Status.
func (Status) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("uri").Unique().Optional(),
		field.Text("caption"),
		field.Text("rendered"),
		field.Uint64("profile_id").Optional(),
		field.String("type").Optional(),
		field.Uint64("in_reply_to_id").Optional(),
		field.Uint64("reblog_of_id").Optional(),
		field.String("url").Optional(),
		field.Bool("is_nsfw").Default(false),
		field.String("scope").Default("public"),
		field.Enum("visibility").Values("public", "unlisted", "private", "direct", "draft").Default("public"),
		field.Bool("reply").Default(false),
		field.Uint64("likes_count").Default(0),
		field.Uint64("reblogs_count").Default(0),
		field.String("language").Optional(),
		field.Uint64("conversation_id").Optional(),
		field.Bool("local").Default(true),
		field.Uint64("application_id").Optional(),
		field.Uint64("in_reply_to_profile_id").Optional(),
		field.JSON("entities", map[string]any{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional(),
		field.String("cw_summary").Optional(),
		field.Uint32("reply_count").Optional(),
		field.Bool("comments_disabled").Default(false),
		field.Uint64("place_id").Optional(),
		field.String("object_url").Unique().Optional(),
		field.Time("edited_at").Optional(),
		field.Bool("trendable").Optional(),
		field.JSON("media_ids", []uint64{}).Optional(),
		field.Int8("pinned_order").Optional(),
	}
}

// Edges of the Status.
func (Status) Edges() []ent.Edge {
	return nil
}

// Annotations of the Status.
func (Status) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "statuses"},
	}
}
