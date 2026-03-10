package schema

import (
	"entgo.io/ent/schema"
	"entgo.io/ent/dialect/entsql"
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// DiscoverCategoryHashtag holds the schema definition for the DiscoverCategoryHashtag entity.
type DiscoverCategoryHashtag struct {
	ent.Schema
}

// Fields of the DiscoverCategoryHashtag.
func (DiscoverCategoryHashtag) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("discover_category_id"),
		field.Uint64("hashtag_id"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the DiscoverCategoryHashtag.
func (DiscoverCategoryHashtag) Edges() []ent.Edge {
	return nil
}

// Annotations of the DiscoverCategoryHashtag.
func (DiscoverCategoryHashtag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "discover_category_hashtags"},
	}
}
