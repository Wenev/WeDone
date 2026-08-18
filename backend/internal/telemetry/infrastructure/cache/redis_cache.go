package cache

import (
	"context"
	"time"

	"github.com/Wenev/WeDone/backend/internal/auth/application/utility"
	"github.com/redis/go-redis/v9"
)

type RedisBoardCache struct {
	client *redis.Client
}

func NewRedisBoardCache(client *redis.Client) *RedisBoardCache {
	return &RedisBoardCache{
		client: client,
	}
}

func (c *RedisBoardCache) Get(ctx context.Context, key string) (string, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", utility.ErrCacheMiss
	}
	if err != nil {
		return "", err
	}
	return val, nil
}

func (c *RedisBoardCache) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	return c.client.Set(ctx, key, value, ttl).Err()
}

func (c *RedisBoardCache) Delete(ctx context.Context, keys ...string) error {
	return c.client.Del(ctx, keys...).Err()
}
