package slotpool

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

// ErrLockNotAcquired is returned when the lock cannot be obtained.
var ErrLockNotAcquired = errors.New("lock not acquired: resource is busy")

// Lock provides Redis-based distributed locking.
type Lock struct {
	client *redis.Client
}

// NewLock creates a new Lock.
func NewLock(client *redis.Client) *Lock {
	return &Lock{client: client}
}

// Acquire tries to acquire a lock on the given key with a TTL.
// Returns ErrLockNotAcquired if the lock is already held.
func (l *Lock) Acquire(ctx context.Context, key string, ttl time.Duration) error {
	ok, err := l.client.SetNX(ctx, "lock:"+key, "1", ttl).Result()
	if err != nil {
		return err
	}
	if !ok {
		return ErrLockNotAcquired
	}
	return nil
}

// Release removes a lock key.
func (l *Lock) Release(ctx context.Context, key string) error {
	return l.client.Del(ctx, "lock:"+key).Err()
}

// WithLock acquires a lock, runs fn, and always releases.
func (l *Lock) WithLock(ctx context.Context, key string, ttl time.Duration, fn func() error) error {
	if err := l.Acquire(ctx, key, ttl); err != nil {
		return err
	}
	defer func() { _ = l.Release(ctx, key) }()
	return fn()
}
