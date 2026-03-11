package schema

import "entgo.io/ent"

// StoryReaction holds the schema definition for the StoryReaction entity.
type StoryReaction struct {
	ent.Schema
}

// Fields of the StoryReaction.
func (StoryReaction) Fields() []ent.Field {
	return nil
}

// Edges of the StoryReaction.
func (StoryReaction) Edges() []ent.Edge {
	return nil
}
