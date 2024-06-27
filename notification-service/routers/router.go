package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/iarsham/task-realtime-app/notification-service/configs"
	"github.com/iarsham/task-realtime-app/notification-service/middlewares"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SetupRouters(r *gin.Engine, chnl *amqp091.Channel, cfg *configs.Config, logger *zap.Logger) {
	r.Use(middlewares.JwtAuthMiddleware(logger, cfg))
	r.GET("/notification", func(ctx *gin.Context) {
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			logger.Error("error while upgrading to websocket", zap.Error(err))
			return
		}
		defer conn.Close()
		msgs, err := chnl.Consume(
			cfg.RabbitMQ.QueueName,
			"",
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			logger.Error("error while consuming queue", zap.Error(err))
			return
		}
		for msg := range msgs {
			err := conn.WriteMessage(websocket.TextMessage, msg.Body)
			if err != nil {
			}
		}
	})
}
