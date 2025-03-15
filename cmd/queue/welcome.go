package queue

import (
	"accounts/cmd/queue/utils"
	email_events "accounts/internal/api/v1/emails/domain/events"
	"accounts/internal/common/controllers"
	"accounts/internal/common/logger"
	"accounts/internal/core/domain/event"
	"accounts/internal/core/settings"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func SendWelcomeEmails(eventBus event.EventBus) {
	fmt.Println("queue accounts v0.0.1")

	//settings.LoadDotEnv()
	//
	//settings.LoadEnvs()

	controlador := controllers.NewEmailController(

		settings.Settings.EMAIL_SENDER,
		settings.Settings.EMAIL_SENDER_PASSWORD,
		settings.Settings.EMAIL_CLIENT,
		"587",
	)

	response := eventBus.Consume("user.activated", "user.activated")

	if response.Err != nil {
		fmt.Println(response.Err)
	}

	entry := logger.FromContext(context.Background())

	go func() {
		for msg := range response.Data {
			// Procesar el mensaje
			var data email_events.UserActivated

			err := json.Unmarshal(msg.Body, &data)
			if err != nil {
				entry.Errorln(err)
			}

			entry.Info("User Name: ", data.UserName)
			entry.Info("Email: ", data.Email)

			body := utils.GenerateBodyWelcome(data.UserName)

			if body.Err != nil {
				entry.Errorln(body.Err)

				// Falatria reencolar el evento
			}

			controlador.SendEmail(data.Email, fmt.Sprintf("Welcome to %s", settings.Settings.APP_NAME), body.Data)

			entry.Info("Email Welcome sent")
		}
	}()

	// Mantener el programa en ejecución para poder consumir mensajes.
	fmt.Println("Esperando mensajes. Para salir presiona CTRL+C")

	for {
		time.Sleep(1 * time.Second)
	}

}
