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

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("accounts v0.0.1")

	// Carga las variables de entorno
	settings.LoadDotEnv()
	settings.LoadEnvs()

	// Define el DSN para la conexión a PostgreSQL
	dsn := settings.Settings.POSTGRES_DSN

	// Conecta a la base de datos Postgres
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
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
