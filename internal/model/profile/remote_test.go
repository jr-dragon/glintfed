package profile

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"glintfed.org/internal/data"
)

func TestUsecase_RemoteUrlExists(t *testing.T) {
	client, cleanup, err := data.NewTestClient(t)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	ctx := context.Background()
	model := NewModel(client)

	// Create test profiles
	url := "https://example.com/users/alice"
	_, err = client.Ent.Profile.Create().
		SetUsername("alice").
		SetRemoteURL(url).
		Save(ctx)
	assert.NoError(t, err)

	t.Run("exists", func(t *testing.T) {
		exists, err := model.RemoteUrlExists(ctx, url)
		assert.NoError(t, err)
		assert.True(t, exists)
	})

	t.Run("not exists", func(t *testing.T) {
		exists, err := model.RemoteUrlExists(ctx, "https://example.com/users/bob")
		assert.NoError(t, err)
		assert.False(t, exists)
	})
}
