package entities

import "time"

type User struct {
	Clave     string    `gorm:"primaryKey;not null" json:"clave"`
	Nombre    string    `gorm:"not null" json:"nombre"`
	Correo    string    `gorm:"not null" json:"correo"`
	Telefono  string    `gorm:"not null" json:"telefono"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
