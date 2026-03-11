package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// DirectMessage holds the schema definition for the DirectMessage entity.
type DirectMessage struct {
	ent.Schema
}

// Fields of the DirectMessage.
func (DirectMessage) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("to_id"),
		field.Uint64("from_id"),
		field.String("type").Default("text").Optional(),
		field.String("from_profile_ids").Optional(),
		field.Bool("group_message").Default(false),
		field.Bool("is_hidden").Default(false),
		field.JSON("meta", map[string]any{}).Optional(),
		field.Uint64("status_id"),
		field.Time("read_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the DirectMessage.
func (DirectMessage) Edges() []ent.Edge {
	return nil
}

// Annotations of the DirectMessage.
func (DirectMessage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "direct_messages"},
	}
}
