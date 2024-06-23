package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/iarsham/task-realtime-app/chat-service/domain"
	"github.com/iarsham/task-realtime-app/chat-service/entities"
	"github.com/iarsham/task-realtime-app/chat-service/helpers"
	"github.com/iarsham/task-realtime-app/chat-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type roomUsecaseImpl struct {
	roomRepository domain.RoomRepository
	cfg            *configs.Config
	logger         *zap.Logger
}

func NewRoomUsecase(roomRepository domain.RoomRepository, cfg *configs.Config, logger *zap.Logger) domain.RoomUsecase {
	return &roomUsecaseImpl{
		roomRepository: roomRepository,
		cfg:            cfg,
		logger:         logger,
	}
}

func (r *roomUsecaseImpl) ListRooms() (*[]models.Room, error) {
	rooms, err := r.roomRepository.List()
	if err != nil {
		r.logger.Error("Error while getting rooms", zap.Error(err))
		return nil, err
	}
	return rooms, nil
}

func (r *roomUsecaseImpl) GetRoomByName(name string) (*models.Room, error) {
	room, err := r.roomRepository.GetByName(name)
	if err != nil {
		r.logger.Error("Error while getting room by name", zap.Error(err))
		return nil, err
	}
	return room, err
}

func (r *roomUsecaseImpl) CreateRoom(room *entities.RoomRequest) (*models.Room, error) {
	createdRoom, err := r.roomRepository.Create(room)
	if err != nil {
		r.logger.Error("Error while creating room", zap.Error(err))
		return nil, err
	}
	return createdRoom, nil
}

func (r *roomUsecaseImpl) GetUserID(ctx *gin.Context) primitive.ObjectID {
	return helpers.GetUserID(ctx.Request)
}
