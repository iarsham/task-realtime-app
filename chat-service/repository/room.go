package repository

import (
	"context"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/iarsham/task-realtime-app/chat-service/domain"
	"github.com/iarsham/task-realtime-app/chat-service/entities"
	"github.com/iarsham/task-realtime-app/chat-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type roomRepositoryImpl struct {
	db  *mongo.Database
	cfg *configs.Config
}

func NewRoomRepository(db *mongo.Database, cfg *configs.Config) domain.RoomRepository {
	return &roomRepositoryImpl{
		db:  db,
		cfg: cfg,
	}
}

func (r *roomRepositoryImpl) List() (*[]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var rooms []models.Room
	cursor, err := r.db.Collection(r.cfg.Mongo.RoomColl).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &rooms); err != nil {
		return nil, err
	}
	return &rooms, nil
}

func (r *roomRepositoryImpl) GetByName(name string) (*models.Room, error) {
	filter := bson.M{"name": name}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return collectRoomDocument(ctx, filter, r.db, r.cfg)
}

func (r *roomRepositoryImpl) Create(room *entities.RoomRequest) (*models.Room, error) {
	room.CreatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := r.db.Collection(r.cfg.Mongo.RoomColl).InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"name": room.Name}
	return collectRoomDocument(ctx, filter, r.db, r.cfg)
}

func collectRoomDocument(ctx context.Context, filter bson.M, db *mongo.Database, cfg *configs.Config) (*models.Room, error) {
	var room models.Room
	err := db.Collection(cfg.Mongo.RoomColl).FindOne(ctx, filter).Decode(&room)
	if err != nil {
		return nil, err
	}
	return &room, nil
}
