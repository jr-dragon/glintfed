package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupPost holds the schema definition for the GroupPost entity.
type GroupPost struct {
	ent.Schema
}

// Fields of the GroupPost.
func (GroupPost) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("group_id"),
		field.Uint64("profile_id").Optional(),
		field.String("type").Optional(),
		field.String("remote_url").Unique().Optional(),
		field.Uint32("reply_count").Optional(),
		field.String("status").Optional(),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Text("caption").Optional(),
		field.String("visibility").Optional(),
		field.Bool("is_nsfw").Default(false),
		field.Uint32("likes_count").Default(0),
		field.Text("cw_summary").Optional(),
		field.JSON("media_ids", []uint64{}).Optional(),
		field.Bool("comments_disabled").Default(false),
	}
}

// Edges of the GroupPost.
func (GroupPost) Edges() []ent.Edge {
	return nil
}

// Annotations of the GroupPost.
func (GroupPost) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_posts"},
	}
}
