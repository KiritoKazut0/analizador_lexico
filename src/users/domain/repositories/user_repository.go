package repositories

import "github.com/KiritoKazut0/analizador-lexico/src/users/domain/entities"

type UserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByID(clave string) (*entities.User, error)
	UpdateUser(user *entities.User) ( *entities.User, error)
	DeleteUser(clave string) error
	GetAllUsers() ([]*entities.User, error) 

}