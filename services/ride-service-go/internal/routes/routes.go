package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/handler"
	"fmt"
)

func RegisterRoutes(
	router *gin.Engine,
	rideHandler *handler.RideHandler,
) {
	fmt.Println("Registering routes...")
	router.POST(
		"/rides",
		rideHandler.CreateRide,
	)
	router.GET(
		"/rides/:id",
		rideHandler.GetRideByID,
	)
}