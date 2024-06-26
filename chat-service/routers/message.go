package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/iarsham/task-realtime-app/chat-service/domain"
	"github.com/iarsham/task-realtime-app/chat-service/handlers"
	"github.com/iarsham/task-realtime-app/chat-service/usecase"
	"go.uber.org/zap"
)

func messageRouters(r *gin.RouterGroup, msgRepo domain.MessageRepository, cfg *configs.Config, logger *zap.Logger) {
	msgAPI := r.Group("/message")
	msgHandler := &handlers.MessageHandler{
		Usecase: usecase.NewMessageUsecase(msgRepo, cfg, logger),
	}
	msgAPI.GET("/:id", msgHandler.ListMessagesHandler)
}
