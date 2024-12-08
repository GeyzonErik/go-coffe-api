package user

import (
	"product-recommendation/internal/domain/user"
	user_repository "product-recommendation/internal/infra/repository"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUserUseCase struct {
	repository user_repository.UserRepository
}

func NewRegisterUserUseCase(repo user_repository.UserRepository) *RegisterUserUseCase {
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
