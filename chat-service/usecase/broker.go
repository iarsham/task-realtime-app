package usecase

import (
	"github.com/iarsham/task-realtime-app/chat-service/domain"
	"go.uber.org/zap"
)

type brokerUsecase struct {
	brokerRepo domain.BrokerRepository
	logger     *zap.Logger
}

func NewBrokerUsecase(brokerRepo domain.BrokerRepository, logger *zap.Logger) domain.BrokerUsecase {
	return &brokerUsecase{
		brokerRepo: brokerRepo,
		logger:     logger,
	}
}

func (b *brokerUsecase) PublishQueue(topic string, message []byte) error {
	if err := b.brokerRepo.Publish(topic, message); err != nil {
		b.logger.Error("failed to publish message", zap.Error(err))
		return err
	}
	return nil
}
