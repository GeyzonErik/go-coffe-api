package user

import (
	"errors"
	"product-recommendation/internal/core/domain/user"
	memory_repository "product-recommendation/internal/core/infra/repository/repository_memory"
)

type FindUserUseCase struct {
	repository memory_repository.UserRepository
}

func NewFindUserUseCase(repo memory_repository.UserRepository) *FindUserUseCase {
	return &FindUserUseCase{
		repository: repo,
	}
}

func (userClient *FindUserUseCase) Execute(email string) (*user.User, error) {
	user, err := userClient.repository.FindOne(email)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return user, nil
}
