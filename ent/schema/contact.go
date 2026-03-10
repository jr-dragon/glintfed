package schema

import "entgo.io/ent"

// Contact holds the schema definition for the Contact entity.
type Contact struct {
	ent.Schema
}

// Fields of the Contact.
func (Contact) Fields() []ent.Field {
	return nil
}

// Edges of the Contact.
func (Contact) Edges() []ent.Edge {
	return nil
}
