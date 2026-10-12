package dservice

import (
	goredis "github.com/redis/go-redis/v9"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/driver/drepository"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/model"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/driver/dhandler/dto"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/event"
)

type DriverService struct {
	repo *drepository.DriverRepository
	redisClient *goredis.Client
	eventBus event.Bus
}

func NewDriverService(
	repo *drepository.DriverRepository,
	redisClient *goredis.Client,
	eventBus event.Bus,
) *DriverService {
	return &DriverService{
		repo: repo,
		redisClient: redisClient,
		eventBus: eventBus,
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