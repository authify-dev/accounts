package main

import (
	"accounts/internal/core/settings"
	postgres_role "accounts/internal/db/postgres/role"
	postgres_users "accounts/internal/db/postgres/users"

	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("accounts v0.0.1")

	settings.LoadDotEnv()

	settings.LoadEnvs()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&postgres_role.RoleModel{})
	db.AutoMigrate(&postgres_users.UserModel{})

}
