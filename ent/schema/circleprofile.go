package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CircleProfile holds the schema definition for the CircleProfile entity.
type CircleProfile struct {
	ent.Schema
}

// Fields of the CircleProfile.
func (CircleProfile) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("owner_id").Optional(),
		field.Uint64("circle_id"),
		field.Uint64("profile_id"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CircleProfile.
func (CircleProfile) Edges() []ent.Edge {
	return nil
}

// Annotations of the CircleProfile.
func (CircleProfile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "circle_profiles"},
	}
}
