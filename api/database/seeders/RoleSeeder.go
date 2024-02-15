package seeders

import (
	"go-blog/api/configs"
	"go-blog/api/entities"
)

type RoleSeeder struct{}

func (u *RoleSeeder) Run() error {
	db := configs.DB

	roles := []entities.Role{
		{NameRole: "Admin"},
		{NameRole: "User"},
	}

	for _, role := range roles {
		if err := db.Create(&role).Error; err != nil {
			panic(err)
		}
	}

	return nil

}
