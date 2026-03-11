package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CuratedRegisterTemplate holds the schema definition for the CuratedRegisterTemplate entity.
type CuratedRegisterTemplate struct {
	ent.Schema
}

// Fields of the CuratedRegisterTemplate.
func (CuratedRegisterTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("name").Optional(),
		field.Text("description").Optional(),
		field.Text("content").Optional(),
		field.Bool("is_active").Default(false),
		field.Uint8("order").Default(10),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the CuratedRegisterTemplate.
func (CuratedRegisterTemplate) Edges() []ent.Edge {
	return nil
}

// Annotations of the CuratedRegisterTemplate.
func (CuratedRegisterTemplate) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "curated_register_templates"},
	}
}
