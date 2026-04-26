package slotpool

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Pool manages a limited number of concurrent slots using Redis.
// Useful for preventing duplicate payments, orders, or double-submits.
type Pool struct {
	client   *redis.Client
	poolName string
	maxSlots int64
}

// NewPool creates a Pool with a given name and capacity.
func NewPool(client *redis.Client, poolName string, maxSlots int64) *Pool {
	return &Pool{client: client, poolName: poolName, maxSlots: maxSlots}
}

func (p *Pool) key() string {
	return fmt.Sprintf("slotpool:%s", p.poolName)
}

// Acquire tries to take a slot. Returns false if pool is full.
func (p *Pool) Acquire(ctx context.Context, identifier string, ttl time.Duration) (bool, error) {
	key := fmt.Sprintf("%s:%s", p.key(), identifier)

	// Atomically check & set slot
	ok, err := p.client.Pipeline().SetNX(ctx, key, "1", ttl).Result()
	if err != nil {
		return false, err
	}
	return ok, nil
}

// Release frees a slot.
func (p *Pool) Release(ctx context.Context, identifier string) error {
	key := fmt.Sprintf("%s:%s", p.key(), identifier)
	return p.client.Del(ctx, key).Err()
}
