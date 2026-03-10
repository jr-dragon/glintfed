package schema

import "entgo.io/ent"

// ImportJob holds the schema definition for the ImportJob entity.
type ImportJob struct {
	ent.Schema
}

// Fields of the ImportJob.
func (ImportJob) Fields() []ent.Field {
	return nil
}

// Edges of the ImportJob.
func (ImportJob) Edges() []ent.Edge {
	return nil
}
