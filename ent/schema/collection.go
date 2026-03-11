package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Collection holds the schema definition for the Collection entity.
type Collection struct {
	ent.Schema
}

// Fields of the Collection.
func (Collection) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("profile_id").Optional(),
		field.String("title").Optional(),
		field.Text("description").Optional(),
		field.Bool("is_nsfw").Default(false),
		field.String("visibility").Default("public"),
		field.Time("published_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Collection.
func (Collection) Edges() []ent.Edge {
	return nil
}

// Annotations of the Collection.
func (Collection) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "collections"},
	}
}
