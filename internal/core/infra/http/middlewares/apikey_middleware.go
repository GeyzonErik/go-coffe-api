package http

import (
	"net/http"
	apikey "product-recommendation/internal/core/domain/api_key"

	"github.com/gin-gonic/gin"
)

func APIKeyMiddleware(repo apikey.Repository) gin.HandlerFunc {
	return func(context *gin.Context) {
		apikey := context.GetHeader("X-API-KEY")

		if apikey == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "missing API Key"})
			context.Abort()
			return
		}

		system, err := repo.GetSystemByKey(apikey)

		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		context.Set("system", system.SystemName)
		context.Next()
	}
}
