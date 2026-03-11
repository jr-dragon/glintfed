package cache

import (
	"context"
	"reflect"
	"time"
)

var drv Driver

func Register(d Driver) {
	drv = d
}

func Has(ctx context.Context, key string) bool {
	return drv.Has(ctx, key)
}

func Get(ctx context.Context, key string) any {
	return drv.Get(ctx, key)
}

func Set(ctx context.Context, key string, val any, ttl time.Duration) error {
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Func {
		t := v.Type()
		if t.NumIn() == 0 && t.NumOut() == 1 {
			results := v.Call(nil)
			val = results[0].Interface()
		}
	}
	
	return drv.Set(ctx, key, val, ttl)
}

func Del(ctx context.Context, key string) error {
	return drv.Del(ctx, key)
}

func Clear(ctx context.Context) error {
	return drv.Clear(ctx)
}

type Driver interface {
	// basic operation
	Has(ctx context.Context, key string) bool
	Get(ctx context.Context, key string) any
	Set(ctx context.Context, key string, val any, ttl time.Duration) (err error)
	Del(ctx context.Context, key string) error

	// clear and close
	Clear(ctx context.Context) error
}
