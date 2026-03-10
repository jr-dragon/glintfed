package schema

import "entgo.io/ent"

// InstanceActor holds the schema definition for the InstanceActor entity.
type InstanceActor struct {
	ent.Schema
}

// Fields of the InstanceActor.
func (InstanceActor) Fields() []ent.Field {
	return nil
}

// Edges of the InstanceActor.
func (InstanceActor) Edges() []ent.Edge {
	return nil
}
