package local

import (
	"accounts/internal/core/domain/event"
	"accounts/internal/utils"
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type ActionFunc = func(context.Context, event.DomainEvent)

type LocalEventBus struct {
	actions map[string]ActionFunc
}

func NewLocalEventBus() *LocalEventBus {
	return &LocalEventBus{
		actions: make(map[string]ActionFunc),
	}
}

func (l *LocalEventBus) AddAction(eventName string, action ActionFunc) {
	l.actions[eventName] = action
}

func (l *LocalEventBus) Publish(events []event.DomainEvent) error {
	ctx := context.Background()

	for _, e := range events {
		eventName := e.EventName()

		action, ok := l.actions[eventName]
		if ok {
			action(ctx, e)
		} else {
			fmt.Print("Event not found: ", eventName)
		}
	}
	return nil
}

func (l *LocalEventBus) Consume(queue, key string) utils.Result[<-chan amqp091.Delivery] {
	return utils.Result[<-chan amqp091.Delivery]{
		Data: nil,
		Err:  nil,
	}
}
