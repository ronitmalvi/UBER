package handler

import (
	"github.com/ronitmalvi/UBER/ride-service-go/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/model"
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

type CreateRideRequest struct {

	RiderID uint `json:"rider_id"`

	Fare float64 `json:"fare"`
}

func (h *RideHandler) CreateRide(
	c *gin.Context,
) {
	fmt.Println("Creating ride...(handler layer)")
	var req CreateRideRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		
		return
	}
	ride := &model.Ride{
		RiderID: req.RiderID,
		Fare:    req.Fare,
	}

	err := h.service.CreateRide(ride)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		ride,
	)
}