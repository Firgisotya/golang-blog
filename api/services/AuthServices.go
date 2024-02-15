package services

import (
	"go-blog/api/entities"
	"go-blog/api/payloads/request"
	"go-blog/api/repositories"

	"gorm.io/gorm"
)

func LoginService(db *gorm.DB, LoginRequest request.LoginRequest) (entities.User, error) {
	repo := repositories.NewAuthRepositories(db)

	user, err := repo.LoginRepositories(LoginRequest)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}
