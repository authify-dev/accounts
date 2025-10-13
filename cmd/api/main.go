package main

import (
	"accounts/internal/api/server"
	"accounts/internal/core/settings"
	"fmt"
	"foundation/domain/logger"
)

func main() {
	fmt.Println("accounts v0.0.1")

	settings.LoadDotEnv()
	settings.LoadEnvs()

	logger.InitLogger(settings.Settings.ENVIRONMENT, settings.Settings.APP_NAME, settings.Settings.LOKI_URL)

	//eventBus := queue.SetUpEventBus()

	//go queue.SendActivationEmails(eventBus)
	//go queue.SendWelcomeEmails(eventBus)
	//go queue.SendResetPasswordEmails(eventBus)

	server.Run()
}
