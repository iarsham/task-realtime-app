package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/iarsham/task-realtime-app/chat-service/domain"
	"github.com/iarsham/task-realtime-app/chat-service/handlers"
	"github.com/iarsham/task-realtime-app/chat-service/middlewares"
	"github.com/iarsham/task-realtime-app/chat-service/usecase"
	"go.uber.org/zap"
)

func roomRouters(r *gin.RouterGroup, roomRepo domain.RoomRepository, cfg *configs.Config, logger *zap.Logger) {
	roomAPI := r.Group("/room")
	roomHandler := &handlers.RoomHandler{
		Usecase: usecase.NewRoomUsecase(roomRepo, cfg, logger),
	}
	roomAPI.GET("/", roomHandler.ListRoomsHandler)
	roomAPI.POST("/", middlewares.JwtAuthMiddleware(logger, cfg), roomHandler.CreateRoomHandler)
}
