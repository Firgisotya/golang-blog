package seeders

import (
	"fmt"
	"go-blog/api/configs"
)

type DBSeeder interface {
	Run() error
}

func Seed() {
	configs.SetupConnection()

	fmt.Println("Running Seeders...")

	seeder := []DBSeeder{
		&RoleSeeder{},
		&UserSeeder{},
	}

	for _, seed := range seeder {
		if err := seed.Run(); err != nil {
			fmt.Println("Error Seed: ", err)
		}
	}

	fmt.Println("Seeders Complete...")

	configs.CloseConnection()
}
