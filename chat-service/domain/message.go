package domain

import (
	"github.com/iarsham/task-realtime-app/chat-service/entities"
	"github.com/iarsham/task-realtime-app/chat-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageRepository interface {
	List(roomID primitive.ObjectID) (*[]models.Message, error)
	Create(message *entities.MessageRequest) (*models.Message, error)
}

type MessageUsecase interface {
	ListRoomMessages(roomID primitive.ObjectID) (*[]models.Message, error)
	CreateMessage(message *entities.MessageRequest) (*models.Message, error)
}
