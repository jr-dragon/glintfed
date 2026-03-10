package schema

import "entgo.io/ent"

// UserAppSettings holds the schema definition for the UserAppSettings entity.
type UserAppSettings struct {
	ent.Schema
}

// Fields of the UserAppSettings.
func (UserAppSettings) Fields() []ent.Field {
	return nil
}

// Edges of the UserAppSettings.
func (UserAppSettings) Edges() []ent.Edge {
	return nil
}
