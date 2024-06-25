package entities

import "time"

type SignupRequest struct {
	Username  string    `json:"username" bson:"username" binding:"required" example:"johndoe"`
	Email     string    `json:"email" bson:"email" binding:"required,email" example:"Kqg8Q@example.com"`
	Password  string    `json:"password" bson:"password" binding:"required,min=8,max=32" example:"1qaz2wsx"`
	CreatedAt time.Time `json:"created_at" bson:"created_at" bson:"created_at" swaggerignore:"true"`
}

type LoginRequest struct {
	Email    string `json:"email" bson:"email" binding:"required,email" example:"Kqg8Q@example.com"`
	Password string `json:"password" bson:"password" binding:"required,min=8,max=32" example:"1qaz2wsx"`
}
