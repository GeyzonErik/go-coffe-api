package user

import (
	"errors"
	"product-recommendation/internal/domain/user"
	user_repository "product-recommendation/internal/infra/repository"
)

type FindUserUseCase struct {
	repository user_repository.UserRepository
}

func NewFindUserUseCase(repo user_repository.UserRepository) *FindUserUseCase {
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
