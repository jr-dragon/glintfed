package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// HashtagRelated holds the schema definition for the HashtagRelated entity.
type HashtagRelated struct {
	ent.Schema
}

// Fields of the HashtagRelated.
func (HashtagRelated) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("hashtag_id").Unique(),
		field.JSON("related_tags", []string{}).Optional(),
		field.Uint64("agg_score").Optional(),
		field.Time("last_calculated_at").Optional(),
		field.Time("last_moderated_at").Optional(),
		field.Bool("skip_refresh").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the HashtagRelated.
func (HashtagRelated) Edges() []ent.Edge {
	return nil
}

// Annotations of the HashtagRelated.
func (HashtagRelated) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "hashtag_related"},
	}
}
