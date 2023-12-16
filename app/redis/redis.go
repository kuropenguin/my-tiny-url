package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func NewRedisClient() *redis.Client {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "tiny-url-aiwe7o.serverless.apne1.cache.amazonaws.com:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
