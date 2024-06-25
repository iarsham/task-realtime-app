package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/iarsham/task-realtime-app/chat-service/handlers"
	"github.com/iarsham/task-realtime-app/chat-service/middlewares"
	"github.com/iarsham/task-realtime-app/chat-service/repository"
	"github.com/iarsham/task-realtime-app/chat-service/usecase"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func roomRouters(r *gin.RouterGroup, mongo *mongo.Database, redis *redis.Client, cfg *configs.Config, logger *zap.Logger) {
	roomAPI := r.Group("/room")
	redisRepo := repository.NewRedisRepository(redis)
	roomRepo := repository.NewRoomRepository(mongo, redisRepo, cfg)
	roomHandler := &handlers.RoomHandler{
		Usecase: usecase.NewRoomUsecase(roomRepo, cfg, logger),
	}
	roomAPI.GET("/", roomHandler.ListRoomsHandler)
	roomAPI.POST("/", middlewares.JwtAuthMiddleware(logger, cfg), roomHandler.CreateRoomHandler)
}
