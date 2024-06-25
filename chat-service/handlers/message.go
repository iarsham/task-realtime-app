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
