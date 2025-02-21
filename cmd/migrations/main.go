package main

import (
	"accounts/internal/core/settings"
	"accounts/internal/db/users/postgres"
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
	db.AutoMigrate(&postgres.UserModel{})
}
