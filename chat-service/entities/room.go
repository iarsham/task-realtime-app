package entities

import "time"

type RoomRequest struct {
	Name      string    `json:"name" binding:"required" bson:"name" example:"warzone"`
	CreatedAt time.Time `json:"created_at" bson:"created_at" swaggerignore:"true"`
}
