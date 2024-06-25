package routers

import (
	"github.com/gin-gonic/gin"
	docs "github.com/iarsham/task-realtime-app/chat-service/api"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/redis/go-redis/v9"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const BaseStr = "/api"

func SetupRouters(r *gin.Engine, mongo *mongo.Database, redis *redis.Client, cfg *configs.Config, logger *zap.Logger) {
	baseAPI := r.Group("/api")
	roomRouters(baseAPI, mongo, redis, cfg, logger)
	messageRouters(baseAPI, mongo, redis, cfg, logger)
	wsRouters(r, mongo, redis, cfg, logger)
	docs.SwaggerInfo.BasePath = BaseStr
	r.GET("/docs/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1)),
	)
}
