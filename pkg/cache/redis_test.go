package cache

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
)

func TestRedisDrv(t *testing.T) {
	db, mock := redismock.NewClientMock()
	drv := NewRedisDriver(db)
	ctx := context.Background()

	t.Run("Has", func(t *testing.T) {
		mock.ExpectExists("key1").SetVal(1)
		if !drv.Has(ctx, "key1") {
			t.Error("expected key1 to exist")
		}

		mock.ExpectExists("key2").SetVal(0)
		if drv.Has(ctx, "key2") {
			t.Error("expected key2 to not exist")
		}

		mock.ExpectExists("key3").SetErr(redis.ErrClosed)
		if drv.Has(ctx, "key3") {
			t.Error("expected false on error")
		}
	})

	t.Run("Get", func(t *testing.T) {
		mock.ExpectGet("key1").SetVal("val1")
		val := drv.Get(ctx, "key1")
		if val != "val1" {
			t.Errorf("expected val1, got %v", val)
		}

		mock.ExpectGet("key2").RedisNil()
		val = drv.Get(ctx, "key2")
		if val != nil {
			t.Errorf("expected nil for non-existent key, got %v", val)
		}
	})

	t.Run("Set", func(t *testing.T) {
		mock.ExpectSet("key1", "val1", time.Minute).SetVal("OK")
		err := drv.Set(ctx, "key1", "val1", time.Minute)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("Del", func(t *testing.T) {
		mock.ExpectDel("key1").SetVal(1)
		err := drv.Del(ctx, "key1")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("Clear", func(t *testing.T) {
		mock.ExpectFlushDB().SetVal("OK")
		err := drv.Clear(ctx)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
