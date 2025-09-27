package main

import (
	"accounts/internal/core/settings"
	apikeys_gorm "accounts/internal/db/postgres/api_keys"
	codes "accounts/internal/db/postgres/codes"
	emails "accounts/internal/db/postgres/emails"
	login_methods "accounts/internal/db/postgres/login_methods"
	oauth_logins "accounts/internal/db/postgres/oauth_logins"
	organizations_pg "accounts/internal/db/postgres/organinizations"
	pending_registrations "accounts/internal/db/postgres/pending_registrations"
	policies_pg "accounts/internal/db/postgres/policies"
	refreshtokens "accounts/internal/db/postgres/refresh_tokens"
	role "accounts/internal/db/postgres/role"
	role_policies_pg "accounts/internal/db/postgres/role_policies"
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

	if settings.Settings.DB_SCHEMA != "" {
		db.Exec("CREATE SCHEMA IF NOT EXISTS " + settings.Settings.DB_SCHEMA)
	}

	// Migrate the schema

	// Migrate roles
	err = db.AutoMigrate(&role.RoleModel{})
	if err != nil {
		panic("failed to migrate roles: " + err.Error())
	}

	// Migrate users
	err = db.AutoMigrate(&users.UserModel{})
	if err != nil {
		panic("failed to migrate users: " + err.Error())
	}

	// Migrate emails
	err = db.AutoMigrate(&emails.EmailModel{})
	if err != nil {
		panic("failed to migrate emails: " + err.Error())
	}

	// Migrate codes
	err = db.AutoMigrate(&codes.CodeModel{})
	if err != nil {
		panic("failed to migrate codes: " + err.Error())
	}

	// Migrate oauth logins
	err = db.AutoMigrate(&oauth_logins.OAuthLoginModel{})
	if err != nil {
		panic("failed to migrate oauth logins: " + err.Error())
	}

	// Migrate login methods
	err = db.AutoMigrate(&login_methods.LoginMethodModel{})
	if err != nil {
		panic("failed to migrate login methods: " + err.Error())
	}

	// Migrate refresh tokens
	err = db.AutoMigrate(&refreshtokens.RefreshTokenModel{})
	if err != nil {
		panic("failed to migrate refresh tokens: " + err.Error())
	}

	// Migrate pending registrations
	err = db.AutoMigrate(&pending_registrations.PendingRegistrationModel{})
	if err != nil {
		panic("failed to migrate pending registrations: " + err.Error())
	}

	// Migrate policies
	err = db.AutoMigrate(&policies_pg.PolicyModel{})
	if err != nil {
		panic("failed to migrate policies: " + err.Error())
	}

	// Migrate role policies
	err = db.AutoMigrate(&role_policies_pg.RolePoliciesModel{})
	if err != nil {
		panic("failed to migrate role policies: " + err.Error())
	}

	// Migrate organizations
	err = db.AutoMigrate(&organizations_pg.OrganizationModel{})
	if err != nil {
		panic("failed to migrate organizations: " + err.Error())
	}

	// Migrate api keys
	err = db.AutoMigrate(&apikeys_gorm.APIKeyModel{})
	if err != nil {
		panic("failed to migrate api keys: " + err.Error())
	}

	fmt.Println("Migrations completed")
}
