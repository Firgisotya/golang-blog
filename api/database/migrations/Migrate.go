package migrations

import (
	"fmt"
	"go-blog/api/configs"
	"go-blog/api/entities"
)

func Migrate() {
	fmt.Println("Running Migrations...")

	configs.SetupConnection()

	for _, entity := range entities.RegisterEntities() {
		// Drop Table
		err := configs.DB.Migrator().DropTable(entity.Entities)
		if err != nil {
			fmt.Println("Error Drop Table: ", err)
		}

		// AutoMigrate
		err = configs.DB.AutoMigrate(entity.Entities)
		if err != nil {
			fmt.Println("Error AutoMigrate: ", err)
		}

	}

	fmt.Println("Migrations Complete...")

	configs.CloseConnection()
}
