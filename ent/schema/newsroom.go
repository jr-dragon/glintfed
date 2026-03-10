package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Newsroom holds the schema definition for the Newsroom entity.
type Newsroom struct {
	ent.Schema
}

// Fields of the Newsroom.
func (Newsroom) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("user_id").Optional(),
		field.String("header_photo_url").Optional(),
		field.String("title").Optional(),
		field.String("slug").Unique().Optional(),
		field.String("category").Default("update"),
		field.Text("summary").Optional(),
		field.Text("body").Optional(),
		field.Text("body_rendered").Optional(),
		field.String("link").Optional(),
		field.Bool("force_modal").Default(false),
		field.Bool("show_timeline").Default(false),
		field.Bool("show_link").Default(false),
		field.Bool("auth_only").Default(true),
		field.Time("published_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Newsroom.
func (Newsroom) Edges() []ent.Edge {
	return nil
}

// Annotations of the Newsroom.
func (Newsroom) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "newsroom"},
	}
}
