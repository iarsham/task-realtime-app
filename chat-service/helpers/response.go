package helpers

import "github.com/iarsham/task-realtime-app/chat-service/models"

type RoomCreated struct {
	Response string `json:"response" example:"room created successfully"`
}

type ListRooms []models.Room

type RoomExists struct {
	Error string `json:"error" example:"room already exists"`
}

type InternalServerError struct {
	Error string `json:"error" example:"internal server error"`
}

type BadRequest struct {
	Error string `json:"error" example:"bad request"`
}

type ListMessages []models.Message
