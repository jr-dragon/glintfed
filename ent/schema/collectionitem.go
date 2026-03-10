package schema

import "entgo.io/ent"

// CollectionItem holds the schema definition for the CollectionItem entity.
type CollectionItem struct {
	ent.Schema
}

// Fields of the CollectionItem.
func (CollectionItem) Fields() []ent.Field {
	return nil
}

// Edges of the CollectionItem.
func (CollectionItem) Edges() []ent.Edge {
	return nil
}
