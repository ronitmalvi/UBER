package dto

type UpdateDriverLocationRequest struct {
	Latitude  float64 `json:"latitude" binding:"required"`               //The DTO simply represents the incoming HTTP request.
	Longitude float64 `json:"longitude" binding:"required"`
}