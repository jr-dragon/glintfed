package fositestore

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStore_CreatePersonalAccessTokens(t *testing.T) {
	store, db := newTestStore(t)
	ctx := context.Background()
	oc := createOauthClient(t, db, 1, true, false, false)
	fc := &FositeClient{oc}

	req := newTestRequester(fc, "99", "read", "write")

	at, rt, err := store.CreatePersonalAccessTokens(ctx, req)
	require.NoError(t, err)
	assert.NotEmpty(t, at)
	assert.NotEmpty(t, rt)

	// CreatePersonalAccessTokens sets req.ID to the access token signature.
	atSig := req.GetID()
	assert.NotEmpty(t, atSig)

	// Verify access token persisted.
	atRecord, err := db.OauthAccessToken.Get(ctx, atSig)
	require.NoError(t, err)
	assert.Equal(t, uint64(1), atRecord.ClientID)
	assert.Equal(t, uint64(99), atRecord.UserID)
	assert.False(t, atRecord.Revoked)

	// Verify refresh token persisted with correct access_token_id.
	rtRecords, err := db.OauthRefreshToken.Query().All(ctx)
	require.NoError(t, err)
	require.Len(t, rtRecords, 1)
	assert.Equal(t, atSig, rtRecords[0].AccessTokenID)
	assert.False(t, rtRecords[0].Revoked)
}

func TestStore_RevokeAccessToken(t *testing.T) {
	store, db := newTestStore(t)
	ctx := context.Background()
	oc := createOauthClient(t, db, 1, false, true, false)
	fc := &FositeClient{oc}

	err := store.CreateAccessTokenSession(ctx, "at-to-revoke", newTestRequester(fc, "1", "read"))
	require.NoError(t, err)

	err = store.RevokeAccessToken(ctx, "at-to-revoke")
	require.NoError(t, err)

	at, err := db.OauthAccessToken.Get(ctx, "at-to-revoke")
	require.NoError(t, err)
	assert.True(t, at.Revoked)
}

func TestStore_RevokeRefreshToken(t *testing.T) {
	store, db := newTestStore(t)
	ctx := context.Background()
	oc := createOauthClient(t, db, 1, false, true, false)
	fc := &FositeClient{oc}

	atReq := newTestRequester(fc, "1", "read")
	err := store.CreateAccessTokenSession(ctx, "at-for-rt-revoke", atReq)
	require.NoError(t, err)
	err = store.CreateRefreshTokenSession(ctx, "rt-to-revoke", "at-for-rt-revoke", atReq)
	require.NoError(t, err)

	err = store.RevokeRefreshToken(ctx, "rt-to-revoke")
	require.NoError(t, err)

	rt, err := db.OauthRefreshToken.Get(ctx, "rt-to-revoke")
	require.NoError(t, err)
	assert.True(t, rt.Revoked)
}

func TestStore_RotateRefreshToken(t *testing.T) {
	store, db := newTestStore(t)
	ctx := context.Background()
	oc := createOauthClient(t, db, 1, false, true, false)
	fc := &FositeClient{oc}

	atReq := newTestRequester(fc, "1", "read")
	err := store.CreateAccessTokenSession(ctx, "at-rotate", atReq)
	require.NoError(t, err)
	err = store.CreateRefreshTokenSession(ctx, "rt-old", "at-rotate", atReq)
	require.NoError(t, err)

	err = store.RotateRefreshToken(ctx, "", "rt-old")
	require.NoError(t, err)

	rt, err := db.OauthRefreshToken.Get(ctx, "rt-old")
	require.NoError(t, err)
	assert.True(t, rt.Revoked)
}
