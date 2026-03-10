package schema

import "entgo.io/ent"

// Mention holds the schema definition for the Mention entity.
type Mention struct {
	ent.Schema
}

// Fields of the Mention.
func (Mention) Fields() []ent.Field {
	return nil
}

// Edges of the Mention.
func (Mention) Edges() []ent.Edge {
	return nil
}
