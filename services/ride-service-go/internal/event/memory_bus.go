package event

import (
	"context"
)

type InMemoryBus struct {
	handlers map[string][]Handler
}

func NewInMemoryBus() *InMemoryBus {

	return &InMemoryBus{
		handlers: make(map[string][]Handler),
	}
}

func (b *InMemoryBus) Subscribe(
	eventName string,
	handler Handler,
) {

	b.handlers[eventName] = append(
		b.handlers[eventName],
		handler,
	)
}

func (b *InMemoryBus) Publish(
    ctx context.Context,
    event Event,
) error {

    eventName := event.Name()

    handlers := b.handlers[eventName]

    for _, handler := range handlers {

        if err := handler(ctx, event); err != nil {
            return err
        }
    }

    return nil
}