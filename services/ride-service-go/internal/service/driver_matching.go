package service

import (
	"context"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/handler/dto"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/redis"
	"strconv"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/errors"
)

func (s *DriverService) FindNearbyDrivers(
    ctx context.Context,
    latitude float64,
    longitude float64,
    radiusKm float64,
) ([]dto.NearbyDriver, error) {

    if err := validateCoordinates(latitude, longitude); err != nil {
        return nil, err
    }

    results, err := redis.FindNearbyDrivers(
        ctx,
        s.redisClient,
        latitude,
        longitude,
        radiusKm,
    )

    if err != nil {
        return nil, err
    }

    nearbyDrivers := make([]dto.NearbyDriver, 0)

    for _, driver := range results {

        driverID, err := strconv.ParseUint(driver.Name, 10, 64)

        if err != nil {
            continue
        }

        nearbyDrivers = append(
            nearbyDrivers,
            dto.NearbyDriver{
                DriverID: uint(driverID),
                Distance: driver.Dist,
            },
        )
    }

    return nearbyDrivers, nil
}

func (s *DriverService) FindBestDriver(
	ctx context.Context,
	latitude float64,
	longitude float64,
	radiusKm float64,
) (*dto.BestDriver, error) {

	if err := validateCoordinates(latitude, longitude); err != nil {
		return nil, err
	}

	results, err := redis.FindNearbyDrivers(
		ctx,
		s.redisClient,
		latitude,
		longitude,
		radiusKm,
	)

	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, errors.ErrNoDriversAvailable
	}

	best := results[0]

	driverID, err := strconv.ParseUint(best.Name, 10, 64)
	if err != nil {
		return nil, err
	}

	return &dto.BestDriver{
		DriverID: uint(driverID),
		Distance: best.Dist,
	}, nil
}