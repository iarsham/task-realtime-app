package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/user-service/domain"
	"github.com/iarsham/task-realtime-app/user-service/entities"
	"github.com/iarsham/task-realtime-app/user-service/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type RegisterHandler struct {
	Usecase domain.RegisterUsecase
}

// RegisterHandler godoc
//
//	@Summary		Register
//	@Description	Register a new user
//	@Accept			json
//	@Produce		json
//	@Tags			Auth
//	@Param			userRequest	body		entities.SignupRequest	true	"register data"
//	@Success		201			{object}	helpers.UserCreated
//	@Failure		400			{object}	helpers.BadRequest
//	@Failure		409			{object}	helpers.EmailAlreadyExists
//	@Failure		500			{object}	helpers.InternalServerError
//	@Router			/register [post]
func (r *RegisterHandler) RegisterHandler(ctx *gin.Context) {
	data := new(entities.SignupRequest)
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, err := r.Usecase.GetUserByUsername(data.Username); !errors.Is(err, mongo.ErrNoDocuments) {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "username already exists"})
		return
	}
	if _, err := r.Usecase.GetUserByUsername(data.Email); !errors.Is(err, mongo.ErrNoDocuments) {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "email already exists"})
		return
	}
	encryptedPass, err := r.Usecase.EncryptPass(data.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	data.Password = encryptedPass
	helpers.Background(func() {
		r.Usecase.CreateUser(data)
	})
	ctx.JSON(http.StatusCreated, gin.H{"response": "User created successfully"})
}
