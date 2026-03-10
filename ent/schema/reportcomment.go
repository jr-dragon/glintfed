package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ReportComment holds the schema definition for the ReportComment entity.
type ReportComment struct {
	ent.Schema
}

// Fields of the ReportComment.
func (ReportComment) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("report_id"),
		field.Uint64("profile_id"),
		field.Uint64("user_id"),
		field.Text("comment"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ReportComment.
func (ReportComment) Edges() []ent.Edge {
	return nil
}

// Annotations of the ReportComment.
func (ReportComment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "report_comments"},
	}
}
