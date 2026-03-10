package schema

import "entgo.io/ent"

// FailedJob holds the schema definition for the FailedJob entity.
type FailedJob struct {
	ent.Schema
}

// Fields of the FailedJob.
func (FailedJob) Fields() []ent.Field {
	return nil
}

// Edges of the FailedJob.
func (FailedJob) Edges() []ent.Edge {
	return nil
}
