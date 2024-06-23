package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/iarsham/task-realtime-app/chat-service/handlers"
	"github.com/iarsham/task-realtime-app/chat-service/middlewares"
	"github.com/iarsham/task-realtime-app/chat-service/repository"
	"github.com/iarsham/task-realtime-app/chat-service/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func SetupRouters(r *gin.Engine, mongo *mongo.Database, cfg *configs.Config, logger *zap.Logger) {
	baseAPI := r.Group("/api/room")
	roomRepo := repository.NewRoomRepository(mongo, cfg)
	roomHandler := &handlers.RoomHandler{
		Usecase: usecase.NewRoomUsecase(roomRepo, cfg, logger),
	}
	baseAPI.GET("/", roomHandler.ListRoomsHandler)
	baseAPI.POST("/", middlewares.JwtAuthMiddleware(logger, cfg), roomHandler.CreateRoomHandler)
}
