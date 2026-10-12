package event

type RideRequested struct {
	RideID uint

	PickupLatitude float64
	PickupLongitude float64
}

func (RideRequested) Name() string {
	return EventRideRequested			//ride.requested event name
}