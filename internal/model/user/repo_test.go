package user

import (
	"context"
	"testing"
	"time"

	"glintfed.org/internal/data"
)

func TestUsecase_GetTotalUsers(t *testing.T) {
	client, cleanup, err := data.NewTestClient(t)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	repo := NewRepo(client)
	ctx := context.Background()

	_, err = client.Ent.User.Create().
		SetUsername("user1").
		SetEmail("user1@example.com").
		SetPassword("password").
		Save(ctx)
	if err != nil {
		t.Fatalf("failed to create user 1: %v", err)
	}

	_, err = client.Ent.User.Create().
		SetUsername("user2").
		SetEmail("user2@example.com").
		SetPassword("password").
		Save(ctx)
	if err != nil {
		t.Fatalf("failed to create user 2: %v", err)
	}

	count, err := repo.GetTotalUsers(ctx)
	if err != nil {
		t.Fatalf("GetTotalUsers failed: %v", err)
	}
	if count != 2 {
		t.Errorf("expected count 2, got %d", count)
	}
}

func TestUsecase_GetMonthActiveUsers(t *testing.T) {
	client, cleanup, err := data.NewTestClient(t)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	repo := NewRepo(client)
	ctx := context.Background()

	now := time.Now()
	withinThreshold := now.Add(-2 * 7 * 24 * time.Hour)   // 2 weeks ago
	outsideThreshold := now.Add(-10 * 7 * 24 * time.Hour) // 10 weeks ago

	// 1. Updated within threshold
	_, err = client.Ent.User.Create().
		SetUsername("active1").
		SetEmail("active1@example.com").
		SetPassword("password").
		SetUpdatedAt(withinThreshold).
		SetLastActiveAt(outsideThreshold).
		Save(ctx)
	if err != nil {
		t.Fatalf("failed to create user 1: %v", err)
	}

	// 2. LastActive within threshold
	_, err = client.Ent.User.Create().
		SetUsername("active2").
		SetEmail("active2@example.com").
		SetPassword("password").
		SetUpdatedAt(outsideThreshold).
		SetLastActiveAt(withinThreshold).
		Save(ctx)
	if err != nil {
		t.Fatalf("failed to create user 2: %v", err)
	}

	// 3. Both outside threshold
	_, err = client.Ent.User.Create().
		SetUsername("inactive").
		SetEmail("inactive@example.com").
		SetPassword("password").
		SetUpdatedAt(outsideThreshold).
		SetLastActiveAt(outsideThreshold).
		Save(ctx)
	if err != nil {
		t.Fatalf("failed to create user 3: %v", err)
	}

	count, err := repo.GetMonthActiveUsers(ctx)
	if err != nil {
		t.Fatalf("GetMonthActiveUsers failed: %v", err)
	}
	if count != 2 {
		t.Errorf("expected count 2, got %d", count)
	}
}

func TestUsecase_GetHalfYearActiveUsers(t *testing.T) {
	client, cleanup, err := data.NewTestClient(t)
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()

	repo := NewRepo(client)
	ctx := context.Background()

	now := time.Now()
	withinThreshold := now.AddDate(0, -3, 0)   // 3 months ago
	outsideThreshold := now.AddDate(0, -12, 0) // 1 year ago

	// 1. Updated within threshold
	_, err = client.Ent.User.Create().
		SetUsername("active1").
		SetEmail("active1@example.com").
		SetPassword("password").
		SetUpdatedAt(withinThreshold).
		SetLastActiveAt(outsideThreshold).
		Save(ctx)
	if err != nil {
		t.Fatalf("failed to create user 1: %v", err)
	}

	// 2. LastActive within threshold
	_, err = client.Ent.User.Create().
		SetUsername("active2").
		SetEmail("active2@example.com").
		SetPassword("password").
		SetUpdatedAt(outsideThreshold).
		SetLastActiveAt(withinThreshold).
		Save(ctx)
	if err != nil {
		t.Fatalf("failed to create user 2: %v", err)
	}

	// 3. Both outside threshold
	_, err = client.Ent.User.Create().
		SetUsername("inactive").
		SetEmail("inactive@example.com").
		SetPassword("password").
		SetUpdatedAt(outsideThreshold).
		SetLastActiveAt(outsideThreshold).
		Save(ctx)
	if err != nil {
		t.Fatalf("failed to create user 3: %v", err)
	}

	count, err := repo.GetHalfYearActiveUsers(ctx)
	if err != nil {
		t.Fatalf("GetHalfYearActiveUsers failed: %v", err)
	}
	if count != 2 {
		t.Errorf("expected count 2, got %d", count)
	}
}
