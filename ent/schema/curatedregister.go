package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CuratedRegister holds the schema definition for the CuratedRegister entity.
type CuratedRegister struct {
	ent.Schema
}

// Fields of the CuratedRegister.
func (CuratedRegister) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("email").Unique().Optional(),
		field.String("username").Unique().Optional(),
		field.String("password").Sensitive().Optional(),
		field.String("ip_address").Optional(),
		field.String("verify_code").Optional(),
		field.Text("reason_to_join").Optional(),
		field.Uint64("invited_by").Optional(),
		field.Bool("is_approved").Default(false),
		field.Bool("is_rejected").Default(false),
		field.Bool("is_awaiting_more_info").Default(false),
		field.Bool("user_has_responded").Default(false),
		field.Bool("is_closed").Default(false),
		field.JSON("autofollow_account_ids", []uint64{}).Optional(),
		field.JSON("admin_notes", []any{}).Optional(),
		field.Uint32("approved_by_admin_id").Optional(),
		field.Time("email_verified_at").Optional(),
		field.Time("admin_notified_at").Optional(),
		field.Time("action_taken_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CuratedRegister.
func (CuratedRegister) Edges() []ent.Edge {
	return nil
}

// Annotations of the CuratedRegister.
func (CuratedRegister) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "curated_registers"},
	}
}
