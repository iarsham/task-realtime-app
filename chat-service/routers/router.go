package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func SetupRouters(r *gin.Engine, mongo *mongo.Database, cfg *configs.Config, logger *zap.Logger) {
	baseAPI := r.Group("/api")
	roomRouters(baseAPI, mongo, cfg, logger)
	messageRouters(baseAPI, mongo, cfg, logger)
	wsRouters(r, mongo, cfg, logger)
}
