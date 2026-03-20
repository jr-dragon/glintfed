package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// CollectionItem holds the schema definition for the CollectionItem entity.
type CollectionItem struct {
	ent.Schema
}

// Fields of the CollectionItem.
func (CollectionItem) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("collection_id"),
		field.Uint32("order").Optional(),
		field.String("object_type").Default("post"),
		field.Uint64("object_id"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CollectionItem.
func (CollectionItem) Edges() []ent.Edge {
	return nil
}

// Annotations of the CollectionItem.
func (CollectionItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "collection_items"},
	}
}
