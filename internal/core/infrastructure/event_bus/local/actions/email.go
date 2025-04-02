package actions

import (
	"accounts/cmd/queue/utils"
	"accounts/internal/common/controllers"
	"accounts/internal/common/logger"
	"accounts/internal/core/domain/event"
	"accounts/internal/core/settings"
	"context"
	"encoding/json"
	"fmt"

	email_events "accounts/internal/api/v1/emails/domain/events"
)

func SendActivationEmail(ctx context.Context, event event.DomainEvent) {
	entry := logger.FromContext(ctx)
	entry.Info("Send Activation Account Email")

	controlador := controllers.NewEmailController(

		settings.Settings.EMAIL_SENDER,
		settings.Settings.EMAIL_SENDER_PASSWORD,
		settings.Settings.EMAIL_CLIENT,
		"587",
	)

	// Se obtiene el map[string]interface{} de ToPrimitive()
	msg := event.ToPrimitive()

	// Convertir el map a bytes en formato JSON
	jsonData, err := json.Marshal(msg)
	if err != nil {
		entry.Errorln("Error al convertir a JSON:", err)
		return
	}

	// Deserializar los bytes JSON en la estructura UserRegistered
	var data email_events.UserRegistered
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		entry.Errorln("Error al deserializar JSON:", err)
		return
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

func SendWelcomeEmail(ctx context.Context, event event.DomainEvent) {
	entry := logger.FromContext(ctx)
	entry.Info("Send Activation Account Email")

	controlador := controllers.NewEmailController(

		settings.Settings.EMAIL_SENDER,
		settings.Settings.EMAIL_SENDER_PASSWORD,
		settings.Settings.EMAIL_CLIENT,
		"587",
	)

	// Se obtiene el map[string]interface{} de ToPrimitive()
	msg := event.ToPrimitive()

	// Convertir el map a bytes en formato JSON
	jsonData, err := json.Marshal(msg)
	if err != nil {
		entry.Errorln("Error al convertir a JSON:", err)
		return
	}

	// Deserializar los bytes JSON en la estructura UserRegistered
	var data email_events.UserActivated

	err = json.Unmarshal(jsonData, &data)
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

func SendResetPasswordEmail(ctx context.Context, event event.DomainEvent) {
	entry := logger.FromContext(ctx)
	entry.Info("Send Reset Password Email")

	controlador := controllers.NewEmailController(

		settings.Settings.EMAIL_SENDER,
		settings.Settings.EMAIL_SENDER_PASSWORD,
		settings.Settings.EMAIL_CLIENT,
		"587",
	)

	// Se obtiene el map[string]interface{} de ToPrimitive()
	msg := event.ToPrimitive()

	// Convertir el map a bytes en formato JSON
	jsonData, err := json.Marshal(msg)
	if err != nil {
		entry.Errorln("Error al convertir a JSON:", err)
		return
	}

	// Deserializar los bytes JSON en la estructura UserRegistered
	var data email_events.UserRegistered
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		entry.Errorln("Error al deserializar JSON:", err)
		return
	}

	entry.Info("User Name: ", data.UserName)
	entry.Info("Email: ", data.Email)
	entry.Info("Code: ", data.CodeVerification)

	body := utils.GenerateBodyActivation(data.UserName, data.CodeVerification)

	if body.Err != nil {
		entry.Errorln(body.Err)

		// Falatria reencolar el evento
	}

	controlador.SendEmail(data.Email, "Código para reseteo de contraseña", body.Data)

	entry.Info("Email sent")
}
