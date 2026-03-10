package schema

import "entgo.io/ent"

// StoryView holds the schema definition for the StoryView entity.
type StoryView struct {
	ent.Schema
}

// Fields of the StoryView.
func (StoryView) Fields() []ent.Field {
	return nil
}

// Edges of the StoryView.
func (StoryView) Edges() []ent.Edge {
	return nil
}
