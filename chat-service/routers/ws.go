package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/iarsham/task-realtime-app/chat-service/domain"
	"github.com/iarsham/task-realtime-app/chat-service/helpers"
	"github.com/iarsham/task-realtime-app/chat-service/middlewares"
	"github.com/iarsham/task-realtime-app/chat-service/ws"
	"go.uber.org/zap"
	"net/http"
)

func wsRouters(r *gin.Engine, msgUsecase domain.MessageUsecase, brokerUsecase domain.BrokerUsecase, cfg *configs.Config, logger *zap.Logger) {
	wsAPI := r.Group("/ws")
	wsAPI.Use(middlewares.JwtAuthMiddleware(logger, cfg))
	pool := ws.NewPool(logger)
	go pool.Run(brokerUsecase)
	wsAPI.GET("/", func(ctx *gin.Context) {
		userID, err := helpers.GetUserID(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ws.ServeWs(pool, userID, msgUsecase, ctx.Writer, ctx.Request)
	})
}
