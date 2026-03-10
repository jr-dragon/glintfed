package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// InstanceActor holds the schema definition for the InstanceActor entity.
type InstanceActor struct {
	ent.Schema
}

// Fields of the InstanceActor.
func (InstanceActor) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Text("private_key").Optional(),
		field.Text("public_key").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the InstanceActor.
func (InstanceActor) Edges() []ent.Edge {
	return nil
}

// Annotations of the InstanceActor.
func (InstanceActor) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "instance_actors"},
	}
}
