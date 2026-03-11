package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ProfileAlias holds the schema definition for the ProfileAlias entity.
type ProfileAlias struct {
	ent.Schema
}

// Fields of the ProfileAlias.
func (ProfileAlias) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("profile_id").Optional(),
		field.String("acct").Optional(),
		field.String("uri").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ProfileAlias.
func (ProfileAlias) Edges() []ent.Edge {
	return nil
}

// Annotations of the ProfileAlias.
func (ProfileAlias) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "profile_aliases"},
	}
}
