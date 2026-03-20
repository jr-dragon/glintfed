package instanceactor

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"glintfed.org/internal/data"
)

func TestModel_Get(t *testing.T) {
	client, _, err := data.NewTestClient(t)
	if err != nil {
		t.Fatal(err)
	}

	m := NewModel(client)
	ctx := context.Background()

	// Test case 1: Empty database
	ia, err := m.Get(ctx)
	assert.Error(t, err)
	assert.Nil(t, ia)

	// Test case 2: Single instance actor
	expected, err := client.Ent.InstanceActor.Create().
		SetPrivateKey("test-private-key").
		SetPublicKey("test-public-key").
		Save(ctx)
	assert.NoError(t, err)

	ia, err = m.Get(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, ia)
	assert.Equal(t, expected.ID, ia.ID)
	assert.Equal(t, "test-public-key", ia.PublicKey)
}
