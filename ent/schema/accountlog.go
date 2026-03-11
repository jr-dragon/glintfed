package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// AccountLog holds the schema definition for the AccountLog entity.
type AccountLog struct {
	ent.Schema
}

// Fields of the AccountLog.
func (AccountLog) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("user_id"),
		field.Uint64("item_id").Optional(),
		field.String("item_type").Optional(),
		field.String("action").Optional(),
		field.String("message").Optional(),
		field.String("link").Optional(),
		field.String("ip_address").Optional(),
		field.String("user_agent").Optional(),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the AccountLog.
func (AccountLog) Edges() []ent.Edge {
	return nil
}

// Annotations of the AccountLog.
func (AccountLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "account_logs"},
	}
}
