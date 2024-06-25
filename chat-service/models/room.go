package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Room struct {
	ID        primitive.ObjectID `json:"id" bson:"_id" example:"667aa959e88fab79e20b728c"`
	Name      string             `json:"name" bson:"name" example:"warzone"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at" example:"2024-06-25T11:36:13.591+00:00"`
}
