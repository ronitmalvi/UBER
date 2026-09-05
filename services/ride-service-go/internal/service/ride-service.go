package service

import (
	"errors"
	"context"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/model"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/repository"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/handler/dto"
)

type RideService struct {					
	repo *repository.RideRepository  //stores the repository for ride service.
	driverMatcher DriverMatcher  //stores the driver matcher for ride service.
}

//Constructor function for RideService.
func NewRideService(
	repo *repository.RideRepository,
	driverMatcher DriverMatcher,
) *RideService {

	return &RideService{
		repo: repo,
		driverMatcher: driverMatcher,
	}
}

//CreateRide creates a new ride in the database.
func (s *RideService) CreateRide(
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
	bestDriver, err := s.driverMatcher.FindBestDriver(
		context.Background(),
		ride.PickupLatitude,
		ride.PickupLongitude,
		5,
	)
	if err != nil {
		return ride, nil
	}
	ride.DriverID = &bestDriver.DriverID
	ride.Status = model.RideDriverAssigned
	if err := s.repo.Update(ride); err != nil {
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


