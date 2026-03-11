package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UserAppSettings holds the schema definition for the UserAppSettings entity.
type UserAppSettings struct {
	ent.Schema
}

// Fields of the UserAppSettings.
func (UserAppSettings) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("user_id").Unique(),
		field.Uint64("profile_id").Unique(),
		field.JSON("common", map[string]any{}).Optional(),
		field.JSON("custom", map[string]any{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the UserAppSettings.
func (UserAppSettings) Edges() []ent.Edge {
	return nil
}

// Annotations of the UserAppSettings.
func (UserAppSettings) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_app_settings"},
	}
}
