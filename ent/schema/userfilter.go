package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UserFilter holds the schema definition for the UserFilter entity.
type UserFilter struct {
	ent.Schema
}

// Fields of the UserFilter.
func (UserFilter) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("user_id"),
		field.Uint64("filterable_id"),
		field.String("filterable_type"),
		field.String("filter_type").Default("block"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the UserFilter.
func (UserFilter) Edges() []ent.Edge {
	return nil
}

// Annotations of the UserFilter.
func (UserFilter) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_filters"},
	}
}
