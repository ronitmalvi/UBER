package model

import "gorm.io/gorm"

type Driver struct {
	gorm.Model

	Name string

	Phone string

	LicenseNumber string

	VehicleNumber string

	VehicleType string

	IsOnline bool

	IsAvailable bool
}