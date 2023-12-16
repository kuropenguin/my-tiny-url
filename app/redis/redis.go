package redis

import (
	"context"
	"crypto/tls"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "my-tiny-url-aiwe7o.serverless.apne1.cache.amazonaws.com:6379", // Redisサーバーのアドレス
		Password: "",                                                             // パスワードがあれば設定（ElastiCacheでは通常不要）
		DB:       0,                                                              // 使用するDB
		TLSConfig: &tls.Config{ // トランジット中の暗号化を使用するためのTLS設定
			MinVersion: tls.VersionTLS12, // TLS1.2以上を使用する設定
		},
	})

	return rdb
}
