package fositestore

import (
	"context"
	"testing"

	"github.com/ory/fosite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStore_CreateAuthorizeCodeSession(t *testing.T) {
	store, db := newTestStore(t)
	ctx := context.Background()
	oc := createOauthClient(t, db, 1, false, true, false)
	fc := &FositeClient{oc}

	t.Run("success", func(t *testing.T) {
		req := newTestRequester(fc, "42", "read", "write")
		err := store.CreateAuthorizeCodeSession(ctx, "auth-code-sig", req)
		require.NoError(t, err)

		ac, err := db.OauthAuthorizationCode.Get(ctx, "auth-code-sig")
		require.NoError(t, err)
		assert.Equal(t, uint64(1), ac.ClientID)
		assert.Equal(t, uint64(42), ac.UserID)
		assert.False(t, ac.Revoked)
	})

	t.Run("invalid client_id returns error", func(t *testing.T) {
		req := fosite.NewRequest()
		req.Client = &badIDClient{id: "not-a-number"}
		req.Session = &fosite.DefaultSession{Subject: "1"}
		err := store.CreateAuthorizeCodeSession(ctx, "bad-client-code", req)
		assert.Error(t, err)
	})

	t.Run("invalid subject (user_id NOT NULL) returns error", func(t *testing.T) {
		req := fosite.NewRequest()
		req.Client = fc
		req.Session = &fosite.DefaultSession{Subject: "not-a-number"}
		err := store.CreateAuthorizeCodeSession(ctx, "bad-subject-code", req)
		assert.Error(t, err)
	})
}

func TestStore_GetAuthorizeCodeSession(t *testing.T) {
	store, db := newTestStore(t)
	ctx := context.Background()
	oc := createOauthClient(t, db, 1, false, true, false)
	fc := &FositeClient{oc}

	req := newTestRequester(fc, "42", "read")
	err := store.CreateAuthorizeCodeSession(ctx, "valid-code", req)
	require.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		session := &fosite.DefaultSession{}
		retrieved, err := store.GetAuthorizeCodeSession(ctx, "valid-code", session)
		require.NoError(t, err)
		assert.Equal(t, "42", retrieved.GetSession().GetSubject())
		assert.True(t, retrieved.GetGrantedScopes().Has("read"))
	})

	t.Run("not found", func(t *testing.T) {
		_, err := store.GetAuthorizeCodeSession(ctx, "nonexistent", &fosite.DefaultSession{})
		assert.ErrorIs(t, err, fosite.ErrNotFound)
	})

	t.Run("revoked returns ErrInvalidatedAuthorizeCode", func(t *testing.T) {
		revReq := newTestRequester(fc, "42", "read")
		err := store.CreateAuthorizeCodeSession(ctx, "revoked-code", revReq)
		require.NoError(t, err)

		err = store.InvalidateAuthorizeCodeSession(ctx, "revoked-code")
		require.NoError(t, err)

		_, err = store.GetAuthorizeCodeSession(ctx, "revoked-code", &fosite.DefaultSession{})
		assert.ErrorIs(t, err, fosite.ErrInvalidatedAuthorizeCode)
	})
}

func TestStore_InvalidateAuthorizeCodeSession(t *testing.T) {
	store, db := newTestStore(t)
	ctx := context.Background()
	oc := createOauthClient(t, db, 1, false, true, false)
	fc := &FositeClient{oc}

	err := store.CreateAuthorizeCodeSession(ctx, "code-to-invalidate", newTestRequester(fc, "1", "read"))
	require.NoError(t, err)

	err = store.InvalidateAuthorizeCodeSession(ctx, "code-to-invalidate")
	require.NoError(t, err)

	ac, err := db.OauthAuthorizationCode.Get(ctx, "code-to-invalidate")
	require.NoError(t, err)
	assert.True(t, ac.Revoked)
}
