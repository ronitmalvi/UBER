package service

import (
	"errors"
	"context"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/model"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/ride/repository"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/ride/handler/dto"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/event"
)

type RideService struct {					
	repo *repository.RideRepository  //stores the repository for ride service.
	driverMatcher DriverMatcher  //stores the driver matcher for ride service.
	eventBus event.Bus  //stores the event bus for ride service.
}

//Constructor function for RideService.
func NewRideService(
	repo *repository.RideRepository,
	driverMatcher DriverMatcher,
	bus event.Bus,
) *RideService {

	return &RideService{
		repo: repo,
		driverMatcher: driverMatcher,
		eventBus: bus,
	}
}

//CreateRide creates a new ride in the database.
func (s *RideService) CreateRide(
	ctx context.Context,
	req *dto.CreateRideRequest,
) (*model.Ride, error) {

	ride := &model.Ride{
		RiderID: req.RiderID,
		PickupLatitude: req.PickupLatitude,
		PickupLongitude: req.PickupLongitude,
		DestinationLatitude: req.DestinationLatitude,
		DestinationLongitude: req.DestinationLongitude,
	}

	if err := s.ValidateRide(ride); err != nil {
		return nil, err
	}

	s.initializeRide(ride)

	s.calculateEstimatedFare(ride)

	if err := s.repo.Create(ride); err != nil {
		return nil, err
	}

	rideRequested := event.RideRequested{
		RideID:            ride.ID,
		PickupLatitude:    ride.PickupLatitude,
		PickupLongitude:   ride.PickupLongitude,
	}

	if err := s.eventBus.Publish(ctx, rideRequested); err != nil {
		return nil, err
	}
	
	return ride, nil
}

func (s *RideService) GetRideByID(
	id uint,
) (*model.Ride, error){
	ride, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return ride, nil
}

func (s *RideService) ValidateRide(
	ride *model.Ride,
) error {
	if ride.RiderID == 0 {
		return errors.New("RiderID is required")
	}

	if ride.PickupLatitude == ride.DestinationLatitude &&
		ride.PickupLongitude == ride.DestinationLongitude {

		return errors.New("pickup and destination cannot be same")
	}

	return nil
}

func (s *RideService) initializeRide(
	ride *model.Ride,
) {
	ride.Status = model.RideRequested
}

func (s *RideService) calculateEstimatedFare(
	ride *model.Ride,
) {

	ride.EstimatedFare = 250
}


