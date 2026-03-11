package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupComment holds the schema definition for the GroupComment entity.
type GroupComment struct {
	ent.Schema
}

// Fields of the GroupComment.
func (GroupComment) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("group_id"),
		field.Uint64("profile_id").Optional(),
		field.Uint64("status_id").Optional(),
		field.Uint64("in_reply_to_id").Optional(),
		field.String("remote_url").Unique().Optional(),
		field.Text("caption").Optional(),
		field.Bool("is_nsfw").Default(false),
		field.String("visibility").Optional(),
		field.Uint32("likes_count").Default(0),
		field.Uint32("replies_count").Default(0),
		field.Text("cw_summary").Optional(),
		field.JSON("media_ids", []uint64{}).Optional(),
		field.String("status").Optional(),
		field.String("type").Default("text").Optional(),
		field.Bool("local").Default(false),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the GroupComment.
func (GroupComment) Edges() []ent.Edge {
	return nil
}

// Annotations of the GroupComment.
func (GroupComment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_comments"},
	}
}
