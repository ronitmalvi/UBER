package service

import (
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