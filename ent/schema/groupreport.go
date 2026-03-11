package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupReport holds the schema definition for the GroupReport entity.
type GroupReport struct {
	ent.Schema
}

// Fields of the GroupReport.
func (GroupReport) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("group_id"),
		field.Uint64("profile_id"),
		field.String("type").Optional(),
		field.String("item_type").Optional(),
		field.String("item_id").Optional(),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.Bool("open").Default(true),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the GroupReport.
func (GroupReport) Edges() []ent.Edge {
	return nil
}

// Annotations of the GroupReport.
func (GroupReport) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_reports"},
	}
}
