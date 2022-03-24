package auth

import (
	"github.com/go-logr/logr"
	"gorm.io/gorm"
)

// Repository exposes the auth related database interactions.
type Repository interface {
	Create(u User) (uint64, error)
	GetByFilter(where map[string]interface{}) (User, error)
}

type repo struct {
	db        *gorm.DB
	l         logr.Logger
	tableName string
}

// NewRepository is a constructor for Auth Reposistory.
func NewRepository(conn *gorm.DB, logger logr.Logger) Repository {
	return &repo{
		db:        conn,
		l:         logger,
		tableName: "users",
	}
}

// Create store a user in datababse.
func (r *repo) Create(u User) (uint64, error) {
	err := r.db.Table(r.tableName).Create(&u).Error
	return u.ID, err
}

// GetByFilter accepts a where clause to filter users from database.
// This function will return only one record if it finds any.
func (r *repo) GetByFilter(where map[string]interface{}) (User, error) {
	var u User

	err := r.db.Table(r.tableName).Where(where).First(&u).Error
	return u, err
}
