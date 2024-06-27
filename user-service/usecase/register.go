package usecase

import (
	"github.com/iarsham/task-realtime-app/user-service/configs"
	"github.com/iarsham/task-realtime-app/user-service/domain"
	"github.com/iarsham/task-realtime-app/user-service/entities"
	"github.com/iarsham/task-realtime-app/user-service/helpers"
	"github.com/iarsham/task-realtime-app/user-service/models"
	"go.uber.org/zap"
)

type registerUsecaseImpl struct {
	userRepository domain.UserRepository
	cfg            *configs.Config
	logger         *zap.Logger
}

func NewRegisterUsecase(userRepository domain.UserRepository, cfg *configs.Config, logger *zap.Logger) domain.RegisterUsecase {
	return &registerUsecaseImpl{
		userRepository: userRepository,
		cfg:            cfg,
		logger:         logger,
	}
}

func (u *registerUsecaseImpl) GetUserByEmail(email string) (*models.Users, error) {
	user, err := u.userRepository.GetUserByEmail(email)
	if err != nil {
		u.logger.Error("Error while getting user", zap.Error(err))
		return nil, err
	}
	return user, nil
}

func (u *registerUsecaseImpl) GetUserByUsername(username string) (*models.Users, error) {
	user, err := u.userRepository.GetUserByUsername(username)
	if err != nil {
		u.logger.Error("Error while getting user", zap.Error(err))
		return nil, err
	}
	return user, nil
}

func (u *registerUsecaseImpl) CreateUser(user *entities.SignupRequest) (*models.Users, error) {
	createdUser, err := u.userRepository.CreateUser(user)
	if err != nil {
		u.logger.Error("Error while creating user", zap.Error(err))
		return nil, err
	}
	return createdUser, nil
}

func (u *registerUsecaseImpl) EncryptPass(plainPass string) (string, error) {
	encryptedPass, err := helpers.EncryptPass(plainPass)
	if err != nil {
		u.logger.Error("Error while encrypting password", zap.Error(err))
		return "", err
	}
	return string(encryptedPass), nil
}
