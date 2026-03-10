package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Bookmark holds the schema definition for the Bookmark entity.
type Bookmark struct {
	ent.Schema
}

// Fields of the Bookmark.
func (Bookmark) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("status_id"),
		field.Uint64("profile_id"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Bookmark.
func (Bookmark) Edges() []ent.Edge {
	return nil

}

// Annotations of the Bookmark.
func (Bookmark) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "bookmarks"},
	}
}
