package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Activity holds the schema definition for the Activity entity.
type Activity struct {
	ent.Schema
}

// Fields of the Activity.
func (Activity) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("to_id").Optional(),
		field.Uint64("from_id").Optional(),
		field.String("object_type").Optional(),
		field.JSON("data", map[string]any{}).Optional(),
		field.Time("processed_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Activity.
func (Activity) Edges() []ent.Edge {
	return nil
}

// Annotations of the Activity.
func (Activity) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "activities"},
	}
}
