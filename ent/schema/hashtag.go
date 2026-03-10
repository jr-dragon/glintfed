package schema

import (
	"entgo.io/ent/schema"
	"entgo.io/ent/dialect/entsql"
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Hashtag holds the schema definition for the Hashtag entity.
type Hashtag struct {
	ent.Schema
}

// Fields of the Hashtag.
func (Hashtag) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("name").Unique(),
		field.String("slug").Unique(),
		field.Bool("can_trend").Optional(),
		field.Bool("can_search").Optional(),
		field.Bool("is_nsfw").Default(false),
		field.Bool("is_banned").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Uint32("cached_count").Optional(),
	}
}

// Edges of the Hashtag.
func (Hashtag) Edges() []ent.Edge {
	return nil
}

// Annotations of the Hashtag.
func (Hashtag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "hashtags"},
	}
}
