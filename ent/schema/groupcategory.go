package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupCategory holds the schema definition for the GroupCategory entity.
type GroupCategory struct {
	ent.Schema
}

// Fields of the GroupCategory.
func (GroupCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("name").Unique(),
		field.String("slug").Unique(),
		field.Bool("active").Default(true),
		field.Uint8("order").Optional(),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the GroupCategory.
func (GroupCategory) Edges() []ent.Edge {
	return nil
}

// Annotations of the GroupCategory.
func (GroupCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_categories"},
	}
}
