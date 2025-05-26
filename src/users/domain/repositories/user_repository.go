package repositories

import entities "github.com/KiritoKazut0/analizador-lexico/src/users/domain/entities"

type UserRepository interface {
	CreateUser(user *entities.User) error
	CreateUsersBatch(users []entities.User) error
	GetUserByID(clave string) (*entities.User, error)
	UpdateUser( clave string ,user *entities.User) (*entities.User, error)
	DeleteUser(clave string) error
	GetAllUsers() ([]entities.User, error)
	GetUsersWithPagination(offset, limit int) ([]entities.User, error)
	GetTotalUsersCount() (int64, error)
	GetUsersAfterTimestamp(timestamp int64) ([]entities.User, error)
}
