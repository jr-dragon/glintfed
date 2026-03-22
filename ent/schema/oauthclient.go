package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// OauthClient holds the schema definition for the OauthClient entity.
type OauthClient struct {
	ent.Schema
}

// Fields of the OauthClient.
func (OauthClient) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("user_id").Optional(),
		field.String("name"),
		field.String("secret").MaxLen(100).Optional(),
		field.String("provider").Optional(),
		field.Text("redirect"),
		field.Bool("personal_access_client"),
		field.Bool("password_client"),
		field.Bool("public").Default(false),
		field.JSON("grant_types", []string{}).Default([]string{"authorization_code", "refresh_token", "password", "client_credentials"}),
		field.JSON("response_types", []string{}).Default([]string{"code", "token"}),
		field.JSON("scopes", []string{}).Default([]string{"read", "write", "follow", "push"}),
		field.JSON("audience", []string{}).Default([]string{}),
		field.Bool("revoked"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the OauthClient.
func (OauthClient) Edges() []ent.Edge {
	return nil
}

// Annotations of the OauthClient.
func (OauthClient) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "oauth_clients"},
	}
}
