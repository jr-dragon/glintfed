package schema

import "entgo.io/ent"

// ProfileAlias holds the schema definition for the ProfileAlias entity.
type ProfileAlias struct {
	ent.Schema
}

// Fields of the ProfileAlias.
func (ProfileAlias) Fields() []ent.Field {
	return nil
}

// Edges of the ProfileAlias.
func (ProfileAlias) Edges() []ent.Edge {
	return nil
}
