package repository

import (
	"errors"
	"gorm.io/gorm"
	"prj/internal/entity"
)

type UserRepository interface {
	List() ([]entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) (UserRepository, error) {
	if db == nil {
		return nil, errors.New("provided db handle to user repo is nil")
	}

	return &userRepository{db: db}, nil
}

func (r *userRepository) List() ([]entity.User, error) {
	var users []entity.User

	return users, nil
}
