package fositestore

import (
	"context"

	"github.com/ory/fosite"

	"glintfed.org/ent"
)

// CreateAuthorizeCodeSession stores a new authorization code session.
//
//	INSERT INTO oauth_authorization_codes (id, request_id, client_id, subject, scopes, session, active, requested_at, expires_at)
//	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
func (s *Store) CreateAuthorizeCodeSession(ctx context.Context, code string, req fosite.Requester) error {
	sessionBytes, err := marshalSession(req.GetSession())
	if err != nil {
		return err
	}
	_, err = s.db.OauthAuthorizationCode.Create().
		SetID(code).
		SetRequestID(req.GetID()).
		SetClientID(req.GetClient().GetID()).
		SetSubject(req.GetSession().GetSubject()).
		SetScopes(req.GetGrantedScopes()).
		SetSession(sessionBytes).
		SetActive(true).
		SetRequestedAt(req.GetRequestedAt()).
		SetExpiresAt(req.GetSession().GetExpiresAt(fosite.AuthorizeCode)).
		Save(ctx)
	return err
}

// GetAuthorizeCodeSession retrieves and hydrates an authorization code session by code.
// Returns fosite.ErrInvalidatedAuthorizeCode if the code has been used (active=false).
//
//	SELECT * FROM oauth_authorization_codes WHERE id = ? LIMIT 1
func (s *Store) GetAuthorizeCodeSession(ctx context.Context, code string, session fosite.Session) (fosite.Requester, error) {
	t, err := s.db.OauthAuthorizationCode.Get(ctx, code)
	if ent.IsNotFound(err) {
		return nil, fosite.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	client, err := s.GetClient(ctx, t.ClientID)
	if err != nil {
		return nil, err
	}
	req, err := requesterFromRow(t.RequestID, t.ClientID, t.Subject, t.Scopes, t.Session, t.RequestedAt, t.ExpiresAt, client, session)
	if err != nil {
		return nil, err
	}
	if !t.Active {
		return req, fosite.ErrInvalidatedAuthorizeCode
	}
	return req, nil
}

// InvalidateAuthorizeCodeSession marks an authorization code as used (active=false).
//
//	UPDATE oauth_authorization_codes SET active = false WHERE id = ?
func (s *Store) InvalidateAuthorizeCodeSession(ctx context.Context, code string) error {
	return s.db.OauthAuthorizationCode.UpdateOneID(code).
		SetActive(false).
		Exec(ctx)
}
