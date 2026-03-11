package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupInteraction holds the schema definition for the GroupInteraction entity.
type GroupInteraction struct {
	ent.Schema
}

// Fields of the GroupInteraction.
func (GroupInteraction) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("group_id"),
		field.Uint64("profile_id"),
		field.String("type").Optional(),
		field.String("item_type").Optional(),
		field.String("item_id").Optional(),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the GroupInteraction.
func (GroupInteraction) Edges() []ent.Edge {
	return nil
}

// Annotations of the GroupInteraction.
func (GroupInteraction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_interactions"},
	}
}
