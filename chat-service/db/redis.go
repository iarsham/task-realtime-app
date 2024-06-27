package db

import (
	"context"
	"fmt"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/redis/go-redis/v9"
)

func OpenRedis(cfg *configs.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}
	return client, nil
}
