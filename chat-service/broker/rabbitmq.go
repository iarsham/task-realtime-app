package broker

import (
	"fmt"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/rabbitmq/amqp091-go"
)

func OpenRabbitMQ(cfg *configs.Config) (*amqp091.Connection, error) {
	dsn := makeDsn(cfg)
	return amqp091.Dial(dsn)
}

func makeDsn(cfg *configs.Config) string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		cfg.RabbitMQ.User,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)
}
