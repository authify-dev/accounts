package queue

import (
	"accounts/cmd/queue/utils"
	"accounts/internal/common/controllers"
	"accounts/internal/common/logger"
	"accounts/internal/core/domain/event"
	"accounts/internal/core/settings"
	"context"
	"encoding/json"
	"fmt"
	"time"

	email_events "accounts/internal/api/v1/emails/domain/events"
)

func SendActivationEmails(eventBus event.EventBus) {
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

	response := eventBus.Consume("user.registered", "user.registered")

	if response.Err != nil {
		fmt.Println(response.Err)
	}

	entry := logger.FromContext(context.Background())

	go func() {
		for msg := range response.Data {
			// Procesar el mensaje
			var data email_events.UserRegistered

			err := json.Unmarshal(msg.Body, &data)
			if err != nil {
				entry.Errorln(err)
			}

			entry.Info("User Name: ", data.UserName)
			entry.Info("Email: ", data.Email)
			entry.Info("Code: ", data.CodeVerification)

			body := utils.GenerateBodyActivation(data.UserName, data.CodeVerification)

			if body.Err != nil {
				entry.Errorln(body.Err)

				// Falatria reencolar el evento
			}

			controlador.SendEmail(data.Email, "Código de verificación", body.Data)

			entry.Info("Email sent")
		}
	}()

	// Mantener el programa en ejecución para poder consumir mensajes.
	fmt.Println("Esperando mensajes. Para salir presiona CTRL+C")

	for {
		time.Sleep(1 * time.Second)
	}

}
