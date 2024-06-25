package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/iarsham/task-realtime-app/chat-service/handlers"
	"github.com/iarsham/task-realtime-app/chat-service/repository"
	"github.com/iarsham/task-realtime-app/chat-service/usecase"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func messageRouters(r *gin.RouterGroup, mongo *mongo.Database, redis *redis.Client, cfg *configs.Config, logger *zap.Logger) {
	msgAPI := r.Group("/message")
	redisRepo := repository.NewRedisRepository(redis)
	msgRepo := repository.NewMessageRepository(mongo, redisRepo, cfg)
	msgHandler := &handlers.MessageHandler{
		Usecase: usecase.NewMessageUsecase(msgRepo, cfg, logger),
	}
	msgAPI.GET("/:id", msgHandler.ListMessagesHandler)
}
