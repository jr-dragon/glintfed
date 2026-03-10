package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// PollVote holds the schema definition for the PollVote entity.
type PollVote struct {
	ent.Schema
}

// Fields of the PollVote.
func (PollVote) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("story_id").Optional(),
		field.Uint64("status_id").Optional(),
		field.Uint64("profile_id"),
		field.Uint64("poll_id"),
		field.Uint32("choice").Default(0),
		field.String("uri").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the PollVote.
func (PollVote) Edges() []ent.Edge {
	return nil
}

// Annotations of the PollVote.
func (PollVote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "poll_votes"},
	}
}
