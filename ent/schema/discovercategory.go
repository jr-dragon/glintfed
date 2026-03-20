package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// DiscoverCategory holds the schema definition for the DiscoverCategory entity.
type DiscoverCategory struct {
	ent.Schema
}

// Fields of the DiscoverCategory.
func (DiscoverCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("name").Optional(),
		field.String("slug").Unique(),
		field.Bool("active").Default(false),
		field.Uint8("order").Default(5),
		field.Uint64("media_id").Unique().Optional(),
		field.Bool("no_nsfw").Default(true),
		field.Bool("local_only").Default(true),
		field.Bool("public_only").Default(true),
		field.Bool("photos_only").Default(true),
		field.Time("active_until").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the DiscoverCategory.
func (DiscoverCategory) Edges() []ent.Edge {
	return nil
}

// Annotations of the DiscoverCategory.
func (DiscoverCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "discover_categories"},
	}
}
