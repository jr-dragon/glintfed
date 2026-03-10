package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// AdminShadowFilter holds the schema definition for the AdminShadowFilter entity.
type AdminShadowFilter struct {
	ent.Schema
}

// Fields of the AdminShadowFilter.
func (AdminShadowFilter) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("admin_id").Optional(),
		field.String("item_type"),
		field.Uint64("item_id"),
		field.Bool("is_local").Default(true),
		field.Text("note").Optional(),
		field.Bool("active").Default(false),
		field.JSON("history", []any{}).Optional(),
		field.JSON("ruleset", map[string]any{}).Optional(),
		field.Bool("prevent_ap_fanout").Default(false),
		field.Bool("prevent_new_dms").Default(false),
		field.Bool("ignore_reports").Default(false),
		field.Bool("ignore_mentions").Default(false),
		field.Bool("ignore_links").Default(false),
		field.Bool("ignore_hashtags").Default(false),
		field.Bool("hide_from_public_feeds").Default(false),
		field.Bool("hide_from_tag_feeds").Default(false),
		field.Bool("hide_embeds").Default(false),
		field.Bool("hide_from_story_carousel").Default(false),
		field.Bool("hide_from_search_autocomplete").Default(false),
		field.Bool("hide_from_search").Default(false),
		field.Bool("requires_login").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the AdminShadowFilter.
func (AdminShadowFilter) Edges() []ent.Edge {
	return nil
}

// Annotations of the AdminShadowFilter.
func (AdminShadowFilter) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "admin_shadow_filters"},
	}
}
