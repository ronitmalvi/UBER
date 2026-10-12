package service

import (
	"context"
	"errors"

	"github.com/ronitmalvi/UBER/ride-service-go/internal/event"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/model"
)

func (s *RideService) HandleDriverAssigned(
	ctx context.Context,
	e event.Event,
) error {
	assigned, ok := e.(event.DriverAssigned)
	if !ok {
		return errors.New("invalid event type for driver assignment")
	}

	ride, err := s.GetRideByID(assigned.RideID)
	if err != nil {
		return err
	}

	ride.DriverID = &assigned.DriverID
	ride.Status = model.RideDriverAssigned

	if err := s.repo.Update(ride); err != nil {
		return err
	}

	return nil
}
