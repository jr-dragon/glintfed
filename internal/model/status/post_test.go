package status

import (
	"context"
	"testing"
	"time"

	"glintfed.org/internal/data"
)

func TestUsecase_GetLocalPostsCount(t *testing.T) {
	client, cleanup, err := data.NewTestClient(t)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	repo := NewRepo(client)
	ctx := context.Background()

	t.Run("CountOnlyLocalNonSharePosts", func(t *testing.T) {
		// 1. Local, not share, not deleted
		_, err = client.Ent.Status.Create().
			SetLocal(true).
			SetType("post").
			SetCaption("caption").
			SetRendered("rendered").
			SetURI("uri1").
			Save(ctx)
		if err != nil {
			t.Fatalf("failed to create status 1: %v", err)
		}

		// 2. Local, share, not deleted
		_, err = client.Ent.Status.Create().
			SetLocal(true).
			SetType("share").
			SetCaption("caption").
			SetRendered("rendered").
			SetURI("uri2").
			Save(ctx)
		if err != nil {
			t.Fatalf("failed to create status 2: %v", err)
		}

		// 3. Remote, not share, not deleted
		_, err = client.Ent.Status.Create().
			SetLocal(false).
			SetType("post").
			SetCaption("caption").
			SetRendered("rendered").
			SetURI("uri3").
			Save(ctx)
		if err != nil {
			t.Fatalf("failed to create status 3: %v", err)
		}

		// 4. Local, not share, deleted
		_, err = client.Ent.Status.Create().
			SetLocal(true).
			SetType("post").
			SetCaption("caption").
			SetRendered("rendered").
			SetURI("uri4").
			SetDeletedAt(time.Now()).
			Save(ctx)
		if err != nil {
			t.Fatalf("failed to create status 4: %v", err)
		}

		count, err := repo.GetLocalPostsCount(ctx)
		if err != nil {
			t.Fatalf("GetLocalPostsCount returned error: %v", err)
		}
		if count != 1 {
			t.Errorf("expected count 1, got %d", count)
		}
	})
}
