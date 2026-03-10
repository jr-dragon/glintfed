package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UserPronoun holds the schema definition for the UserPronoun entity.
type UserPronoun struct {
	ent.Schema
}

// Fields of the UserPronoun.
func (UserPronoun) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("user_id").Unique().Optional(),
		field.Uint64("profile_id").Unique(),
		field.JSON("pronouns", []string{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the UserPronoun.
func (UserPronoun) Edges() []ent.Edge {
	return nil
}

// Annotations of the UserPronoun.
func (UserPronoun) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_pronouns"},
	}
}
