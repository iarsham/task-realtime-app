package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/iarsham/task-realtime-app/chat-service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type MessageHandler struct {
	Usecase domain.MessageUsecase
}

// ListMessagesHandler godoc
//
//	@Summary		Get all Messages
//	@Description	Get all Messages from the database or cache and return it
//	@Produce		json
//	@Tags			Message
//	@Param			id	path		string	true	"room id"	example("667aa959e88fab79e20b728c")
//	@Success		200	{object}	helpers.ListMessages
//	@Failure		400	{object}	helpers.BadRequest
//	@Failure		500	{object}	helpers.InternalServerError
//	@router			/room/{id} [get]
func (m *MessageHandler) ListMessagesHandler(ctx *gin.Context) {
	roomID := ctx.Param("id")
	objID, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messages, err := m.Usecase.ListRoomMessages(objID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	ctx.JSON(http.StatusOK, messages)
}
