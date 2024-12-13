package main

import (
	"fmt"
	"net/http"

	user "product-recommendation/internal/application/usecases"
	httpServer "product-recommendation/internal/infra/http"
	user_repository "product-recommendation/internal/infra/repository"
	apikey_repository "product-recommendation/internal/infra/repository/apikey"

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

func main() {
	fmt.Println("Microservice de Recomendação iniciado!")

	userRepo := user_repository.NewInMemoryUserRepository()
	apiKeyRepo := apikey_repository.NewInMemoryAPIKeyRepository()

	registerUser := user.NewRegisterUserUseCase(userRepo)
	listUsers := user.NewListUsersUseCase(userRepo)
	findUser := user.NewFindUserUseCase(userRepo)

	userHandler := httpServer.NewUserHandler(registerUser, listUsers, findUser)
	apiKeyHandler := httpServer.NewAPIKeyHandler(apiKeyRepo)

	router := gin.Default()
	router.Use(errorHandler())

	router.POST("/create/apiKey", apiKeyHandler.CreateAPIKey)

	usersGroup := router.Group("/users")
	usersGroup.Use(httpServer.APIKeyMiddleware(apiKeyRepo))
	usersGroup.POST("", userHandler.Register)
	usersGroup.GET("", userHandler.List)
	usersGroup.GET("/:email", userHandler.Find)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
