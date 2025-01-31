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

//	@title			Real-Time Task
//	@version		0.1.0
//	@description	API Server for Real-Time Task
//	@termsOfService	http://swagger.io/terms/
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	Arsham Roshannejad
//	@contact.url	arsham.cloudarshamdev2001@gmail.com
//	@contact.email	arshamdev2001@gmail.com
//	@license.name	MIT
//	@license.url	https://www.mit.edu/~amini/LICENSE.md
//	@host			localhost:8000
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
