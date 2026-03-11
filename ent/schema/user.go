package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("profile_id").Unique().Optional(),
		field.String("name").Optional(),
		field.String("username").Unique().Optional(),
		field.String("email").Unique(),
		field.String("status").Optional(),
		field.String("language").Optional(),
		field.String("password").Sensitive(),
		field.String("remember_token").MaxLen(100).Optional(),
		field.Bool("is_admin").Default(false),
		field.Time("email_verified_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional(),
		field.Time("last_active_at").Optional(),
		field.Bool("two_fa_enabled").StorageKey("2fa_enabled").Default(false),
		field.String("two_fa_secret").StorageKey("2fa_secret").Optional(),
		field.JSON("two_fa_backup_codes", []string{}).StorageKey("2fa_backup_codes").Optional(),
		field.Time("two_fa_setup_at").StorageKey("2fa_setup_at").Optional(),
		field.Time("delete_after").Optional(),
		field.Bool("has_interstitial").Default(false),
		field.String("guid").Unique().Optional(),
		field.String("domain").Optional(),
		field.String("register_source").Default("web").Optional(),
		field.String("app_register_token").Optional(),
		field.String("app_register_ip").Optional(),
		field.Bool("has_roles").Default(false),
		field.Uint32("parent_id").Optional(),
		field.Uint8("role_id").Optional(),
		field.String("expo_token").Optional(),
		field.Bool("notify_like").Default(true),
		field.Bool("notify_follow").Default(true),
		field.Bool("notify_mention").Default(true),
		field.Bool("notify_comment").Default(true),
		field.Uint64("storage_used").Default(0),
		field.Time("storage_used_updated_at").Optional(),
		field.Bool("notify_enabled").Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
	}
}
