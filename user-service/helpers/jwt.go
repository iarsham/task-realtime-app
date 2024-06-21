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

func IsTokenValid(reqToken string, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetClaims(token *jwt.Token) (map[string]interface{}, error) {
	claimsMap := make(map[string]interface{})
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		for k, v := range claims {
			claimsMap[k] = v
		}
		return claimsMap, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}
