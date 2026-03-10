package schema

import "entgo.io/ent"

// OauthClient holds the schema definition for the OauthClient entity.
type OauthClient struct {
	ent.Schema
}

// Fields of the OauthClient.
func (OauthClient) Fields() []ent.Field {
	return nil
}

// Edges of the OauthClient.
func (OauthClient) Edges() []ent.Edge {
	return nil
}
