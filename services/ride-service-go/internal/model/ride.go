package model

import "gorm.io/gorm"

type RideStatus string

const (
	RideRequested       RideStatus = "REQUESTED"
	RideSearchingDriver RideStatus = "SEARCHING_DRIVER"
	RideDriverAssigned  RideStatus = "DRIVER_ASSIGNED"
	RideDriverArrived   RideStatus = "DRIVER_ARRIVED"
	RideStarted         RideStatus = "STARTED"
	RideCompleted       RideStatus = "COMPLETED"
	RideCancelled       RideStatus = "CANCELLED"
)

type Ride struct {

    gorm.Model

    RiderID uint

    DriverID *uint  // Pointer to allow null value when no driver is assigned

    PickupLatitude float64
    PickupLongitude float64

    DestinationLatitude float64
    DestinationLongitude float64

    EstimatedFare float64

    FinalFare *float64  // Pointer to allow null value when final fare is not yet calculated

    Status RideStatus
}

func (r *Ride) UpdateStatus(status RideStatus) {
	r.Status = status
}