package entities

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type User struct {
    ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
    Nombre    string    `gorm:"not null" json:"nombre"`
    Correo    string    `gorm:"not null" json:"correo"`
    Telefono  string    `gorm:"not null" json:"telefono"`
	Clave     string    `gorm:"not null" json:"clave"`
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
    user.ID = uuid.New()
    return
}
