package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Notification struct {
	Content   string             `json:"content" bson:"content"`
	RoomID    primitive.ObjectID `json:"room_id" bson:"room_id"`
	SenderID  primitive.ObjectID `json:"sender_id" bson:"sender_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
