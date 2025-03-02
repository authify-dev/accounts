package rabbit

import (
	"accounts/internal/core/domain/event"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

/*

import amqplib from "amqplib";

export type Settings = {
	username: string;
	password: string;
	vhost: string;
	connection: {
		hostname: string;
		port: number;
	};
};

export class RabbitMqConnection {
	private amqpConnection?: amqplib.Connection;
	private amqpChannel?: amqplib.ConfirmChannel;
	private readonly settings: Settings = {
		username: "codely",
		password: "codely",
		vhost: "/",
		connection: {
			hostname: "localhost",
			port: 5672,
		},
	};

	async connect(): Promise<void> {
		this.amqpConnection = await this.amqpConnect();
		this.amqpChannel = await this.amqpChannelConnect();
	}

	async close(): Promise<void> {
		await this.channel().close();

		await this.connection().close();
	}

	async publish(
		exchange: string,
		routingKey: string,
		content: Buffer,
		options: {
			messageId: string;
			contentType: string;
			contentEncoding: string;
			priority?: number;
			headers?: unknown;
		},
	): Promise<void> {
		if (!this.amqpChannel) {
			await this.connect();
		}

		return new Promise((resolve: Function, reject: Function) => {
			this.channel().publish(exchange, routingKey, content, options, (error: unknown) =>
				error ? reject(error) : resolve(),
			);
		});
	}

	private connection(): amqplib.Connection {
		if (!this.amqpConnection) {
			throw new Error("RabbitMQ not connected");
		}

		return this.amqpConnection;
	}

	private channel(): amqplib.ConfirmChannel {
		if (!this.amqpChannel) {
			throw new Error("RabbitMQ channel not connected");
		}

		return this.amqpChannel;
	}

	private async amqpConnect(): Promise<amqplib.Connection> {
		const connection = await amqplib.connect({
			protocol: "amqp",
			hostname: this.settings.connection.hostname,
			port: this.settings.connection.port,
			username: this.settings.username,
			password: this.settings.password,
			vhost: this.settings.vhost,
		});

		connection.on("error", (error: unknown) => {
			throw error;
		});

		return connection;
	}

	private async amqpChannelConnect(): Promise<amqplib.ConfirmChannel> {
		const channel = await this.connection().createConfirmChannel();
		await channel.prefetch(1);

		return channel;
	}
}
*/

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
