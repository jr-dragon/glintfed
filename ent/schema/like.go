package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Like holds the schema definition for the Like entity.
type Like struct {
	ent.Schema
}

// Fields of the Like.
func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("profile_id"),
		field.Uint64("status_id"),
		field.Uint64("status_profile_id").Optional(),
		field.Bool("is_comment").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Like.
func (Like) Edges() []ent.Edge {
	return nil
}

// Annotations of the Like.
func (Like) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "likes"},
	}
}
