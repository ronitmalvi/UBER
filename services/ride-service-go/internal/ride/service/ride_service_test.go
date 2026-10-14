package service

import (
	"context"
	"errors"
	"testing"

	"github.com/ronitmalvi/UBER/ride-service-go/internal/event"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/ride/handler/dto"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/model"
)

// fakeRideRepository replaces the real database in unit tests.
type fakeRideRepository struct {
	rides     map[uint]*model.Ride
	nextID    uint
	createErr error
	updateErr error
	getErr    error
}

func newFakeRideRepository() *fakeRideRepository {
	return &fakeRideRepository{
		rides:  make(map[uint]*model.Ride),
		nextID: 1,
	}
}

func (f *fakeRideRepository) Create(ride *model.Ride) error {
	if f.createErr != nil {
		return f.createErr
	}

	if ride.ID == 0 {
		ride.ID = f.nextID
		f.nextID++
	}

	f.rides[ride.ID] = ride
	return nil
}

func (f *fakeRideRepository) Update(ride *model.Ride) error {
	if f.updateErr != nil {
		return f.updateErr
	}

	f.rides[ride.ID] = ride
	return nil
}

func (f *fakeRideRepository) GetByID(id uint) (*model.Ride, error) {
	if f.getErr != nil {
		return nil, f.getErr
	}

	ride, exists := f.rides[id]
	if !exists {
		return nil, errors.New("ride not found")
	}

	return ride, nil
}

func TestCreateRideSuccessPublishesEvent(t *testing.T) {
	repo := newFakeRideRepository()
	bus := event.NewInMemoryBus()

	eventReceived := false

	bus.Subscribe(event.EventRideRequested, func(
		ctx context.Context,
		e event.Event,
	) error {
		rideEvent, ok := e.(event.RideRequested)
		if !ok {
			t.Fatal("expected RideRequested event")
		}

		eventReceived = true

		if rideEvent.RideID != 1 {
			t.Errorf("expected RideID 1, got %d", rideEvent.RideID)
		}

		return nil
	})

	svc := NewRideService(repo, bus)

	req := &dto.CreateRideRequest{
		RiderID:              10,
		PickupLatitude:       28.6315,
		PickupLongitude:      77.2167,
		DestinationLatitude:  28.5355,
		DestinationLongitude: 77.3910,
	}

	ride, err := svc.CreateRide(context.Background(), req)
	if err != nil {
		t.Fatalf("expected ride creation to succeed, got: %v", err)
	}

	if ride.ID != 1 {
		t.Errorf("expected ride ID 1, got %d", ride.ID)
	}

	if ride.Status != model.RideRequested {
		t.Errorf("expected REQUESTED status, got %v", ride.Status)
	}

	if !eventReceived {
		t.Fatal("expected ride.requested event to be published")
	}

	if ride.EstimatedFare <= 0 {
		t.Errorf("expected positive estimated fare, got %v", ride.EstimatedFare)
	}
}

func TestCreateRideRejectsMissingRiderID(t *testing.T) {
	repo := newFakeRideRepository()
	bus := event.NewInMemoryBus()
	svc := NewRideService(repo, bus)

	req := &dto.CreateRideRequest{
		RiderID:              0,
		PickupLatitude:       28.6315,
		PickupLongitude:      77.2167,
		DestinationLatitude:  28.5355,
		DestinationLongitude: 77.3910,
	}

	ride, err := svc.CreateRide(context.Background(), req)

	if err == nil {
		t.Fatal("expected validation error for missing RiderID")
	}

	if ride != nil {
		t.Fatal("expected no ride to be returned for invalid input")
	}

	if len(repo.rides) != 0 {
		t.Fatal("invalid request should not create a ride")
	}
}

func TestCreateRideRejectsSamePickupAndDestination(t *testing.T) {
	repo := newFakeRideRepository()
	bus := event.NewInMemoryBus()
	svc := NewRideService(repo, bus)

	req := &dto.CreateRideRequest{
		RiderID:              10,
		PickupLatitude:       28.6315,
		PickupLongitude:      77.2167,
		DestinationLatitude:  28.6315,
		DestinationLongitude: 77.2167,
	}

	_, err := svc.CreateRide(context.Background(), req)

	if err == nil {
		t.Fatal("expected validation error for identical coordinates")
	}

	if len(repo.rides) != 0 {
		t.Fatal("invalid request should not create a ride")
	}
}

func TestCreateRideReturnsRepositoryError(t *testing.T) {
	repo := newFakeRideRepository()
	repo.createErr = errors.New("database unavailable")

	bus := event.NewInMemoryBus()
	svc := NewRideService(repo, bus)

	req := &dto.CreateRideRequest{
		RiderID:              10,
		PickupLatitude:       28.6315,
		PickupLongitude:      77.2167,
		DestinationLatitude:  28.5355,
		DestinationLongitude: 77.3910,
	}

	_, err := svc.CreateRide(context.Background(), req)

	if err == nil {
		t.Fatal("expected repository error")
	}

	if err.Error() != "database unavailable" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestHandleDriverAssignedUpdatesRide(t *testing.T) {
	repo := newFakeRideRepository()
	bus := event.NewInMemoryBus()

	ride := &model.Ride{
		RiderID:              10,
		PickupLatitude:       28.6315,
		PickupLongitude:      77.2167,
		DestinationLatitude:  28.5355,
		DestinationLongitude: 77.3910,
		Status:               model.RideRequested,
	}

	if err := repo.Create(ride); err != nil {
		t.Fatalf("failed to create fake ride: %v", err)
	}

	svc := NewRideService(repo, bus)

	err := svc.HandleDriverAssigned(
		context.Background(),
		event.DriverAssigned{
			RideID:   ride.ID,
			DriverID: 42,
		},
	)
	if err != nil {
		t.Fatalf("expected assignment handler to succeed, got: %v", err)
	}

	updatedRide, err := repo.GetByID(ride.ID)
	if err != nil {
		t.Fatalf("failed to retrieve ride: %v", err)
	}

	if updatedRide.DriverID == nil {
		t.Fatal("expected DriverID to be set")
	}

	if *updatedRide.DriverID != 42 {
		t.Errorf("expected DriverID 42, got %d", *updatedRide.DriverID)
	}

	if updatedRide.Status != model.RideDriverAssigned {
		t.Errorf(
			"expected DRIVER_ASSIGNED status, got %v",
			updatedRide.Status,
		)
	}
}
