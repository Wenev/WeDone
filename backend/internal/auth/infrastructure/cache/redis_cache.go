package cache

import (
	"context"
	"time"

	"github.com/Wenev/WeDone/backend/internal/auth/application/utility"
	"github.com/redis/go-redis/v9"
)

type RedisUserCache struct {
	client *redis.Client
}

func NewRedisUserCache(client *redis.Client) *RedisUserCache {
	return &RedisUserCache{
		client: client,
	}
}

func (c *RedisUserCache) Get(ctx context.Context, key string) (string, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", utility.ErrCacheMiss
	}
	if err != nil {
		return "", err
	}
	return val, nil
}

func (c *RedisUserCache) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	return c.client.Set(ctx, key, value, ttl).Err()
}

func (c *RedisUserCache) Delete(ctx context.Context, keys ...string) error {
	return c.client.Del(ctx, keys...).Err()
}
