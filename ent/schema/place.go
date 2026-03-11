package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Place holds the schema definition for the Place entity.
type Place struct {
	ent.Schema
}

// Fields of the Place.
func (Place) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("slug"),
		field.String("name"),
		field.String("state").Optional(),
		field.String("country"),
		field.JSON("aliases", map[string]any{}).Optional(),
		field.Float("lat").Optional(),
		field.Float("long").Optional(),
		field.Int8("score").Default(0),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Uint64("cached_post_count").Optional(),
		field.Time("last_checked_at").Optional(),
	}
}

// Edges of the Place.
func (Place) Edges() []ent.Edge {
	return nil
}

// Annotations of the Place.
func (Place) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "places"},
	}
}
