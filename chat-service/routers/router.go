package routers

import (
	"github.com/gin-gonic/gin"
	docs "github.com/iarsham/task-realtime-app/chat-service/api"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/iarsham/task-realtime-app/chat-service/repository"
	"github.com/redis/go-redis/v9"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const BaseStr = "/api"

func SetupRouters(r *gin.Engine, mongo *mongo.Database, redis *redis.Client, cfg *configs.Config, logger *zap.Logger) {
	baseAPI := r.Group("/api")
	redisRepo := repository.NewRedisRepository(redis)
	roomRepo := repository.NewRoomRepository(mongo, redisRepo, cfg)
	msgRepo := repository.NewMessageRepository(mongo, redisRepo, cfg)
	roomRouters(baseAPI, roomRepo, cfg, logger)
	messageRouters(baseAPI, msgRepo, cfg, logger)
	wsRouters(r, msgRepo, cfg, logger)
	docs.SwaggerInfo.BasePath = BaseStr
	r.GET("/docs/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1)),
	)
}
