package redis

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/kuropenguin/my-tiny-url/app/config"
	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		// Addr:     "my-tiny-url-aiwe7o.serverless.apne1.cache.amazonaws.com:6379", // Redisサーバーのアドレス
		Addr:     fmt.Sprintf("%s:%d", config.Reids.Host, config.Reids.Port),
		Password: config.Reids.Password,
		DB:       config.Reids.DB, // 使用するDB
		TLSConfig: &tls.Config{ // トランジット中の暗号化を使用するためのTLS設定
			MinVersion: tls.VersionTLS12, // TLS1.2以上を使用する設定
		},
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(rdb.Ping(ctx).Err())
	}

	return rdb
}
