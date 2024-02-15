package repositories

import (
	"errors"
	"go-blog/api/entities"
	"go-blog/api/payloads/request"

	"gorm.io/gorm"
)

type AuthRepositories interface {
	LoginRepositories(request.LoginRequest) (entities.User, error)
}

type authRepositories struct {
	db *gorm.DB
}

func NewAuthRepositories(db *gorm.DB) AuthRepositories {
	return &authRepositories{db}
}

func (r *authRepositories) LoginRepositories(LoginRequest request.LoginRequest) (entities.User, error) {
	var user entities.User

	if err := r.db.Preload("Role").Where("email = ?", LoginRequest.Email).First(&user).Error; err != nil {
		return user, errors.New("Record not found")
	}

	return user, nil
}
