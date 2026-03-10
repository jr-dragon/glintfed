package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// LiveStream holds the schema definition for the LiveStream entity.
type LiveStream struct {
	ent.Schema
}

// Fields of the LiveStream.
func (LiveStream) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("profile_id"),
		field.String("stream_id").Unique().Optional(),
		field.String("stream_key").Optional(),
		field.String("visibility").Optional(),
		field.String("name").Optional(),
		field.Text("description").Optional(),
		field.String("thumbnail_path").Optional(),
		field.JSON("settings", map[string]any{}).Optional(),
		field.Bool("live_chat").Default(true),
		field.JSON("mod_ids", []uint64{}).Optional(),
		field.Bool("discoverable").Optional(),
		field.Time("live_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the LiveStream.
func (LiveStream) Edges() []ent.Edge {
	return nil
}

// Annotations of the LiveStream.
func (LiveStream) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "live_streams"},
	}
}
