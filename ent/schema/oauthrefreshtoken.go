package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// OauthRefreshToken holds the schema definition for the OauthRefreshToken entity.
type OauthRefreshToken struct{ ent.Schema }

// Fields of the OauthRefreshToken.
func (OauthRefreshToken) Fields() []ent.Field {
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

// Edges of the OauthRefreshToken.
func (OauthRefreshToken) Edges() []ent.Edge { return nil }

// Annotations of the OauthRefreshToken.
func (OauthRefreshToken) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "oauth_refresh_tokens"}}
}
