package oauth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSvc_Token_PasswordGrant(t *testing.T) {
	env := newTestEnv(t)
	env.seedOauthClient(t, 1, false, true)

	tests := []struct {
		name           string
		form           url.Values
		authFunc       func(ctx context.Context, username, password string) (uint64, error)
		expectedStatus int
		checkBody      func(t *testing.T, body map[string]any)
	}{
		{
			name:           "missing username",
			form:           url.Values{"grant_type": {"password"}, "password": {"pass"}, "client_id": {"1"}},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "missing password",
			form:           url.Values{"grant_type": {"password"}, "username": {"user"}, "client_id": {"1"}},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "authentication failed",
			form: url.Values{
				"grant_type": {"password"},
				"username":   {"user"},
				"password":   {"wrong"},
				"client_id":  {"1"},
			},
			authFunc: func(ctx context.Context, username, password string) (uint64, error) {
				return 0, errors.New("invalid credentials")
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "missing client_id",
			form: url.Values{
				"grant_type": {"password"},
				"username":   {"user"},
				"password":   {"pass"},
			},
			authFunc: func(ctx context.Context, username, password string) (uint64, error) {
				return 42, nil
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "unknown client_id",
			form: url.Values{
				"grant_type": {"password"},
				"username":   {"user"},
				"password":   {"pass"},
				"client_id":  {"999"},
			},
			authFunc: func(ctx context.Context, username, password string) (uint64, error) {
				return 42, nil
			},
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name: "success - default scopes",
			form: url.Values{
				"grant_type": {"password"},
				"username":   {"user"},
				"password":   {"pass"},
				"client_id":  {"1"},
			},
			authFunc: func(ctx context.Context, username, password string) (uint64, error) {
				return 42, nil
			},
			expectedStatus: http.StatusOK,
			checkBody: func(t *testing.T, body map[string]any) {
				assert.NotEmpty(t, body["access_token"])
				assert.NotEmpty(t, body["refresh_token"])
				assert.Equal(t, "Bearer", body["token_type"])
				assert.NotNil(t, body["expires_in"])
			},
		},
		{
			name: "success - explicit scopes",
			form: url.Values{
				"grant_type": {"password"},
				"username":   {"user"},
				"password":   {"pass"},
				"client_id":  {"1"},
				"scope":      {"read write"},
			},
			authFunc: func(ctx context.Context, username, password string) (uint64, error) {
				return 42, nil
			},
			expectedStatus: http.StatusOK,
			checkBody: func(t *testing.T, body map[string]any) {
				assert.NotEmpty(t, body["access_token"])
				assert.Equal(t, "read write", body["scope"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.authFunc != nil {
				env.auth.AuthenticateFunc = tt.authFunc
			} else {
				env.auth.AuthenticateFunc = func(ctx context.Context, username, password string) (uint64, error) {
					t.Fatal("Authenticate should not be called in this test case")
					return 0, nil
				}
			}

			req := httptest.NewRequest(http.MethodPost, "/oauth/token", strings.NewReader(tt.form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()

			env.svc.Token(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			if tt.checkBody != nil {
				var body map[string]any
				require.NoError(t, json.Unmarshal(w.Body.Bytes(), &body))
				tt.checkBody(t, body)
			}
		})
	}
}

func TestSvc_Token_NonPasswordGrant_MissingClient(t *testing.T) {
	env := newTestEnv(t)

	// A refresh_token grant without a valid client should fail.
	form := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {"some-token"},
		"client_id":     {"999"},
	}
	req := httptest.NewRequest(http.MethodPost, "/oauth/token", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	env.svc.Token(w, req)

	// fosite should write an error response (not 200).
	assert.NotEqual(t, http.StatusOK, w.Code)
}
