package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	*redis.Client
}

func New() *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	rdb.Ping(context.Background())
	return &Client{rdb}
}
