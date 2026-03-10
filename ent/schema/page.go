package schema

import "entgo.io/ent"

// Page holds the schema definition for the Page entity.
type Page struct {
	ent.Schema
}

// Fields of the Page.
func (Page) Fields() []ent.Field {
	return nil
}

// Edges of the Page.
func (Page) Edges() []ent.Edge {
	return nil
}
