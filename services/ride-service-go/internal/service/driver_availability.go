package service

import (
	"context"
	"fmt"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/redis"
)


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

func (s *DriverService) MarkBusy(ctx context.Context, driverID uint) error {
	driver, err := s.repo.GetByID(driverID)
	if err != nil {
		return fmt.Errorf("failed to retrieve driver: %w", err)
	}

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