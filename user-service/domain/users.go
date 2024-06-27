package domain

import (
	"github.com/iarsham/task-realtime-app/user-service/entities"
	"github.com/iarsham/task-realtime-app/user-service/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	GetUserById(id primitive.ObjectID) (*models.Users, error)
	GetUserByEmail(email string) (*models.Users, error)
	GetUserByUsername(username string) (*models.Users, error)
	CreateUser(user *entities.SignupRequest) (*models.Users, error)
}
