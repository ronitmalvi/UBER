package dservice

import (
	"context"
	"errors"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/event"
)

func (s *DriverService) HandleRideRequested(
	ctx context.Context,
	e event.Event,
) error {

	rideRequested, ok := e.(event.RideRequested)
	if !ok {
		return errors.New("invalid event type")
	}

	return s.AssignDriverToRide(
		ctx,
		rideRequested,
	)
}