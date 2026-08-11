package service

import (
	"fmt"
	"context"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/redis"
	goredis "github.com/redis/go-redis/v9"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/repository"
)

type DriverService struct {
	repo *repository.DriverRepository
	redisClient *goredis.Client
}

func NewDriverService(
	repo *repository.DriverRepository,
	redisClient *goredis.Client,
) *DriverService {
	return &DriverService{
		repo: repo,
		redisClient: redisClient,
	}
}

func (s *DriverService) UpdateLocation(
	ctx context.Context,
	driverID uint,
	latitude float64,
	longitude float64,
) error {

	if err := validateCoordinates(latitude, longitude); err != nil {
		return err
	}

	return redis.AddDriverLocation(
		ctx,
		s.redisClient,
		driverID,
		latitude,
		longitude,
	)
}

func validateCoordinates(
	latitude float64,
	longitude float64,
) error {

	if latitude < -90 || latitude > 90 {
		return fmt.Errorf("latitude must be between -90 and 90")
	}

	if longitude < -180 || longitude > 180 {
		return fmt.Errorf("longitude must be between -180 and 180")
	}

	return nil
}