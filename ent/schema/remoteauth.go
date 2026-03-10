package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// RemoteAuth holds the schema definition for the RemoteAuth entity.
type RemoteAuth struct {
	ent.Schema
}

// Fields of the RemoteAuth.
func (RemoteAuth) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("software").Optional(),
		field.String("domain").Optional(),
		field.String("webfinger").Unique().Optional(),
		field.Uint32("instance_id").Optional(),
		field.Uint64("user_id").Unique().Optional(),
		field.Uint32("client_id").Optional(),
		field.String("ip_address").Optional(),
		field.Text("bearer_token").Optional(),
		field.JSON("verify_credentials", map[string]any{}).Optional(),
		field.Time("last_successful_login_at").Optional(),
		field.Time("last_verify_credentials_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the RemoteAuth.
func (RemoteAuth) Edges() []ent.Edge {
	return nil
}

// Annotations of the RemoteAuth.
func (RemoteAuth) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "remote_auths"},
	}
}
