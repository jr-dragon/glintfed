package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Avatar holds the schema definition for the Avatar entity.
type Avatar struct {
	ent.Schema
}

// Fields of the Avatar.
func (Avatar) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("profile_id").Unique(),
		field.String("media_path").Optional(),
		field.String("remote_url").Optional(),
		field.String("cdn_url").Optional(),
		field.Bool("is_remote").Optional(),
		field.Uint32("size").Optional(),
		field.Uint32("change_count").Default(0),
		field.Time("last_fetched_at").Optional(),
		field.Time("last_processed_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Avatar.
func (Avatar) Edges() []ent.Edge {
	return nil

}

// Annotations of the Avatar.
func (Avatar) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "avatars"},
	}
}
