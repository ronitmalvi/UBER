package event

import (
	"context"
	"errors"
	"testing"
)

func TestPublishCallsSubscribedHandler(t *testing.T) {
	bus := NewInMemoryBus()

	called := false

	bus.Subscribe(EventRideRequested, func(
		ctx context.Context,
		e Event,
	) error {
		called = true

		rideEvent, ok := e.(RideRequested)
		if !ok {
			t.Fatal("expected RideRequested event")
		}

		if rideEvent.RideID != 123 {
			t.Errorf(
				"expected RideID 123, got %d",
				rideEvent.RideID,
			)
		}

		return nil
	})

	err := bus.Publish(context.Background(), RideRequested{
		RideID:         123,
		PickupLatitude: 28.6315,
		PickupLongitude: 77.2167,
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !called {
		t.Fatal("expected subscribed handler to be called")
	}
}

func TestPublishReturnsHandlerError(t *testing.T) {
	bus := NewInMemoryBus()

	expectedErr := errors.New("handler failed")

	bus.Subscribe(EventRideRequested, func(
		ctx context.Context,
		e Event,
	) error {
		return expectedErr
	})

	err := bus.Publish(context.Background(), RideRequested{
		RideID: 123,
	})

	if !errors.Is(err, expectedErr) {
		t.Fatalf(
			"expected error %v, got %v",
			expectedErr,
			err,
		)
	}
}

func TestPublishWithoutSubscribers(t *testing.T) {
	bus := NewInMemoryBus()

	err := bus.Publish(context.Background(), RideRequested{
		RideID: 123,
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestEventNames(t *testing.T) {
	rideEvent := RideRequested{}
	if rideEvent.Name() != EventRideRequested {
		t.Errorf(
			"expected %q, got %q",
			EventRideRequested,
			rideEvent.Name(),
		)
	}

	driverEvent := DriverAssigned{}
	if driverEvent.Name() != EventDriverAssigned {
		t.Errorf(
			"expected %q, got %q",
			EventDriverAssigned,
			driverEvent.Name(),
		)
	}
}
