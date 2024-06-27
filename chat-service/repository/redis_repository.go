package repository

import (
	"context"
	"github.com/iarsham/task-realtime-app/chat-service/domain"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisRepositoryImpl struct {
	redis *redis.Client
}

func NewRedisRepository(redis *redis.Client) domain.RedisRepository {
	return &redisRepositoryImpl{
		redis: redis,
	}
}

func (r *redisRepositoryImpl) Get(key string) ([]byte, error) {
	return r.redis.Get(context.Background(), key).Bytes()
}

func (r *redisRepositoryImpl) Set(key string, value interface{}) error {
	return r.redis.Set(context.Background(), key, value, time.Hour).Err()
}

func (r *redisRepositoryImpl) Del(key string) error {
	return r.redis.Del(context.Background(), key).Err()
}
