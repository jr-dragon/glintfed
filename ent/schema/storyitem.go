package schema

import "entgo.io/ent"

// StoryItem holds the schema definition for the StoryItem entity.
type StoryItem struct {
	ent.Schema
}

// Fields of the StoryItem.
func (StoryItem) Fields() []ent.Field {
	return nil
}

// Edges of the StoryItem.
func (StoryItem) Edges() []ent.Edge {
	return nil
}
