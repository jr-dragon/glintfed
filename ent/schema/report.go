package schema

import "entgo.io/ent"

// Report holds the schema definition for the Report entity.
type Report struct {
	ent.Schema
}

// Fields of the Report.
func (Report) Fields() []ent.Field {
	return nil
}

// Edges of the Report.
func (Report) Edges() []ent.Edge {
	return nil
}
