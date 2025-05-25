package repositories

import entities "github.com/KiritoKazut0/analizador-lexico/src/users/domain/entities"

type UserCacheRepository interface {
	SetUsers(users []entities.User) error
	GetUsers() ([]entities.User, error)
	SetUsersPaginated(page int, users []entities.User) error
	GetUsersPaginated(page int) ([]entities.User, error)
	DeleteUsersCache() error
	SetLastUpdateTimestamp(timestamp int64) error
	GetLastUpdateTimestamp() (int64, error)
}