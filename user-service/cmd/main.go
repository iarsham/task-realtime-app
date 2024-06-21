package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/user-service/configs"
	"github.com/iarsham/task-realtime-app/user-service/db"
	"github.com/iarsham/task-realtime-app/user-service/logger"
	"github.com/iarsham/task-realtime-app/user-service/routers"
	"go.uber.org/zap"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		panic(err)
	}

	logs, err := logger.NewZapLog(cfg.App.Debug)
	if err != nil {
		panic(err)
	}
	defer logs.Sync()

	mongo, err := db.OpenDB(cfg)
	if err != nil {
		logs.Panic("mongo connection failed", zap.Error(err))
	}
	defer mongo.Disconnect(context.Background())

	if !cfg.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.ForceConsoleColor()
	r := gin.Default()
	routers.SetupRouters(r, mongo.Database(cfg.Mongo.MongoDB), cfg, logs)

	addr := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)
	logs.Fatal("server start failed", zap.Error(r.Run(addr)))
}
