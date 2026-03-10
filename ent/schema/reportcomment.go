package schema

import "entgo.io/ent"

// ReportComment holds the schema definition for the ReportComment entity.
type ReportComment struct {
	ent.Schema
}

// Fields of the ReportComment.
func (ReportComment) Fields() []ent.Field {
	return nil
}

// Edges of the ReportComment.
func (ReportComment) Edges() []ent.Edge {
	return nil
}
