package http

import (
	"fmt"
	"net/http"

	api_key "product-recommendation/internal/core/application/apikey"
	"product-recommendation/internal/core/application/user"
	handlers "product-recommendation/internal/core/infra/http/handlers"
	"product-recommendation/internal/core/infra/repository/repository_memory"
	routes "product-recommendation/internal/interfaces/http"

	"github.com/gin-gonic/gin"
)

func errorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		if len(context.Errors) > 0 {
			err := context.Errors.Last()
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

func InitializeDependencies() (*handlers.UserHandler, *api_key.APIKeyHandler, *repository_memory.InMemoryAPIKeyRepository) {
	userRepo := repository_memory.NewInMemoryUserRepository()
	apiKeyRepo := repository_memory.NewInMemoryAPIKeyRepository()

	registerUser := user.NewRegisterUserUseCase(userRepo)
	listUsers := user.NewListUsersUseCase(userRepo)
	findUser := user.NewFindUserUseCase(userRepo)

	userHandler := handlers.NewUserHandler(registerUser, listUsers, findUser)
	apiKeyHandler := api_key.NewAPIKeyHandler(apiKeyRepo)

	return userHandler, apiKeyHandler, apiKeyRepo
}

func StartServer() {
	fmt.Println("Microservice de Recomendação iniciado!")

	userHandler, apiKeyHandler, apiKeyRepo := InitializeDependencies()

	router := gin.Default()
	router.Use(errorHandler())

	routes.SetupRoutes(router, userHandler, apiKeyHandler, apiKeyRepo)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
