package schema

import "entgo.io/ent"

// Bookmark holds the schema definition for the Bookmark entity.
type Bookmark struct {
	ent.Schema
}

// Fields of the Bookmark.
func (Bookmark) Fields() []ent.Field {
	return nil
}

// Edges of the Bookmark.
func (Bookmark) Edges() []ent.Edge {
	return nil
}
