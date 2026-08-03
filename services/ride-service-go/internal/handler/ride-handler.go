package handler

import (
	"github.com/ronitmalvi/UBER/ride-service-go/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/handler/dto"
)

type RideHandler struct {
	service *service.RideService			//A RideHandler contains a RideService.
}

func NewRideHandler(						//Constructor function for RideHandler.
	service *service.RideService,
) *RideHandler {

	return &RideHandler{
		service: service,
	}
}

func (h *RideHandler) CreateRide(
	c *gin.Context,
) {
	fmt.Println("Creating ride...(handler layer)")
	var req dto.CreateRideRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": err.Error(),
			},
		)
		
		return
	}
	

	ride, err := h.service.CreateRide(&req);

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	response := dto.CreateRideResponse{
		ID:      ride.ID,
		RiderID: ride.RiderID,
		Status:  string(ride.Status),
		Fare:    ride.EstimatedFare,
	}

	c.JSON(http.StatusCreated, response)
}