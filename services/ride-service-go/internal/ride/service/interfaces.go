package service

import (
	"context"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/driver/dhandler/dto"
)

type DriverMatcher interface {
	FindBestDriver(
		ctx context.Context,
		latitude float64,
		longitude float64,
		radiusKm float64,
	) (*dto.BestDriver, error)
	MarkBusy(
		ctx context.Context,
		driverID uint,
	) error
}