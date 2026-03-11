package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ProfileMigration holds the schema definition for the ProfileMigration entity.
type ProfileMigration struct {
	ent.Schema
}

// Fields of the ProfileMigration.
func (ProfileMigration) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("profile_id"),
		field.String("acct").Optional(),
		field.Uint64("followers_count").Default(0),
		field.Uint64("target_profile_id").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ProfileMigration.
func (ProfileMigration) Edges() []ent.Edge {
	return nil
}

// Annotations of the ProfileMigration.
func (ProfileMigration) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "profile_migrations"},
	}
}
