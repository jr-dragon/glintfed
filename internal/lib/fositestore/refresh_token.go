package fositestore

import (
	"context"

	"github.com/ory/fosite"

	"glintfed.org/ent"
	"glintfed.org/ent/oauthrefreshtoken"
)

// CreateRefreshTokenSession stores a new refresh token session.
//
//	INSERT INTO oauth_refresh_tokens (id, request_id, client_id, subject, scopes, session, active, requested_at, expires_at)
//	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
func (s *Store) CreateRefreshTokenSession(ctx context.Context, signature string, _ string, req fosite.Requester) error {
	sessionBytes, err := marshalSession(req.GetSession())
	if err != nil {
		return err
	}
	_, err = s.db.OauthRefreshToken.Create().
		SetID(signature).
		SetRequestID(req.GetID()).
		SetClientID(req.GetClient().GetID()).
		SetSubject(req.GetSession().GetSubject()).
		SetScopes(req.GetGrantedScopes()).
		SetSession(sessionBytes).
		SetActive(true).
		SetRequestedAt(req.GetRequestedAt()).
		SetExpiresAt(req.GetSession().GetExpiresAt(fosite.RefreshToken)).
		Save(ctx)
	return err
}

// GetRefreshTokenSession retrieves and hydrates a refresh token session by signature.
//
//	SELECT * FROM oauth_refresh_tokens WHERE id = ? LIMIT 1
func (s *Store) GetRefreshTokenSession(ctx context.Context, signature string, session fosite.Session) (fosite.Requester, error) {
	t, err := s.db.OauthRefreshToken.Get(ctx, signature)
	if ent.IsNotFound(err) {
		return nil, fosite.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	if !t.Active {
		return nil, fosite.ErrTokenSignatureMismatch
	}
	client, err := s.GetClient(ctx, t.ClientID)
	if err != nil {
		return nil, err
	}
	return requesterFromRow(t.RequestID, t.ClientID, t.Subject, t.Scopes, t.Session, t.RequestedAt, t.ExpiresAt, client, session)
}

// DeleteRefreshTokenSession removes a refresh token session by signature.
//
//	DELETE FROM oauth_refresh_tokens WHERE id = ?
func (s *Store) DeleteRefreshTokenSession(ctx context.Context, signature string) error {
	return s.db.OauthRefreshToken.DeleteOneID(signature).Exec(ctx)
}

// DeleteRefreshTokens removes all refresh tokens for a given client ID.
//
//	DELETE FROM oauth_refresh_tokens WHERE client_id = ?
func (s *Store) DeleteRefreshTokens(ctx context.Context, clientID string) error {
	_, err := s.db.OauthRefreshToken.Delete().
		Where(oauthrefreshtoken.ClientID(clientID)).
		Exec(ctx)
	return err
}
