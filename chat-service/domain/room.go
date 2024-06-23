package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/chat-service/entities"
	"github.com/iarsham/task-realtime-app/chat-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomRepository interface {
	GetByName(name string) (*models.Room, error)
	List() (*[]models.Room, error)
	Create(room *entities.RoomRequest) (*models.Room, error)
}

type RoomUsecase interface {
	ListRooms() (*[]models.Room, error)
	GetRoomByName(name string) (*models.Room, error)
	CreateRoom(room *entities.RoomRequest) (*models.Room, error)
	GetUserID(ctx *gin.Context) primitive.ObjectID
}
