package schema

import "entgo.io/ent"

// EmailVerification holds the schema definition for the EmailVerification entity.
type EmailVerification struct {
	ent.Schema
}

// Fields of the EmailVerification.
func (EmailVerification) Fields() []ent.Field {
	return nil
}

// Edges of the EmailVerification.
func (EmailVerification) Edges() []ent.Edge {
	return nil
}
