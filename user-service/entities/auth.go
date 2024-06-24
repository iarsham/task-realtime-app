package entities

import "time"

type SignupRequest struct {
	Username  string    `json:"username" bson:"username" binding:"required"`
	Email     string    `json:"email" bson:"email" binding:"required,email"`
	Password  string    `json:"password" bson:"password" binding:"required,min=8,max=32"`
	CreatedAt time.Time `json:"created_at" bson:"created_at" bson:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email" bson:"email" binding:"required,email"`
	Password string `json:"password" bson:"password" binding:"required,min=8,max=32"`
}
