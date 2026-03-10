package schema

import "entgo.io/ent"

// RemoteReport holds the schema definition for the RemoteReport entity.
type RemoteReport struct {
	ent.Schema
}

// Fields of the RemoteReport.
func (RemoteReport) Fields() []ent.Field {
	return nil
}

// Edges of the RemoteReport.
func (RemoteReport) Edges() []ent.Edge {
	return nil
}
