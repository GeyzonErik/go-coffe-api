package user

import (
	"product-recommendation/internal/core/domain/user"
	memory_repository "product-recommendation/internal/core/infra/repository/repository_memory"
)

type ListUsersUseCase struct {
	repository memory_repository.UserRepository
}

func NewListUsersUseCase(repo memory_repository.UserRepository) *ListUsersUseCase {
	return &ListUsersUseCase{repository: repo}
}

func (userClient *ListUsersUseCase) Execute() ([]*user.User, error) {
	return userClient.repository.FindAll()
}
