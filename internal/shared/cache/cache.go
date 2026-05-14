package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Cache wraps the Redis client with helper methods.
type Cache struct {
	client *redis.Client
}

// New creates a new Cache instance.
func New(client *redis.Client) *Cache {
	return &Cache{client: client}
}

// Set stores a key with a TTL.
func (c *Cache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return c.client.Set(ctx, key, value, ttl).Err()
}

// Get retrieves a value by key.
func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

// Delete removes a key.
func (c *Cache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

// Remember retrieves the value from cache or executes fn to populate it.
func (c *Cache) Remember(ctx context.Context, key string, ttl time.Duration, fn func() (string, error)) (string, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err == nil {
		return val, nil
	}
	if err != redis.Nil {
		return "", err
	}

	val, err = fn()
	if err != nil {
		return "", err
	}

	_ = c.client.Set(ctx, key, val, ttl).Err()
	return val, nil
}

// Increment atomically increments a key value and sets TTL if new.
func (c *Cache) Increment(ctx context.Context, key string, ttl time.Duration) (int64, error) {
	count, err := c.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	// Only set TTL on first increment (when count becomes 1)
	if count == 1 {
		c.client.Expire(ctx, key, ttl)
	}
	return count, nil
}

// Client returns the underlying Redis client.
func (c *Cache) Client() *redis.Client {
	return c.client
}
