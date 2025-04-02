package event

import (
	"accounts/internal/utils"
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
)

type EventBus interface {
	Publish(events []DomainEvent) error
	Consume(queue, key string) utils.Result[<-chan amqp091.Delivery]
}

type SettingsEventBus struct {
	Username   string
	Password   string
	VHost      string
	Connection struct {
		Hostname string
		Port     int
	}
}

type OptionsEventBus struct {
	MessageID       string
	ContentType     string
	ContentEncoding string
	Headers         map[string]interface{}
	Priority        uint8
}

func ToBytes(data map[string]interface{}) utils.Result[[]byte] {
	body, err := json.Marshal(data)
	if err != nil {
		return utils.Result[[]byte]{Err: err}
	}
	return utils.Result[[]byte]{Data: body}
}

type OptionsQueue struct {
	Name string
	Key  string
}
