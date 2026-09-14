package event

type RideRequested struct {
	RideID uint

	PickupLatitude float64
	PickupLongitude float64
}

func (RideRequested) Name() string {
	return "ride.requested"				//ride.requested event name
}