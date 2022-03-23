package auth

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Model
	UserName string `gorm:"userName"`
	FullName string `gorm:"fullName"`
	Password string `gorm:"password"`
}

type Model struct {
	ID        uint           `gorm:"id"`
	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at"`
}
