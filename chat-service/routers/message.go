package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/chat-service/domain"
	"github.com/iarsham/task-realtime-app/chat-service/handlers"
)

func messageRouters(r *gin.RouterGroup, msgUsecase domain.MessageUsecase) {
	msgAPI := r.Group("/message")
	msgHandler := &handlers.MessageHandler{
		Usecase: msgUsecase,
	}
	msgAPI.GET("/:id", msgHandler.ListMessagesHandler)
}
