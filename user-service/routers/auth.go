package routers

import (
	"github.com/gin-gonic/gin"
	docs "github.com/iarsham/task-realtime-app/user-service/api"
	"github.com/iarsham/task-realtime-app/user-service/configs"
	"github.com/iarsham/task-realtime-app/user-service/handlers"
	"github.com/iarsham/task-realtime-app/user-service/repository"
	"github.com/iarsham/task-realtime-app/user-service/usecase"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const BaseStr = "/api/auth"

func SetupRouters(r *gin.Engine, mongo *mongo.Database, cfg *configs.Config, logger *zap.Logger) {
	baseAPI := r.Group(BaseStr)
	userRepo := repository.NewUsersRepository(mongo, cfg)
	registerHandler := &handlers.RegisterHandler{
		Usecase: usecase.NewRegisterUsecase(userRepo, cfg, logger),
	}
	loginHandler := &handlers.LoginHandler{
		Usecase: usecase.NewLoginUsecase(userRepo, cfg, logger),
	}
	baseAPI.POST("/register", registerHandler.RegisterHandler)
	baseAPI.POST("/login", loginHandler.LoginHandler)
	docs.SwaggerInfo.BasePath = BaseStr
	r.GET("/docs/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.DefaultModelsExpandDepth(-1)),
	)
}
