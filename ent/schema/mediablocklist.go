package schema

import (
	"entgo.io/ent/schema"
	"entgo.io/ent/dialect/entsql"
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// MediaBlocklist holds the schema definition for the MediaBlocklist entity.
type MediaBlocklist struct {
	ent.Schema
}

// Fields of the MediaBlocklist.
func (MediaBlocklist) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("sha256").Unique().Optional(),
		field.String("sha512").Unique().Optional(),
		field.String("name").Optional(),
		field.Text("description").Optional(),
		field.Bool("active").Default(true),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the MediaBlocklist.
func (MediaBlocklist) Edges() []ent.Edge {
	return nil
}

// Annotations of the MediaBlocklist.
func (MediaBlocklist) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "media_blocklists"},
	}
}
