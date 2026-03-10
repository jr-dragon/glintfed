package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupStore holds the schema definition for the GroupStore entity.
type GroupStore struct {
	ent.Schema
}

// Fields of the GroupStore.
func (GroupStore) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("group_id").Optional(),
		field.String("store_key"),
		field.JSON("store_value", map[string]any{}).Optional(),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the GroupStore.
func (GroupStore) Edges() []ent.Edge {
	return nil
}

// Annotations of the GroupStore.
func (GroupStore) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_stores"},
	}
}
