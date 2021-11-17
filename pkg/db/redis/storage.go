package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

type Client struct {
	*redis.Client
}

func CreateClient(db int) *Client {
	return &Client{
		Client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       db,
		}),
	}
}

func (c *Client) Get(key string) (string, error) {
	return c.Client.Get(ctx, key).Result()
}

func (c *Client) Set(key string, value string, ttl int) {
	err := c.Client.Set(ctx, key, value, time.Duration(time.Duration(ttl).Milliseconds())).Err()
	if err != nil {
		fmt.Printf("redis error: %s", err.Error())
	}
}
