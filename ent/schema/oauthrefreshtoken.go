package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// OauthRefreshToken holds the schema definition for the OauthRefreshToken entity.
type OauthRefreshToken struct{ ent.Schema }

// Fields of the OauthRefreshToken.
func (OauthRefreshToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(100).Unique(),
		field.String("access_token_id").MaxLen(100),
		field.Bool("revoked"),
		field.Time("expires_at").Optional(),
	}
}

// Edges of the OauthRefreshToken.
func (OauthRefreshToken) Edges() []ent.Edge { return nil }

// Indexes of the OauthRefreshToken.
func (OauthRefreshToken) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("access_token_id"),
	}
}

// Annotations of the OauthRefreshToken.
func (OauthRefreshToken) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "oauth_refresh_tokens"}}
}
