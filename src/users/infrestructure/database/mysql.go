package database

import (
	"time"
	"github.com/KiritoKazut0/analizador-lexico/src/users/domain/entities"
	"gorm.io/gorm"
)

type MysqlRepository struct {
	db *gorm.DB
}

func NewUserMysqlRepository(db *gorm.DB) *MysqlRepository {
	return &MysqlRepository{db: db}
}

func (m *MysqlRepository) GetAllUsers() ([]entities.User, error) {
	var users []entities.User
	result := m.db.Find(&users)
	return users, result.Error
}


func (m *MysqlRepository) GetUserByID(clave string) (*entities.User, error) {
	var user entities.User
	result := m.db.Where("clave = ?", clave).First(&user)
	return &user, result.Error
}

func (m *MysqlRepository) CreateUser(user *entities.User) error {
	return m.db.Create(user).Error
}

func (m *MysqlRepository) CreateUsersBatch(users []entities.User) error {
	const batchSize = 1000 // Insertar en lotes de 1000

	for i := 0; i < len(users); i += batchSize {
		end := i + batchSize
		if end > len(users) {
			end = len(users)
		}

		batch := users[i:end]
		if err := m.db.CreateInBatches(batch, batchSize).Error; err != nil {
			return err
		}
	}

	return nil
}

func (m *MysqlRepository) UpdateUser(user *entities.User) (*entities.User, error) {
	err := m.db.Model(&entities.User{}).Where("clave = ?", user.Clave).Updates(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *MysqlRepository) DeleteUser(clave string) error {
	return m.db.Where("clave = ?", clave).Delete(&entities.User{}).Error
}

func (m *MysqlRepository) GetUsersWithPagination(offset, limit int) ([]entities.User, error) {
	var users []entities.User
	result := m.db.Offset(offset).Limit(limit).Find(&users)
	return users, result.Error
}

func (m *MysqlRepository) GetTotalUsersCount() (int64, error) {
	var count int64
	result := m.db.Model(&entities.User{}).Count(&count)
	return count, result.Error
}

func (m *MysqlRepository) GetUsersAfterTimestamp(timestamp int64) ([]entities.User, error) {
	var users []entities.User
	timeAfter := time.Unix(timestamp, 0)
	result := m.db.Where("updated_at > ?", timeAfter).Find(&users)
	return users, result.Error
}