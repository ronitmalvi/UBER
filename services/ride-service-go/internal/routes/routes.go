package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ronitmalvi/UBER/ride-service-go/internal/handler"
	"fmt"
)

func RegisterRoutes(
	router *gin.Engine,
	rideHandler *handler.RideHandler,
	driverHandler *handler.DriverHandler,
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
	router.PUT(
		"/drivers/:id/location",
		driverHandler.UpdateLocation,
	)
	router.POST(
		"/drivers/:id/online",
		driverHandler.GoOnline,
	)

	router.POST(
		"/drivers/:id/offline",
		driverHandler.GoOffline,
	)

	router.POST(
		"/drivers",
		driverHandler.CreateDriver,
	)

	router.POST(
		"/drivers/nearby",
		driverHandler.FindNearbyDrivers,
	)
}