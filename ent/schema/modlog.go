package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ModLog holds the schema definition for the ModLog entity.
type ModLog struct {
	ent.Schema
}

// Fields of the ModLog.
func (ModLog) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("user_id"),
		field.String("user_username").Optional(),
		field.Uint64("object_uid").Optional(),
		field.Uint64("object_id").Optional(),
		field.String("object_type").Optional(),
		field.String("action").Optional(),
		field.Text("message").Optional(),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.String("access_level").Default("admin").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ModLog.
func (ModLog) Edges() []ent.Edge {
	return nil
}

// Annotations of the ModLog.
func (ModLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "mod_logs"},
	}
}
