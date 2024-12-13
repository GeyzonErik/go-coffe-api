package user

import (
	"product-recommendation/internal/core/domain/user"
	memory_repository "product-recommendation/internal/core/infra/repository/repository_memory"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUserUseCase struct {
	repository memory_repository.UserRepository
}

func NewRegisterUserUseCase(repo memory_repository.UserRepository) *RegisterUserUseCase {
	return &RegisterUserUseCase{repository: repo}
}

func (registerUser *RegisterUserUseCase) Execute(id, name, email, password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err
	}

	newUser, err := user.NewUser(id, name, email, string(bytes))

	if err != nil {
		return err
	}

	return registerUser.repository.Save(newUser)
}
