package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ConfigCache holds the schema definition for the ConfigCache entity.
type ConfigCache struct {
	ent.Schema
}

// Fields of the ConfigCache.
func (ConfigCache) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("k").Unique(),
		field.Text("v").Optional(),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ConfigCache.
func (ConfigCache) Edges() []ent.Edge {
	return nil
}

// Annotations of the ConfigCache.
func (ConfigCache) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "config_cache"},
	}
}
