package event

type DriverAssigned struct {
	RideID   uint
	DriverID uint
}

func (DriverAssigned) Name() string {
	return EventDriverAssigned		//driver.assigned event name
}