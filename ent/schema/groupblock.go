package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupBlock holds the schema definition for the GroupBlock entity.
type GroupBlock struct {
	ent.Schema
}

// Fields of the GroupBlock.
func (GroupBlock) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("group_id"),
		field.Uint64("admin_id").Optional(),
		field.Uint64("profile_id").Optional(),
		field.Uint64("instance_id").Optional(),
		field.String("name").Optional(),
		field.String("reason").Optional(),
		field.Bool("is_user"),
		field.Bool("moderated").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the GroupBlock.
func (GroupBlock) Edges() []ent.Edge {
	return nil
}

// Annotations of the GroupBlock.
func (GroupBlock) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_blocks"},
	}
}
