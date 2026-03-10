package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Contact holds the schema definition for the Contact entity.
type Contact struct {
	ent.Schema
}

// Fields of the Contact.
func (Contact) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("user_id"),
		field.Bool("response_requested").Default(false),
		field.Text("message"),
		field.Text("response"),
		field.Time("read_at").Optional(),
		field.Time("responded_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Contact.
func (Contact) Edges() []ent.Edge {
	return nil
}

// Annotations of the Contact.
func (Contact) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "contacts"},
	}
}
