package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CuratedRegisterActivity holds the schema definition for the CuratedRegisterActivity entity.
type CuratedRegisterActivity struct {
	ent.Schema
}

// Fields of the CuratedRegisterActivity.
func (CuratedRegisterActivity) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint32("register_id").Optional(),
		field.Uint32("admin_id").Optional(),
		field.Uint32("reply_to_id").Optional(),
		field.String("secret_code").Optional(),
		field.String("type").Optional(),
		field.String("title").Optional(),
		field.String("link").Optional(),
		field.Text("message").Optional(),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.Bool("from_admin").Default(false),
		field.Bool("from_user").Default(false),
		field.Bool("admin_only_view").Default(true),
		field.Bool("action_required").Default(false),
		field.Time("admin_notified_at").Optional(),
		field.Time("action_taken_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CuratedRegisterActivity.
func (CuratedRegisterActivity) Edges() []ent.Edge {
	return nil
}

// Annotations of the CuratedRegisterActivity.
func (CuratedRegisterActivity) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "curated_register_activities"},
	}
}
