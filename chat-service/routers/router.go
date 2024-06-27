package routers

import (
	"github.com/gin-gonic/gin"
	docs "github.com/iarsham/task-realtime-app/chat-service/api"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/iarsham/task-realtime-app/chat-service/repository"
	"github.com/iarsham/task-realtime-app/chat-service/usecase"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const BaseStr = "/api"

func SetupRouters(r *gin.Engine, mongo *mongo.Database, redis *redis.Client, msgBroker *amqp091.Connection, chnl *amqp091.Channel, cfg *configs.Config, logger *zap.Logger) {
	baseAPI := r.Group("/api")
	redisRepo := repository.NewRedisRepository(redis)
	brokerRepo := repository.NewBrokerRepository(msgBroker, chnl)
	roomRepo := repository.NewRoomRepository(mongo, redisRepo, cfg)
	msgRepo := repository.NewMessageRepository(mongo, redisRepo, cfg)
	roomUsecase := usecase.NewRoomUsecase(roomRepo, cfg, logger)
	msgUsecase := usecase.NewMessageUsecase(msgRepo, cfg, logger)
	brokerUsecase := usecase.NewBrokerUsecase(brokerRepo, logger)
	roomRouters(baseAPI, roomUsecase, cfg, logger)
	messageRouters(baseAPI, msgUsecase)
	wsRouters(r, msgUsecase, brokerUsecase, cfg, logger)
	docs.SwaggerInfo.BasePath = BaseStr
	r.GET("/docs/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1)),
	)
}
