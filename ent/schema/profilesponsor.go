package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ProfileSponsor holds the schema definition for the ProfileSponsor entity.
type ProfileSponsor struct {
	ent.Schema
}

// Fields of the ProfileSponsor.
func (ProfileSponsor) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("profile_id").Unique(),
		field.JSON("sponsors", map[string]any{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ProfileSponsor.
func (ProfileSponsor) Edges() []ent.Edge {
	return nil
}

// Annotations of the ProfileSponsor.
func (ProfileSponsor) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "profile_sponsors"},
	}
}
