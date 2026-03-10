package schema

import "entgo.io/ent"

// ModeratedProfile holds the schema definition for the ModeratedProfile entity.
type ModeratedProfile struct {
	ent.Schema
}

// Fields of the ModeratedProfile.
func (ModeratedProfile) Fields() []ent.Field {
	return nil
}

// Edges of the ModeratedProfile.
func (ModeratedProfile) Edges() []ent.Edge {
	return nil
}
