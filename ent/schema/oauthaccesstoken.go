package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// OauthAccessToken holds the schema definition for the OauthAccessToken entity.
type OauthAccessToken struct{ ent.Schema }

// Fields of the OauthAccessToken.
func (OauthAccessToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(100).Unique(),
		field.Uint64("user_id").Optional(),
		field.Uint64("client_id"),
		field.String("name").MaxLen(191).Optional(),
		field.Text("scopes").Optional(),
		field.Bool("revoked"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("expires_at").Optional(),
	}
}

// Edges of the OauthAccessToken.
func (OauthAccessToken) Edges() []ent.Edge { return nil }

// Indexes of the OauthAccessToken.
func (OauthAccessToken) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}

// Annotations of the OauthAccessToken.
func (OauthAccessToken) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "oauth_access_tokens"}}
}
