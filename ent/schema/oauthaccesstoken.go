package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// OauthAccessToken holds the schema definition for the OauthAccessToken entity.
type OauthAccessToken struct{ ent.Schema }

// Fields of the OauthAccessToken.
func (OauthAccessToken) Fields() []ent.Field {
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

// Edges of the OauthAccessToken.
func (OauthAccessToken) Edges() []ent.Edge { return nil }

// Annotations of the OauthAccessToken.
func (OauthAccessToken) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "oauth_access_tokens"}}
}
