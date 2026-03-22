package fositestore

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ory/fosite"

	"glintfed.org/ent"
	"glintfed.org/ent/oauthaccesstoken"
	"glintfed.org/ent/oauthrefreshtoken"
)

// CreateRefreshTokenSession stores a new refresh token session.
// req.GetID() is used as access_token_id; callers must set req.SetID(atSignature) before calling.
//
//	INSERT INTO oauth_refresh_tokens (id, access_token_id, revoked, expires_at) VALUES (...)
func (s *Store) CreateRefreshTokenSession(ctx context.Context, signature string, req fosite.Requester) error {
	c := s.db.OauthRefreshToken.Create().
		SetID(signature).
		SetAccessTokenID(req.GetID()).
		SetRevoked(false)

	if t := req.GetSession().GetExpiresAt(fosite.RefreshToken); !t.IsZero() {
		c = c.SetExpiresAt(t)
	}

	_, err := c.Save(ctx)
	return err
}

// GetRefreshTokenSession retrieves and hydrates a refresh token session by signature.
// It JOINs the access token table to recover user/client/scopes.
//
//	SELECT rt.*, at.user_id, at.client_id, at.scopes FROM oauth_refresh_tokens rt
//	JOIN oauth_access_tokens at ON rt.access_token_id = at.id
//	WHERE rt.id = ? LIMIT 1
func (s *Store) GetRefreshTokenSession(ctx context.Context, signature string, session fosite.Session) (fosite.Requester, error) {
	rt, err := s.db.OauthRefreshToken.Get(ctx, signature)
	if ent.IsNotFound(err) {
		return nil, fosite.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	if rt.Revoked {
		return nil, fosite.ErrTokenSignatureMismatch
	}

	at, err := s.db.OauthAccessToken.Get(ctx, rt.AccessTokenID)
	if ent.IsNotFound(err) {
		return nil, fosite.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	client, err := s.GetClient(ctx, strconv.FormatUint(at.ClientID, 10))
	if err != nil {
		return nil, err
	}

	scopes := unmarshalScopes(at.Scopes)

	if ds, ok := session.(*fosite.DefaultSession); ok {
		ds.Subject = strconv.FormatUint(at.UserID, 10)
		if !rt.ExpiresAt.IsZero() {
			ds.SetExpiresAt(fosite.RefreshToken, rt.ExpiresAt)
		}
	}

	req := fosite.NewRequest()
	req.SetID(signature)
	req.Client = client
	req.SetRequestedScopes(fosite.Arguments(scopes))
	for _, sc := range scopes {
		req.GrantScope(sc)
	}
	req.Session = session
	return req, nil
}

// DeleteRefreshTokenSession removes a refresh token session by signature.
//
//	DELETE FROM oauth_refresh_tokens WHERE id = ?
func (s *Store) DeleteRefreshTokenSession(ctx context.Context, signature string) error {
	return s.db.OauthRefreshToken.DeleteOneID(signature).Exec(ctx)
}

// DeleteRefreshTokens revokes all refresh tokens associated with a given client ID
// by looking up the client's access tokens first.
//
//	UPDATE oauth_refresh_tokens SET revoked = true WHERE access_token_id IN (
//	  SELECT id FROM oauth_access_tokens WHERE client_id = ?
//	)
func (s *Store) DeleteRefreshTokens(ctx context.Context, clientID string) error {
	cid, err := strconv.ParseUint(clientID, 10, 64)
	if err != nil {
		return fmt.Errorf("fositestore: invalid client_id %q: %w", clientID, err)
	}
	ats, err := s.db.OauthAccessToken.Query().
		Where(oauthaccesstoken.ClientID(cid)).
		Select(oauthaccesstoken.FieldID).
		All(ctx)
	if err != nil {
		return err
	}
	if len(ats) == 0 {
		return nil
	}
	atIDs := make([]string, len(ats))
	for i, at := range ats {
		atIDs[i] = at.ID
	}
	_, err = s.db.OauthRefreshToken.Delete().
		Where(oauthrefreshtoken.AccessTokenIDIn(atIDs...)).
		Exec(ctx)
	return err
}
