package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UserRoles holds the schema definition for the UserRoles entity.
type UserRoles struct {
	ent.Schema
}

// Fields of the UserRoles.
func (UserRoles) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("profile_id").Unique().Optional(),
		field.Uint64("user_id").Unique(),
		field.JSON("roles", []string{}).Optional(),
		field.JSON("meta", map[string]any{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the UserRoles.
func (UserRoles) Edges() []ent.Edge {
	return nil
}

// Annotations of the UserRoles.
func (UserRoles) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_roles"},
	}
}
