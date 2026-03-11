package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Conversation holds the schema definition for the Conversation entity.
type Conversation struct {
	ent.Schema
}

// Fields of the Conversation.
func (Conversation) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("to_id"),
		field.Uint64("from_id"),
		field.Uint64("dm_id").Optional(),
		field.Uint64("status_id").Optional(),
		field.String("type").Optional(),
		field.Bool("is_hidden").Default(false),
		field.Bool("has_seen").Default(false),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Conversation.
func (Conversation) Edges() []ent.Edge {
	return nil
}

// Annotations of the Conversation.
func (Conversation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "conversations"},
	}
}
