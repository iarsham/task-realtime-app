package repository

import (
	"context"
	"github.com/iarsham/task-realtime-app/chat-service/domain"
	"github.com/rabbitmq/amqp091-go"
	"time"
)

type brokerRepositoryImpl struct {
	conn *amqp091.Connection
	chnl *amqp091.Channel
}

func NewBrokerRepository(conn *amqp091.Connection, chnl *amqp091.Channel) domain.BrokerRepository {
	return &brokerRepositoryImpl{
		conn: conn,
		chnl: chnl,
	}
}

func (b *brokerRepositoryImpl) Publish(topic string, message []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	queue, err := b.chnl.QueueDeclare(
		topic,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	return b.chnl.PublishWithContext(
		ctx,
		"",
		queue.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
}
