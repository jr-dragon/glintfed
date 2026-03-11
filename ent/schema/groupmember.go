package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupMember holds the schema definition for the GroupMember entity.
type GroupMember struct {
	ent.Schema
}

// Fields of the GroupMember.
func (GroupMember) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("group_id"),
		field.Uint64("profile_id"),
		field.String("role").Default("member"),
		field.Bool("local_group").Default(false),
		field.Bool("local_profile").Default(false),
		field.Bool("join_request").Default(false),
		field.Time("approved_at").Optional(),
		field.Time("rejected_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the GroupMember.
func (GroupMember) Edges() []ent.Edge {
	return nil
}

// Annotations of the GroupMember.
func (GroupMember) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_members"},
	}
}
