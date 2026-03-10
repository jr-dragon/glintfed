package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// AutospamCustomTokens holds the schema definition for the AutospamCustomTokens entity.
type AutospamCustomTokens struct {
	ent.Schema
}

// Fields of the AutospamCustomTokens.
func (AutospamCustomTokens) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("token"),
		field.Int("weight").Default(1),
		field.Bool("is_spam").Default(true),
		field.Text("note").Optional(),
		field.String("category").Optional(),
		field.Bool("active").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the AutospamCustomTokens.
func (AutospamCustomTokens) Edges() []ent.Edge {
	return nil
}

// Annotations of the AutospamCustomTokens.
func (AutospamCustomTokens) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "autospam_custom_tokens"},
	}
}
