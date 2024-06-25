package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func SetupRouters(r *gin.Engine, mongo *mongo.Database, redis *redis.Client, cfg *configs.Config, logger *zap.Logger) {
	baseAPI := r.Group("/api")
	roomRouters(baseAPI, mongo, redis, cfg, logger)
	messageRouters(baseAPI, mongo, redis, cfg, logger)
	wsRouters(r, mongo, redis, cfg, logger)
}
