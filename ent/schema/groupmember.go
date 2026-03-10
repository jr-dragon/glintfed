package schema

import "entgo.io/ent"

// GroupMember holds the schema definition for the GroupMember entity.
type GroupMember struct {
	ent.Schema
}

// Fields of the GroupMember.
func (GroupMember) Fields() []ent.Field {
	return nil
}

// Edges of the GroupMember.
func (GroupMember) Edges() []ent.Edge {
	return nil
}
