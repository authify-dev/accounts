package server

import (
	"accounts/internal/api/health"
	"accounts/internal/api/router"
	apikeys "accounts/internal/api/v1/api_keys"
	"accounts/internal/api/v1/emails"
	"accounts/internal/api/v1/oauth_logins"
	"accounts/internal/api/v1/organizations"
	"accounts/internal/api/v1/pending_registrations"
	"accounts/internal/api/v1/policies"
	refreshtokens "accounts/internal/api/v1/refresh_tokens"
	"accounts/internal/api/v1/role_policies"
	"accounts/internal/api/v1/roles"
	"accounts/internal/api/v1/users"
	"os"

	"accounts/internal/core/settings"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Server *fiber.App

func Run() {

	app := setUpRouter()

	if _, inLambda := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME"); inLambda {
		fmt.Println("Running in Lambda")
		lambda.Start(ginadapter.NewV2(app).ProxyWithContext)
		return
	}

	app.Run(fmt.Sprintf(":%d", settings.Settings.PORT))
}

func setUpRouter() *gin.Engine {

	app := router.NewRouter()

	// app.Use(middlewares.TraceMiddleware())
	// //app.Use(middlewares.CatcherMiddleware)
	// app.Use(middlewares.LoggerMiddleware())

	db, err := gorm.Open(postgres.Open(settings.Settings.POSTGRES_DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	health.SetupHealthModule(app)
	roles.SetupRolesModule(app, db)
	users.SetupUsersModule(app, db)
	emails.SetupEmailsModule(app, db)
	refreshtokens.SetupRefreshTokensModule(app, db)
	oauth_logins.SetupOAuthModule(app, db)
	pending_registrations.SetupPendingRegistrationsModule(app, db)
	policies.SetupPoliciesModule(app, db)
	role_policies.SetupRolePoliciesModule(app, db)
	organizations.SetupOrganizationsModule(app, db)
	apikeys.SetupAPIKeysModule(app, db)
	return app
}
