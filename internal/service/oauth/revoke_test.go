package oauth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/ory/fosite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"glintfed.org/internal/lib/fositestore"
)

func TestSvc_Revoke(t *testing.T) {
	env := newTestEnv(t)
	// Use a public client (personalAccess=true) so fosite doesn't require client_secret for revocation.
	oc := env.seedOauthClient(t, 1, true, false)
	fc := &fositestore.FositeClient{OauthClient: oc}

	// Issue a real token so we have something to revoke.
	req := fosite.NewRequest()
	req.Client = fc
	req.Session = &fosite.DefaultSession{
		Subject: "42",
		ExpiresAt: map[fosite.TokenType]time.Time{
			fosite.AccessToken:  time.Now().Add(time.Hour),
			fosite.RefreshToken: time.Now().Add(2 * time.Hour),
		},
	}
	req.SetRequestedScopes(fosite.Arguments{"read"})
	req.GrantScope("read")
	req.RequestedAt = time.Now()

	at, _, err := env.store.CreatePersonalAccessTokens(context.Background(), req)
	require.NoError(t, err)

	t.Run("revoke valid access token returns 200", func(t *testing.T) {
		form := url.Values{
			"token":           {at},
			"token_type_hint": {"access_token"},
			"client_id":       {"1"},
		}
		r := httptest.NewRequest(http.MethodPost, "/oauth/revoke", strings.NewReader(form.Encode()))
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		env.svc.Revoke(w, r)

		// RFC 7009: revocation always returns 200 (even for invalid tokens).
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("revoke with no token still returns 200", func(t *testing.T) {
		form := url.Values{"client_id": {"1"}}
		r := httptest.NewRequest(http.MethodPost, "/oauth/revoke", strings.NewReader(form.Encode()))
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		w := httptest.NewRecorder()

		env.svc.Revoke(w, r)

		// fosite returns 200 for missing token per RFC 7009.
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
