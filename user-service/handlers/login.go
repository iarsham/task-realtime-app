package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/user-service/domain"
	"github.com/iarsham/task-realtime-app/user-service/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type LoginHandler struct {
	Usecase domain.LoginUsecase
}

func (l *LoginHandler) LoginHandler(ctx *gin.Context) {
	data := new(entities.LoginRequest)
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := l.Usecase.GetUserByEmail(data.Email)
	if errors.Is(err, mongo.ErrNoDocuments) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	if err := l.Usecase.ValidatePass(user.Password, data.Password); errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}
	accessToken, err := l.Usecase.CreateAccessToken(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"access-token": accessToken})
}
