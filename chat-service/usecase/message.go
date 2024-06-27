package usecase

import (
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/iarsham/task-realtime-app/chat-service/domain"
	"github.com/iarsham/task-realtime-app/chat-service/entities"
	"github.com/iarsham/task-realtime-app/chat-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type messageUsecaseImpl struct {
	messageRepository domain.MessageRepository
	cfg               *configs.Config
	logger            *zap.Logger
}

func NewMessageUsecase(messageRepository domain.MessageRepository, cfg *configs.Config, logger *zap.Logger) domain.MessageUsecase {
	return &messageUsecaseImpl{
		messageRepository: messageRepository,
		cfg:               cfg,
		logger:            logger,
	}
}

func (m *messageUsecaseImpl) ListRoomMessages(roomID primitive.ObjectID) (*[]models.Message, error) {
	messages, err := m.messageRepository.List(roomID)
	if err != nil {
		m.logger.Error("failed to list messages", zap.Error(err))
		return nil, err
	}
	return messages, nil
}

func (m *messageUsecaseImpl) CreateMessage(message *entities.MessageRequest) (*models.Message, error) {
	messages, err := m.messageRepository.Create(message)
	if err != nil {
		m.logger.Error("failed to create message", zap.Error(err))
		return nil, err
	}
	return messages, nil
}
