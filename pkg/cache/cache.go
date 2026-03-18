package cache

import (
	"context"
	"errors"
	"fmt"
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
		return errors.New("val cannot be function, use SetFunc instead")
	}

	return drv.Set(ctx, key, val, ttl)
}

func SetFunc(ctx context.Context, key string, valf func() any, ttl time.Duration) error {
	val := valf()
	if err, iserr := val.(error); iserr {
		return fmt.Errorf("valf returns err: %w", err)
	}

	return Set(ctx, key, val, ttl)
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
