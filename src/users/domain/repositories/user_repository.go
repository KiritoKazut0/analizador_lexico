package repositories

import "github.com/KiritoKazut0/analizador-lexico/src/users/domain/entities"

type UserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByID(id int) (*entities.User, error)
	UpdateUser(user *entities.User) error
	DeleteUser(id int) error
	GetAllUsers() ([]*entities.User, error) 

}