package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisDrv struct {
	client *redis.Client
}

func NewRedisDriver(client *redis.Client) Driver {
	return &redisDrv{
		client: client,
	}
}

func (d *redisDrv) Has(ctx context.Context, key string) bool {
	n, err := d.client.Exists(ctx, key).Result()
	if err != nil {
		return false
	}
	return n > 0
}

func (d *redisDrv) Get(ctx context.Context, key string) any {
	val, err := d.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		return nil
	}
	return val
}

func (d *redisDrv) Set(ctx context.Context, key string, val any, ttl time.Duration) error {
	return d.client.Set(ctx, key, val, ttl).Err()
}

func (d *redisDrv) Del(ctx context.Context, key string) error {
	return d.client.Del(ctx, key).Err()
}

func (d *redisDrv) Clear(ctx context.Context) error {
	return d.client.FlushDB(ctx).Err()
}
