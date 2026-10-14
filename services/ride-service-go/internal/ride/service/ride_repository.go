package service

import "github.com/ronitmalvi/UBER/ride-service-go/internal/model"

type RideRepository interface {
	Create(ride *model.Ride) error
	Update(ride *model.Ride) error
	GetByID(id uint) (*model.Ride, error)
}
