package service

import (
	"errors"
	"fmt"

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

	if ride.Fare <= 0 {
		fmt.Println("fare must be greater than zero")
		return errors.New("fare must be greater than zero")
	}

	ride.Status = model.RideRequested

	return s.repo.Create(ride)
}


