package schema

import "entgo.io/ent"

// ImportData holds the schema definition for the ImportData entity.
type ImportData struct {
	ent.Schema
}

// Fields of the ImportData.
func (ImportData) Fields() []ent.Field {
	return nil
}

// Edges of the ImportData.
func (ImportData) Edges() []ent.Edge {
	return nil
}
