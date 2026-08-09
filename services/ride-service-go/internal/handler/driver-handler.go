package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/handler/dto"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/service"
)
type DriverHandler struct {
	service *service.DriverService
}

func NewDriverHandler(
	service *service.DriverService,
) *DriverHandler{
	return &DriverHandler{
		service: service,
	}
}

func (h *DriverHandler) UpdateLocation(c *gin.Context) {
	driverIDParam := c.Param("id")

	driverID, err := strconv.ParseUint(driverIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid driver ID"})
		return
	}

	var req dto.UpdateDriverLocationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = h.service.UpdateLocation(
		c.Request.Context(),
		uint(driverID),
		req.Latitude,
		req.Longitude,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update driver location"})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "driver location updated successfully",
		},
	)
}