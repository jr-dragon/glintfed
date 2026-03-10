package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// EmailVerification holds the schema definition for the EmailVerification entity.
type EmailVerification struct {
	ent.Schema
}

// Fields of the EmailVerification.
func (EmailVerification) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("user_id"),
		field.String("email").Optional(),
		field.String("user_token"),
		field.String("random_token"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the EmailVerification.
func (EmailVerification) Edges() []ent.Edge {
	return nil
}

// Annotations of the EmailVerification.
func (EmailVerification) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "email_verifications"},
	}
}
