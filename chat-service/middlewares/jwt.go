package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/iarsham/task-realtime-app/chat-service/configs"
	"github.com/iarsham/task-realtime-app/chat-service/helpers"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func JwtAuthMiddleware(logger *zap.Logger, cfg *configs.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
			return
		}
		authToken := strings.Split(authHeader, " ")
		if len(authToken) != 2 || strings.ToLower(authToken[0]) != "bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
			return
		}
		token, err := helpers.IsTokenValid(authToken[1], cfg.App.SecretKey)
		if err != nil {
			switch {
			case errors.Is(err, jwt.ErrTokenExpired):
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is expired"})
			default:
				logger.Error("invalid token: failed to parse token", zap.Any("error", err))
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is invalid"})
			}
			return
		}
		claims, err := helpers.GetClaims(token)
		if err != nil {
			logger.Error("invalid claims: failed to parse token", zap.Any("error", err))
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
			return
		}
		if claims["sub"] != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "refresh token not allowed"})
			return
		}
		ctx.Set("user_id", claims["user_id"])
		ctx.Set("username", claims["username"])
		ctx.Set("email", claims["email"])
		ctx.Next()
	}
}
