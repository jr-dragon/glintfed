package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupLimit holds the schema definition for the GroupLimit entity.
type GroupLimit struct {
	ent.Schema
}

// Fields of the GroupLimit.
func (GroupLimit) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("group_id"),
		field.Uint64("profile_id"),
		field.JSON("limits", map[string]any{}).Optional(),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the GroupLimit.
func (GroupLimit) Edges() []ent.Edge {
	return nil
}

// Annotations of the GroupLimit.
func (GroupLimit) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_limits"},
	}
}
