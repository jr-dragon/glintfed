package schema

import "entgo.io/ent"

// RemoteAuth holds the schema definition for the RemoteAuth entity.
type RemoteAuth struct {
	ent.Schema
}

// Fields of the RemoteAuth.
func (RemoteAuth) Fields() []ent.Field {
	return nil
}

// Edges of the RemoteAuth.
func (RemoteAuth) Edges() []ent.Edge {
	return nil
}
