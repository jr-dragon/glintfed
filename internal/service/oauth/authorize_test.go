package oauth

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSvc_Authorize_InvalidRequest(t *testing.T) {
	env := newTestEnv(t)
	env.svc.loginURL = "https://login.example.com/login"

	// Missing required params — fosite rejects before we reach the redirect.
	req := httptest.NewRequest(http.MethodGet, "/oauth/authorize", nil)
	w := httptest.NewRecorder()

	env.svc.Authorize(w, req)

	assert.NotEqual(t, http.StatusFound, w.Code)
}

func TestSvc_Authorize_RedirectsToLoginURL(t *testing.T) {
	env := newTestEnv(t)
	env.svc.loginURL = "https://login.example.com/login"
	env.seedOauthClient(t, 1, false, true)

	query := "response_type=code&client_id=1&redirect_uri=https%3A%2F%2Fexample.com&scope=read&state=random-state-value"
	req := httptest.NewRequest(http.MethodGet, "/oauth/authorize?"+query, nil)
	w := httptest.NewRecorder()

	env.svc.Authorize(w, req)

	require.Equal(t, http.StatusSeeOther, w.Code)

	location := w.Header().Get("Location")
	require.NotEmpty(t, location)

	parsed, err := url.Parse(location)
	require.NoError(t, err)

	// Redirects to loginURL.
	assert.Equal(t, "login.example.com", parsed.Host)
	assert.Equal(t, "/login", parsed.Path)

	// Includes `next` pointing back to /oauth/authorize with original params.
	next := parsed.Query().Get("next")
	require.NotEmpty(t, next)
	assert.True(t, strings.Contains(next, "/oauth/authorize"))
	assert.True(t, strings.Contains(next, "client_id=1"))
}

func TestSvc_Authorize_NoLoginURL(t *testing.T) {
	env := newTestEnv(t)
	env.svc.loginURL = "" // not configured
	env.seedOauthClient(t, 1, false, true)

	query := "response_type=code&client_id=1&redirect_uri=https%3A%2F%2Fexample.com&state=random-state-value"
	req := httptest.NewRequest(http.MethodGet, "/oauth/authorize?"+query, nil)
	w := httptest.NewRecorder()

	env.svc.Authorize(w, req)

	// Should return an error (not a redirect to loginURL) when login_url is not configured.
	location := w.Header().Get("Location")
	if location != "" {
		parsed, _ := url.Parse(location)
		assert.NotEqual(t, "login.example.com", parsed.Host)
	}
}
