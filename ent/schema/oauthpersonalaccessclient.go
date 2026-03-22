package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// OauthPersonalAccessClient holds the schema definition for the OauthPersonalAccessClient entity.
type OauthPersonalAccessClient struct{ ent.Schema }

// Fields of the OauthPersonalAccessClient.
func (OauthPersonalAccessClient) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.Uint64("client_id"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the OauthPersonalAccessClient.
func (OauthPersonalAccessClient) Edges() []ent.Edge { return nil }

// Annotations of the OauthPersonalAccessClient.
func (OauthPersonalAccessClient) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "oauth_personal_access_clients"}}
}
