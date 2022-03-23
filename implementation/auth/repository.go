package auth

import (
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	Create(u User) (uint, error)
}

type repo struct {
	db gorm.DB
	l  log.Logger
}

func NewAuthRepository(conn gorm.DB, logger log.Logger) Repository {
	return &repo{
		db: conn,
		l:  logger,
	}
}

func (r *repo) Create(u User) (uint, error) {
	err := r.db.Table("users").Create(&u).Error
	return u.ID, err
}
