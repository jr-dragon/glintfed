package oauth

import (
	"context"
	"net/http"

	"github.com/ory/fosite"

	"glintfed.org/internal/data"
	"glintfed.org/internal/lib/fositestore"
)

// Service defines the OAuth2 HTTP handlers.
type Service interface {
	Authorize(w http.ResponseWriter, r *http.Request)
	Token(w http.ResponseWriter, r *http.Request)
	Revoke(w http.ResponseWriter, r *http.Request)
}

//go:generate go tool moq -rm -out mock_user_authenticator.go . UserAuthenticator
type UserAuthenticator interface {
	// Authenticate verifies username/password and returns the user ID.
	Authenticate(ctx context.Context, username, password string) (uint64, error)
}

type svc struct {
	provider fosite.OAuth2Provider
	store    *fositestore.Store
	auth     UserAuthenticator
	appURL   string
}

// New creates a new OAuth service.
func New(provider fosite.OAuth2Provider, store *fositestore.Store, auth UserAuthenticator, cfg *data.Config) Service {
	return &svc{
		provider: provider,
		store:    store,
		auth:     auth,
		appURL:   cfg.App.Url,
	}
}
