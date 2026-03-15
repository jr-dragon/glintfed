package cache

import (
	"context"
	"testing"
)

func TestMemoryDrv_Basic(t *testing.T) {
	ctx := context.Background()
	drv := NewMemoryDriver()

	// Test Set and Get
	err := drv.Set(ctx, "key1", "value1", 0)
	if err != nil {
		t.Fatalf("failed to set key1: %v", err)
	}

	val := drv.Get(ctx, "key1")
	if val != "value1" {
		t.Errorf("expected value1, got %v", val)
	}

	// Test Has
	if !drv.Has(ctx, "key1") {
		t.Errorf("expected key1 to exist")
	}

	if drv.Has(ctx, "non-existent") {
		t.Errorf("expected non-existent key to not exist")
	}

	// Test Del
	err = drv.Del(ctx, "key1")
	if err != nil {
		t.Fatalf("failed to delete key1: %v", err)
	}

	if drv.Has(ctx, "key1") {
		t.Errorf("expected key1 to be deleted")
	}
}

func TestMemoryDrv_Clear(t *testing.T) {
	ctx := context.Background()
	drv := NewMemoryDriver()

	drv.Set(ctx, "k1", "v1", 0)
	drv.Set(ctx, "k2", "v2", 0)

	err := drv.Clear(ctx)
	if err != nil {
		t.Fatalf("failed to clear: %v", err)
	}

	if drv.Has(ctx, "k1") || drv.Has(ctx, "k2") {
		t.Errorf("expected cache to be empty after clear")
	}
}
