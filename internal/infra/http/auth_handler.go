package http

import (
	"net/http"

	auth "product-recommendation/internal/application/services"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthHandler struct {
	authService *auth.AuthService
	userHandler *UserHandler
}

func NewAuthHandler(authService *auth.AuthService, userHandler *UserHandler) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userHandler: userHandler,
	}
}

func (handler *AuthHandler) Login(context *gin.Context) {
	var request LoginRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "json invalido"})
		return
	}

	if err := validate.Struct(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, findUserErr := handler.userHandler.findOneUseCase.Execute(request.Email)

	if findUserErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": findUserErr.Error()})
		return
	}

	token, err := handler.authService.GenerateToken(user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "falha ao gerar token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}
