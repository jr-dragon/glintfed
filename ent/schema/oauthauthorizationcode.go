package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// OauthAuthorizationCode holds the schema definition for the OauthAuthorizationCode entity.
type OauthAuthorizationCode struct{ ent.Schema }

// Fields of the OauthAuthorizationCode.
func (OauthAuthorizationCode) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("request_id"),
		field.String("client_id"),
		field.String("subject"),
		field.JSON("scopes", []string{}),
		field.Bytes("session"),
		field.Bool("active").Default(true),
		field.Time("requested_at"),
		field.Time("expires_at"),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

// Edges of the OauthAuthorizationCode.
func (OauthAuthorizationCode) Edges() []ent.Edge { return nil }

// Annotations of the OauthAuthorizationCode.
func (OauthAuthorizationCode) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "oauth_authorization_codes"}}
}
