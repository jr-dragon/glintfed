package ent_test

import (
	"context"
	"testing"

	"glintfed.org/ent"
	"glintfed.org/ent/enttest"
	"glintfed.org/ent/story"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStory_Url_And_Permalink(t *testing.T) {
	appurl := "https://glintfed.test"
	ctx := context.Background()

	// 使用 enttest 建立記憶體資料庫連線
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	// 建立測試資料
	p, err := client.Profile.Create().SetUsername("alice").Save(ctx)
	require.NoError(t, err)

	s, err := client.Story.Create().
		SetProfileID(p.ID).
		SetID(123).
		SetDuration(15).
		Save(ctx)
	require.NoError(t, err)

	t.Run("with_preloaded_profile", func(t *testing.T) {
		sWithEdge, err := client.Story.Query().
			Where(story.ID(s.ID)).
			WithProfile().
			Only(ctx)
		require.NoError(t, err)

		assert.Equal(t, appurl+"/stories/alice/123", sWithEdge.Url(appurl))
		assert.Equal(t, appurl+"/stories/alice/123/activity", sWithEdge.Permalink(appurl))
	})

	t.Run("with_missing_edge_triggering_query", func(t *testing.T) {
		sWithoutEdge, err := client.Story.Query().
			Where(story.ID(s.ID)).
			Only(ctx)
		require.NoError(t, err)

		assert.Nil(t, sWithoutEdge.Edges.Profile)

		// 這會觸發 s.QueryProfile().FirstX()
		assert.Equal(t, appurl+"/stories/alice/123", sWithoutEdge.Url(appurl))
		assert.Equal(t, appurl+"/stories/alice/123/activity", sWithoutEdge.Permalink(appurl))
	})

	t.Run("manual_struct_with_edge", func(t *testing.T) {
		manualS := &ent.Story{
			ID: 456,
			Edges: ent.StoryEdges{
				Profile: &ent.Profile{
					Username: "bob",
				},
			},
		}
		assert.Equal(t, appurl+"/stories/bob/456", manualS.Url(appurl))
		assert.Equal(t, appurl+"/stories/bob/456/activity", manualS.Permalink(appurl))
	})
}

func TestStory_MediaUrl(t *testing.T) {
	appurl := "https://glintfed.test"

	t.Run("with_explicit_media_url", func(t *testing.T) {
		s := &ent.Story{
			MediaURL: "https://external.test/media.mp4",
		}
		assert.Equal(t, "https://external.test/media.mp4", s.MediaUrl(appurl))
	})

	t.Run("with_path_and_base_url", func(t *testing.T) {
		s := &ent.Story{
			Path: "2024/01/01/xyz.mp4",
		}
		assert.Equal(t, appurl+"/storage/2024/01/01/xyz.mp4", s.MediaUrl(appurl))
	})

	t.Run("with_path_and_base_url_trailing_slash", func(t *testing.T) {
		s := &ent.Story{
			Path: "2024/01/01/xyz.mp4",
		}
		assert.Equal(t, appurl+"/storage/2024/01/01/xyz.mp4", s.MediaUrl(appurl+"/"))
	})
}
