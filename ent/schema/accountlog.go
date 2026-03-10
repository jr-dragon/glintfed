package schema

import "entgo.io/ent"

// AccountLog holds the schema definition for the AccountLog entity.
type AccountLog struct {
	ent.Schema
}

// Fields of the AccountLog.
func (AccountLog) Fields() []ent.Field {
	return nil
}

// Edges of the AccountLog.
func (AccountLog) Edges() []ent.Edge {
	return nil
}
