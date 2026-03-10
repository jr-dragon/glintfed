package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint32("category_id").Default(1),
		field.Uint64("profile_id").Optional(),
		field.String("status").Optional(),
		field.String("name").Optional(),
		field.Text("description").Optional(),
		field.Text("rules").Optional(),
		field.Bool("local").Default(true),
		field.String("remote_url").Optional(),
		field.String("inbox_url").Optional(),
		field.Bool("is_private").Default(false),
		field.Bool("local_only").Default(false),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.Uint32("member_count").Optional(),
		field.Bool("recommended").Default(false),
		field.Bool("discoverable").Default(false),
		field.Bool("activitypub").Default(false),
		field.Bool("is_nsfw").Default(false),
		field.Bool("dms").Default(false),
		field.Bool("autospam").Default(false),
		field.Bool("verified").Default(false),
		field.Time("last_active_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return nil
}

// Annotations of the Group.
func (Group) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "groups"},
	}
}
