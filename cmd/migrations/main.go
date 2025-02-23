package main

import (
	"accounts/internal/core/settings"
	codes "accounts/internal/db/postgres/codes"
	emails "accounts/internal/db/postgres/emails"
	login_methods "accounts/internal/db/postgres/login_methods"
	oauth_logins "accounts/internal/db/postgres/oauth_logins"
	refreshtokens "accounts/internal/db/postgres/refresh_tokens"
	role "accounts/internal/db/postgres/role"
	users "accounts/internal/db/postgres/users"

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
	db.AutoMigrate(&role.RoleModel{})
	db.AutoMigrate(&users.UserModel{})
	db.AutoMigrate(&emails.EmailModel{})
	db.AutoMigrate(&codes.CodeModel{})
	db.AutoMigrate(&oauth_logins.OAuthLoginModel{})
	db.AutoMigrate(&login_methods.LoginMethodModel{})
	db.AutoMigrate(&refreshtokens.RefreshTokenModel{})

}
