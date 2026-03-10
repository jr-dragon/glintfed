package schema

import "entgo.io/ent"

// CustomEmoji holds the schema definition for the CustomEmoji entity.
type CustomEmoji struct {
	ent.Schema
}

// Fields of the CustomEmoji.
func (CustomEmoji) Fields() []ent.Field {
	return nil
}

// Edges of the CustomEmoji.
func (CustomEmoji) Edges() []ent.Edge {
	return nil
}
