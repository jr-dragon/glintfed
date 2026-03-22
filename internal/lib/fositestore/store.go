package fositestore

import (
	"context"
	"time"

	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"
	"github.com/ory/fosite/handler/oauth2"

	"glintfed.org/ent"
	"glintfed.org/internal/data"
)

// Store implements fosite storage interfaces backed by ent ORM.
type Store struct {
	db       *ent.Client
	strategy *oauth2.HMACSHAStrategy
}

// New creates a new Store with an HMAC strategy derived from the configuration.
func New(client *data.Client, cfg *data.Config) *Store {
	globalSecret := []byte(cfg.App.Auth.OAuth.HMACSecret)
	fositeCfg := &fosite.Config{GlobalSecret: globalSecret}
	strategy := compose.NewOAuth2HMACStrategy(fositeCfg)
	return &Store{
		db:       client.Ent,
		strategy: strategy,
	}
}

// Strategy returns the underlying HMAC strategy, used by the OAuth2Provider.
func (s *Store) Strategy() *oauth2.HMACSHAStrategy {
	return s.strategy
}

// ClientAssertionJWTValid — JWT assertion JTI replay protection (not implemented; always allows).
func (s *Store) ClientAssertionJWTValid(_ context.Context, _ string) error {
	return nil
}

// SetClientAssertionJWT — stores JTI (no-op).
func (s *Store) SetClientAssertionJWT(_ context.Context, _ string, _ time.Time) error {
	return nil
}

// RevokeAccessToken revokes an access token by its token signature (id).
//
//	UPDATE oauth_access_tokens SET revoked = true WHERE id = ?
func (s *Store) RevokeAccessToken(ctx context.Context, requestID string) error {
	return s.db.OauthAccessToken.UpdateOneID(requestID).SetRevoked(true).Exec(ctx)
}

// RevokeRefreshToken revokes a refresh token by its token signature (id).
//
//	UPDATE oauth_refresh_tokens SET revoked = true WHERE id = ?
func (s *Store) RevokeRefreshToken(ctx context.Context, requestID string) error {
	return s.db.OauthRefreshToken.UpdateOneID(requestID).SetRevoked(true).Exec(ctx)
}

// RevokeRefreshTokenMaybeGracePeriod is the same as RevokeRefreshToken (no grace period implemented).
func (s *Store) RevokeRefreshTokenMaybeGracePeriod(ctx context.Context, requestID string, _ string) error {
	return s.RevokeRefreshToken(ctx, requestID)
}

// RotateRefreshToken revokes the old refresh token by signature.
//
//	UPDATE oauth_refresh_tokens SET revoked = true WHERE id = ?
func (s *Store) RotateRefreshToken(ctx context.Context, _ string, refreshTokenSignature string) error {
	return s.db.OauthRefreshToken.UpdateOneID(refreshTokenSignature).SetRevoked(true).Exec(ctx)
}

// CreatePersonalAccessTokens directly issues an access token and refresh token for the given requester
// using the HMAC strategy without going through the full OAuth2 authorization flow.
// It sets req.ID to the access token signature so that CreateRefreshTokenSession can use it as access_token_id.
func (s *Store) CreatePersonalAccessTokens(ctx context.Context, req fosite.Requester) (accessToken, refreshToken string, err error) {
	at, atSig, err := s.strategy.GenerateAccessToken(ctx, req)
	if err != nil {
		return "", "", err
	}
	rt, rtSig, err := s.strategy.GenerateRefreshToken(ctx, req)
	if err != nil {
		return "", "", err
	}

	if err = s.CreateAccessTokenSession(ctx, atSig, req); err != nil {
		return "", "", err
	}
	// Set req ID to atSig so CreateRefreshTokenSession can use req.GetID() as access_token_id.
	req.SetID(atSig)
	if err = s.CreateRefreshTokenSession(ctx, rtSig, req); err != nil {
		return "", "", err
	}
	return at, rt, nil
}

// Authenticate verifies a username/password pair (used by ROPC grant).
// This is intentionally not implemented — password grant is handled at the service layer.
func (s *Store) Authenticate(_ context.Context, _, _ string) (string, error) {
	return "", fosite.ErrNotFound
}
