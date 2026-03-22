package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// OauthAuthorizationCode holds the schema definition for the OauthAuthorizationCode entity.
type OauthAuthorizationCode struct{ ent.Schema }

// Fields of the OauthAuthorizationCode.
func (OauthAuthorizationCode) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(100).Unique(),
		field.Uint64("user_id"),
		field.Uint64("client_id"),
		field.Text("scopes").Optional(),
		field.Bool("revoked"),
		field.Time("expires_at").Optional(),
	}
}

// Edges of the OauthAuthorizationCode.
func (OauthAuthorizationCode) Edges() []ent.Edge { return nil }

// Indexes of the OauthAuthorizationCode.
func (OauthAuthorizationCode) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}

// Annotations of the OauthAuthorizationCode.
func (OauthAuthorizationCode) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "oauth_auth_codes"}}
}
