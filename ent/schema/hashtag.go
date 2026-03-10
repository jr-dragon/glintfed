package schema

import "entgo.io/ent"

// Hashtag holds the schema definition for the Hashtag entity.
type Hashtag struct {
	ent.Schema
}

// Fields of the Hashtag.
func (Hashtag) Fields() []ent.Field {
	return nil
}

// Edges of the Hashtag.
func (Hashtag) Edges() []ent.Edge {
	return nil
}
