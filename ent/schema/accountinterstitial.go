package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// AccountInterstitial holds the schema definition for the AccountInterstitial entity.
type AccountInterstitial struct {
	ent.Schema
}

// Fields of the AccountInterstitial.
func (AccountInterstitial) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("user_id").Optional(),
		field.String("type").Optional(),
		field.String("view").Optional(),
		field.Uint64("item_id").Optional(),
		field.String("item_type").Optional(),
		field.Bool("is_spam").Optional(),
		field.Bool("in_violation").Optional(),
		field.Uint32("violation_id").Optional(),
		field.Bool("email_notify").Optional(),
		field.Bool("has_media").Default(false).Optional(),
		field.String("blurhash").Optional(),
		field.Text("message").Optional(),
		field.Text("violation_header").Optional(),
		field.Text("violation_body").Optional(),
		field.JSON("meta", map[string]any{}).Optional(),
		field.Text("appeal_message").Optional(),
		field.Time("appeal_requested_at").Optional(),
		field.Time("appeal_handled_at").Optional(),
		field.Time("read_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Uint8("severity_index").Optional(),
		field.Uint64("thread_id").Unique().Optional(),
		field.Time("emailed_at").Optional(),
	}
}

// Edges of the AccountInterstitial.
func (AccountInterstitial) Edges() []ent.Edge {
	return nil
}

// Annotations of the AccountInterstitial.
func (AccountInterstitial) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "account_interstitials"},
	}
}
