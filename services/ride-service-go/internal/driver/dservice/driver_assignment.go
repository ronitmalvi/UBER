package dservice

import (
	"context"
	"errors"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/event"
)

var ErrNoDriversAvailable = errors.New("no drivers available")

func (s *DriverService) AssignDriverToRide(
	ctx context.Context,
	rideRequested event.RideRequested,
) error {

	bestDriver, err := s.FindBestDriver(
		ctx,
		rideRequested.PickupLatitude,
		rideRequested.PickupLongitude,
		5,
	)

	if err != nil {
		if errors.Is(err, ErrNoDriversAvailable) {
			return nil
		}

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