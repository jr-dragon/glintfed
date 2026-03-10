package schema

import "entgo.io/ent"

// ProfileMigration holds the schema definition for the ProfileMigration entity.
type ProfileMigration struct {
	ent.Schema
}

// Fields of the ProfileMigration.
func (ProfileMigration) Fields() []ent.Field {
	return nil
}

// Edges of the ProfileMigration.
func (ProfileMigration) Edges() []ent.Edge {
	return nil
}
