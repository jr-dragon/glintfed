package schema

import "entgo.io/ent"

// AppRegister holds the schema definition for the AppRegister entity.
type AppRegister struct {
	ent.Schema
}

// Fields of the AppRegister.
func (AppRegister) Fields() []ent.Field {
	return nil
}

// Edges of the AppRegister.
func (AppRegister) Edges() []ent.Edge {
	return nil
}
