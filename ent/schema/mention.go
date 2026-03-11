package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Mention holds the schema definition for the Mention entity.
type Mention struct {
	ent.Schema
}

// Fields of the Mention.
func (Mention) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("status_id"),
		field.Uint64("profile_id"),
		field.Bool("local").Default(true),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Mention.
func (Mention) Edges() []ent.Edge {
	return nil
}

// Annotations of the Mention.
func (Mention) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "mentions"},
	}
}
