package fositestore

import (
	"context"

	"github.com/ory/fosite"

	"glintfed.org/ent"
	"glintfed.org/ent/oauthaccesstoken"
)

// CreateAccessTokenSession stores a new access token session.
//
//	INSERT INTO oauth_access_tokens (id, request_id, client_id, subject, scopes, session, active, requested_at, expires_at)
//	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
func (s *Store) CreateAccessTokenSession(ctx context.Context, signature string, req fosite.Requester) error {
	sessionBytes, err := marshalSession(req.GetSession())
	if err != nil {
		return err
	}
	_, err = s.db.OauthAccessToken.Create().
		SetID(signature).
		SetRequestID(req.GetID()).
		SetClientID(req.GetClient().GetID()).
		SetSubject(req.GetSession().GetSubject()).
		SetScopes(req.GetGrantedScopes()).
		SetSession(sessionBytes).
		SetActive(true).
		SetRequestedAt(req.GetRequestedAt()).
		SetExpiresAt(req.GetSession().GetExpiresAt(fosite.AccessToken)).
		Save(ctx)
	return err
}

// GetAccessTokenSession retrieves and hydrates an access token session by signature.
//
//	SELECT * FROM oauth_access_tokens WHERE id = ? LIMIT 1
func (s *Store) GetAccessTokenSession(ctx context.Context, signature string, session fosite.Session) (fosite.Requester, error) {
	t, err := s.db.OauthAccessToken.Get(ctx, signature)
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

// DeleteAccessTokenSession removes an access token session by signature.
//
//	DELETE FROM oauth_access_tokens WHERE id = ?
func (s *Store) DeleteAccessTokenSession(ctx context.Context, signature string) error {
	return s.db.OauthAccessToken.DeleteOneID(signature).Exec(ctx)
}

// DeleteAccessTokens removes all access tokens for a given client ID.
//
//	DELETE FROM oauth_access_tokens WHERE client_id = ?
func (s *Store) DeleteAccessTokens(ctx context.Context, clientID string) error {
	_, err := s.db.OauthAccessToken.Delete().
		Where(oauthaccesstoken.ClientID(clientID)).
		Exec(ctx)
	return err
}
