package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/notification-service/broker"
	"github.com/iarsham/task-realtime-app/notification-service/configs"
	"github.com/iarsham/task-realtime-app/notification-service/logger"
	"github.com/iarsham/task-realtime-app/notification-service/routers"
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

	msgBroker, err := broker.OpenRabbitMQ(cfg)
	if err != nil {
		logs.Fatal("failed to connect to rabbitmq", zap.Error(err))
	}
	defer msgBroker.Close()

	chnl, err := msgBroker.Channel()
	if err != nil {
		logs.Fatal("failed to open channel", zap.Error(err))
	}
	defer chnl.Close()

	if !cfg.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.ForceConsoleColor()
	r := gin.Default()
	routers.SetupRouters(r, chnl, cfg, logs)

	addr := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)
	logs.Fatal("server start failed", zap.Error(r.Run(addr)))
}
