package user

import (
	"product-recommendation/internal/domain/user"
	user_repository "product-recommendation/internal/infra/repository"
)

type ListUsersUseCase struct {
	repository user_repository.UserRepository
}

func NewListUsersUseCase(repo user_repository.UserRepository) *ListUsersUseCase {
	return &ListUsersUseCase{repository: repo}
}

func (userClient *ListUsersUseCase) Execute() ([]*user.User, error) {
	return userClient.repository.FindAll()
}
