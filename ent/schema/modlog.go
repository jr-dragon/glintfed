package schema

import "entgo.io/ent"

// ModLog holds the schema definition for the ModLog entity.
type ModLog struct {
	ent.Schema
}

// Fields of the ModLog.
func (ModLog) Fields() []ent.Field {
	return nil
}

// Edges of the ModLog.
func (ModLog) Edges() []ent.Edge {
	return nil
}
