package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/chat-service/ws"
	"go.uber.org/zap"
)

func wsRouters(r *gin.Engine, logger *zap.Logger) {
	wsAPI := r.Group("/ws")
	pool := ws.NewPool(logger)
	go pool.Run()
	wsAPI.GET("/", func(c *gin.Context) {
		ws.ServeWs(pool, c.Writer, c.Request)
	})
}
