package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupInvitation holds the schema definition for the GroupInvitation entity.
type GroupInvitation struct {
	ent.Schema
}

// Fields of the GroupInvitation.
func (GroupInvitation) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("group_id"),
		field.Uint64("from_profile_id"),
		field.Uint64("to_profile_id"),
		field.String("role").Optional(),
		field.Bool("to_local").Default(true),
		field.Bool("from_local").Default(true),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the GroupInvitation.
func (GroupInvitation) Edges() []ent.Edge {
	return nil
}

// Annotations of the GroupInvitation.
func (GroupInvitation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_invitations"},
	}
}
