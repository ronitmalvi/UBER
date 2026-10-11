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

	bestDriver, err := s.FindBestDriver(
		ctx,
		rideRequested.PickupLatitude,
		rideRequested.PickupLongitude,
		5,
	)
	if err != nil {
		return err
	}

	if err := s.MarkBusy(
		ctx,
		bestDriver.DriverID,
	); err != nil {
		return err
	}

	driverAssigned := event.DriverAssigned{
		RideID:   rideRequested.RideID,
		DriverID: bestDriver.DriverID,
	}

	if err := s.eventBus.Publish(
		ctx,
		driverAssigned,
	); err != nil {
		return err
	}

	return nil
}