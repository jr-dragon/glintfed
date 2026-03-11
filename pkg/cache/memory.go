package cache

import (
	"context"
	"reflect"
	"sync"
	"time"
)

type memoryDrv struct {
	mu      sync.RWMutex
	storage map[string]any
}

func NewMemoryDriver() Driver {
	return &memoryDrv{
		storage: make(map[string]any),
	}
}

func (d *memoryDrv) Has(_ context.Context, key string) bool {
	d.mu.RLock()
	defer d.mu.RUnlock()
	_, ok := d.storage[key]
	return ok
}

func (d *memoryDrv) Get(_ context.Context, key string) any {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.storage[key]
}

func (d *memoryDrv) Set(_ context.Context, key string, val any, _ time.Duration) (err error) {
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Func {
		t := v.Type()
		if t.NumIn() == 0 && t.NumOut() == 1 {
			results := v.Call(nil)
			val = results[0].Interface()
		}
	}

	d.mu.Lock()
	defer d.mu.Unlock()
	d.storage[key] = val
	return nil
}

func (d *memoryDrv) Del(_ context.Context, key string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.storage, key)
	return nil
}

func (d *memoryDrv) Clear(_ context.Context) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.storage = make(map[string]any)
	return nil
}
