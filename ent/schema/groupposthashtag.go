package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupPostHashtag holds the schema definition for the GroupPostHashtag entity.
type GroupPostHashtag struct {
	ent.Schema
}

// Fields of the GroupPostHashtag.
func (GroupPostHashtag) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("hashtag_id"),
		field.Uint64("group_id"),
		field.Uint64("profile_id"),
		field.Uint64("status_id").Optional(),
		field.String("status_visibility").Optional(),
		field.Bool("nsfw").Default(false),
	}
}

// Edges of the GroupPostHashtag.
func (GroupPostHashtag) Edges() []ent.Edge {
	return nil
}

// Annotations of the GroupPostHashtag.
func (GroupPostHashtag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_post_hashtags"},
	}
}
