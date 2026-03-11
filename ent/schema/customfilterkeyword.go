package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CustomFilterKeyword holds the schema definition for the CustomFilterKeyword entity.
type CustomFilterKeyword struct {
	ent.Schema
}

// Fields of the CustomFilterKeyword.
func (CustomFilterKeyword) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("custom_filter_id"),
		field.String("keyword"),
		field.Bool("whole_word").Default(true),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CustomFilterKeyword.
func (CustomFilterKeyword) Edges() []ent.Edge {
	return nil
}

// Annotations of the CustomFilterKeyword.
func (CustomFilterKeyword) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "custom_filter_keywords"},
	}
}
