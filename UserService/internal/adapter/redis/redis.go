package redis

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/KaminurOrynbek/BiznesAsh/internal/entity"
	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(redisURL string) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // No password for default Redis setup
		DB:       0,  // Default DB
	})

	// Test the connection
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	return &RedisClient{client: client}
}

func (r *RedisClient) GetUser(ctx context.Context, key string) (*entity.User, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil // Cache miss
	}
	if err != nil {
		return nil, err
	}

	var user entity.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *RedisClient) SetUser(ctx context.Context, key string, user *entity.User, expiration time.Duration) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, data, expiration).Err()
}

func (r *RedisClient) DeleteUser(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
