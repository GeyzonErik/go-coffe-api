package http

import (
	api_key "product-recommendation/internal/core/application/apikey"
	handlers "product-recommendation/internal/core/infra/http/handlers"
	http_middleware "product-recommendation/internal/core/infra/http/middlewares"
	memory_repo "product-recommendation/internal/core/infra/repository/repository_memory"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(route *gin.Engine, userHandler *handlers.UserHandler, apiKeyHandler *api_key.APIKeyHandler, apiKeyRepo *memory_repo.InMemoryAPIKeyRepository) {
	apiKeyMiddleware := http_middleware.APIKeyMiddleware(apiKeyRepo)

	v1 := route.Group("/v1")
	{
		v1.POST("/create/apiKey", apiKeyHandler.CreateAPIKey)

		usersGroup := v1.Group("/users")
		usersGroup.Use(apiKeyMiddleware)
		{
			usersGroup.POST("", userHandler.Register)
			usersGroup.GET("", userHandler.List)
			usersGroup.GET("/:email", userHandler.Find)
		}
	}
}
