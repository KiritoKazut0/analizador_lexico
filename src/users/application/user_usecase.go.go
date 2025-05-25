package application

import (
	entities "github.com/KiritoKazut0/analizador-lexico/src/users/domain/entities"
	repositories "github.com/KiritoKazut0/analizador-lexico/src/users/domain/repositories"
	"math"
	"time"
)

type UserUseCase struct {
	repository      repositories.UserRepository
	cacheRepository repositories.UserCacheRepository
}

func NewUserUseCase(repo repositories.UserRepository, cacheRepo repositories.UserCacheRepository) *UserUseCase {
	return &UserUseCase{
		repository:      repo,
		cacheRepository: cacheRepo,
	}
}

func (u *UserUseCase) CreateUser(user *entities.User) error {
	err := u.repository.CreateUser(user)
	if err != nil {
		return err
	}

	// Invalidar cache después de crear un usuario
	u.cacheRepository.DeleteUsersCache()
	u.cacheRepository.SetLastUpdateTimestamp(time.Now().Unix())

	return nil
}

func (u *UserUseCase) CreateUsersBatch(users []entities.User) error {
	err := u.repository.CreateUsersBatch(users)
	if err != nil {
		return err
	}

	// Invalidar cache después de crear usuarios en lote
	u.cacheRepository.DeleteUsersCache()
	u.cacheRepository.SetLastUpdateTimestamp(time.Now().Unix())

	return nil
}

func (u *UserUseCase) GetAllUsersPaginated(page, perPage int) (*entities.PaginatedUsersResponse, error) {
	// Intentar obtener datos paginados desde cache
	cachedUsers, err := u.cacheRepository.GetUsersPaginated(page)
	if err == nil && cachedUsers != nil {
		// Obtener total count desde cache o base de datos
		totalCount, err := u.repository.GetTotalUsersCount()
		if err != nil {
			return nil, err
		}

		totalPages := int(math.Ceil(float64(totalCount) / float64(perPage)))

		return &entities.PaginatedUsersResponse{
			Users:       cachedUsers,
			CurrentPage: page,
			TotalPages:  totalPages,
			TotalCount:  totalCount,
			PerPage:     perPage,
		}, nil
	}

	// Si no hay datos en cache, obtener desde base de datos
	offset := (page - 1) * perPage
	users, err := u.repository.GetUsersWithPagination(offset, perPage)
	if err != nil {
		return nil, err
	}

	totalCount, err := u.repository.GetTotalUsersCount()
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(perPage)))

	// Guardar en cache
	u.cacheRepository.SetUsersPaginated(page, users)

	return &entities.PaginatedUsersResponse{
		Users:       users,
		CurrentPage: page,
		TotalPages:  totalPages,
		TotalCount:  totalCount,
		PerPage:     perPage,
	}, nil
}

func (u *UserUseCase) GetAllUsers() ([]entities.User, error) {
	// Intentar obtener desde cache primero
	cachedUsers, err := u.cacheRepository.GetUsers()
	if err == nil && cachedUsers != nil {
		return cachedUsers, nil
	}

	// Si no hay datos en cache, verificar si hay datos nuevos desde la última actualización
	lastUpdate, err := u.cacheRepository.GetLastUpdateTimestamp()
	if err == nil && lastUpdate > 0 {
		// Obtener solo usuarios nuevos
		newUsers, err := u.repository.GetUsersAfterTimestamp(lastUpdate)
		if err != nil {
			return nil, err
		}

		if len(newUsers) > 0 {
			// Combinar con usuarios existentes en cache si los hay
			if cachedUsers != nil {
				allUsers := append(cachedUsers, newUsers...)
				u.cacheRepository.SetUsers(allUsers)
				return allUsers, nil
			}
		} else if cachedUsers != nil {
			// No hay datos nuevos, devolver cache existente
			return cachedUsers, nil
		}
	}

	// Obtener todos los usuarios desde la base de datos
	users, err := u.repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	// Guardar en cache
	u.cacheRepository.SetUsers(users)
	u.cacheRepository.SetLastUpdateTimestamp(time.Now().Unix())

	return users, nil
}

func (u *UserUseCase) GetUserByID(clave string) (*entities.User, error) {
	return u.repository.GetUserByID(clave)
}

func (u *UserUseCase) UpdateUser(user *entities.User) (*entities.User, error) {
	updatedUser, err := u.repository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	// Invalidar cache después de actualizar
	u.cacheRepository.DeleteUsersCache()
	u.cacheRepository.SetLastUpdateTimestamp(time.Now().Unix())

	return updatedUser, nil
}

func (u *UserUseCase) DeleteUser(clave string) error {
	err := u.repository.DeleteUser(clave)
	if err != nil {
		return err
	}

	// Invalidar cache después de eliminar
	u.cacheRepository.DeleteUsersCache()
	u.cacheRepository.SetLastUpdateTimestamp(time.Now().Unix())

	return nil
}
