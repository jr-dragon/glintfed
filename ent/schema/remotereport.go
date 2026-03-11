package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// RemoteReport holds the schema definition for the RemoteReport entity.
type RemoteReport struct {
	ent.Schema
}

// Fields of the RemoteReport.
func (RemoteReport) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.JSON("status_ids", []uint64{}).Optional(),
		field.Text("comment").Optional(),
		field.Uint64("account_id").Optional(),
		field.String("uri").Optional(),
		field.Uint32("instance_id").Optional(),
		field.Time("action_taken_at").Optional(),
		field.JSON("report_meta", map[string]any{}).Optional(),
		field.JSON("action_taken_meta", map[string]any{}).Optional(),
		field.Uint64("action_taken_by_account_id").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the RemoteReport.
func (RemoteReport) Edges() []ent.Edge {
	return nil
}

// Annotations of the RemoteReport.
func (RemoteReport) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "remote_reports"},
	}
}
