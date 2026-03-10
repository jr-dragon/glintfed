package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ModeratedProfile holds the schema definition for the ModeratedProfile entity.
type ModeratedProfile struct {
	ent.Schema
}

// Fields of the ModeratedProfile.
func (ModeratedProfile) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("profile_url").Unique().Optional(),
		field.Uint64("profile_id").Unique().Optional(),
		field.String("domain").Optional(),
		field.Text("note").Optional(),
		field.Bool("is_banned").Default(false),
		field.Bool("is_nsfw").Default(false),
		field.Bool("is_unlisted").Default(false),
		field.Bool("is_noautolink").Default(false),
		field.Bool("is_nodms").Default(false),
		field.Bool("is_notrending").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ModeratedProfile.
func (ModeratedProfile) Edges() []ent.Edge {
	return nil
}

// Annotations of the ModeratedProfile.
func (ModeratedProfile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "moderated_profiles"},
	}
}
