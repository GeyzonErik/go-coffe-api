package apikey_repository

import (
	"errors"
	"math/rand"
	apikey "product-recommendation/internal/domain/api-key"
	"sync"

	"github.com/google/uuid"
)

type InMemoryAPIKeyRepository struct {
	keys map[string]*apikey.APIKey
	mu   sync.Mutex
}

func NewInMemoryAPIKeyRepository() *InMemoryAPIKeyRepository {
	return &InMemoryAPIKeyRepository{
		keys: make(map[string]*apikey.APIKey),
	}
}

func (repo *InMemoryAPIKeyRepository) GetSystemByKey(apiKey string) (*apikey.APIKey, error) {
	if system, exists := repo.keys[apiKey]; exists {
		return system, nil
	}

	return nil, errors.New("invalid API key")
}

func (repo *InMemoryAPIKeyRepository) CreateAPIKey(systemName string) (*apikey.APIKey, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	apikey := &apikey.APIKey{
		Key:        generateRandomKey(),
		SystemName: systemName,
	}

	repo.keys[apikey.Key] = apikey

	return apikey, nil
}

func generateRandomKey() string {

	return "Key-" + uuid.New().String() + RandStringBytes(14)
}

func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
