package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserID(ctx *gin.Context) (primitive.ObjectID, error) {
	strUserID, ok := ctx.Get("user_id")
	if !ok {
		return primitive.NilObjectID, errors.New("userid not exists in context")
	}
	return primitive.ObjectIDFromHex(strUserID.(string))
}
