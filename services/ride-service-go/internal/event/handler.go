package event

import "context"

type Handler func(
	ctx context.Context,
	event Event,
) error