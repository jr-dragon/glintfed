package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupLike holds the schema definition for the GroupLike entity.
type GroupLike struct {
	ent.Schema
}

// Fields of the GroupLike.
func (GroupLike) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("group_id"),
		field.Uint64("profile_id"),
		field.Uint64("status_id").Optional(),
		field.Uint64("comment_id").Optional(),
		field.Bool("local").Default(true),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the GroupLike.
func (GroupLike) Edges() []ent.Edge {
	return nil
}

// Annotations of the GroupLike.
func (GroupLike) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_likes"},
	}
}
