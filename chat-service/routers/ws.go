package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/iarsham/task-realtime-app/chat-service/handlers"
	"github.com/iarsham/task-realtime-app/chat-service/usecase"
	"go.uber.org/zap"
)

func wsRouters(r *gin.Engine, cfg *configs.Config, logger *zap.Logger) {
	wsAPI := r.Group("/ws")
	wsHandler := &handlers.WsHandler{
		Usecase: usecase.NewWebSocketUsecase(cfg, logger),
	}
	wsAPI.GET("/", wsHandler.WebSocketHandler)
}
