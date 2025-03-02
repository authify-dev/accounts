package main

import (
	email_events "accounts/internal/api/v1/emails/domain/events"
	"accounts/internal/common/controllers"
	"accounts/internal/core/domain/event"
	"accounts/internal/core/infrastructure/event_bus/rabbit"
	"accounts/internal/core/settings"
	"encoding/json"
	"fmt"
	"time"
	// Asegúrate de que el paquete "queue" esté en el path correcto.
)

func eventBus() event.EventBus {

	connection := rabbit.NewRabbitMqConnection(
		event.SettingsEventBus{
			Username: settings.Settings.USER_EVENT_BUS,
			Password: settings.Settings.PASSWORD_EVENT_BUS,
			VHost:    settings.Settings.VHOST_EVENT_BUS,
			Connection: struct {
				Hostname string
				Port     int
			}{
				Hostname: settings.Settings.HOST_EVENT_BUS,
				Port:     settings.Settings.PORT_EVENT_BUS,
			},
		},
	)

	connection.Connect()

	event_bus := rabbit.NewRabbitEventBus(
		connection,
	)

	return event_bus
}

func main() {
	fmt.Println("accounts v0.0.1")

	settings.LoadDotEnv()

	settings.LoadEnvs()

	controlador := controllers.NewEmailController(

		settings.Settings.EMAIL_SENDER,
		settings.Settings.EMAIL_SENDER_PASSWORD,
		settings.Settings.EMAIL_CLIENT,
		"587",
	)

	eventBus := eventBus()

	response := eventBus.Consume()

	if response.Err != nil {
		fmt.Println(response.Err)
	}

	go func() {
		for msg := range response.Data {
			// Procesar el mensaje
			var data email_events.UserRegistered

			err := json.Unmarshal(msg.Body, &data)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("User Name: ", data.UserName)
			fmt.Println("Email: ", data.Email)
			fmt.Println("Code: ", data.CodeVerification)

			controlador.SendEmail(data.Email, "Código de verificación", "Código de verificación: "+data.CodeVerification)

		}
	}()

	// Mantener el programa en ejecución para poder consumir mensajes.
	fmt.Println("Esperando mensajes. Para salir presiona CTRL+C")

	for {
		time.Sleep(1 * time.Second)
	}

}
