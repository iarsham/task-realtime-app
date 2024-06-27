package usecase

import (
	"github.com/iarsham/task-realtime-app/user-service/configs"
	"github.com/iarsham/task-realtime-app/user-service/domain"
	"github.com/iarsham/task-realtime-app/user-service/helpers"
	"github.com/iarsham/task-realtime-app/user-service/models"
	"go.uber.org/zap"
)

type loginUsecaseImpl struct {
	userRepository domain.UserRepository
	cfg            *configs.Config
	logger         *zap.Logger
}

func NewLoginUsecase(userRepository domain.UserRepository, cfg *configs.Config, logger *zap.Logger) domain.LoginUsecase {
	return &loginUsecaseImpl{
		userRepository: userRepository,
		cfg:            cfg,
		logger:         logger,
	}
}

func (u *loginUsecaseImpl) GetUserByEmail(email string) (*models.Users, error) {
	user, err := u.userRepository.GetUserByEmail(email)
	if err != nil {
		u.logger.Error("Error while getting user", zap.Error(err))
		return nil, err
	}
	return user, nil
}

func (u *loginUsecaseImpl) ValidatePass(hashedPass string, plainPass string) error {
	err := helpers.ValidatePass(hashedPass, plainPass)
	if err != nil {
		u.logger.Error("Error while validating password", zap.Error(err))
		return err
	}
	return nil
}

func (u *loginUsecaseImpl) CreateAccessToken(user *models.Users) (string, error) {
	token, err := helpers.CreateAccessToken(user, u.cfg.App.SecretKey, u.cfg.App.TokenExpireHour)
	if err != nil {
		u.logger.Error("Error while creating access token", zap.Error(err))
		return "", err
	}
	return token, nil
}
