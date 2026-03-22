package oauth

import (
	"context"
	"testing"
	"time"

	"github.com/ory/fosite"
	"github.com/stretchr/testify/require"

	"glintfed.org/ent"
	"glintfed.org/internal/data"
	"glintfed.org/internal/lib/fositestore"
)

type testEnv struct {
	svc      *svc
	store    *fositestore.Store
	entDB    *ent.Client
	auth     *UserAuthenticatorMock
	provider fosite.OAuth2Provider
}

// newTestEnv creates a full OAuth service environment backed by in-memory SQLite.
func newTestEnv(t *testing.T) *testEnv {
	t.Helper()
	client, _, err := data.NewTestClient(t)
	require.NoError(t, err)

	cfg := &data.Config{}
	cfg.App.Url = "https://example.com"
	cfg.App.Auth.OAuth.HMACSecret = "test-hmac-secret-32-bytes-long!!"
	cfg.App.Auth.OAuth.AccessTokenLifespanDays = 365
	cfg.App.Auth.OAuth.RefreshTokenLifespanDays = 400

	store := fositestore.New(client, cfg)
	provider := fositestore.NewOAuth2Provider(store, cfg)
	auth := &UserAuthenticatorMock{}

	s := &svc{
		provider:        provider,
		store:           store,
		auth:            auth,
		appURL:          cfg.App.Url,
		accessTokenTTL:  365 * 24 * time.Hour,
		refreshTokenTTL: 400 * 24 * time.Hour,
	}
	return &testEnv{svc: s, store: store, entDB: client.Ent, auth: auth, provider: provider}
}

// seedOauthClient inserts an OauthClient fixture.
// personalAccess=true creates a public client that does not require client_secret for revocation.
func (e *testEnv) seedOauthClient(t *testing.T, id uint64, personalAccess, password bool) *ent.OauthClient {
	t.Helper()
	c, err := e.entDB.OauthClient.Create().
		SetID(id).
		SetName("Test Client").
		SetSecret("client-secret").
		SetRedirect("https://example.com").
		SetPersonalAccessClient(personalAccess).
		SetPasswordClient(password).
		SetRevoked(false).
		Save(context.Background())
	require.NoError(t, err)
	return c
}
