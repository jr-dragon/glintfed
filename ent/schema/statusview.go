package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// StatusView holds the schema definition for the StatusView entity.
type StatusView struct {
	ent.Schema
}

// Fields of the StatusView.
func (StatusView) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("status_id").Optional(),
		field.Uint64("status_profile_id").Optional(),
		field.Uint64("profile_id").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the StatusView.
func (StatusView) Edges() []ent.Edge {
	return nil
}

// Annotations of the StatusView.
func (StatusView) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "status_views"},
	}
}
