package status

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"glintfed.org/ent/status"
	"glintfed.org/internal/data"
)

func TestUsecase_ObjectUrlExists(t *testing.T) {
	client, cleanup, err := data.NewTestClient(t)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	ctx := context.Background()
	model := NewModel(client)

	// Create test status
	url := "https://example.com/statuses/1"
	_, err = client.Ent.Status.Create().
		SetObjectURL(url).
		SetCaption("test status").
		SetRendered("test status").
		SetIsNsfw(false).
		SetScope("public").
		SetVisibility(status.VisibilityPublic).
		SetReply(false).
		SetLikesCount(0).
		SetReblogsCount(0).
		SetLocal(true).
		Save(ctx)
	assert.NoError(t, err)

	t.Run("exists", func(t *testing.T) {
		exists, err := model.ObjectUrlExists(ctx, url)
		assert.NoError(t, err)
		assert.True(t, exists)
	})

	t.Run("not exists", func(t *testing.T) {
		exists, err := model.ObjectUrlExists(ctx, "https://example.com/statuses/2")
		assert.NoError(t, err)
		assert.False(t, exists)
	})
}
