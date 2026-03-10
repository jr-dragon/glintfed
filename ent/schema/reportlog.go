package schema

import "entgo.io/ent"

// ReportLog holds the schema definition for the ReportLog entity.
type ReportLog struct {
	ent.Schema
}

// Fields of the ReportLog.
func (ReportLog) Fields() []ent.Field {
	return nil
}

// Edges of the ReportLog.
func (ReportLog) Edges() []ent.Edge {
	return nil
}
