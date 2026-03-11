package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GroupEvent holds the schema definition for the GroupEvent entity.
type GroupEvent struct {
	ent.Schema
}

// Fields of the GroupEvent.
func (GroupEvent) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("group_id").Optional(),
		field.Uint64("profile_id").Optional(),
		field.String("name").Optional(),
		field.String("type"),
		field.JSON("tags", []string{}).Optional(),
		field.JSON("location", map[string]any{}).Optional(),
		field.Text("description").Optional(),
		field.JSON("metadata", map[string]any{}).Optional(),
		field.Bool("open").Default(false),
		field.Bool("comments_open").Default(false),
		field.Bool("show_guest_list").Default(false),
		field.Time("start_at").Optional(),
		field.Time("end_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the GroupEvent.
func (GroupEvent) Edges() []ent.Edge {
	return nil
}

// Annotations of the GroupEvent.
func (GroupEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "group_events"},
	}
}
