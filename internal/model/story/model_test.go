package story

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"glintfed.org/internal/data"
)

func TestModel_GetByUsernameAndID(t *testing.T) {
	client, cleanup, err := data.NewTestClient(t)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	ctx := context.Background()
	model := NewModel(client)

	// Create test profile
	p, err := client.Ent.Profile.Create().
		SetUsername("testuser").
		SetIsPrivate(false).
		SetIsSuggestable(true).
		Save(ctx)
	assert.NoError(t, err)

	// Create test story
	s, err := client.Ent.Story.Create().
		SetProfile(p).
		SetActive(true).
		SetBearcapToken("token").
		SetMime("image/jpeg").
		SetType("photo").
		SetDuration(15).
		Save(ctx)
	assert.NoError(t, err)

	t.Run("found", func(t *testing.T) {
		st, err := model.GetByUsernameAndID(ctx, "testuser", s.ID)
		if assert.NoError(t, err) && assert.NotNil(t, st) {
			assert.Equal(t, s.ID, st.ID)
			assert.NotNil(t, st.Edges.Profile)
			assert.Equal(t, p.ID, st.Edges.Profile.ID)
		}
	})

	t.Run("wrong username", func(t *testing.T) {
		st, err := model.GetByUsernameAndID(ctx, "wronguser", s.ID)
		assert.Error(t, err)
		assert.Nil(t, st)
	})

	t.Run("inactive story", func(t *testing.T) {
		s2, err := client.Ent.Story.Create().
			SetProfile(p).
			SetActive(false).
			SetDuration(15).
			Save(ctx)
		assert.NoError(t, err)

		st, err := model.GetByUsernameAndID(ctx, "testuser", s2.ID)
		assert.Error(t, err)
		assert.Nil(t, st)
	})
}
