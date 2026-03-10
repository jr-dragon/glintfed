package schema

import "entgo.io/ent"

// PollVote holds the schema definition for the PollVote entity.
type PollVote struct {
	ent.Schema
}

// Fields of the PollVote.
func (PollVote) Fields() []ent.Field {
	return nil
}

// Edges of the PollVote.
func (PollVote) Edges() []ent.Edge {
	return nil
}
