package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       1,
})

func GetTokenFromStorage(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

func SetTokenToStorage(key string, value string, ttl int) {
	err := rdb.Set(ctx, key, value, time.Duration(time.Duration(ttl).Milliseconds())).Err()
	if err != nil {
		fmt.Printf("redis error: %s", err.Error())
	}
}
