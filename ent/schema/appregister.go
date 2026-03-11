package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// AppRegister holds the schema definition for the AppRegister entity.
type AppRegister struct {
	ent.Schema
}

// Fields of the AppRegister.
func (AppRegister) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("email"),
		field.String("verify_code"),
		field.Time("email_delivered_at").Optional(),
		field.Time("email_verified_at").Optional(),
		field.Uint32("uses").Default(0),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the AppRegister.
func (AppRegister) Edges() []ent.Edge {
	return nil
}

// Annotations of the AppRegister.
func (AppRegister) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_registers"},
	}
}
