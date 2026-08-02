package model

import "gorm.io/gorm"

type RideStatus string

const (
	RideRequested RideStatus = "REQUESTED"
)

type Ride struct {

    gorm.Model

    RiderID uint

    DriverID *uint

    PickupLatitude float64
    PickupLongitude float64

    DestinationLatitude float64
    DestinationLongitude float64

    EstimatedFare float64

    FinalFare *float64

    Status RideStatus
}