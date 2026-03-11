package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Portfolio holds the schema definition for the Portfolio entity.
type Portfolio struct {
	ent.Schema
}

// Fields of the Portfolio.
func (Portfolio) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("user_id").Unique().Optional(),
		field.Uint64("profile_id").Unique(),
		field.Bool("active").Optional(),
		field.Bool("show_captions").Default(true).Optional(),
		field.Bool("show_license").Default(true).Optional(),
		field.Bool("show_location").Default(true).Optional(),
		field.Bool("show_timestamp").Default(true).Optional(),
		field.Bool("show_link").Default(true).Optional(),
		field.String("profile_source").Default("recent").Optional(),
		field.Bool("show_avatar").Default(true).Optional(),
		field.Bool("show_bio").Default(true).Optional(),
		field.String("profile_layout").Default("grid").Optional(),
		field.String("profile_container").Default("fixed").Optional(),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Portfolio.
func (Portfolio) Edges() []ent.Edge {
	return nil
}

// Annotations of the Portfolio.
func (Portfolio) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "portfolios"},
	}
}
