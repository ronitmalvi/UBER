package event

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

