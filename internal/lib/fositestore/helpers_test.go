package fositestore

import (
	"context"
	"testing"
	"time"

	"github.com/ory/fosite"
	"github.com/stretchr/testify/require"

	"glintfed.org/ent"
	"glintfed.org/internal/data"
)

// newTestStore creates a Store backed by in-memory SQLite for testing.
func newTestStore(t *testing.T) (*Store, *ent.Client) {
	t.Helper()
	client, _, err := data.NewTestClient(t)
	require.NoError(t, err)

	cfg := &data.Config{}
	cfg.App.Auth.OAuth.HMACSecret = "test-hmac-secret-32-bytes-long!!"

	return New(client, cfg), client.Ent
}

// createOauthClient inserts an OauthClient fixture into the DB.
func createOauthClient(t *testing.T, db *ent.Client, id uint64, personalAccess, password, revoked bool) *ent.OauthClient {
	t.Helper()
	c, err := db.OauthClient.Create().
		SetID(id).
		SetName("Test Client").
		SetSecret("client-secret").
		SetRedirect("https://example.com\nhttps://example.com/callback").
		SetPersonalAccessClient(personalAccess).
		SetPasswordClient(password).
		SetRevoked(revoked).
		Save(context.Background())
	require.NoError(t, err)
	return c
}

// newTestRequester builds a fosite.Requester with the given FositeClient.
func newTestRequester(fc *FositeClient, subject string, scopes ...string) fosite.Requester {
	req := fosite.NewRequest()
	req.Client = fc
	req.Session = &fosite.DefaultSession{
		Subject: subject,
		ExpiresAt: map[fosite.TokenType]time.Time{
			fosite.AccessToken:  time.Now().Add(time.Hour),
			fosite.RefreshToken: time.Now().Add(2 * time.Hour),
		},
	}
	req.SetRequestedScopes(fosite.Arguments(scopes))
	for _, s := range scopes {
		req.GrantScope(s)
	}
	req.RequestedAt = time.Now()
	return req
}
