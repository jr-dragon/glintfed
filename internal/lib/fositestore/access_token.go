package fositestore

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ory/fosite"

	"glintfed.org/ent"
	"glintfed.org/ent/oauthaccesstoken"
)

// CreateAccessTokenSession stores a new access token session.
//
//	INSERT INTO oauth_access_tokens (id, user_id, client_id, name, scopes, revoked, expires_at) VALUES (...)
func (s *Store) CreateAccessTokenSession(ctx context.Context, signature string, req fosite.Requester) error {
	clientID, err := strconv.ParseUint(req.GetClient().GetID(), 10, 64)
	if err != nil {
		return fmt.Errorf("fositestore: invalid client_id %q: %w", req.GetClient().GetID(), err)
	}

	c := s.db.OauthAccessToken.Create().
		SetID(signature).
		SetClientID(clientID).
		SetScopes(marshalScopes(req.GetGrantedScopes())).
		SetRevoked(false)

	// subject is empty for client_credentials flow (no user); non-empty must be a valid uint64.
	if sub := req.GetSession().GetSubject(); sub != "" {
		userID, err := strconv.ParseUint(sub, 10, 64)
		if err != nil {
			return fmt.Errorf("fositestore: invalid subject %q: %w", sub, err)
		}
		c = c.SetUserID(userID)
	}

	if t := req.GetSession().GetExpiresAt(fosite.AccessToken); !t.IsZero() {
		c = c.SetExpiresAt(t)
	}

	_, err = c.Save(ctx)
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
	if t.Revoked {
		return nil, fosite.ErrTokenSignatureMismatch
	}

	client, err := s.GetClient(ctx, strconv.FormatUint(t.ClientID, 10))
	if err != nil {
		return nil, err
	}

	scopes := unmarshalScopes(t.Scopes)

	if ds, ok := session.(*fosite.DefaultSession); ok {
		ds.Subject = strconv.FormatUint(t.UserID, 10)
		if !t.ExpiresAt.IsZero() {
			ds.SetExpiresAt(fosite.AccessToken, t.ExpiresAt)
		}
	}

	req := fosite.NewRequest()
	req.SetID(signature)
	req.Client = client
	req.SetRequestedScopes(fosite.Arguments(scopes))
	for _, sc := range scopes {
		req.GrantScope(sc)
	}
	req.RequestedAt = t.CreatedAt
	req.Session = session
	return req, nil
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
	cid, err := strconv.ParseUint(clientID, 10, 64)
	if err != nil {
		return fmt.Errorf("fositestore: invalid client_id %q: %w", clientID, err)
	}
	_, err = s.db.OauthAccessToken.Delete().
		Where(oauthaccesstoken.ClientID(cid)).
		Exec(ctx)
	return err
}
