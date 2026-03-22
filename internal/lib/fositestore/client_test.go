package fositestore

import (
	"context"
	"testing"

	"github.com/ory/fosite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"glintfed.org/ent"
)

func TestStore_GetClient(t *testing.T) {
	store, db := newTestStore(t)
	ctx := context.Background()

	createOauthClient(t, db, 1, false, true, false) // active
	createOauthClient(t, db, 2, false, false, true) // revoked
	createOauthClient(t, db, 3, true, false, false) // personal access

	t.Run("success", func(t *testing.T) {
		client, err := store.GetClient(ctx, "1")
		require.NoError(t, err)
		assert.Equal(t, "1", client.GetID())
	})

	t.Run("not found", func(t *testing.T) {
		_, err := store.GetClient(ctx, "999")
		assert.ErrorIs(t, err, fosite.ErrNotFound)
	})

	t.Run("invalid id returns not found", func(t *testing.T) {
		_, err := store.GetClient(ctx, "not-a-number")
		assert.Error(t, err)
	})

	t.Run("revoked client returns ErrInvalidClient", func(t *testing.T) {
		_, err := store.GetClient(ctx, "2")
		assert.ErrorIs(t, err, fosite.ErrInvalidClient)
	})
}

func TestFositeClient_GetID(t *testing.T) {
	fc := &FositeClient{&ent.OauthClient{}}
	fc.ID = 42
	assert.Equal(t, "42", fc.GetID())
}

func TestFositeClient_GetHashedSecret(t *testing.T) {
	fc := &FositeClient{&ent.OauthClient{Secret: "mysecret"}}
	assert.Equal(t, []byte("mysecret"), fc.GetHashedSecret())
}

func TestFositeClient_GetRedirectURIs(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		fc := &FositeClient{&ent.OauthClient{Redirect: ""}}
		assert.Nil(t, fc.GetRedirectURIs())
	})

	t.Run("single URI", func(t *testing.T) {
		fc := &FositeClient{&ent.OauthClient{Redirect: "https://example.com"}}
		assert.Equal(t, []string{"https://example.com"}, fc.GetRedirectURIs())
	})

	t.Run("multiple URIs separated by newline", func(t *testing.T) {
		fc := &FositeClient{&ent.OauthClient{Redirect: "https://a.com\nhttps://b.com"}}
		assert.Equal(t, []string{"https://a.com", "https://b.com"}, fc.GetRedirectURIs())
	})
}

func TestFositeClient_GetGrantTypes(t *testing.T) {
	t.Run("password_client=true includes password grant", func(t *testing.T) {
		fc := &FositeClient{&ent.OauthClient{PasswordClient: true}}
		types := fc.GetGrantTypes()
		assert.True(t, types.Has("password"))
		assert.True(t, types.Has("authorization_code"))
		assert.True(t, types.Has("refresh_token"))
	})

	t.Run("password_client=false excludes password grant", func(t *testing.T) {
		fc := &FositeClient{&ent.OauthClient{PasswordClient: false}}
		types := fc.GetGrantTypes()
		assert.False(t, types.Has("password"))
		assert.True(t, types.Has("authorization_code"))
	})
}

func TestFositeClient_IsPublic(t *testing.T) {
	t.Run("personal_access_client=true is public", func(t *testing.T) {
		fc := &FositeClient{&ent.OauthClient{PersonalAccessClient: true}}
		assert.True(t, fc.IsPublic())
	})

	t.Run("personal_access_client=false is confidential", func(t *testing.T) {
		fc := &FositeClient{&ent.OauthClient{PersonalAccessClient: false}}
		assert.False(t, fc.IsPublic())
	})
}

func TestFositeClient_GetScopes(t *testing.T) {
	fc := &FositeClient{&ent.OauthClient{}}
	scopes := fc.GetScopes()
	assert.True(t, scopes.Has("read"))
	assert.True(t, scopes.Has("write"))
	assert.True(t, scopes.Has("follow"))
	assert.True(t, scopes.Has("push"))
}
