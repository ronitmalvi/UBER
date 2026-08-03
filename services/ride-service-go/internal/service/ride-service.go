package service

import (
	"errors"
	// "math"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/model"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/repository"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/handler/dto"
)

type RideService struct {						//stores the repository for ride service.
	repo *repository.RideRepository
}

//Constructor function for RideService.
func NewRideService(
	repo *repository.RideRepository,
) *RideService {

	return &RideService{
		repo: repo,
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


