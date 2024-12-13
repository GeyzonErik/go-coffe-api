package main

import (
	"fmt"
	"net/http"
	"os"

	auth "product-recommendation/internal/application/services"
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

	jwtSecret := os.Getenv("JWT_SECRET")
	authSerive := auth.NewAuthSerive(jwtSecret)

	userRepo := user_repository.NewInMemoryUserRepository()
	registerUser := user.NewRegisterUserUseCase(userRepo)
	listUsers := user.NewListUsersUseCase(userRepo)
	findUser := user.NewFindUserUseCase(userRepo)

	userHandler := httpServer.NewUserHandler(registerUser, listUsers, findUser)
	authHandler := httpServer.NewAuthHandler(authSerive, userHandler)

	router := gin.Default()
	router.Use(errorHandler())

	router.POST("/login", authHandler.Login)

	router.POST("/users", userHandler.Register)

	usersGroup := router.Group("/users")
	usersGroup.Use(httpServer.AuthMiddleware(authSerive))
	usersGroup.GET("/", userHandler.List)
	usersGroup.GET("/:email", userHandler.Find)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
