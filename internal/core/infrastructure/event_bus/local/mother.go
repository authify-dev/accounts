package local

import "accounts/internal/core/infrastructure/event_bus/local/actions"

func MockEventBus() *LocalEventBus {
	eb := NewLocalEventBus()

	eb.AddAction("user.registered", actions.SendActivationEmail)
	eb.AddAction("user.activated", actions.SendWelcomeEmail)
	eb.AddAction("user.reset_password", actions.SendResetPasswordEmail)

	return eb
}
