package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupActivityGraph holds the schema definition for the GroupActivityGraph entity.
type GroupActivityGraph struct {
	ent.Schema
}

// Fields of the GroupActivityGraph.
func (GroupActivityGraph) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Int64("instance_id").Optional(),
		field.Int64("actor_id").Optional(),
		field.String("verb").Optional(),
		field.String("id_url").Unique().Optional(),
		field.JSON("payload", map[string]any{}).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the GroupActivityGraph.
func (GroupActivityGraph) Edges() []ent.Edge {
	return nil
}

// Annotations of the GroupActivityGraph.
func (GroupActivityGraph) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_activity_graphs"},
	}
}
