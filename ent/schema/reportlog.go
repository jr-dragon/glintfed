package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ReportLog holds the schema definition for the ReportLog entity.
type ReportLog struct {
	ent.Schema
}

// Fields of the ReportLog.
func (ReportLog) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("profile_id"),
		field.Uint64("item_id").Optional(),
		field.String("item_type").Optional(),
		field.String("action").Optional(),
		field.Bool("system_message").Default(false),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ReportLog.
func (ReportLog) Edges() []ent.Edge {
	return nil
}

// Annotations of the ReportLog.
func (ReportLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "report_logs"},
	}
}
