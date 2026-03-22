package media

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"glintfed.org/internal/data"
)

func TestModel_GetCDNUrl(t *testing.T) {
	client, cleanup, err := data.NewTestClient(t)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	model := NewModel(client)
	ctx := context.Background()

	path := "test/path/to/media.jpg"
	cdnURL := "https://cdn.example.com/media.jpg"

	t.Run("success", func(t *testing.T) {
		_, err := client.Ent.Media.Create().
			SetMediaPath(path).
			SetCdnURL(cdnURL).
			Save(ctx)
		assert.NoError(t, err)

		url, err := model.GetCDNUrl(ctx, path)
		assert.NoError(t, err)
		assert.Equal(t, cdnURL, url)
	})

	t.Run("not found", func(t *testing.T) {
		url, err := model.GetCDNUrl(ctx, "nonexistent")
		assert.Error(t, err)
		assert.Empty(t, url)
	})

	t.Run("cdn url is nil", func(t *testing.T) {
		pathNoCDN := "test/path/no_cdn.jpg"
		_, err := client.Ent.Media.Create().
			SetMediaPath(pathNoCDN).
			Save(ctx)
		assert.NoError(t, err)

		url, err := model.GetCDNUrl(ctx, pathNoCDN)
		assert.Error(t, err)
		assert.Empty(t, url)
	})
}
