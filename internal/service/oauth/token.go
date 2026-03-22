package oauth

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/ory/fosite"

	"glintfed.org/internal/lib/logs"
	"glintfed.org/internal/service/internal"
)

// Token handles POST /oauth/token for all grant types.
func (s *svc) Token(w http.ResponseWriter, r *http.Request) {
	ctx, span := internal.T.Start(r.Context(), "OAuth.Token")
	defer span.End()

	// Handle password grant type manually to avoid needing the ROPC grant handler.
	// This is intentional: we validate credentials ourselves and issue tokens directly.
	if r.FormValue("grant_type") == "password" {
		s.handlePasswordGrant(w, r.WithContext(ctx))
		return
	}

	accessReq, err := s.provider.NewAccessRequest(ctx, r, &fosite.DefaultSession{})
	if err != nil {
		s.provider.WriteAccessError(ctx, w, accessReq, err)
		return
	}

	accessResp, err := s.provider.NewAccessResponse(ctx, accessReq)
	if err != nil {
		s.provider.WriteAccessError(ctx, w, accessReq, err)
		return
	}

	s.provider.WriteAccessResponse(ctx, w, accessReq, accessResp)
}

// handlePasswordGrant processes the resource owner password credentials grant type.
// Credentials are validated manually; then tokens are issued via fositestore.
func (s *svc) handlePasswordGrant(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		writeTokenError(w, "invalid_request", "username and password are required", http.StatusBadRequest)
		return
	}

	userID, err := s.auth.Authenticate(ctx, username, password)
	if err != nil {
		slog.ErrorContext(ctx, "password grant: authentication failed", logs.ErrAttr(err))
		writeTokenError(w, "invalid_grant", "invalid username or password", http.StatusUnauthorized)
		return
	}

	// Parse requested scopes from form; fall back to default scopes.
	scopeStr := r.FormValue("scope")
	scopes := []string{"read", "write", "follow", "push"}
	if scopeStr != "" {
		scopes = splitScopes(scopeStr)
	}

	// Look up the client to validate client_id.
	clientID := r.FormValue("client_id")
	if clientID == "" {
		writeTokenError(w, "invalid_client", "client_id is required", http.StatusBadRequest)
		return
	}

	client, err := s.store.GetClient(ctx, clientID)
	if err != nil {
		writeTokenError(w, "invalid_client", "unknown client", http.StatusUnauthorized)
		return
	}

	subject := fmt.Sprintf("%d", userID)
	now := time.Now()

	session := &fosite.DefaultSession{
		Subject:  subject,
		Username: username,
		ExpiresAt: map[fosite.TokenType]time.Time{
			fosite.AccessToken:  now.Add(s.accessTokenTTL),
			fosite.RefreshToken: now.Add(s.refreshTokenTTL),
		},
	}

	req := fosite.NewRequest()
	req.ID = tokenRequestID()
	req.Client = client
	req.RequestedAt = now
	req.Session = session
	req.SetRequestedScopes(fosite.Arguments(scopes))
	for _, scope := range scopes {
		req.GrantScope(scope)
	}

	accessToken, refreshToken, err := s.store.CreatePersonalAccessTokens(ctx, req)
	if err != nil {
		slog.ErrorContext(ctx, "password grant: failed to create tokens", logs.ErrAttr(err))
		writeTokenError(w, "server_error", "failed to issue tokens", http.StatusInternalServerError)
		return
	}

	resp := map[string]any{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"token_type":    "Bearer",
		"expires_in":    int64(s.accessTokenTTL.Seconds()),
		"scope":         scopeStr,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.ErrorContext(ctx, "password grant: failed to encode response", logs.ErrAttr(err))
	}
}

func writeTokenError(w http.ResponseWriter, code, description string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"error":             code,
		"error_description": description,
	})
}

func splitScopes(s string) []string {
	if s == "" {
		return nil
	}
	var scopes []string
	start := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			if i > start {
				scopes = append(scopes, s[start:i])
			}
			start = i + 1
		}
	}
	return scopes
}

func tokenRequestID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)
}
