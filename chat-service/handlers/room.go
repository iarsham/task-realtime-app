package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/chat-service/domain"
	"github.com/iarsham/task-realtime-app/chat-service/entities"
	"github.com/iarsham/task-realtime-app/chat-service/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type RoomHandler struct {
	Usecase domain.RoomUsecase
}

// ListRoomsHandler godoc
//
//	@Summary		Get all Rooms
//	@Description	Get all Rooms from the database or cache and return it
//	@Produce		json
//	@Tags			Room
//	@Success		200	{object}	helpers.ListRooms
//	@Failure		500	{object}	helpers.InternalServerError
//	@router			/room [get]
func (r *RoomHandler) ListRoomsHandler(ctx *gin.Context) {
	rooms, err := r.Usecase.ListRooms()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, rooms)
}

// CreateRoomHandler godoc
//
//	@Summary		Create Room
//	@Description	Create a Room with the provided data
//	@Accept			json
//	@Produce		json
//	@Tags			Room
//	@Param			templateRequest	body		entities.RoomRequest	true	"room data"
//	@Success		201				{object}	helpers.RoomCreated
//	@Failure		400				{object}	helpers.BadRequest
//	@Failure		408				{object}	helpers.RoomExists
//	@Failure		500				{object}	helpers.InternalServerError
//	@router			/room [post]
func (r *RoomHandler) CreateRoomHandler(ctx *gin.Context) {
	data := new(entities.RoomRequest)
	if err := ctx.ShouldBindJSON(data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, err := r.Usecase.GetRoomByName(data.Name); !errors.Is(err, mongo.ErrNoDocuments) {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "room already exists"})
		return
	}
	helpers.Background(func() {
		r.Usecase.CreateRoom(data)
	})
	ctx.JSON(http.StatusCreated, gin.H{"response": "room created successfully"})
}
