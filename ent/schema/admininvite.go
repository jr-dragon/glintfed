package schema

import "entgo.io/ent"

// AdminInvite holds the schema definition for the AdminInvite entity.
type AdminInvite struct {
	ent.Schema
}

// Fields of the AdminInvite.
func (AdminInvite) Fields() []ent.Field {
	return nil
}

// Edges of the AdminInvite.
func (AdminInvite) Edges() []ent.Edge {
	return nil
}
