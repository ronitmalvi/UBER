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

func (h *DriverHandler) CreateDriver(c *gin.Context) {
	var req dto.CreateDriverRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	driver, err := h.service.CreateDriver(&req)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "driver created successfully",
			"driver":  driver,
		},
	)
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

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "driver location updated successfully",
		},
	)
}

func (h *DriverHandler) GoOnline(c *gin.Context) {

	driverID, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid driver ID"},
		)
		return
	}

	err = h.service.GoOnline(
		c.Request.Context(),
		uint(driverID),
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "driver is now online",
		},
	)
}

func (h *DriverHandler) GoOffline(c *gin.Context) {

	driverID, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid driver ID"},
		)
		return
	}

	err = h.service.GoOffline(
		c.Request.Context(),
		uint(driverID),
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "driver is now offline",
		},
	)
}