package seeders

import (
	"go-blog/api/configs"
	"go-blog/api/entities"
	"go-blog/api/helpers"
)

type UserSeeder struct{}

func (u *UserSeeder) Run() error {
	db := configs.DB

	hashedPassword, _ := helpers.HashPassword("password")

	users := []entities.User{
		{RoleID: 1, FirstName: "Admin", LastName: "Admin", Username: "admin", Email: "admin@gmail.com", NoTelp: "08123456789", Photo: "admin.jpg", Password: string(hashedPassword)},
		{RoleID: 2, FirstName: "User", LastName: "User", Username: "user", Email: "user@gmail.com", NoTelp: "08123456789", Photo: "user.jpg", Password: string(hashedPassword)},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			panic(err)
		}
	}

	return nil
}
