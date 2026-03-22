package oauth

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"glintfed.org/internal/data"
	"glintfed.org/internal/lib/fositestore"
)

func TestUsecase_CreateTokens(t *testing.T) {
	client, _, err := data.NewTestClient(t)
	require.NoError(t, err)

	cfg := &data.Config{}
	cfg.App.Auth.OAuth.HMACSecret = "test-hmac-secret-32-bytes-long!!"
	cfg.App.Auth.OAuth.PersonalClientID = "1"
	cfg.App.Auth.OAuth.AccessTokenLifespan = 30 * 24 * time.Hour

	store := fositestore.New(client, cfg)

	// Create the personal access client that CreateTokens will look up.
	_, err = client.Ent.OauthClient.Create().
		SetID(1).
		SetName("Personal Access Client").
		SetSecret("personal-secret").
		SetRedirect("").
		SetPersonalAccessClient(true).
		SetPasswordClient(false).
		SetRevoked(false).
		Save(context.Background())
	require.NoError(t, err)

	uc := NewUsecase(store, cfg)

	t.Run("success", func(t *testing.T) {
		result, err := uc.CreateTokens(context.Background(), 42, []string{"read", "write"})
		require.NoError(t, err)
		assert.NotEmpty(t, result.AccessToken)
		assert.NotEmpty(t, result.RefreshToken)
		assert.Equal(t, "1", result.ClientID)
		assert.Equal(t, int64((30 * 24 * time.Hour).Seconds()), result.ExpiresIn)
		assert.Equal(t, "personal-secret", result.ClientSecret)
	})

	t.Run("different users get distinct tokens", func(t *testing.T) {
		r1, err := uc.CreateTokens(context.Background(), 10, []string{"read"})
		require.NoError(t, err)
		r2, err := uc.CreateTokens(context.Background(), 20, []string{"read"})
		require.NoError(t, err)
		assert.NotEqual(t, r1.AccessToken, r2.AccessToken)
		assert.NotEqual(t, r1.RefreshToken, r2.RefreshToken)
	})

	t.Run("unknown personal client returns error", func(t *testing.T) {
		cfg2 := *cfg
		cfg2.App.Auth.OAuth.PersonalClientID = "999"
		uc2 := NewUsecase(store, &cfg2)

		_, err := uc2.CreateTokens(context.Background(), 42, []string{"read"})
		assert.Error(t, err)
	})
}

func TestUsecase_DefaultTTL(t *testing.T) {
	client, _, err := data.NewTestClient(t)
	require.NoError(t, err)

	// Zero AccessTokenLifespan should fall back to 365 days.
	cfg := &data.Config{}
	cfg.App.Auth.OAuth.HMACSecret = "test-hmac-secret-32-bytes-long!!"
	cfg.App.Auth.OAuth.PersonalClientID = "1"
	cfg.App.Auth.OAuth.AccessTokenLifespan = 0

	store := fositestore.New(client, cfg)
	uc := NewUsecase(store, cfg)

	assert.Equal(t, 365*24*time.Hour, uc.tokenTTL)
}
