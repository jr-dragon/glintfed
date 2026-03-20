package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Instance holds the schema definition for the Instance entity.
type Instance struct {
	ent.Schema
}

// Fields of the Instance.
func (Instance) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("domain").Unique(),
		field.Bool("active_deliver").Optional(),
		field.String("url").Optional(),
		field.String("name").Optional(),
		field.String("admin_url").Optional(),
		field.String("limit_reason").Optional(),
		field.Bool("unlisted").Default(false),
		field.Bool("auto_cw").Default(false),
		field.Bool("banned").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.String("software").Optional(),
		field.Uint32("user_count").Optional(),
		field.Uint32("status_count").Optional(),
		field.Time("last_crawled_at").Optional(),
		field.Time("actors_last_synced_at").Optional(),
		field.Text("notes").Optional(),
		field.Bool("manually_added").Default(false),
		field.String("base_domain").Optional(),
		field.Bool("ban_subdomains").Optional(),
		field.String("ip_address").Optional(),
		field.Bool("list_limitation").Default(false),
		field.Bool("valid_nodeinfo").Optional(),
		field.Time("nodeinfo_last_fetched").Optional(),
		field.Bool("delivery_timeout").Default(false),
		field.Time("delivery_next_after").Optional(),
		field.String("shared_inbox").Optional(),
	}
}

// Edges of the Instance.
func (Instance) Edges() []ent.Edge {
	return nil
}

// Annotations of the Instance.
func (Instance) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "instances"},
	}
}
