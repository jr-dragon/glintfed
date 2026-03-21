package oauth

import (
	"context"

	"glintfed.org/internal/lib/errs"
)

// TokenResult contains the OAuth token details issued after successful authentication.
type TokenResult struct {
	AccessToken  string
	RefreshToken string
	ClientID     string
	ClientSecret string
	ExpiresIn    int64
}

// CreateTokens creates an OAuth access token and a refresh token for the given user ID
// with the specified scopes, and returns the resulting token details.
func CreateTokens(ctx context.Context, userID uint64, scopes []string) (*TokenResult, error) {
	return nil, errs.Todo
}
