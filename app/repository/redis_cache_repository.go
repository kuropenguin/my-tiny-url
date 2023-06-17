package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func NewCacheRedisRepository(rdb *redis.Client) ICacheRepository {
	return &CacheRedisRepository{
		cacheStorage: rdb,
	}
}

type CacheRedisRepository struct {
	cacheStorage *redis.Client
}

func (r *CacheRedisRepository) Save(key, val string) error {
	result := r.cacheStorage.Set(context.Background(), key, val, 0)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (r *CacheRedisRepository) Get(key string) (string, error) {
	result := r.cacheStorage.Get(context.Background(), key)
	if result.Err() != nil {
		return "", result.Err()
	}
	if result.Val() == "" {
		return "", ErrCacheNotFound
	}
	return result.String(), nil
}
