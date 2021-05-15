package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client     *redis.Client
	expiration time.Duration
}

func NewRedisCache(client *redis.Client, defaultExpiration time.Duration) Cache {
	return &RedisCache{
		client:     client,
		expiration: defaultExpiration,
	}
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	originalUrl, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return originalUrl, nil
}

func (r *RedisCache) Set(ctx context.Context, key, value string, expiration time.Duration) error {
	expiration = minDruation(r.expiration, expiration)

	if _, err := r.client.Set(ctx, key, value, expiration).Result(); err != nil {
		return err
	}

	return nil
}

func minDruation(a, b time.Duration) time.Duration {
	if a <= b {
		return a
	}

	return b
}

func (r *RedisCache) Delete(ctx context.Context, key string) error {
	if _, err := r.client.Del(ctx, key).Result(); err != nil {
		return err
	}

	return nil
}
