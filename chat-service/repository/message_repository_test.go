package repository

import (
	"github.com/iarsham/task-realtime-app/chat-service/entities"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestMessageInsertOne(t *testing.T) {
	msgRepo := NewMessageRepository(DB, Cfg)
	data := &entities.MessageRequest{
		Content:  "Hello Sir",
		RoomID:   primitive.NewObjectID(),
		SenderID: primitive.NewObjectID(),
	}
	msg, err := msgRepo.Create(data)
	assert.Nilf(t, err, "Error while inserting message")
	assert.Equal(t, "Hello Sir", msg.Content)
}

func TestMessageList(t *testing.T) {
	msgRepo := NewMessageRepository(DB, Cfg)
	msgs, err := msgRepo.List(primitive.NewObjectID())
	assert.Nilf(t, err, "Error while listing rooms")
	assert.Equal(t, 0, len(*msgs))
}
