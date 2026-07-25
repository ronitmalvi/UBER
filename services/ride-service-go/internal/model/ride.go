package model

import "gorm.io/gorm"

type RideStatus string

const (
	RideRequested RideStatus = "REQUESTED"
)

type Ride struct {
	gorm.Model

	RiderID uint

	Status RideStatus

	Fare float64
}