package schema

import (
	"entgo.io/ent/schema"
	"entgo.io/ent/dialect/entsql"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// FailedJob holds the schema definition for the FailedJob entity.
type FailedJob struct {
	ent.Schema
}

// Fields of the FailedJob.
func (FailedJob) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("uuid").Unique().Optional(),
		field.Text("connection"),
		field.Text("queue"),
		field.Text("payload"),
		field.Text("exception"),
		field.Time("failed_at").Default(time.Now),
	}
}

// Edges of the FailedJob.
func (FailedJob) Edges() []ent.Edge {
	return nil
}

// Annotations of the FailedJob.
func (FailedJob) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "failed_jobs"},
	}
}
