package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Report holds the schema definition for the Report entity.
type Report struct {
	ent.Schema
}

// Fields of the Report.
func (Report) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("profile_id"),
		field.Uint64("user_id").Optional(),
		field.Uint64("object_id"),
		field.String("object_type").Optional(),
		field.Uint64("reported_profile_id").Optional(),
		field.String("type").Optional(),
		field.String("message").Optional(),
		field.Time("admin_seen").Optional(),
		field.Bool("not_interested").Default(false),
		field.Bool("spam").Default(false),
		field.Bool("nsfw").Default(false),
		field.Bool("abusive").Default(false),
		field.JSON("meta", map[string]any{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Report.
func (Report) Edges() []ent.Edge {
	return nil
}

// Annotations of the Report.
func (Report) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "reports"},
	}
}
