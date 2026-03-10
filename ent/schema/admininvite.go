package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// AdminInvite holds the schema definition for the AdminInvite entity.
type AdminInvite struct {
	ent.Schema
}

// Fields of the AdminInvite.
func (AdminInvite) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("name").Optional(),
		field.String("invite_code").Unique(),
		field.Text("description").Optional(),
		field.Text("message").Optional(),
		field.Uint32("max_uses").Optional(),
		field.Uint32("uses").Default(0),
		field.Bool("skip_email_verification").Default(false),
		field.Time("expires_at").Optional(),
		field.JSON("used_by", []any{}).Optional(),
		field.Uint32("admin_user_id").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the AdminInvite.
func (AdminInvite) Edges() []ent.Edge {
	return nil
}

// Annotations of the AdminInvite.
func (AdminInvite) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "admin_invites"},
	}
}
