package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupRole holds the schema definition for the GroupRole entity.
type GroupRole struct {
	ent.Schema
}

// Fields of the GroupRole.
func (GroupRole) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("group_id"),
		field.String("name"),
		field.String("slug").Optional(),
		field.Text("abilities").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the GroupRole.
func (GroupRole) Edges() []ent.Edge {
	return nil
}

// Annotations of the GroupRole.
func (GroupRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_roles"},
	}
}
