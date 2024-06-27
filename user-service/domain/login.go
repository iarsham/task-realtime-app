package domain

import "github.com/iarsham/task-realtime-app/user-service/models"

type LoginUsecase interface {
	GetUserByEmail(email string) (*models.Users, error)
	ValidatePass(hashedPass string, plainPass string) error
	CreateAccessToken(user *models.Users) (string, error)
}
