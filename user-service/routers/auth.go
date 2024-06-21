package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/user-service/configs"
	"github.com/iarsham/task-realtime-app/user-service/handlers"
	"github.com/iarsham/task-realtime-app/user-service/repository"
	"github.com/iarsham/task-realtime-app/user-service/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func SetupRouters(r *gin.Engine, mongo *mongo.Database, cfg *configs.Config, logger *zap.Logger) {
	baseAPI := r.Group("/api/auth")
	userRepo := repository.NewUsersRepository(mongo, cfg)
	registerHandler := &handlers.RegisterHandler{
		Usecase: usecase.NewRegisterUsecase(userRepo, cfg, logger),
	}
	loginHandler := &handlers.LoginHandler{
		Usecase: usecase.NewLoginUsecase(userRepo, cfg, logger),
	}
	baseAPI.POST("/register", registerHandler.RegisterHandler)
	baseAPI.POST("/login", loginHandler.LoginHandler)
}
