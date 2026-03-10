package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Poll holds the schema definition for the Poll entity.
type Poll struct {
	ent.Schema
}

// Fields of the Poll.
func (Poll) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("story_id").Optional(),
		field.Uint64("status_id").Optional(),
		field.Uint64("group_id").Optional(),
		field.Uint64("profile_id"),
		field.JSON("poll_options", map[string]any{}).Optional(),
		field.JSON("cached_tallies", map[string]any{}).Optional(),
		field.Bool("multiple").Default(false),
		field.Bool("hide_totals").Default(false),
		field.Uint32("votes_count").Default(0),
		field.Time("last_fetched_at").Optional(),
		field.Time("expires_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Poll.
func (Poll) Edges() []ent.Edge {
	return nil
}

// Annotations of the Poll.
func (Poll) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "polls"},
	}
}
