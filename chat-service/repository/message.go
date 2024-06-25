package repository

import (
	"context"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/iarsham/task-realtime-app/chat-service/domain"
	"github.com/iarsham/task-realtime-app/chat-service/entities"
	"github.com/iarsham/task-realtime-app/chat-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type messageRepositoryImpl struct {
	db  *mongo.Database
	cfg *configs.Config
}

func NewMessageRepository(db *mongo.Database, cfg *configs.Config) domain.MessageRepository {
	return &messageRepositoryImpl{
		db:  db,
		cfg: cfg,
	}
}

func (m *messageRepositoryImpl) List(roomID primitive.ObjectID) (*[]models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var messages []models.Message
	cursor, err := m.db.Collection(m.cfg.Mongo.MessageColl).Find(ctx, bson.M{"room_id": roomID})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &messages); err != nil {
		return nil, err
	}
	return &messages, nil
}

func (m *messageRepositoryImpl) Create(message *entities.MessageRequest) (*models.Message, error) {
	message.CreatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	inserted, err := m.db.Collection(m.cfg.Mongo.MessageColl).InsertOne(ctx, message)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": inserted.InsertedID}
	return collectMessageDocument(ctx, filter, m.db, m.cfg)
}

func collectMessageDocument(ctx context.Context, filter bson.M, db *mongo.Database, cfg *configs.Config) (*models.Message, error) {
	var message models.Message
	err := db.Collection(cfg.Mongo.MessageColl).FindOne(ctx, filter).Decode(&message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}
