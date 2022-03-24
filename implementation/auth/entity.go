package auth

import (
	"time"

	"gorm.io/gorm"
)

// User represents the user model that will be stored in DB and use for internal business logic
type User struct {
	Model
	UserName string `gorm:"user_name"`
	FullName string `gorm:"full_name"`
	Password string `gorm:"password"`
}

// Model is a helper struct for embedding common fields
type Model struct {
	ID        uint64         `gorm:"id"`
	CreatedAt time.Time      `gorm:"created_at"`
	UpdatedAt time.Time      `gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at"`
}
