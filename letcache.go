package ciutils

import (
	"context"
	"errors"
)

// MemoryCache implements the Cache interface for single-threaded use.
type MemoryCache struct {
	store map[string][]byte
}

// NewMemoryCache creates a new instance of MemoryCache.
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		store: make(map[string][]byte),
	}
}

// Get returns certificate data for the specified key.
func (c *MemoryCache) Get(ctx context.Context, key string) ([]byte, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	data, exists := c.store[key]
	if !exists {
		return nil, errors.New("key not found in cache")
	}

	return data, nil
}

// Put stores data in the cache under the specified key.
func (c *MemoryCache) Put(ctx context.Context, key string, data []byte) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	c.store[key] = data
	return nil
}

// Delete removes certificate data from the cache for the specified key.
func (c *MemoryCache) Delete(ctx context.Context, key string) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	delete(c.store, key)
	return nil
}