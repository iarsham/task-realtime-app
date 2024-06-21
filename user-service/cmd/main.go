package main

import (
	"context"
	"github.com/iarsham/task-realtime-app/user-service/configs"
	"github.com/iarsham/task-realtime-app/user-service/db"
	"github.com/iarsham/task-realtime-app/user-service/logger"
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
}
