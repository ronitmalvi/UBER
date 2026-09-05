package service

import (
	"context"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/handler/dto"
)

type DriverMatcher interface {
	FindBestDriver(
		ctx context.Context,
		latitude float64,
		longitude float64,
		radiusKm float64,
	) (*dto.BestDriver, error)
}