package dto

type CreateRideResponse struct {
	ID       uint   `json:"id"`
	RiderID  uint   `json:"rider_id"`
	Status   string `json:"status"`
	Fare     float64 `json:"fare"`
}