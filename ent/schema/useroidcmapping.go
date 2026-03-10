package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UserOidcMapping holds the schema definition for the UserOidcMapping entity.
type UserOidcMapping struct {
	ent.Schema
}

// Fields of the UserOidcMapping.
func (UserOidcMapping) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("user_id"),
		field.String("oidc_id").Unique(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the UserOidcMapping.
func (UserOidcMapping) Edges() []ent.Edge {
	return nil
}

// Annotations of the UserOidcMapping.
func (UserOidcMapping) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_oidc_mappings"},
	}
}
