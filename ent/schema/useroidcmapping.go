package schema

import "entgo.io/ent"

// UserOidcMapping holds the schema definition for the UserOidcMapping entity.
type UserOidcMapping struct {
	ent.Schema
}

// Fields of the UserOidcMapping.
func (UserOidcMapping) Fields() []ent.Field {
	return nil
}

// Edges of the UserOidcMapping.
func (UserOidcMapping) Edges() []ent.Edge {
	return nil
}
