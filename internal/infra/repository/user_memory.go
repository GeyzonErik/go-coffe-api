package user_repository

import (
	"errors"
	"product-recommendation/internal/domain/user"
	"sync"
)

type UserRepository interface {
	Save(*user.User) error
	FindOne(email string) (*user.User, error)
	FindAll() ([]*user.User, error)
}

type InMemoryUserRepository struct {
	data map[string]*user.User
	mu   sync.RWMutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		data: make(map[string]*user.User),
	}
}

func (request *InMemoryUserRepository) Save(user *user.User) error {
	request.mu.Lock()
	defer request.mu.Unlock()
	request.data[user.ID] = user
	return nil
}

func (request *InMemoryUserRepository) FindAll() ([]*user.User, error) {
	request.mu.Lock()
	defer request.mu.Unlock()

	users := make([]*user.User, 0, len(request.data))

	for _, user := range request.data {
		users = append(users, user)
	}

	return users, nil
}

func (request *InMemoryUserRepository) FindOne(email string) (*user.User, error) {
	request.mu.Lock()
	defer request.mu.Unlock()

	for _, user := range request.data {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.New("usuário não encontrado")
}
