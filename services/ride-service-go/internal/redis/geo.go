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