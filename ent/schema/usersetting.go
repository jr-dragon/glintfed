package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UserSetting holds the schema definition for the UserSetting entity.
type UserSetting struct {
	ent.Schema
}

// Fields of the UserSetting.
func (UserSetting) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("user_id").Unique(),
		field.String("role").Default("user"),
		field.Bool("crawlable").Default(true),
		field.Bool("show_guests").Default(true),
		field.Bool("show_discover").Default(true),
		field.Bool("public_dm").Default(false),
		field.Bool("hide_cw_search").Default(true),
		field.Bool("hide_blocked_search").Default(true),
		field.Bool("always_show_cw").Default(false),
		field.Bool("compose_media_descriptions").Default(false),
		field.Bool("reduce_motion").Default(false),
		field.Bool("optimize_screen_reader").Default(false),
		field.Bool("high_contrast_mode").Default(false),
		field.Bool("video_autoplay").Default(false),
		field.Bool("send_email_new_follower").Default(false),
		field.Bool("send_email_new_follower_request").Default(true),
		field.Bool("send_email_on_share").Default(false),
		field.Bool("send_email_on_like").Default(false),
		field.Bool("send_email_on_mention").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Bool("show_profile_followers").Default(true),
		field.Bool("show_profile_follower_count").Default(true),
		field.Bool("show_profile_following").Default(true),
		field.Bool("show_profile_following_count").Default(true),
		field.JSON("compose_settings", map[string]any{}).Optional(),
		field.JSON("other", map[string]any{}).Optional(),
		field.Bool("show_atom").Default(true),
	}
}

// Edges of the UserSetting.
func (UserSetting) Edges() []ent.Edge {
	return nil
}

// Annotations of the UserSetting.
func (UserSetting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_settings"},
	}
}
