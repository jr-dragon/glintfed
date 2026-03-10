package schema

import "entgo.io/ent"

// ProfileSponsor holds the schema definition for the ProfileSponsor entity.
type ProfileSponsor struct {
	ent.Schema
}

// Fields of the ProfileSponsor.
func (ProfileSponsor) Fields() []ent.Field {
	return nil
}

// Edges of the ProfileSponsor.
func (ProfileSponsor) Edges() []ent.Edge {
	return nil
}
