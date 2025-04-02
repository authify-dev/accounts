package queue

import (
	"accounts/internal/core/domain/event"
	"accounts/internal/core/infrastructure/event_bus/rabbit"
	"accounts/internal/core/settings"
)

func SetUpEventBus() event.EventBus {

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
		"domain_events",
	)

	return event_bus
}
