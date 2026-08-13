package service

import (
	"fmt"
	"context"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/redis"
	goredis "github.com/redis/go-redis/v9"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/repository"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/model"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/handler/dto"
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

func (s *DriverService) CreateDriver(
	req *dto.CreateDriverRequest,
) (*model.Driver, error) {
	driver := &model.Driver{
		Name:          req.Name,
		Phone:         req.Phone,
		LicenseNumber: req.LicenseNumber,
		VehicleNumber: req.VehicleNumber,
		VehicleType:   req.VehicleType,
		IsOnline:      false,
		IsAvailable:   false,
	}
	if err := s.repo.Create(driver); err != nil {
		return nil, err
	}

	return driver, nil
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

func (s *DriverService) GoOnline(ctx context.Context, driverID uint) error {
	// Update the driver's online status in the database
	driver, err := s.repo.GetByID(driverID)
	if err != nil {
		return fmt.Errorf("failed to retrieve driver: %w", err)
	}

	driver.IsOnline = true
	driver.IsAvailable = true

	return s.repo.Update(driver)
}

func (s *DriverService) GoOffline(
	ctx context.Context,
	driverID uint,
) error {

	driver, err := s.repo.GetByID(driverID)

	if err != nil {
		return err
	}

	driver.IsOnline = false
	driver.IsAvailable = false

	if err := s.repo.Update(driver); err != nil {
		return err
	}

	return redis.RemoveDriverLocation(
		ctx,
		s.redisClient,
		driverID,
	)
}

func (s *DriverService) FindNearbyDrivers(
	ctx context.Context,
	latitude float64,
	longitude float64,
	radiusKm float64,
) ([]goredis.GeoLocation, error) {

	if err := validateCoordinates(latitude, longitude); err != nil {
		return nil, err
	}

	if radiusKm <= 0 {
		return nil, fmt.Errorf("radius must be greater than zero")
	}

	return redis.FindNearbyDrivers(
		ctx,
		s.redisClient,
		latitude,
		longitude,
		radiusKm,
	)
}