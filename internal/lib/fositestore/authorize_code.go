package fositestore

import (
	"context"
	"strconv"

	"github.com/ory/fosite"

	"glintfed.org/ent"
)

// CreateAuthorizeCodeSession stores a new authorization code session.
//
//	INSERT INTO oauth_auth_codes (id, user_id, client_id, scopes, revoked, expires_at) VALUES (...)
func (s *Store) CreateAuthorizeCodeSession(ctx context.Context, code string, req fosite.Requester) error {
	clientID, _ := strconv.ParseUint(req.GetClient().GetID(), 10, 64)
	userID, _ := strconv.ParseUint(req.GetSession().GetSubject(), 10, 64)

	c := s.db.OauthAuthorizationCode.Create().
		SetID(code).
		SetUserID(userID).
		SetClientID(clientID).
		SetScopes(marshalScopes(req.GetGrantedScopes())).
		SetRevoked(false)

	if t := req.GetSession().GetExpiresAt(fosite.AuthorizeCode); !t.IsZero() {
		c = c.SetExpiresAt(t)
	}

	_, err := c.Save(ctx)
	return err
}

// GetAuthorizeCodeSession retrieves and hydrates an authorization code session by code.
// Returns fosite.ErrInvalidatedAuthorizeCode if the code has been used (revoked=true).
//
//	SELECT * FROM oauth_auth_codes WHERE id = ? LIMIT 1
func (s *Store) GetAuthorizeCodeSession(ctx context.Context, code string, session fosite.Session) (fosite.Requester, error) {
	ac, err := s.db.OauthAuthorizationCode.Get(ctx, code)
	if ent.IsNotFound(err) {
		return nil, fosite.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	client, err := s.GetClient(ctx, strconv.FormatUint(ac.ClientID, 10))
	if err != nil {
		return nil, err
	}

	scopes := unmarshalScopes(ac.Scopes)

	if ds, ok := session.(*fosite.DefaultSession); ok {
		ds.Subject = strconv.FormatUint(ac.UserID, 10)
		if !ac.ExpiresAt.IsZero() {
			ds.SetExpiresAt(fosite.AuthorizeCode, ac.ExpiresAt)
		}
	}

	req := fosite.NewRequest()
	req.SetID(code)
	req.Client = client
	req.SetRequestedScopes(fosite.Arguments(scopes))
	for _, sc := range scopes {
		req.GrantScope(sc)
	}
	req.Session = session

	if ac.Revoked {
		return req, fosite.ErrInvalidatedAuthorizeCode
	}
	return req, nil
}

// InvalidateAuthorizeCodeSession marks an authorization code as used (revoked=true).
//
//	UPDATE oauth_auth_codes SET revoked = true WHERE id = ?
func (s *Store) InvalidateAuthorizeCodeSession(ctx context.Context, code string) error {
	return s.db.OauthAuthorizationCode.UpdateOneID(code).
		SetRevoked(true).
		Exec(ctx)
}
