package rabbit

import (
	"accounts/internal/core/domain/event"
	"accounts/internal/utils"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMqConnection struct {
	amqpConnection *amqp091.Connection
	amqpChannel    *amqp091.Channel
	settings       event.SettingsEventBus
}

func NewRabbitMqConnection(settings event.SettingsEventBus) *RabbitMqConnection {

	return &RabbitMqConnection{
		settings: settings,
	}
}

func (r *RabbitMqConnection) amqpConnect() (*amqp091.Connection, error) {

	dsn := fmt.Sprintf(
		"amqp://%s:%s@%s:%d",
		r.settings.Username,
		r.settings.Password,
		r.settings.Connection.Hostname,
		r.settings.Connection.Port,
	)

	conn, err := amqp091.Dial(dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (r *RabbitMqConnection) amqpChannelConnect() (*amqp091.Channel, error) {

	ch, err := r.amqpConnection.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func (r *RabbitMqConnection) Connect() error {

	conn, err := r.amqpConnect()
	if err != nil {
		return err
	}

	r.amqpConnection = conn

	ch, err := r.amqpChannelConnect()
	if err != nil {
		return err
	}

	r.amqpChannel = ch

	return nil
}

func (r *RabbitMqConnection) Close() error {

	if err := r.amqpChannel.Close(); err != nil {
		return err
	}

	if err := r.amqpConnection.Close(); err != nil {
		return err
	}

	return nil
}

func (r *RabbitMqConnection) Publish(
	exchangeName string,
	routingKey string,
	data []byte,
	opetions event.OptionsEventBus,
) error {
	err := r.amqpChannel.ExchangeDeclare(
		exchangeName, // nombre del exchange
		"topic",      // tipo de exchange
		true,         // durable
		false,        // auto-delete
		false,        // internal
		false,        // no-wait
		nil,
	)
	if err != nil {
		return err
	}

	err = r.amqpChannel.Publish(
		exchangeName,
		routingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType:     opetions.ContentType,
			Body:            data,
			Headers:         opetions.Headers,
			MessageId:       opetions.MessageID,
			ContentEncoding: opetions.ContentEncoding,
			Priority:        opetions.Priority,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitMqConnection) Consume(
	exchangeName string,
	options event.OptionsQueue,
) utils.Result[<-chan amqp091.Delivery] {
	err := r.amqpChannel.ExchangeDeclare(
		exchangeName, // nombre del exchange
		"topic",      // tipo de exchange
		true,         // durable
		false,        // auto-delete
		false,        // internal
		false,        // no-wait
		nil,
	)
	if err != nil {
		return utils.Result[<-chan amqp091.Delivery]{Err: err}
	}

	queue, err := r.amqpChannel.QueueDeclare(
		options.Name,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return utils.Result[<-chan amqp091.Delivery]{Err: err}
	}

	err = r.amqpChannel.QueueBind(
		queue.Name,
		options.Key,
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		return utils.Result[<-chan amqp091.Delivery]{Err: err}
	}

	deliveries, err := r.amqpChannel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return utils.Result[<-chan amqp091.Delivery]{Err: err}
	}

	return utils.Result[<-chan amqp091.Delivery]{Data: deliveries}
}
