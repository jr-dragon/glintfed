package schema

import (
	"entgo.io/ent/schema"
	"entgo.io/ent/dialect/entsql"
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Page holds the schema definition for the Page entity.
type Page struct {
	ent.Schema
}

// Fields of the Page.
func (Page) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("root").Optional(),
		field.String("slug").Unique().Optional(),
		field.String("title").Optional(),
		field.Uint32("category_id").Optional(),
		field.Text("content").Optional(),
		field.String("template").Default("layouts.app"),
		field.Bool("active").Default(false),
		field.Bool("cached").Default(true),
		field.Time("active_until").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Page.
func (Page) Edges() []ent.Edge {
	return nil
}

// Annotations of the Page.
func (Page) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "pages"},
	}
}
