package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	maxAge = time.Duration(30 * 24 * 60 * 60 * time.Second)
)

func NewCacheRedisRepository(rdb *redis.Client) ICacheRepository {
	return &CacheRedisRepository{
		cacheStorage: rdb,
	}
}

type CacheRedisRepository struct {
	cacheStorage *redis.Client
}

func (r *CacheRedisRepository) Set(key, val string) error {
	result := r.cacheStorage.Set(context.Background(), key, val, maxAge)
	fmt.Println(result.Err())
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (r *CacheRedisRepository) Get(key string) (string, error) {
	result := r.cacheStorage.Get(context.Background(), key)
	if result.Err() == redis.Nil {
		return "", ErrCacheNotFound
	}
	if result.Err() != nil {
		return "", result.Err()
	}
	return result.Val(), nil
}
