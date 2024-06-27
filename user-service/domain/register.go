package domain

import (
	"github.com/iarsham/task-realtime-app/user-service/entities"
	"github.com/iarsham/task-realtime-app/user-service/models"
)

type RegisterUsecase interface {
	GetUserByEmail(email string) (*models.Users, error)
	GetUserByUsername(username string) (*models.Users, error)
	CreateUser(user *entities.SignupRequest) (*models.Users, error)
	EncryptPass(plainPass string) (string, error)
}
