package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("user_id").Optional(),
		field.String("domain").Optional(),
		field.String("username").Optional(),
		field.String("status").Optional(),
		field.String("name").Optional(),
		field.Text("bio").Optional(),
		field.Bool("unlisted").Default(false),
		field.Bool("cw").Default(false),
		field.Bool("no_autolink").Default(false),
		field.String("location").Optional(),
		field.String("website").Optional(),
		field.String("profile_layout").Optional(),
		field.String("header_bg").Optional(),
		field.String("post_layout").Optional(),
		field.Bool("is_private").Default(false),
		field.String("shared_inbox").StorageKey("sharedInbox").Optional(),
		field.String("inbox_url").Optional(),
		field.String("outbox_url").Optional(),
		field.String("key_id").Unique().Optional(),
		field.String("follower_url").Optional(),
		field.String("following_url").Optional(),
		field.Text("private_key").Optional(),
		field.Text("public_key").Optional(),
		field.String("remote_url").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional(),
		field.Time("delete_after").Optional(),
		field.Bool("is_suggestable").Default(false),
		field.Time("last_fetched_at").Optional(),
		field.Uint32("status_count").Default(0).Optional(),
		field.Uint32("followers_count").Default(0).Optional(),
		field.Uint32("following_count").Default(0).Optional(),
		field.String("webfinger").Unique().Optional(),
		field.String("avatar_url").Optional(),
		field.Time("last_status_at").Optional(),
		field.Uint64("moved_to_profile_id").Optional(),
		field.Bool("indexable").Default(false),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return nil
}

// Annotations of the Profile.
func (Profile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "profiles"},
	}
}
