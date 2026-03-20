package instance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"glintfed.org/internal/data"
)

func TestUsecase_GetBlockedDomains(t *testing.T) {
	client, cleanup, err := data.NewTestClient(t)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	ctx := context.Background()
	model := NewModel(client)

	// Create test instances
	_, err = client.Ent.Instance.Create().
		SetDomain("banned.com").
		SetBanned(true).
		Save(ctx)
	assert.NoError(t, err)

	_, err = client.Ent.Instance.Create().
		SetDomain("allowed.com").
		SetBanned(false).
		Save(ctx)
	assert.NoError(t, err)

	t.Run("success", func(t *testing.T) {
		domains, err := model.GetBlockedDomains(ctx)
		assert.NoError(t, err)
		assert.Contains(t, domains, "banned.com")
		assert.NotContains(t, domains, "allowed.com")
		assert.Len(t, domains, 1)
	})
}
