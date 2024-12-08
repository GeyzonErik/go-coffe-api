package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	useCase "product-recommendation/internal/application/usecases"
	handler "product-recommendation/internal/infra/http"
	user_repository "product-recommendation/internal/infra/repository"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_Register(test *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := user_repository.NewInMemoryUserRepository()
	registerUseCase := useCase.NewRegisterUserUseCase(repo)
	listUseCase := useCase.NewListUsersUseCase(repo)
	handler := handler.NewUserHandler(registerUseCase, listUseCase)

	router := gin.Default()
	router.POST("/users", handler.Register)

	body := map[string]string{
		"id":       "1234",
		"name":     "João das Neves",
		"email":    "mail@mail.com",
		"password": "123456",
	}
	jsonBody, _ := json.Marshal(body)

	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBody))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(test, http.StatusCreated, response.Code)
	assert.Contains(test, response.Body.String(), "Usuário registrado com sucesso")

}

func TestUserHandler_List(test *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := user_repository.NewInMemoryUserRepository()
	registerUseCase := useCase.NewRegisterUserUseCase(repo)
	listUseCase := useCase.NewListUsersUseCase(repo)
	handler := handler.NewUserHandler(registerUseCase, listUseCase)

	router := gin.Default()
	router.GET("/users", handler.List)

	request, _ := http.NewRequest("GET", "/users", nil)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	assert.Equal(test, http.StatusOK, response.Code)

}
