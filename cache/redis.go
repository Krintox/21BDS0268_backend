package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func SetCache(ctx context.Context, key string, value string) error {
	err := Rdb.Set(ctx, key, value, 5*time.Minute).Err()
	return err
}

func GetCache(ctx context.Context, key string) (string, error) {
	val, err := Rdb.Get(ctx, key).Result()
	return val, err
}
