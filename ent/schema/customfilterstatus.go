package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CustomFilterStatus holds the schema definition for the CustomFilterStatus entity.
type CustomFilterStatus struct {
	ent.Schema
}

// Fields of the CustomFilterStatus.
func (CustomFilterStatus) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("custom_filter_id"),
		field.Uint64("status_id"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CustomFilterStatus.
func (CustomFilterStatus) Edges() []ent.Edge {
	return nil
}

// Annotations of the CustomFilterStatus.
func (CustomFilterStatus) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "custom_filter_statuses"},
	}
}
