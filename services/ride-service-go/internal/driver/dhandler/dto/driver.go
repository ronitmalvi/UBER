package dto

type UpdateDriverLocationRequest struct {
	Latitude  float64 `json:"latitude" binding:"required"`               //The DTO simply represents the incoming HTTP request.
	Longitude float64 `json:"longitude" binding:"required"`
}

type CreateDriverRequest struct {
	Name          string `json:"name" binding:"required"`
	Phone         string `json:"phone" binding:"required"`
	LicenseNumber string `json:"license_number" binding:"required"`
	VehicleNumber string `json:"vehicle_number" binding:"required"`
	VehicleType   string `json:"vehicle_type" binding:"required"`
}