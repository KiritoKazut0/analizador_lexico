package application

import (
	"github.com/KiritoKazut0/analizador-lexico/src/users/domain/entities"
	"github.com/KiritoKazut0/analizador-lexico/src/users/domain/repositories"
)

type UserUseCase struct {
	repository repositories.UserRepository
}

func (u *UserUseCase) CreateUser(user *entities.User) error {
	return u.repository.CreateUser(user)
}

func (u *UserUseCase) GetAllUsers() ([]*entities.User, error) {
	return u.repository.GetAllUsers()
}

func (u *UserUseCase) GetUserByID(id int) (*entities.User, error) {
	return u.repository.GetUserByID(id)
}

func (u *UserUseCase) UpdateUser(user *entities.User) error {
	return u.repository.UpdateUser(user)
}

func (u *UserUseCase) DeleteUser(id int) error {
	return u.repository.DeleteUser(id)
}



