package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// NewRedis creates and pings a new Redis client.
func NewRedis(addr string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis ping failed: %w", err)
	}

	return client, nil
}
