package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("profile_id"),
		field.Uint64("actor_id").Optional(),
		field.Uint64("item_id").Optional(),
		field.String("item_type").Optional(),
		field.String("action").Optional(),
		field.Time("read_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return nil
}

// Annotations of the Notification.
func (Notification) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "notifications"},
	}
}
