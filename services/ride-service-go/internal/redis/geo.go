package redis

import (
	"context"
	"strconv"

	goredis "github.com/redis/go-redis/v9"
)

func AddDriverLocation(
	ctx context.Context,
	client *goredis.Client,
	driverID uint,
	latitude float64,
	longitude float64,
) error {

	return client.GeoAdd(
		ctx,
		"drivers",                //key for the Redis geospatial index
		&goredis.GeoLocation{
			Name:      strconv.FormatUint(uint64(driverID), 10),
			Longitude: longitude,
			Latitude:  latitude,
		},
	).Err()
}

func RemoveDriverLocation(
	ctx context.Context,
	client *goredis.Client,
	driverID uint,
) error {

	return client.ZRem(
		ctx,
		"drivers",
		strconv.FormatUint(uint64(driverID), 10),
	).Err()
}

func FindNearbyDrivers(
	ctx context.Context,
	client *goredis.Client,
	latitude float64,
	longitude float64,
	radiusKm float64,
) ([]goredis.GeoLocation, error) {

	results, err := client.GeoSearchLocation(
		ctx,
		"drivers",
		&goredis.GeoSearchLocationQuery{
			GeoSearchQuery: goredis.GeoSearchQuery{
				Longitude:  longitude,
				Latitude:   latitude,
				Radius:     radiusKm,
				RadiusUnit: "km",
				Sort:      "ASC",
			},
			WithDist: true,
			WithCoord: true,
			
		},
	).Result()

	if err != nil {
		return nil, err
	}

	return results, nil
}