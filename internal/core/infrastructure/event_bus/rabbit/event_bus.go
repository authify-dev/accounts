package rabbit

import (
	"accounts/internal/core/domain/event"
	"accounts/internal/utils"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitEventBus struct {
	connection *RabbitMqConnection
}

func NewRabbitEventBus(connection *RabbitMqConnection) *RabbitEventBus {
	return &RabbitEventBus{
		connection: connection,
	}
}

func (r *RabbitEventBus) Publish(events []event.DomainEvent) error {
	for _, e := range events {

		buffer := event.ToBytes(e.ToPrimitive())
		if buffer.Err != nil {
			return buffer.Err
		}

		err := r.connection.Publish(
			"domain_events",
			e.EventName(),
			buffer.Data,
			event.OptionsEventBus{
				MessageID:       e.EventID(),
				ContentType:     "application/json",
				ContentEncoding: "utf-8",
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RabbitEventBus) Consume() utils.Result[<-chan amqp091.Delivery] {
	return r.connection.Consume(
		"domain_events",
		event.OptionsQueue{
			Name: "user_events",
			Key:  "user.registered",
		},
	)
}
