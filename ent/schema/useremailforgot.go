package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UserEmailForgot holds the schema definition for the UserEmailForgot entity.
type UserEmailForgot struct {
	ent.Schema
}

// Fields of the UserEmailForgot.
func (UserEmailForgot) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("user_id"),
		field.String("ip_address").Optional(),
		field.String("user_agent").Optional(),
		field.String("referrer").Optional(),
		field.Time("email_sent_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the UserEmailForgot.
func (UserEmailForgot) Edges() []ent.Edge {
	return nil
}

// Annotations of the UserEmailForgot.
func (UserEmailForgot) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_email_forgots"},
	}
}
