package entities

import "time"

type RoomRequest struct {
	Name      string    `json:"name" binding:"required" bson:"name"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
