package fositestore

import (
	"context"
	"testing"

	"github.com/ory/fosite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStore_CreateAndGetRefreshTokenSession(t *testing.T) {
	store, db := newTestStore(t)
	ctx := context.Background()
	oc := createOauthClient(t, db, 1, false, true, false)
	fc := &FositeClient{oc}

	// Create the associated access token first.
	atSig := "at-sig-for-rt"
	atReq := newTestRequester(fc, "55", "read", "write")
	err := store.CreateAccessTokenSession(ctx, atSig, atReq)
	require.NoError(t, err)

	rtSig := "rt-sig-001"
	err = store.CreateRefreshTokenSession(ctx, rtSig, atSig, atReq)
	require.NoError(t, err)

	t.Run("access_token_id is stored correctly", func(t *testing.T) {
		rt, err := db.OauthRefreshToken.Get(ctx, rtSig)
		require.NoError(t, err)
		assert.Equal(t, atSig, rt.AccessTokenID)
		assert.False(t, rt.Revoked)
	})

	t.Run("GetRefreshTokenSession reconstructs session from JOIN", func(t *testing.T) {
		session := &fosite.DefaultSession{}
		retrieved, err := store.GetRefreshTokenSession(ctx, rtSig, session)
		require.NoError(t, err)
		assert.Equal(t, "55", retrieved.GetSession().GetSubject())
		assert.True(t, retrieved.GetGrantedScopes().Has("read"))
		assert.True(t, retrieved.GetGrantedScopes().Has("write"))
	})

	t.Run("GetRefreshTokenSession not found", func(t *testing.T) {
		_, err := store.GetRefreshTokenSession(ctx, "nonexistent-rt", &fosite.DefaultSession{})
		assert.ErrorIs(t, err, fosite.ErrNotFound)
	})

	t.Run("GetRefreshTokenSession revoked", func(t *testing.T) {
		// Revoke the refresh token.
		err := store.RevokeRefreshToken(ctx, rtSig)
		require.NoError(t, err)

		_, err = store.GetRefreshTokenSession(ctx, rtSig, &fosite.DefaultSession{})
		assert.ErrorIs(t, err, fosite.ErrTokenSignatureMismatch)
	})
}

func TestStore_DeleteRefreshTokenSession(t *testing.T) {
	store, db := newTestStore(t)
	ctx := context.Background()
	oc := createOauthClient(t, db, 1, false, true, false)
	fc := &FositeClient{oc}

	atReq := newTestRequester(fc, "1", "read")
	err := store.CreateAccessTokenSession(ctx, "at-del", atReq)
	require.NoError(t, err)

	err = store.CreateRefreshTokenSession(ctx, "rt-del", "at-del", atReq)
	require.NoError(t, err)

	err = store.DeleteRefreshTokenSession(ctx, "rt-del")
	require.NoError(t, err)

	count, err := db.OauthRefreshToken.Query().Count(ctx)
	require.NoError(t, err)
	assert.Equal(t, 0, count)
}

func TestStore_DeleteRefreshTokens(t *testing.T) {
	store, db := newTestStore(t)
	ctx := context.Background()
	oc := createOauthClient(t, db, 1, false, true, false)
	fc := &FositeClient{oc}

	// Create two access tokens and their refresh tokens.
	for i, atSig := range []string{"at-r1", "at-r2"} {
		req := newTestRequester(fc, "1", "read")
		err := store.CreateAccessTokenSession(ctx, atSig, req)
		require.NoError(t, err)
		err = store.CreateRefreshTokenSession(ctx, "rt-r"+string(rune('1'+i)), atSig, req)
		require.NoError(t, err)
	}

	t.Run("invalid client_id returns error", func(t *testing.T) {
		err := store.DeleteRefreshTokens(ctx, "not-a-number")
		assert.Error(t, err)
	})

	t.Run("success revokes all refresh tokens for client", func(t *testing.T) {
		err := store.DeleteRefreshTokens(ctx, "1")
		require.NoError(t, err)

		// Tokens should be deleted (DeleteRefreshTokens deletes, not revokes).
		count, err := db.OauthRefreshToken.Query().Count(ctx)
		require.NoError(t, err)
		assert.Equal(t, 0, count)
	})
}
