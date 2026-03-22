package fositestore

import (
	"context"
	"encoding/json"

	"github.com/ory/fosite"

	"glintfed.org/ent"
)

// CreatePKCERequestSession stores a new PKCE session.
//
//	INSERT INTO oauth_pkce (id, request_id, client_id, subject, scopes, session, active, requested_at, expires_at)
//	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
func (s *Store) CreatePKCERequestSession(ctx context.Context, signature string, req fosite.Requester) error {
	sessionBytes, err := json.Marshal(req.GetSession())
	if err != nil {
		return err
	}
	_, err = s.db.OauthPkce.Create().
		SetID(signature).
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

// GetPKCERequestSession retrieves and hydrates a PKCE session by signature.
//
//	SELECT * FROM oauth_pkce WHERE id = ? LIMIT 1
func (s *Store) GetPKCERequestSession(ctx context.Context, signature string, session fosite.Session) (fosite.Requester, error) {
	t, err := s.db.OauthPkce.Get(ctx, signature)
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

	if err = json.Unmarshal(t.Session, session); err != nil {
		return nil, err
	}

	r := fosite.NewRequest()
	r.ID = t.RequestID
	r.Client = client
	r.RequestedAt = t.RequestedAt
	r.Session = session
	r.SetRequestedScopes(fosite.Arguments(t.Scopes))
	for _, sc := range t.Scopes {
		r.GrantScope(sc)
	}
	return r, nil
}

// DeletePKCERequestSession removes a PKCE session by signature.
//
//	DELETE FROM oauth_pkce WHERE id = ?
func (s *Store) DeletePKCERequestSession(ctx context.Context, signature string) error {
	return s.db.OauthPkce.DeleteOneID(signature).Exec(ctx)
}
