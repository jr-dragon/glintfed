package fositestore

import (
	"context"
	"testing"

	"github.com/ory/fosite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"glintfed.org/ent"
)

func TestStore_CreateAccessTokenSession(t *testing.T) {
	store, db := newTestStore(t)
	ctx := context.Background()
	oc := createOauthClient(t, db, 1, false, true, false)
	fc := &FositeClient{oc}

	t.Run("success with user subject", func(t *testing.T) {
		req := newTestRequester(fc, "42", "read", "write")
		err := store.CreateAccessTokenSession(ctx, "sig-with-user", req)
		require.NoError(t, err)

		at, err := db.OauthAccessToken.Get(ctx, "sig-with-user")
		require.NoError(t, err)
		assert.Equal(t, uint64(1), at.ClientID)
		assert.Equal(t, uint64(42), at.UserID)
		assert.False(t, at.Revoked)
		assert.Equal(t, `["read","write"]`, at.Scopes)
	})

	t.Run("success without subject (client_credentials)", func(t *testing.T) {
		req := newTestRequester(fc, "", "read")
		err := store.CreateAccessTokenSession(ctx, "sig-no-user", req)
		require.NoError(t, err)

		at, err := db.OauthAccessToken.Get(ctx, "sig-no-user")
		require.NoError(t, err)
		assert.Equal(t, uint64(1), at.ClientID)
		assert.False(t, at.Revoked)
	})

	t.Run("invalid client_id returns error", func(t *testing.T) {
		req := fosite.NewRequest()
		req.Client = &badIDClient{id: "not-a-number"}
		req.Session = &fosite.DefaultSession{Subject: "1"}
		err := store.CreateAccessTokenSession(ctx, "sig-bad-client", req)
		assert.Error(t, err)
	})

	t.Run("invalid subject returns error", func(t *testing.T) {
		req := fosite.NewRequest()
		req.Client = fc
		req.Session = &fosite.DefaultSession{Subject: "not-a-number"}
		err := store.CreateAccessTokenSession(ctx, "sig-bad-sub", req)
		assert.Error(t, err)
	})
}

func TestStore_GetAccessTokenSession(t *testing.T) {
	store, db := newTestStore(t)
	ctx := context.Background()
	oc := createOauthClient(t, db, 1, false, true, false)
	fc := &FositeClient{oc}

	req := newTestRequester(fc, "99", "read", "write")
	err := store.CreateAccessTokenSession(ctx, "valid-sig", req)
	require.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		session := &fosite.DefaultSession{}
		retrieved, err := store.GetAccessTokenSession(ctx, "valid-sig", session)
		require.NoError(t, err)
		assert.Equal(t, "99", retrieved.GetSession().GetSubject())
		assert.True(t, retrieved.GetGrantedScopes().Has("read"))
		assert.True(t, retrieved.GetGrantedScopes().Has("write"))
	})

	t.Run("not found", func(t *testing.T) {
		_, err := store.GetAccessTokenSession(ctx, "nonexistent-sig", &fosite.DefaultSession{})
		assert.ErrorIs(t, err, fosite.ErrNotFound)
	})

	t.Run("revoked returns error", func(t *testing.T) {
		_, err := db.OauthAccessToken.Create().
			SetID("revoked-sig").
			SetClientID(1).
			SetRevoked(true).
			Save(ctx)
		require.NoError(t, err)

		_, err = store.GetAccessTokenSession(ctx, "revoked-sig", &fosite.DefaultSession{})
		assert.ErrorIs(t, err, fosite.ErrTokenSignatureMismatch)
	})
}

func TestStore_DeleteAccessTokenSession(t *testing.T) {
	store, db := newTestStore(t)
	ctx := context.Background()
	oc := createOauthClient(t, db, 1, false, true, false)
	fc := &FositeClient{oc}

	err := store.CreateAccessTokenSession(ctx, "to-delete", newTestRequester(fc, "1", "read"))
	require.NoError(t, err)

	err = store.DeleteAccessTokenSession(ctx, "to-delete")
	require.NoError(t, err)

	_, err = db.OauthAccessToken.Get(ctx, "to-delete")
	assert.True(t, ent.IsNotFound(err))
}

func TestStore_DeleteAccessTokens(t *testing.T) {
	store, db := newTestStore(t)
	ctx := context.Background()
	oc := createOauthClient(t, db, 1, false, true, false)
	fc := &FositeClient{oc}

	for _, sig := range []string{"tok-a", "tok-b"} {
		err := store.CreateAccessTokenSession(ctx, sig, newTestRequester(fc, "1", "read"))
		require.NoError(t, err)
	}

	t.Run("invalid client_id returns error", func(t *testing.T) {
		err := store.DeleteAccessTokens(ctx, "not-a-number")
		assert.Error(t, err)
	})

	t.Run("success removes all tokens for client", func(t *testing.T) {
		err := store.DeleteAccessTokens(ctx, "1")
		require.NoError(t, err)

		count, err := db.OauthAccessToken.Query().Count(ctx)
		require.NoError(t, err)
		assert.Equal(t, 0, count)
	})
}

// badIDClient is a minimal fosite.Client stub that returns a non-numeric ID.
type badIDClient struct{ id string }

func (b *badIDClient) GetID() string                      { return b.id }
func (b *badIDClient) GetHashedSecret() []byte            { return nil }
func (b *badIDClient) GetRedirectURIs() []string          { return nil }
func (b *badIDClient) GetGrantTypes() fosite.Arguments    { return nil }
func (b *badIDClient) GetResponseTypes() fosite.Arguments { return nil }
func (b *badIDClient) GetScopes() fosite.Arguments        { return nil }
func (b *badIDClient) IsPublic() bool                     { return false }
func (b *badIDClient) GetAudience() fosite.Arguments      { return nil }
