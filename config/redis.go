package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	Client *redis.Client
}

func SetupRedis(cfg *Config) (*RedisService, error) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.REDISHOST,
		DB:   0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis at %s: %w", cfg.REDISHOST, err)
	}
	log.Printf("Successfully connected to Redis at %s", cfg.REDISHOST)
	return &RedisService{
		Client: client,
	}, nil
}

func (r *RedisService) SetDataInRedis(key string, value []byte, expTime time.Duration) error {
	err := r.Client.Set(context.Background(), key, value, expTime).Err()
	if err != nil {
		return err
	}
	return nil
}

// func (r *RedisService) GetFromRedis(key string) (string, error) {
// 	fmt.Println("This is redis")
// 	jsonData, err := r.Client.Get(context.Background(), key).Result()
// 	if err != nil {
// 		return "", err
// 	}
// 	fmt.Println(jsonData)
// 	return jsonData, nil
// }

func (r *RedisService) GetFromRedis(key string) (string, error) {
	jsonData, err := r.Client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key %s does not exist", key)
	} else if err != nil {
		return "", err
	}
	return jsonData, nil
}
