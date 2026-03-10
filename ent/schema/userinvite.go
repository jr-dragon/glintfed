package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UserInvite holds the schema definition for the UserInvite entity.
type UserInvite struct {
	ent.Schema
}

// Fields of the UserInvite.
func (UserInvite) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("user_id"),
		field.Uint64("profile_id"),
		field.String("email").Unique(),
		field.Text("message").Optional(),
		field.String("key"),
		field.String("token"),
		field.Time("valid_until").Optional(),
		field.Time("used_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the UserInvite.
func (UserInvite) Edges() []ent.Edge {
	return nil
}

// Annotations of the UserInvite.
func (UserInvite) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_invites"},
	}
}
