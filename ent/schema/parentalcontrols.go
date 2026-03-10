package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ParentalControls holds the schema definition for the ParentalControls entity.
type ParentalControls struct {
	ent.Schema
}

// Fields of the ParentalControls.
func (ParentalControls) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint32("parent_id"),
		field.Uint32("child_id").Unique().Optional(),
		field.String("email").Unique().Optional(),
		field.String("verify_code").Optional(),
		field.Time("email_sent_at").Optional(),
		field.Time("email_verified_at").Optional(),
		field.JSON("permissions", map[string]any{}).Optional(),
		field.Time("deleted_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ParentalControls.
func (ParentalControls) Edges() []ent.Edge {
	return nil
}

// Annotations of the ParentalControls.
func (ParentalControls) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "parental_controls"},
	}
}
