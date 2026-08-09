package service

import (
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

	return redis.AddDriverLocation(
		ctx,
		s.redisClient,
		driverID,
		latitude,
		longitude,
	)
}