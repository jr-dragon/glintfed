package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CustomFilter holds the schema definition for the CustomFilter entity.
type CustomFilter struct {
	ent.Schema
}

// Fields of the CustomFilter.
func (CustomFilter) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("profile_id"),
		field.Text("phrase"),
		field.Int("action").Default(0),
		field.JSON("context", []string{}).Optional(),
		field.Time("expires_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CustomFilter.
func (CustomFilter) Edges() []ent.Edge {
	return nil
}

// Annotations of the CustomFilter.
func (CustomFilter) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "custom_filters"},
	}
}
