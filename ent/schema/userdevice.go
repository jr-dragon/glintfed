package schema

import "entgo.io/ent"

// UserDevice holds the schema definition for the UserDevice entity.
type UserDevice struct {
	ent.Schema
}

// Fields of the UserDevice.
func (UserDevice) Fields() []ent.Field {
	return nil
}

// Edges of the UserDevice.
func (UserDevice) Edges() []ent.Edge {
	return nil
}
