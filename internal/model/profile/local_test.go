package profile

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"glintfed.org/internal/data"
)

func TestUsecase_GetByUsername(t *testing.T) {
	client, cleanup, err := data.NewTestClient(t)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	ctx := context.Background()
	model := NewModel(client)

	// Create test profiles
	_, err = client.Ent.Profile.Create().
		SetUsername("alice").
		Save(ctx)
	assert.NoError(t, err)

	_, err = client.Ent.Profile.Create().
		SetUsername("bob").
		SetDomain("example.com").
		Save(ctx)
	assert.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		p, err := model.GetByUsername(ctx, "alice")
		assert.NoError(t, err)
		assert.Equal(t, "alice", p.Username)
	})
}
