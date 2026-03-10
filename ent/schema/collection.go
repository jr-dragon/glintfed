package schema

import "entgo.io/ent"

// Collection holds the schema definition for the Collection entity.
type Collection struct {
	ent.Schema
}

// Fields of the Collection.
func (Collection) Fields() []ent.Field {
	return nil
}

// Edges of the Collection.
func (Collection) Edges() []ent.Edge {
	return nil
}
