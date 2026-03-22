package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// OauthPkce holds the schema definition for the OauthPkce entity.
type OauthPkce struct{ ent.Schema }

// Fields of the OauthPkce.
func (OauthPkce) Fields() []ent.Field {
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

// Edges of the OauthPkce.
func (OauthPkce) Edges() []ent.Edge { return nil }

// Annotations of the OauthPkce.
func (OauthPkce) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "oauth_pkce"}}
}
