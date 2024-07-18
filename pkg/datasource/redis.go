package datasource

import (
	"context"
	"i3/config"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value any, exp time.Duration) (string, error)
	Del(ctx context.Context, key ...string) (int64, error)
}

type re struct {
	client *redis.Client
}

func NewRedis() Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     config.ReadConfig().RedisAddr,
		Username: config.ReadConfig().RedisUsername,
		Password: config.ReadConfig().RedisPassword,
	})

	return &re{
		client: client,
	}
}

func (r *re) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *re) Set(ctx context.Context, key string, value any, exp time.Duration) (string, error) {
	return r.client.Set(ctx, key, value, exp).Result()
}

func (r *re) Del(ctx context.Context, key ...string) (int64, error) {
	return r.client.Del(ctx, key...).Result()
}
