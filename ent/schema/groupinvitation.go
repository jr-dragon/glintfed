package schema

import "entgo.io/ent"

// GroupInvitation holds the schema definition for the GroupInvitation entity.
type GroupInvitation struct {
	ent.Schema
}

// Fields of the GroupInvitation.
func (GroupInvitation) Fields() []ent.Field {
	return nil
}

// Edges of the GroupInvitation.
func (GroupInvitation) Edges() []ent.Edge {
	return nil
}
