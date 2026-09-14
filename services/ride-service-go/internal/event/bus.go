package event

import "context"

type Bus interface {
	Publish(
		ctx context.Context,
		event Event,
	) error

	Subscribe(
		eventName string,
		handler Handler,
	)
}