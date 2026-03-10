package schema

import "entgo.io/ent"

// UserDomainBlock holds the schema definition for the UserDomainBlock entity.
type UserDomainBlock struct {
	ent.Schema
}

// Fields of the UserDomainBlock.
func (UserDomainBlock) Fields() []ent.Field {
	return nil
}

// Edges of the UserDomainBlock.
func (UserDomainBlock) Edges() []ent.Edge {
	return nil
}
