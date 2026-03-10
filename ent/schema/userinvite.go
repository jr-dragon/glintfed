package schema

import "entgo.io/ent"

// UserInvite holds the schema definition for the UserInvite entity.
type UserInvite struct {
	ent.Schema
}

// Fields of the UserInvite.
func (UserInvite) Fields() []ent.Field {
	return nil
}

// Edges of the UserInvite.
func (UserInvite) Edges() []ent.Edge {
	return nil
}
