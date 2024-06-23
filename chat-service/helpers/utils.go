package helpers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func GetUserID(r *http.Request) primitive.ObjectID {
	return r.Context().Value("user_id").(primitive.ObjectID)
}
