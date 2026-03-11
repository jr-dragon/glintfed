package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Circle holds the schema definition for the Circle entity.
type Circle struct {
	ent.Schema
}

// Fields of the Circle.
func (Circle) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("profile_id"),
		field.String("name").Optional(),
		field.Text("description").Optional(),
		field.String("scope").Default("public"),
		field.Bool("bcc").Default(false),
		field.Bool("active").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Circle.
func (Circle) Edges() []ent.Edge {
	return nil
}

// Annotations of the Circle.
func (Circle) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "circles"},
	}
}
