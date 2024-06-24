package helpers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/iarsham/task-realtime-app/user-service/models"
	"time"
)

func CreateAccessToken(user *models.Users, secretKey string, expire int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expire)).Unix()
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      exp,
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
}
