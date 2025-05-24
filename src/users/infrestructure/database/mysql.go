package database

import (
	"github.com/KiritoKazut0/analizador-lexico/src/users/domain/entities"
	"gorm.io/gorm"
)

type MysqlRepository struct {
	db * gorm.DB
}


func NewUserMysqlRepository(db * gorm.DB) *MysqlRepository {
	return &MysqlRepository{ db: db }
}

func (m *MysqlRepository) GetAllUsers() ([]entities.User, error){
	 var users  []entities.User
	 result := m.db.Find(&users)
	 return users, result.Error
}

func (m *MysqlRepository) GetUserByID(clave string) (entities.User, error) {
	var user entities.User
	result := m.db.First(&user, clave)
	return user, result.Error
}

func (m *MysqlRepository) CreateUser(user *entities.User) error {
	return m.db.Create(user).Error
}


func (m *MysqlRepository) UpdateUser(user *entities.User) error {
    return m.db.Model(&entities.User{}).Where("clave = ?", user.Clave).Updates(user).Error
}

func (m *MysqlRepository) DeleteUser(clave string) error {
	return m.db.Delete( &entities.User{} , clave).Error
}
