package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupHashtag holds the schema definition for the GroupHashtag entity.
type GroupHashtag struct {
	ent.Schema
}

// Fields of the GroupHashtag.
func (GroupHashtag) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("name").Unique(),
		field.String("formatted").Optional(),
		field.Bool("recommended").Default(false),
		field.Bool("sensitive").Default(false),
		field.Bool("banned").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the GroupHashtag.
func (GroupHashtag) Edges() []ent.Edge {
	return nil
}

// Annotations of the GroupHashtag.
func (GroupHashtag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_hashtags"},
	}
}
