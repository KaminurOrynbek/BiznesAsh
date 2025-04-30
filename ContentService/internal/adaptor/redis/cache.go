package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

func NewRedisClient(addr, password string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("Redis ping failed: %v", err)
	}
	return rdb
}
