package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ImportJob holds the schema definition for the ImportJob entity.
type ImportJob struct {
	ent.Schema
}

// Fields of the ImportJob.
func (ImportJob) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("profile_id"),
		field.String("service").Default("instagram"),
		field.String("uuid").Optional(),
		field.String("storage_path").Optional(),
		field.Uint8("stage").Default(0),
		field.Text("media_json").Optional(),
		field.Time("completed_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ImportJob.
func (ImportJob) Edges() []ent.Edge {
	return nil
}

// Annotations of the ImportJob.
func (ImportJob) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "import_jobs"},
	}
}
