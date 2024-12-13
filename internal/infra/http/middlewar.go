package http

import (
	"net/http"
	auth "product-recommendation/internal/application/services"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService *auth.AuthService) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")

		if authHeader == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			context.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token == authHeader {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			context.Abort()
			return
		}

		userID, err := authService.ValidateToken(token)

		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired access token"})
			context.Abort()
			return
		}

		context.Set("UserID", userID)
		context.Next()
	}
}
