package main

import (
	"fmt"
	"net/http"

	user "product-recommendation/internal/application/usecases"
	httpServer "product-recommendation/internal/infra/http"
	user_repository "product-recommendation/internal/infra/repository"

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
	registerUser := user.NewRegisterUserUseCase(userRepo)
	listUsers := user.NewListUsersUseCase(userRepo)

	userHandler := httpServer.NewUserHandler(registerUser, listUsers)

	router := gin.Default()
	router.Use(errorHandler())

	router.POST("/users", userHandler.Register)
	router.GET("/users", userHandler.List)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
