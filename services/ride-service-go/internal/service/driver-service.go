package service

import (
	"github.com/ronitmalvi/UBER/ride-service-go/internal/repository"
)

type DriverService struct {
	repository *repository.DriverRepository
}

func NewDriverService(
	repository *repository.DriverRepository,
) *DriverService {
	return &DriverService{
		repository: repository,
	}
}