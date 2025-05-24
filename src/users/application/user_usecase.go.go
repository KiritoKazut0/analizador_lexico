package application

import (
	"github.com/KiritoKazut0/analizador-lexico/src/users/domain/entities"
	"github.com/KiritoKazut0/analizador-lexico/src/users/domain/repositories"
)

type UserUseCase struct {
	repository repositories.UserRepository
}

func NewUserUseCase(repo repositories.UserRepository) *UserUseCase {
	return &UserUseCase{
		repository: repo,
	}
}

func (u *UserUseCase) CreateUser(user *entities.User) error {
	return u.repository.CreateUser(user)
}

func (u *UserUseCase) GetAllUsers() ([]entities.User, error) {
	return u.repository.GetAllUsers()
}

func (u *UserUseCase) GetUserByID(clave string) (*entities.User, error) {
	return u.repository.GetUserByID(clave)
}

func (u *UserUseCase) UpdateUser(user *entities.User) (*entities.User, error) {
	return u.repository.UpdateUser(user)
}

func (u *UserUseCase) DeleteUser(clave string) error {
	return u.repository.DeleteUser(clave)
}


