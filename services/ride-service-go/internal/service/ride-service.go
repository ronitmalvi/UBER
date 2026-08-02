package service

import (
	"errors"
	"fmt"
	// "math"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/model"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/repository"
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
	ride *model.Ride,
) error {

	// distance := math.Sqrt(
	// 		math.Pow(
	// 			ride.DestinationLatitude-ride.PickupLatitude,
	// 			2,
	// 		) +
	// 		math.Pow(
	// 			ride.DestinationLongitude-ride.PickupLongitude,
	// 			2,
	// 		),
	// 	)

	ride.EstimatedFare = 250

	if ride.EstimatedFare <= 0 {
		fmt.Println("fare must be greater than zero")
		return errors.New("fare must be greater than zero")
	}

	ride.Status = model.RideRequested
	fmt.Println("Creating ride in the service layer...")
	return s.repo.Create(ride)
}


