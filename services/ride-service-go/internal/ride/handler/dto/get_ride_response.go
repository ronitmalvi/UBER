package dto

type GetRideResponse struct {
	ID                     uint    `json:"id"`
	RiderID                uint    `json:"rider_id"`
	DriverID               *uint   `json:"driver_id"`
	Status                 string  `json:"status"`
	PickupLatitude         float64 `json:"pickup_latitude"`
	PickupLongitude        float64 `json:"pickup_longitude"`
	DestinationLatitude    float64 `json:"destination_latitude"`
	DestinationLongitude   float64 `json:"destination_longitude"`
	EstimatedFare          float64 `json:"estimated_fare"`
	FinalFare              *float64 `json:"final_fare"`
}