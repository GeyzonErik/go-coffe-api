package http

import (
	"net/http"
	apikey "product-recommendation/internal/domain/api-key"

	"github.com/gin-gonic/gin"
)

type APIKeyHandler struct {
	repository apikey.Repository
}

func NewAPIKeyHandler(repo apikey.Repository) *APIKeyHandler {
	return &APIKeyHandler{repository: repo}
}

type CreateAPIKeyRequest struct {
	SystemName string `json:"systemName" binding:"required"`
}

func (handler *APIKeyHandler) CreateAPIKey(context *gin.Context) {
	var request CreateAPIKeyRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err": "invalid request"})
		return
	}

	apikey, err := handler.repository.CreateAPIKey(request.SystemName)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create your API Key"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"apiKey":     apikey.Key,
		"systemName": apikey.SystemName,
	})
}
