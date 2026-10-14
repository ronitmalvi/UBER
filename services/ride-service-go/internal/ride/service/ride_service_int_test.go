package service_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ronitmalvi/UBER/ride-service-go/internal/config"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/database"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/driver/drepository"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/driver/dservice"
	//driverDTO "github.com/ronitmalvi/UBER/ride-service-go/internal/driver/dhandler/dto"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/event"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/model"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/redis"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/ride/handler/dto"
	rideRepo "github.com/ronitmalvi/UBER/ride-service-go/internal/ride/repository"
	rideService "github.com/ronitmalvi/UBER/ride-service-go/internal/ride/service"
)

func TestRideAssignmentIntegration(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") != "1" {
		t.Skip("set RUN_INTEGRATION_TESTS=1 to run integration tests")
	}

	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		t.Fatalf("connect PostgreSQL: %v", err)
	}

	redisClient, err := redis.Connect(cfg)
	if err != nil {
		t.Fatalf("connect Redis: %v", err)
	}
	t.Cleanup(func() {
		_ = redisClient.Close()
	})

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bus := event.NewInMemoryBus()

	driverRepository := drepository.NewDriverRepository(db)
	driverService := dservice.NewDriverService(driverRepository, redisClient, bus)

	ridesRepository := rideRepo.NewRideRepository(db)
	ridesService := rideService.NewRideService(ridesRepository, bus)

	bus.Subscribe(event.EventRideRequested, driverService.HandleRideRequested)
	bus.Subscribe(event.EventDriverAssigned, ridesService.HandleDriverAssigned)

	// Use a unique test driver to avoid conflicting with existing records.
	suffix := time.Now().UnixNano()
	driver := &model.Driver{
		Name:          fmt.Sprintf("Integration Driver %d", suffix),
		Phone:         fmt.Sprintf("test-%d", suffix),
		LicenseNumber: fmt.Sprintf("test-license-%d", suffix),
		VehicleNumber: fmt.Sprintf("test-vehicle-%d", suffix),
		VehicleType:   "CAR",
		IsOnline:      false,
		IsAvailable:   false,
	}

	if err := driverRepository.Create(driver); err != nil {
		t.Fatalf("create test driver: %v", err)
	}

	// Clean up only this test's records.
	t.Cleanup(func() {
		_ = redis.RemoveDriverLocation(context.Background(), redisClient, driver.ID)
		_ = db.Delete(&model.Ride{}, "rider_id = ?", uint(987654321)).Error
		_ = db.Delete(&model.Driver{}, driver.ID).Error
	})

	if err := driverService.GoOnline(ctx, driver.ID); err != nil {
		t.Fatalf("bring driver online: %v", err)
	}

	const pickupLat = 28.6315
	const pickupLon = 77.2167

	if err := driverService.UpdateLocation(ctx, driver.ID, pickupLat, pickupLon); err != nil {
		t.Fatalf("update driver location: %v", err)
	}

	req := &dto.CreateRideRequest{
		RiderID:              987654321,
		PickupLatitude:       pickupLat,
		PickupLongitude:      pickupLon,
		DestinationLatitude:  28.5355,
		DestinationLongitude: 77.3910,
	}

	createdRide, err := ridesService.CreateRide(ctx, req)
	if err != nil {
		t.Fatalf("create and assign ride: %v", err)
	}

	persistedRide, err := ridesService.GetRideByID(createdRide.ID)
	if err != nil {
		t.Fatalf("get persisted ride: %v", err)
	}

	if persistedRide.Status != model.RideDriverAssigned {
		t.Fatalf("expected assigned status, got %v", persistedRide.Status)
	}

	if persistedRide.DriverID == nil {
		t.Fatal("expected DriverID to be set, got nil")
	}

	if *persistedRide.DriverID != driver.ID {
		t.Fatalf(
			"expected driver ID %d, got %d",
			driver.ID,
			*persistedRide.DriverID,
		)
	}

	persistedDriver, err := driverRepository.GetByID(driver.ID)
	if err != nil {
		t.Fatalf("get persisted driver: %v", err)
	}

	if persistedDriver.IsAvailable {
		t.Fatal("expected assigned driver to be unavailable")
	}
}
